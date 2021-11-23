// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mall

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

// ShowClient is the client API for Show service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShowClient interface {
	// Sends a greeting
	GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...grpc.CallOption) (*BannerByIdReply, error)
}

type showClient struct {
	cc grpc.ClientConnInterface
}

func NewShowClient(cc grpc.ClientConnInterface) ShowClient {
	return &showClient{cc}
}

func (c *showClient) GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...grpc.CallOption) (*BannerByIdReply, error) {
	out := new(BannerByIdReply)
	err := c.cc.Invoke(ctx, "/mall.Show/GetBannerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShowServer is the server API for Show service.
// All implementations must embed UnimplementedShowServer
// for forward compatibility
type ShowServer interface {
	// Sends a greeting
	GetBannerById(context.Context, *BannerByIdRequest) (*BannerByIdReply, error)
	mustEmbedUnimplementedShowServer()
}

// UnimplementedShowServer must be embedded to have forward compatible implementations.
type UnimplementedShowServer struct {
}

func (UnimplementedShowServer) GetBannerById(context.Context, *BannerByIdRequest) (*BannerByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBannerById not implemented")
}
func (UnimplementedShowServer) mustEmbedUnimplementedShowServer() {}

// UnsafeShowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShowServer will
// result in compilation errors.
type UnsafeShowServer interface {
	mustEmbedUnimplementedShowServer()
}

func RegisterShowServer(s grpc.ServiceRegistrar, srv ShowServer) {
	s.RegisterService(&Show_ServiceDesc, srv)
}

func _Show_GetBannerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShowServer).GetBannerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.Show/GetBannerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShowServer).GetBannerById(ctx, req.(*BannerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Show_ServiceDesc is the grpc.ServiceDesc for Show service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Show_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mall.Show",
	HandlerType: (*ShowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBannerById",
			Handler:    _Show_GetBannerById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mall/app.proto",
}
