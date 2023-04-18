package bridge

type Chat struct {
	chatId int
}

func NewChat(chatId int) *Chat {
	return &Chat{chatId: chatId}
}
