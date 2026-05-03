package servicecore

// CapabilityProfile documents enterprise package targets for this service.
type CapabilityProfile struct {
	ServiceName         string
	InboundHTTP         bool
	InboundGRPC         bool
	InboundEvent        bool
	OutboundGRPCClient  bool
	OutboundEvent       bool
	HasDBOrMigration    bool
	RecommendedPackages []string
}

func EnterpriseCapabilityProfile() CapabilityProfile {
	return CapabilityProfile{
		ServiceName:         "user-coupon-service",
		InboundHTTP:         true,
		InboundGRPC:         true,
		InboundEvent:        true,
		OutboundGRPCClient:  true,
		OutboundEvent:       false,
		HasDBOrMigration:    true,
		RecommendedPackages: []string{"engine-core/messaging", "engine-core/messaging/idempotency", "engine-core/persistence", "engine-core/runtime/app", "engine-core/transport/grpc", "middleware/auth", "middleware/metrics", "middleware/ratelimit", "middleware/recovery", "middleware/requestid", "middleware/retry", "middleware/timeout", "middleware/tracing", "middleware/validation", "tx/uow + inbox/outbox/dlq"},
	}
}
