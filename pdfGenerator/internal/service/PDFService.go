package service

import (
	"pdfGenerator/internal/repository"
	"pdfGenerator/internal/utils"
)

type PDFService struct {
	repository *repository.UserRepository
}

func NewPDFService() *PDFService {
	return &PDFService{repository: repository.NewUserRepository()}
}

func (s *PDFService) GeneratePDF(chatID int64) []byte {
	users := s.repository.FetchUsersFromChat(chatID)
	pdf := utils.NewPDF(users, chatID)
	return pdf.GetBytesPdf()
}
