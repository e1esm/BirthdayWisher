package bridge

type Chat struct {
	ChatId int64
}

func NewChat(chatId int64) *Chat {
	return &Chat{ChatId: chatId}
}
