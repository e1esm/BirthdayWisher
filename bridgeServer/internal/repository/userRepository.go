package repository

import (
	"bridgeServer/internal/model"
	"errors"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveUser(user *model.User) {
	var retrievedUser model.User
	err := r.db.First(&retrievedUser, user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.db.Debug().Create(user)
	} else {
		r.db.Debug().Save(user)
	}
}

func (r *UserRepository) FindUsers() []model.User {
	users := make([]model.User, 0, 10)
	r.db.Preload("CurrentChat").Where("extract(month from date) = extract(month from current_date) and extract(day from date) = extract(day from current_date)").Find(&users)
	return users
}

func (r *UserRepository) SoonBirthdaysOfUsers(chatId int64) *bot_to_server_proto.ChatBirthdaysResponse {
	users := make([]model.User, 0, 10)
	r.db.Debug().Preload("CurrentChat", "chat_id = ?", chatId).Where("extract(month from date) = extract(month from current_date) and extract(day from date) > extract(day from current_date)").Find(&users)
	return chatBirthdaysResponseBuilder(users, chatId)

}

func chatBirthdaysResponseBuilder(users []model.User, id int64) *bot_to_server_proto.ChatBirthdaysResponse {
	var chatBirthdaysResponse bot_to_server_proto.ChatBirthdaysResponse
	chatBirthdaysResponse.ChatID = id
	birthdayResponseArr := make([]*bot_to_server_proto.ChatBirthdaysResponse_BirthdaysResponse, 0, 10)
	for i := 0; i < len(users); i++ {
		birthdayResponseArr = append(birthdayResponseArr,
			&bot_to_server_proto.ChatBirthdaysResponse_BirthdaysResponse{BirthdayDate: users[i].Date.String(),
				Username: users[i].Username})
	}
	chatBirthdaysResponse.SoonBirthdays = birthdayResponseArr
	return &chatBirthdaysResponse
}
