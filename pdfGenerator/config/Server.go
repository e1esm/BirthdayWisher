package config

import (
	"context"
	"github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"
	"pdfGenerator/internal/service"
	"pdfGenerator/internal/utils"
	"time"
)

type Server struct {
	PDFService service.PDFService
	gen_proto.PDFGenerationServiceServer
}

func NewServer() *Server {
	return &Server{PDFService: *service.NewPDFService()}
}

func (s *Server) QueryForPDF(ctx context.Context, request *gen_proto.PDFRequest) (*gen_proto.PDFResponse, error) {
	start := time.Now()
	bytes := s.PDFService.GeneratePDF(request.ChatID)
	elapsed := time.Since(start).Seconds()
	utils.GrpcRequestDuration.Observe(elapsed)
	return &gen_proto.PDFResponse{Data: bytes}, nil
}
