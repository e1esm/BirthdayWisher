// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protobuf/bridge_to_PDF-Generator/pdfGeneration.proto

package gen_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PDFGenerationService_QueryForPDF_FullMethodName = "/pdf_generation_service.PDFGenerationService/QueryForPDF"
)

// PDFGenerationServiceClient is the client API for PDFGenerationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PDFGenerationServiceClient interface {
	QueryForPDF(ctx context.Context, in *PDFRequest, opts ...grpc.CallOption) (*PDFResponse, error)
}

type pDFGenerationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPDFGenerationServiceClient(cc grpc.ClientConnInterface) PDFGenerationServiceClient {
	return &pDFGenerationServiceClient{cc}
}

func (c *pDFGenerationServiceClient) QueryForPDF(ctx context.Context, in *PDFRequest, opts ...grpc.CallOption) (*PDFResponse, error) {
	out := new(PDFResponse)
	err := c.cc.Invoke(ctx, PDFGenerationService_QueryForPDF_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PDFGenerationServiceServer is the server API for PDFGenerationService service.
// All implementations must embed UnimplementedPDFGenerationServiceServer
// for forward compatibility
type PDFGenerationServiceServer interface {
	QueryForPDF(context.Context, *PDFRequest) (*PDFResponse, error)
	mustEmbedUnimplementedPDFGenerationServiceServer()
}

// UnimplementedPDFGenerationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPDFGenerationServiceServer struct {
}

func (UnimplementedPDFGenerationServiceServer) QueryForPDF(context.Context, *PDFRequest) (*PDFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryForPDF not implemented")
}
func (UnimplementedPDFGenerationServiceServer) mustEmbedUnimplementedPDFGenerationServiceServer() {}

// UnsafePDFGenerationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PDFGenerationServiceServer will
// result in compilation errors.
type UnsafePDFGenerationServiceServer interface {
	mustEmbedUnimplementedPDFGenerationServiceServer()
}

func RegisterPDFGenerationServiceServer(s grpc.ServiceRegistrar, srv PDFGenerationServiceServer) {
	s.RegisterService(&PDFGenerationService_ServiceDesc, srv)
}

func _PDFGenerationService_QueryForPDF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PDFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDFGenerationServiceServer).QueryForPDF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PDFGenerationService_QueryForPDF_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDFGenerationServiceServer).QueryForPDF(ctx, req.(*PDFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PDFGenerationService_ServiceDesc is the grpc.ServiceDesc for PDFGenerationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PDFGenerationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdf_generation_service.PDFGenerationService",
	HandlerType: (*PDFGenerationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryForPDF",
			Handler:    _PDFGenerationService_QueryForPDF_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/bridge_to_PDF-Generator/pdfGeneration.proto",
}
