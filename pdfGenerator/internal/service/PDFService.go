package service

import "pdfGenerator/internal/repository"

type PDFService struct {
	repository *repository.UserRepository
}

func NewPDFService() *PDFService {
	return &PDFService{repository: repository.NewUserRepository()}
}

func (s *PDFService) GeneratePDF(chatID int64) []byte {
	return []byte{}
}
