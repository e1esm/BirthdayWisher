package bridge

type User struct {
	userId      int64
	date        string
	currentChat Chat
}

func NewUser(id int64, date string, chat Chat) *User {
	return &User{userId: id, date: date, currentChat: chat}
}
