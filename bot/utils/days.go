package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type DaysInMonth struct {
	month int
	days  []int
}

func newDays(month int, year int) DaysInMonth {

	lastDay, err := strconv.Atoi(endOfMonth(month, year).Format(time.DateOnly)[8:len(time.DateOnly)])
	if err != nil {
		Logger.Error("Couldn't have extracted last day of the current month", zap.String(fmt.Sprintf("%d", month), endOfMonth(month, year).Format(time.DateOnly)))
	}
	var currentDaysInMonth DaysInMonth
	currentDaysInMonth.month = month
	currentDaysInMonth.days = make([]int, lastDay)
	for i := 0; i < lastDay; i++ {
		currentDaysInMonth.days[i] = i + 1
	}

	return currentDaysInMonth
}

func beginningOfMonth(month, year int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
}

func endOfMonth(month, year int) time.Time {
	return beginningOfMonth(month, year).AddDate(0, 1, 0).Add(-time.Second)
}

func GenerateDaysButtons(month int, year int) [][]tgbotapi.InlineKeyboardButton {
	monthInfo := newDays(month, year)
	initialLen := 5
	dayButtons := make([][]tgbotapi.InlineKeyboardButton, initialLen)

	for i := 0; i < 5; i++ {
		dayButtons[i] = make([]tgbotapi.InlineKeyboardButton, 0, 7)
		for j := 0; j < 7; j++ {
			index := i*7 + j
			if index >= len(monthInfo.days) {
				break
			}
			date := fmt.Sprintf("%d", monthInfo.days[index])
			dayButtons[i] = append(dayButtons[i], tgbotapi.NewInlineKeyboardButtonData(date, date))
		}
	}

	return dayButtons

}
