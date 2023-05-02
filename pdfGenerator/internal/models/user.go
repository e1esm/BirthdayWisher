package models

type User struct {
	Username string
	Date     string
}

func NewUser(Username string, Date string) *User {
	return &User{Username: Username, Date: Date}
}
