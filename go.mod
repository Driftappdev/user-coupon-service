module dift_user_insentive/user-coupon-service

go 1.25.1

require (
	github.com/driftappdev/libpackage/filemods/middleware/adminshield/admin-middleware v0.0.0
	github.com/driftappdev/libpackage/goauth v0.0.0
	github.com/driftappdev/libpackage/gocircuit v0.0.0
	github.com/driftappdev/libpackage/goerror v0.0.0
	github.com/driftappdev/libpackage/gologger v0.0.0
	github.com/driftappdev/libpackage/gometrics v0.0.0
	github.com/driftappdev/libpackage/goratelimit v0.0.0
	github.com/driftappdev/libpackage/goretry v0.0.0
	github.com/driftappdev/libpackage/gosanitizer v0.0.0
	github.com/driftappdev/libpackage/gotimeout v0.0.0
	github.com/driftappdev/libpackage/gotracing v0.0.0
	github.com/driftappdev/libpackage/logmid/logging-middleware v0.0.0
	github.com/driftappdev/libpackage/resilience/cache v0.0.0
	github.com/driftappdev/libpackage/resilience/pagination v0.0.0
	github.com/driftappdev/libpackage/resilience/validate v0.0.0
	github.com/driftappdev/libpackage/resilience/validator v0.0.0
	github.com/gin-gonic/gin v1.12.0
	github.com/lib/pq v1.10.9
	github.com/nats-io/nats.go v1.49.0
	google.golang.org/grpc v1.79.2
	google.golang.org/protobuf v1.36.11
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.15.0 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.1 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.2 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nkeys v0.4.12 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	go.mongodb.org/mongo-driver/v2 v2.5.0 // indirect
	golang.org/x/arch v0.22.0 // indirect
	golang.org/x/crypto v0.48.0 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)

replace github.com/driftappdev/libpackage/filemods/middleware/adminshield/admin-middleware => ../../libpackage/middleware/adminshield/admin-middleware

replace github.com/driftappdev/libpackage/goauth => ../../libpackage/goauth

replace github.com/driftappdev/libpackage/resilience/cache => ../../libpackage/resilience/cache

replace github.com/driftappdev/libpackage/gocircuit => ../../libpackage/gocircuit

replace github.com/driftappdev/libpackage/goerror => ../../libpackage/goerror

replace github.com/driftappdev/libpackage/gologger => ../../libpackage/gologger

replace github.com/driftappdev/libpackage/logmid/logging-middleware => ../../libpackage/logging-middleware

replace github.com/driftappdev/libpackage/gometrics => ../../libpackage/gometrics

replace github.com/driftappdev/libpackage/resilience/pagination => ../../libpackage/resilience/pagination

replace github.com/driftappdev/libpackage/goratelimit => ../../libpackage/goratelimit

replace github.com/driftappdev/libpackage/goretry => ../../libpackage/goretry

replace github.com/driftappdev/libpackage/gosanitizer => ../../libpackage/gosanitizer

replace github.com/driftappdev/libpackage/gotimeout => ../../libpackage/gotimeout

replace github.com/driftappdev/libpackage/gotracing => ../../libpackage/gotracing

replace github.com/driftappdev/libpackage/resilience/validate => ../../libpackage/resilience/validate

replace github.com/driftappdev/libpackage/resilience/validator => ../../libpackage/resilience/validator
