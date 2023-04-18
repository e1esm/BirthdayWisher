package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"reflect"
	"strings"
)

func (r *BirthdayRouter) help(message tgbotapi.Message) {
	switcherType := reflect.TypeOf((*Switcher)(nil)).Elem()
	methods := "hey + \n"
	strBuilder := strings.Builder{}
	for i := 0; i < switcherType.NumMethod(); i++ {
		strBuilder.WriteString(switcherType.Method(i).Name)
		strBuilder.WriteString("\n")
	}
	methods = strBuilder.String()
	log.Println(methods)
	msg := tgbotapi.NewMessage(message.Chat.ID, methods)
	r.bot.Send(msg)
}
