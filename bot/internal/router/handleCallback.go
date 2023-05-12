package router

import (
	"BirthdayWisherBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
)

var YearIfOmit = "1000"

func (r *BirthdayRouter) handleCallback(update tgbotapi.Update) {
	callback := update.CallbackQuery
	RWapInstance.Mutex.RLock()
	stateCFG, ok := RWapInstance.UserStateConfigs[callback.From.ID]
	RWapInstance.Mutex.RUnlock()

	if len(callback.Data) == 3 && stateCFG.CurrentState == YEAR {
		intRepr, _ := strconv.Atoi(callback.Data)
		stateCFG.Offset -= intRepr
		RWapInstance.Mutex.Lock()
		RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		RWapInstance.Mutex.Unlock()
		r.addYear(update, &update.CallbackQuery.Message.MessageID)
		return
	}
	if len(callback.Data) == 1 && stateCFG.CurrentState == YEAR {
		stateCFG.Date += YearIfOmit
		RWapInstance.Mutex.Lock()
		RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		RWapInstance.Mutex.Unlock()

		//addMonth...
	}

	if ok {
		switch stateCFG.CurrentState {
		case YEAR:
			stateCFG.Date += callback.Data
			RWapInstance.Mutex.Lock()
			RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			RWapInstance.Mutex.Unlock()
			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, stateCFG.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				utils.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}

		}
	}
}
