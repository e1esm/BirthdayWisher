package bridge

type Chat struct {
	chatId int64
}

func NewChat(chatId int64) *Chat {
	return &Chat{chatId: chatId}
}
