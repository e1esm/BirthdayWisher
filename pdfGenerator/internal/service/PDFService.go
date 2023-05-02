package service

import (
	"log"
	"pdfGenerator/internal/repository"
)

type PDFService struct {
	repository *repository.UserRepository
}

func NewPDFService() *PDFService {
	return &PDFService{repository: repository.NewUserRepository()}
}

func (s *PDFService) GeneratePDF(chatID int64) []byte {
	users := s.repository.FetchUsersFromChat(chatID)
	log.Println(users)
	return []byte{}
}
