package router

import (
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"time"
)

type State int

const (
	START State = iota
	YEAR
	MONTH
	DAY
)

const ARR_LENGTH = 12
const OFFSET_STEP = 12

var years []int

func init() {
	years = make([]int, 0, 12)
	y := time.Now().Year()
	for i := 0; i < ARR_LENGTH; i++ {
		years = append(years, y-i)
	}
	RWapInstance = RWMap{UserStateConfigs: make(map[int64]StateConfig), Mutex: &sync.RWMutex{}}
}

var RWapInstance RWMap

type RWMap struct {
	UserStateConfigs map[int64]StateConfig
	Mutex            *sync.RWMutex
}

type StateConfig struct {
	Date         string
	CurrentState State
	Offset       int
}

func (r *BirthdayRouter) addDate(message tgbotapi.Message) {
	currentStateConfig := StateConfig{
		CurrentState: START,
		Offset:       0,
	}
	RWapInstance.Mutex.Lock()
	RWapInstance.UserStateConfigs[message.From.ID] = currentStateConfig
	RWapInstance.Mutex.Unlock()
	r.addYear(message)
}

func (r *BirthdayRouter) addYear(message tgbotapi.Message) {

	RWapInstance.Mutex.RLock()
	v, _ := RWapInstance.UserStateConfigs[message.From.ID]
	RWapInstance.Mutex.RUnlock()
	v.CurrentState = YEAR
	RWapInstance.Mutex.Lock()
	RWapInstance.UserStateConfigs[message.From.ID] = v
	RWapInstance.Mutex.Unlock()

	arrRows := make([][]tgbotapi.InlineKeyboardButton, 4)
	currentOffset := 0

	for i := 0; i < 4; i++ {
		arrRows[i] = make([]tgbotapi.InlineKeyboardButton, 3)
		for j := 0; j < 3; j++ {
			arrRows[i][j] = tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d", years[currentOffset]), fmt.Sprintf("%d", years[currentOffset]))
			currentOffset++
		}
	}
	arrRows[len(arrRows)-1] = []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(emoji.LeftArrow.String(), fmt.Sprintf("-12")),
		tgbotapi.NewInlineKeyboardButtonData(emoji.CrossMark.String(), fmt.Sprintf("0")),
		tgbotapi.NewInlineKeyboardButtonData(emoji.RightArrow.String(), fmt.Sprintf("-12")),
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите год")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		arrRows[0],
		arrRows[1],
		arrRows[2],
		arrRows[3],
	)

	RWapInstance.Mutex.RLock()
	v, _ = RWapInstance.UserStateConfigs[message.From.ID]
	RWapInstance.Mutex.RUnlock()
	v.Offset = currentOffset
	RWapInstance.Mutex.Lock()
	RWapInstance.UserStateConfigs[message.From.ID] = v
	RWapInstance.Mutex.Unlock()

	r.bot.Send(msg)

}

//RIGHT ARROW - -offset
//LEFT ARROW - +offset
//MIDDLECROSS - 0
