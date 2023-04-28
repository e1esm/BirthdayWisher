package main

import (
	route "BirthdayWisherBot/internal/router"
	"BirthdayWisherBot/internal/service"
	"github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panic("Couldn't have read env file")
	}
	token := os.Getenv("BOT_TOKEN")
	address := os.Getenv("BRIDGE_SERVER_CONTAINER_NAME")
	port := os.Getenv("GRPC_PORT")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	log.Println(address + port)
	conn, err := grpc.Dial(address+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer conn.Close()

	client := gen_proto.NewCongratulationServiceClient(conn)
	location, err := time.LoadLocation("Europe/Moscow")
	var scheduler *gocron.Scheduler
	if err != nil {
		log.Println("UTC Time")
		scheduler = gocron.NewScheduler(time.UTC)
	} else {
		log.Println("Moscow time")
		scheduler = gocron.NewScheduler(location)
	}
	router := route.NewBirthdayRouter(bot, *service.NewBridgeConnectorService(client), scheduler)
	scheduler.Every(1).Day().At("00:00").Do(router.DailyBirthdayChecker)
	router.Scheduler.StartAsync()
	for update := range updates {
		router.HandleUpdate(update)
	}
}
