package main

import (
	"bridgeServer/config"
	"bridgeServer/internal/service"
	"fmt"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't have opened env file: %v", err)
	}

	cfg := config.NewConfig(dbConfiguration(), GRPCConfiguration())
	gptService := service.NewGPTService(cfg.Client)
	for i := 0; i < 10; i++ {
		gptService.GetCongratulation("egor")
	}
}

func GRPCConfiguration() gen_proto.CongratulationServiceClient {
	port := os.Getenv("GRPC_PORT")
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")
	conn, err := grpc.Dial(name+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := gen_proto.NewCongratulationServiceClient(conn)

	return client
}

func dbConfiguration() *gorm.DB {
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_CONTAINER_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
