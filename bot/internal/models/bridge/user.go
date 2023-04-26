package bridge

type User struct {
	UserId      int64
	Username    string
	Date        string
	CurrentChat Chat
}

func NewUser(id int64, date string, chat Chat, username string) *User {
	return &User{UserId: id, Date: date, CurrentChat: chat, Username: username}
}
