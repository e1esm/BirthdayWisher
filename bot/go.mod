module BirthdayWisherBot

go 1.20

require (
	github.com/e1esm/protobuf/bot_to_server/gen_proto v1.0.0
	github.com/go-co-op/gocron v1.22.3
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/joho/godotenv v1.5.1
	github.com/prometheus/client_golang v1.15.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0

)

replace github.com/e1esm/protobuf/bot_to_server/gen_proto => ./../protobuf/bot_to_server/gen_proto

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
