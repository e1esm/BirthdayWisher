package config

import (
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	Client gen_proto.CongratulationServiceClient
}

func NewConfig(db *gorm.DB, client gen_proto.CongratulationServiceClient) *Config {
	return &Config{DB: db, Client: client}
}
