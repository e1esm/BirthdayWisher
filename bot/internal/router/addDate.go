package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type State int

const (
	YEAR State = iota
	MONTH
	DAY
)

func init() {

}

type StateConfig struct {
	Date         string
	CurrentState State
}

func (r *BirthdayRouter) addDate(message tgbotapi.Message) {
	r.addYear(message)
}

func (r *BirthdayRouter) addYear(message tgbotapi.Message) {

}
