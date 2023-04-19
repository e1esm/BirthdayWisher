package main

import (
	"CongratulationsGenerator/internal/service"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error while reading from env file")
	}
	token := os.Getenv("AI_TOKEN")
	port := os.Getenv("GRPC_PORT")
	client := openai.NewClient(token)

	server, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Couldn't have started the server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCongratulationServiceServer()
	service.NewOpenAIService(client).Query("hey")

}
