package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type InlineButtonMonths struct {
	monthsInlineButtons [][]tgbotapi.InlineKeyboardButton
}

func GetMonthButtons() [][]tgbotapi.InlineKeyboardButton {
	var buttons InlineButtonMonths
	buttons.monthsInlineButtons = make([][]tgbotapi.InlineKeyboardButton, 4)
	currentMonth := 1
	for i := 0; i < 4; i++ {
		buttons.monthsInlineButtons[i] = make([]tgbotapi.InlineKeyboardButton, 3)
		for j := 0; j < 3; j++ {
			buttons.monthsInlineButtons[i][j] = tgbotapi.NewInlineKeyboardButtonData(time.Month(currentMonth).String(), fmt.Sprintf("%d", currentMonth))
			currentMonth++
		}
	}

	return buttons.monthsInlineButtons
}
