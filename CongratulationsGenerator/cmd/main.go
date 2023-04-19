package main

import (
	"CongratulationsGenerator/internal/service"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error while reading from env file")
	}
	token := os.Getenv("AI_TOKEN")
	client := openai.NewClient(token)
	service.NewOpenAIService(client).Query("hey")

}
