package bridge

type User struct {
	userId      int
	date        string
	currentChat Chat
}

func NewUser(id int, date string, chat Chat) *User {
	return &User{userId: id, date: date, currentChat: chat}
}
