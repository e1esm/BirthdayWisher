package main

import (
	route "BirthdayWisherBot/internal/router"
	"BirthdayWisherBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panic("Couldn't have read env file")
	}
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	router := route.NewBirthdayRouter(bot, service.BridgeConnectorService{})

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {

			}
			/*
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			*/
		}
	}
}
