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

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...grpc.CallOption) (*Banner, error)
	GetBannerByName(ctx context.Context, in *BannerByNameRequest, opts ...grpc.CallOption) (*Banner, error)
	GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest, opts ...grpc.CallOption) (*Themes, error)
	//
	GetThemeWithSpu(ctx context.Context, in *ThemeWithSpuRequest, opts ...grpc.CallOption) (*ThemeSpu, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, "/mall.App/GetBannerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetBannerByName(ctx context.Context, in *BannerByNameRequest, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, "/mall.App/GetBannerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest, opts ...grpc.CallOption) (*Themes, error) {
	out := new(Themes)
	err := c.cc.Invoke(ctx, "/mall.App/GetThemeByNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetThemeWithSpu(ctx context.Context, in *ThemeWithSpuRequest, opts ...grpc.CallOption) (*ThemeSpu, error) {
	out := new(ThemeSpu)
	err := c.cc.Invoke(ctx, "/mall.App/GetThemeWithSpu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	GetBannerById(context.Context, *BannerByIdRequest) (*Banner, error)
	GetBannerByName(context.Context, *BannerByNameRequest) (*Banner, error)
	GetThemeByNames(context.Context, *ThemeByNamesRequest) (*Themes, error)
	//
	GetThemeWithSpu(context.Context, *ThemeWithSpuRequest) (*ThemeSpu, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) GetBannerById(context.Context, *BannerByIdRequest) (*Banner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBannerById not implemented")
}
func (UnimplementedAppServer) GetBannerByName(context.Context, *BannerByNameRequest) (*Banner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBannerByName not implemented")
}
func (UnimplementedAppServer) GetThemeByNames(context.Context, *ThemeByNamesRequest) (*Themes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThemeByNames not implemented")
}
func (UnimplementedAppServer) GetThemeWithSpu(context.Context, *ThemeWithSpuRequest) (*ThemeSpu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThemeWithSpu not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_GetBannerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetBannerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.App/GetBannerById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetBannerById(ctx, req.(*BannerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetBannerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetBannerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.App/GetBannerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetBannerByName(ctx, req.(*BannerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetThemeByNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThemeByNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetThemeByNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.App/GetThemeByNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetThemeByNames(ctx, req.(*ThemeByNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetThemeWithSpu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThemeWithSpuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetThemeWithSpu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.App/GetThemeWithSpu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetThemeWithSpu(ctx, req.(*ThemeWithSpuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mall.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBannerById",
			Handler:    _App_GetBannerById_Handler,
		},
		{
			MethodName: "GetBannerByName",
			Handler:    _App_GetBannerByName_Handler,
		},
		{
			MethodName: "GetThemeByNames",
			Handler:    _App_GetThemeByNames_Handler,
		},
		{
			MethodName: "GetThemeWithSpu",
			Handler:    _App_GetThemeWithSpu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mall/app.proto",
}
