module congratulationsGenerator

go 1.20

require (
	github.com/e1esm/protobuf/bridge_to_API/gen_proto v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/joho/godotenv v1.5.1
	github.com/sashabaranov/go-openai v1.8.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.54.0
)

replace github.com/e1esm/protobuf/bridge_to_API/gen_proto => ./../protobuf/bridge_to_API/gen_proto

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
