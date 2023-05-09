package service

import pdf_client "github.com/e1esm/protobuf/bridge_to_PDF-Generator/gen_proto"

type PDFService struct {
	client pdf_client.PDFGenerationServiceClient
}

func NewPDFService(client pdf_client.PDFGenerationServiceClient) *PDFService {
	return &PDFService{client: client}
}
