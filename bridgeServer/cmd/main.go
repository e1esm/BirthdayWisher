package main

import (
	"bridgeServer/internal/service"
	"github.com/e1esm/protobuf/bridge_to_API/gen_proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Couldn't have opened env file: %v", err)
	}
	port := os.Getenv("GRPC_PORT")
	name := os.Getenv("AI_CALLER_CONTAINER_NAME")

	conn, err := grpc.Dial(name+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := gen_proto.NewCongratulationServiceClient(conn)
	gptService := service.NewGPTService(client)
	for i := 0; i < 10; i++ {
		gptService.GetCongratulation("egor")
	}
}
