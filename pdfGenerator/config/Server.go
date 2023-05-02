package config

import (
	"context"
	"github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"
	"pdfGenerator/internal/service"
)

type Server struct {
	PDFService service.PDFService
	gen_proto.PDFGenerationServiceServer
}

func NewServer() *Server {
	return &Server{PDFService: *service.NewPDFService()}
}

func (s *Server) QueryForPDF(ctx context.Context, request *gen_proto.PDFRequest) (*gen_proto.PDFResponse, error) {
	s.PDFService.GeneratePDF(request.ChatID)
	return nil, nil
}
