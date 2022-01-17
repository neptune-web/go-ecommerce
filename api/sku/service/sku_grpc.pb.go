// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// SkuClient is the client API for Sku service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkuClient interface {
	GetSkuById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SkuVO, error)
	GetSpuByTheme(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SpuByThemeReply, error)
}

type skuClient struct {
	cc grpc.ClientConnInterface
}

func NewSkuClient(cc grpc.ClientConnInterface) SkuClient {
	return &skuClient{cc}
}

func (c *skuClient) GetSkuById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SkuVO, error) {
	out := new(SkuVO)
	err := c.cc.Invoke(ctx, "/order.service.Sku/GetSkuById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skuClient) GetSpuByTheme(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SpuByThemeReply, error) {
	out := new(SpuByThemeReply)
	err := c.cc.Invoke(ctx, "/order.service.Sku/GetSpuByTheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkuServer is the server API for Sku service.
// All implementations must embed UnimplementedSkuServer
// for forward compatibility
type SkuServer interface {
	GetSkuById(context.Context, *IdRequest) (*SkuVO, error)
	GetSpuByTheme(context.Context, *IdRequest) (*SpuByThemeReply, error)
	mustEmbedUnimplementedSkuServer()
}

// UnimplementedSkuServer must be embedded to have forward compatible implementations.
type UnimplementedSkuServer struct {
}

func (UnimplementedSkuServer) GetSkuById(context.Context, *IdRequest) (*SkuVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkuById not implemented")
}
func (UnimplementedSkuServer) GetSpuByTheme(context.Context, *IdRequest) (*SpuByThemeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpuByTheme not implemented")
}
func (UnimplementedSkuServer) mustEmbedUnimplementedSkuServer() {}

// UnsafeSkuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkuServer will
// result in compilation errors.
type UnsafeSkuServer interface {
	mustEmbedUnimplementedSkuServer()
}

func RegisterSkuServer(s grpc.ServiceRegistrar, srv SkuServer) {
	s.RegisterService(&Sku_ServiceDesc, srv)
}

func _Sku_GetSkuById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkuServer).GetSkuById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.service.Sku/GetSkuById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkuServer).GetSkuById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sku_GetSpuByTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkuServer).GetSpuByTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.service.Sku/GetSpuByTheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkuServer).GetSpuByTheme(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sku_ServiceDesc is the grpc.ServiceDesc for Sku service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sku_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.service.Sku",
	HandlerType: (*SkuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSkuById",
			Handler:    _Sku_GetSkuById_Handler,
		},
		{
			MethodName: "GetSpuByTheme",
			Handler:    _Sku_GetSpuByTheme_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sku/service/sku.proto",
}
