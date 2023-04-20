module bridgeServer

go 1.20

require (
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.54.0
	github.com/e1esm/protobuf/bridge_to_API/gen_proto v1.0.0
)

replace github.com/e1esm/protobuf/bridge_to_API/gen_proto => ./../protobuf/bridge_to_API/gen_proto

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
