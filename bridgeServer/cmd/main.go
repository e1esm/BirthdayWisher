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
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't have opened env file: %v", err)
	}

	cfg := config.NewConfig(dbConfiguration(), GRPCClientConfiguration())
	repositories := repository.NewRepositories(repository.NewUserRepository(cfg.DB), repository.NewChatRepository(cfg.DB))
	chatService := service.NewChatService(repositories)
	userService := service.NewUserService(repositories)
	serverImpl := config.NewServer(userService, chatService)

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
	/*
		gptService := service.NewGPTService(cfg.Client)
		for i := 0; i < 10; i++ {
			gptService.GetCongratulation("egor")
		}

	*/
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(model.User{}, model.Chat{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
