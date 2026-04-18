package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"dift_user_insentive/user-coupon-service/config"

	grpcadapter "dift_user_insentive/user-coupon-service/internal/adapter/grpc"
	httpadapter "dift_user_insentive/user-coupon-service/internal/adapter/http"
	ordercompletedconsumer "dift_user_insentive/user-coupon-service/internal/adapter/order_completed_consumer"
	orderfailedconsumer "dift_user_insentive/user-coupon-service/internal/adapter/order_failed_consumer"
	repositoryadapter "dift_user_insentive/user-coupon-service/internal/adapter/repository"
	natsadapter "dift_user_insentive/user-coupon-service/internal/adapter/user_coupon_comsumer"
	worker "dift_user_insentive/user-coupon-service/internal/adapter/worker"

	couponhandler "dift_user_insentive/user-coupon-service/internal/service_logic/handler/coupon_event"
	orderhandler "dift_user_insentive/user-coupon-service/internal/service_logic/handler/order"
	assignservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/coupon_event"
	usageservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/order_complete"
	releaseservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/order_failed"
	queryservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/query"
	reserveservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/reserve_service"
	expireservice "dift_user_insentive/user-coupon-service/internal/service_logic/service/worker"
	servicecore "dift_user_insentive/user-coupon-service/internal/servicecore"
	route "dift_user_insentive/user-coupon-service/route"

	grpcinfra "dift_user_insentive/user-coupon-service/internal/integration/grpc_server"
	httpinfra "dift_user_insentive/user-coupon-service/internal/integration/http"
	natsinfra "dift_user_insentive/user-coupon-service/internal/integration/nats"
	dbinfra "dift_user_insentive/user-coupon-service/internal/integration/postgres"

	pb "dift_user_insentive/user-coupon-service/proto/pb/usercoupon"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.LoadConfig()

	db, err := dbinfra.New(dbinfra.Config{
		DSN:             cfg.Database.DSN,
		MaxOpenConns:    cfg.Database.MaxOpenConns,
		MaxIdleConns:    cfg.Database.MaxIdleConns,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbinfra.Close(db); err != nil {
			log.Println("failed to close database:", err)
		}
	}()

	repo := repositoryadapter.NewUserCouponRepository(db)
	idemRepo := repositoryadapter.NewIdempotencyRepository(db)
	usageRepo := repositoryadapter.NewUsageRepository(db)

	assignSvc := assignservice.NewCouponAssignService(
		repo,
	)
	usageSvc := usageservice.NewCouponUsageService(
		repo,
		idemRepo,
		usageRepo,
	)
	releaseSvc := releaseservice.NewCouponReleaseService(repo)
	querySvc := queryservice.NewCouponQueryService(repo)
	reserveSvc := reserveservice.NewCouponReserveService(repo)
	expireSvc := expireservice.NewCouponExpireService(repo)

	couponEventHandler := couponhandler.NewCouponEventHandler(assignSvc)
	orderCompletedHandler := orderhandler.NewOrderCompletedHandler(usageSvc)
	orderFailedHandler := orderhandler.NewOrderFailedHandler(releaseSvc)

	queryHandler := httpadapter.New(querySvc)

	mw := httpadapter.NewMiddlewareSet("user-coupon-service")
	router := gin.New()
	servicecore.UseDefaultMiddlewares(router)
	router.Use(
		httpadapter.Recover(mw),
		httpadapter.AccessLog(mw),
		httpadapter.RateLimit(mw),
	)

	route.RegisterRoutes(router, queryHandler)

	httpServer := httpinfra.New(cfg.HTTP.Port, router)
	httpServer.Start()

	grpcHandler := grpcadapter.New(querySvc, reserveSvc)

	grpcPort := strings.TrimPrefix(cfg.GRPC.Address, ":")
	grpcServer := grpcinfra.NewGRPCServer(grpcPort)

	grpcServer.RegisterService(func(s *grpc.Server) {
		pb.RegisterUserCouponServiceServer(s, grpcHandler)
	})

	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Println("gRPC server stopped with error:", err)
		}
	}()

	natsClient, err := natsinfra.NewClient(cfg.NATS.URL)
	if err != nil {
		log.Fatal(err)
	}

	jsConsumer := natsinfra.NewJetStreamConsumer(natsClient.Js)

	if err := natsinfra.EnsureStream(
		natsClient.Js,
		cfg.NATS.Stream,
		[]string{
			cfg.NATS.Subject,
			"order.completed",
			"order.failed",
		},
	); err != nil {
		log.Fatal(err)
	}

	couponConsumer := natsadapter.New(couponEventHandler)

	if err := jsConsumer.Subscribe(
		ctx,
		cfg.NATS.Subject,
		"user-coupon-consumer",
		couponConsumer.Handle,
	); err != nil {
		log.Fatal(err)
	}

	orderCompletedConsumer := ordercompletedconsumer.NewOrderCompletedConsumer(
		orderCompletedHandler,
	)

	if err := jsConsumer.Subscribe(
		ctx,
		"order.completed",
		"user-coupon-order-completed",
		orderCompletedConsumer.Handle,
	); err != nil {
		log.Fatal(err)
	}

	orderFailedConsumer := orderfailedconsumer.NewOrderFailedConsumer(
		orderFailedHandler,
	)

	if err := jsConsumer.Subscribe(
		ctx,
		"order.failed",
		"user-coupon-order-failed",
		orderFailedConsumer.Handle,
	); err != nil {
		log.Fatal(err)
	}

	expiryWorker := worker.New(expireSvc)
	expiryWorker.Start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("shutting down...")

	expiryWorker.Stop()

	grpcServer.Stop(ctx)
	if err := httpServer.Stop(ctx); err != nil {
		log.Println("failed to shutdown HTTP server:", err)
	}

	natsClient.Close()
}
