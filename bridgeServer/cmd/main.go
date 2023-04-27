package main

import (
	"bridgeServer/config"
	"bridgeServer/internal/model"
	"bridgeServer/internal/repository"
	"bridgeServer/internal/service"
	"fmt"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't have opened env file: %v", err)
	}

	cfg := config.NewConfig(dbConfiguration(), GRPCClientConfiguration())
	defer cfg.DB.DB()
	userRepository := repository.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepository)
	gptService := service.NewGPTService(cfg.Client)
	serverImpl := config.NewServer(userService, gptService, cfg)

	port := os.Getenv("GRPC_PORT")
	address := os.Getenv("BRIDGE_SERVER_CONTAINER_NAME")
	server, err := net.Listen("tcp", address+port)
	if err != nil {
		log.Fatalf("%s", err)
	}
	grpcServer := grpc.NewServer()
	bot_to_server_proto.RegisterCongratulationServiceServer(grpcServer, serverImpl)
	if err = grpcServer.Serve(server); err != nil {
		log.Fatalf("%s", err)
	}

}

func GRPCClientConfiguration() gen_proto.CongratulationServiceClient {
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			localization, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				log.Fatal(err)
			}
			return time.Now().In(localization)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	//db.Exec("SET TIME ZONE 'Europe/Moscow'")
	err = db.AutoMigrate(model.User{}, model.Chat{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
