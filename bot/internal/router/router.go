package router

import (
	"BirthdayWisherBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Switcher interface {
	add(message tgbotapi.Message)
	addFull(message tgbotapi.Message)
	change(message tgbotapi.Message)
	list(message tgbotapi.Message)
	soon(message tgbotapi.Message)
}

type BirthdayRouter struct {
	bot           *tgbotapi.BotAPI
	bridgeService service.BridgeConnectorService
}

func NewBirthdayRouter(bot *tgbotapi.BotAPI, connectorService service.BridgeConnectorService) *BirthdayRouter {
	return &BirthdayRouter{bot: bot, bridgeService: connectorService}
}
