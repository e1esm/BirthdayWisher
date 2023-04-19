package main

import (
	route "BirthdayWisherBot/internal/router"
	"BirthdayWisherBot/internal/service"
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
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
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	router := route.NewBirthdayRouter(bot, service.BridgeConnectorService{})

	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Seconds().Do(func() {
		log.Println("Working fine")
	})
	s.StartAsync()
	for update := range updates {

		router.HandleUpdate(update)
	}
}
