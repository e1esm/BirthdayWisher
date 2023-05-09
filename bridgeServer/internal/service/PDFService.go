package service

import (
	"context"
	pdf_client "github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"
)

type PDFService struct {
	client pdf_client.PDFGenerationServiceClient
}

func NewPDFService(client pdf_client.PDFGenerationServiceClient) *PDFService {
	return &PDFService{client: client}
}

func (s *PDFService) QueryForPDF(ctx context.Context, request *pdf_client.PDFRequest) (*pdf_client.PDFResponse, error) {
	return s.client.QueryForPDF(ctx, request)
}
