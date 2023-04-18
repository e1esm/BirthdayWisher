package main

import (
	route "BirthdayWisherBot/internal/router"
	"BirthdayWisherBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6161612970:AAE4wWcup6gjYwChdmzesPEOhw195X3y98M")
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
