package bridge

type User struct {
	UserId      int64
	Date        string
	CurrentChat Chat
}

func NewUser(id int64, date string, chat Chat) *User {
	return &User{UserId: id, Date: date, CurrentChat: chat}
}
