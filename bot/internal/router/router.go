package router

import (
	"BirthdayWisherBot/internal/service"
	"github.com/go-co-op/gocron"
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
	bot              *tgbotapi.BotAPI
	bridgeService    service.BridgeConnectorService
	ConnectorService service.BridgeConnectorService
	scheduler        *gocron.Scheduler
}

func NewBirthdayRouter(bot *tgbotapi.BotAPI, connectorService service.BridgeConnectorService, scheduler *gocron.Scheduler) *BirthdayRouter {
	return &BirthdayRouter{bot: bot, bridgeService: connectorService, scheduler: scheduler}
}
