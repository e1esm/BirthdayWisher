package repository

import (
	"bridgeServer/internal/model"
	"bridgeServer/utils"
	"errors"
	bot_to_server_proto "github.com/e1esm/protobuf/bot_to_server/gen_proto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) DeleteUser(userID, chatID int64) error {
	var amountOfRows int
	r.db.Raw("SELECT COUNT(*) FROM chats where user_id = ?", userID).Find(&amountOfRows)
	var deletionErr error
	if amountOfRows > 1 {
		r.db.Select("CurrentChat").Where("chat_id = ? && user_id = ?", chatID, userID).Delete(&model.User{})
	} else if amountOfRows == 1 {
		r.db.Unscoped().Select("CurrentChat").Delete(&model.User{ID: userID})
	} else {
		deletionErr = errors.New("there's no such user in the database")
	}
	return deletionErr
}

func (r *UserRepository) SaveUser(user *model.User) {
	var retrievedUser model.User
	err := r.db.First(&retrievedUser, user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Logger.Info("Created user", zap.String("user", user.Username))
		r.db.Debug().Create(user)
	} else {
		utils.Logger.Info("Updated user", zap.String("user", user.Username))
		//r.db.Model(user).Omit("CurrentChat").Update("date", user.Date)
		r.db.Model(user).Select("CurrentChat")
	}
}

func (r *UserRepository) FindUsers() []model.User {
	users := make([]model.User, 0, 10)
	r.db.Preload("CurrentChat").Where("extract(month from date) = extract(month from current_date) and extract(day from date) = extract(day from current_date)").Find(&users)
	return users
}

func (r *UserRepository) SoonBirthdaysOfUsers(chatId int64) *bot_to_server_proto.ChatBirthdaysResponse {
	users := make([]model.User, 0, 10)
	r.db.Preload("CurrentChat", "chat_id = ?", chatId).Where("extract(month from date) = extract(month from current_date) and extract(day from date) > extract(day from current_date)").Find(&users)
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
