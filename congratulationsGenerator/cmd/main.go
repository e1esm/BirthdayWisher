package main

import (
	"congratulationsGenerator/internal/service"
	"congratulationsGenerator/utils"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error while reading from env file")
	}

	token := os.Getenv("AI_TOKEN")
	port := os.Getenv("GRPC_PORT")
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")
	client := openai.NewClient(token)
	server, err := net.Listen("tcp", name+port)
	defer server.Close()
	if err != nil {
		utils.Logger.Fatal("Address can't be listened", zap.String("error", err.Error()))
	}

	grpcServer := grpc.NewServer()
	aiService := service.NewOpenAIService(client)
	gen_proto.RegisterCongratulationServiceServer(grpcServer, aiService)

	if err := grpcServer.Serve(server); err != nil {
		utils.Logger.Fatal("Server can't be started", zap.String("error", err.Error()))
	}

}
