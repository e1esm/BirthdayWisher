package repository

type Repositories struct {
	UserRepository *UserRepository
	ChatRepository *ChatRepository
}

func NewRepositories(repository *UserRepository, chatRepository *ChatRepository) *Repositories {
	return &Repositories{UserRepository: repository, ChatRepository: chatRepository}
}
