module bridgeServer

go 1.20

require (
	github.com/e1esm/protobuf/bot_to_server/gen_proto v1.0.0
	github.com/e1esm/protobuf/bridge_to_API/gen_proto v1.0.0
	github.com/joho/godotenv v1.5.1
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
	gorm.io/driver/postgres v1.5.0
	gorm.io/gorm v1.25.0
)

replace github.com/e1esm/protobuf/bot_to_server/gen_proto => ./../protobuf/bot_to_server/gen_proto

replace github.com/e1esm/protobuf/bridge_to_API/gen_proto => ./../protobuf/bridge_to_API/gen_proto

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
