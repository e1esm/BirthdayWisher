package config

import (
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	Client gen_proto.CongratulationServiceClient
	bot_to_server_proto.CongratulationServiceServer
}

func NewConfig(db *gorm.DB, client gen_proto.CongratulationServiceClient) *Config {
	return &Config{DB: db, Client: client}
}
