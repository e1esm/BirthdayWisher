package main

import (
	"congratulationsGenerator/internal/service"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
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
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")
	client := openai.NewClient(token)
	server, err := net.Listen("tcp", name+port)
	defer server.Close()
	if err != nil {
		log.Fatalf("Couldn't have started the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	aiService := service.NewOpenAIService(client)
	gen_proto.RegisterCongratulationServiceServer(grpcServer, aiService)

	log.Printf("Server started at: %v", server.Addr())
	if err := grpcServer.Serve(server); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
