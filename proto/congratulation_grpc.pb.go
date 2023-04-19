// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/congratulation.proto

package proto

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
	CongratulationService_QueryForCongratulation_FullMethodName = "/congratulation_service.CongratulationService/QueryForCongratulation"
)

// CongratulationServiceClient is the client API for CongratulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CongratulationServiceClient interface {
	QueryForCongratulation(ctx context.Context, in *CongratulationRequest, opts ...grpc.CallOption) (*CongratulationResponse, error)
}

type congratulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCongratulationServiceClient(cc grpc.ClientConnInterface) CongratulationServiceClient {
	return &congratulationServiceClient{cc}
}

func (c *congratulationServiceClient) QueryForCongratulation(ctx context.Context, in *CongratulationRequest, opts ...grpc.CallOption) (*CongratulationResponse, error) {
	out := new(CongratulationResponse)
	err := c.cc.Invoke(ctx, CongratulationService_QueryForCongratulation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CongratulationServiceServer is the server API for CongratulationService service.
// All implementations must embed UnimplementedCongratulationServiceServer
// for forward compatibility
type CongratulationServiceServer interface {
	QueryForCongratulation(context.Context, *CongratulationRequest) (*CongratulationResponse, error)
	mustEmbedUnimplementedCongratulationServiceServer()
}

// UnimplementedCongratulationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCongratulationServiceServer struct {
}

func (UnimplementedCongratulationServiceServer) QueryForCongratulation(context.Context, *CongratulationRequest) (*CongratulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryForCongratulation not implemented")
}
func (UnimplementedCongratulationServiceServer) mustEmbedUnimplementedCongratulationServiceServer() {}

// UnsafeCongratulationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CongratulationServiceServer will
// result in compilation errors.
type UnsafeCongratulationServiceServer interface {
	mustEmbedUnimplementedCongratulationServiceServer()
}

func RegisterCongratulationServiceServer(s grpc.ServiceRegistrar, srv CongratulationServiceServer) {
	s.RegisterService(&CongratulationService_ServiceDesc, srv)
}

func _CongratulationService_QueryForCongratulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CongratulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CongratulationServiceServer).QueryForCongratulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CongratulationService_QueryForCongratulation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CongratulationServiceServer).QueryForCongratulation(ctx, req.(*CongratulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CongratulationService_ServiceDesc is the grpc.ServiceDesc for CongratulationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CongratulationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "congratulation_service.CongratulationService",
	HandlerType: (*CongratulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryForCongratulation",
			Handler:    _CongratulationService_QueryForCongratulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/congratulation.proto",
}
