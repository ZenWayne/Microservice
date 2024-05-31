// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/NFT/NFT.proto

package NFT

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

// NFTServiceClient is the client API for NFTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NFTServiceClient interface {
	AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
}

type nFTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNFTServiceClient(cc grpc.ClientConnInterface) NFTServiceClient {
	return &nFTServiceClient{cc}
}

func (c *nFTServiceClient) AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error) {
	out := new(AddCollectionResponse)
	err := c.cc.Invoke(ctx, "/NFT.NFTService/AddCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/NFT.NFTService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NFTServiceServer is the server API for NFTService service.
// All implementations must embed UnimplementedNFTServiceServer
// for forward compatibility
type NFTServiceServer interface {
	AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	mustEmbedUnimplementedNFTServiceServer()
}

// UnimplementedNFTServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNFTServiceServer struct {
}

func (UnimplementedNFTServiceServer) AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollection not implemented")
}
func (UnimplementedNFTServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedNFTServiceServer) mustEmbedUnimplementedNFTServiceServer() {}

// UnsafeNFTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NFTServiceServer will
// result in compilation errors.
type UnsafeNFTServiceServer interface {
	mustEmbedUnimplementedNFTServiceServer()
}

func RegisterNFTServiceServer(s grpc.ServiceRegistrar, srv NFTServiceServer) {
	s.RegisterService(&NFTService_ServiceDesc, srv)
}

func _NFTService_AddCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).AddCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NFT.NFTService/AddCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).AddCollection(ctx, req.(*AddCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NFT.NFTService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NFTService_ServiceDesc is the grpc.ServiceDesc for NFTService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NFTService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NFT.NFTService",
	HandlerType: (*NFTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCollection",
			Handler:    _NFTService_AddCollection_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _NFTService_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/NFT/NFT.proto",
}
