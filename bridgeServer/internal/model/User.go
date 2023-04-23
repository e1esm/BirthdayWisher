package model

import (
	"google.golang.org/genproto/googleapis/type/date"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	userId      int64
	date        date.Date
	currentChat []Chat
}
