module enhanced_riskctrl

go 1.15

require (
	github.com/ethereum/go-ethereum v1.13.14
	github.com/gogf/gf/contrib/drivers/mysql/v2 v2.7.0 // indirect
	github.com/gogf/gf/contrib/drivers/pgsql/v2 v2.7.0
	github.com/gogf/gf/contrib/nosql/redis/v2 v2.7.0
	github.com/gogf/gf/contrib/trace/jaeger/v2 v2.7.0
	github.com/gogf/gf/v2 v2.7.0
	github.com/golang/protobuf v1.5.3
	github.com/mpcsdk/mpcCommon v0.0.0
	github.com/nats-io/nats.go v1.33.1
	github.com/nats-rpc/nrpc v0.0.0-20231018091755-18e69326f052
	go.opentelemetry.io/otel/trace v1.14.0
	google.golang.org/protobuf v1.31.0
)

replace github.com/mpcsdk/mpcCommon v0.0.0 => ./mpcCommon
