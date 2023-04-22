package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	userId      int64
	date        string
	currentChat Chat
}
