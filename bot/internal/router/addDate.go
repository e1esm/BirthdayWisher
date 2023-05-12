package router

import (
	"BirthdayWisherBot/utils"
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"log"
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

func init() {
	RWapInstance = RWMap{UserStateConfigs: make(map[int64]StateConfig), Mutex: &sync.RWMutex{}}
}

var RWapInstance RWMap

type RWMap struct {
	UserStateConfigs map[int64]StateConfig
	Mutex            *sync.RWMutex
}

type StateConfig struct {
	Date         string
	ChatID       int64
	UserID       int64
	MessageID    int
	CurrentState State
	Offset       int
}

func (r *BirthdayRouter) addDate(update tgbotapi.Update) {
	currentStateConfig := StateConfig{
		CurrentState: START,
		UserID:       update.SentFrom().ID,
		ChatID:       update.FromChat().ID,
		MessageID:    update.Message.MessageID,
		Offset:       0,
	}
	RWapInstance.Mutex.Lock()
	RWapInstance.UserStateConfigs[currentStateConfig.UserID] = currentStateConfig
	RWapInstance.Mutex.Unlock()
	log.Println(currentStateConfig)
	r.addYear(update, nil)
}

func (r *BirthdayRouter) addYear(update tgbotapi.Update, messageID *int) {

	RWapInstance.Mutex.RLock()
	v, _ := RWapInstance.UserStateConfigs[update.SentFrom().ID]
	RWapInstance.Mutex.RUnlock()
	v.CurrentState = YEAR
	RWapInstance.Mutex.Lock()
	RWapInstance.UserStateConfigs[v.UserID] = v
	RWapInstance.Mutex.Unlock()

	arrRows := make([][]tgbotapi.InlineKeyboardButton, 5)
	currentOffset := v.Offset

	currentYear := time.Now().Year()
	for i := 0; i < 4; i++ {
		arrRows[i] = make([]tgbotapi.InlineKeyboardButton, 3)
		for j := 0; j < 3; j++ {
			arrRows[i][j] = tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d", currentYear-currentOffset), fmt.Sprintf("%d", currentYear-currentOffset))
			currentOffset++
		}
	}
	arrRows[len(arrRows)-1] = []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(emoji.LeftArrow.String(), "-12"),
		tgbotapi.NewInlineKeyboardButtonData(emoji.CrossMark.String(), fmt.Sprintf("0")),
		tgbotapi.NewInlineKeyboardButtonData(emoji.RightArrow.String(), fmt.Sprintf("+12")),
	}
	wasMessageIDFound := false
	RWapInstance.Mutex.RLock()
	v, _ = RWapInstance.UserStateConfigs[v.UserID]
	RWapInstance.Mutex.RUnlock()
	RWapInstance.Mutex.Lock()
	if messageID != nil {
		v.MessageID = *messageID
		wasMessageIDFound = true
	}
	RWapInstance.UserStateConfigs[v.UserID] = v
	RWapInstance.Mutex.Unlock()

	markup := tgbotapi.NewInlineKeyboardMarkup(arrRows...)

	if wasMessageIDFound {
		msg := tgbotapi.NewEditMessageReplyMarkup(v.ChatID, *messageID, markup)
		r.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(v.ChatID, "Выберите год")
		msg.ReplyMarkup = markup
		if _, err := r.bot.Send(msg); err != nil {
			utils.Logger.Error(err.Error(), zap.Int64("chatID", v.ChatID))
		}
	}

}
