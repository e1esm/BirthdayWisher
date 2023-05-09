package config

import (
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	pdf_client "github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"
	"gorm.io/gorm"
)

type Config struct {
	DB        *gorm.DB
	Client    gen_proto.CongratulationServiceClient
	PDFClient pdf_client.PDFGenerationServiceClient
}

func NewConfig(db *gorm.DB, client gen_proto.CongratulationServiceClient, pdfClient pdf_client.PDFGenerationServiceClient) *Config {
	return &Config{DB: db, Client: client, PDFClient: pdfClient}
}
