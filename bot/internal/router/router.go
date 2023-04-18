package router

import (
	"BirthdayWisherBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Switcher interface {
	add(message *tgbotapi.Update)
	addFull(message *tgbotapi.Update)
	change(message *tgbotapi.Update)
	list(message *tgbotapi.Update)
	soon(message *tgbotapi.Update)
}

type BirthdayRouter struct {
	bot           *tgbotapi.BotAPI
	bridgeService service.BridgeConnectorService
}

func NewBirthdayRouter(bot *tgbotapi.BotAPI, connectorService service.BridgeConnectorService) *BirthdayRouter {
	return &BirthdayRouter{bot: bot, bridgeService: connectorService}
}
