// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package app

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	GetActivityByName(ctx context.Context, in *ActivityByNameRequest, opts ...grpc.CallOption) (*Activity, error)
	GetActivityWithCoupon(ctx context.Context, in *ActivityWithCouponRequest, opts ...grpc.CallOption) (*ActivityCoupon, error)
	GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllCategory, error)
	GetGridCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GridCategories, error)
	GetTagByType(ctx context.Context, in *TagByTypeRequest, opts ...grpc.CallOption) (*Tags, error)
	GetCouponByCategory(ctx context.Context, in *CouponByCategoryRequest, opts ...grpc.CallOption) (*Coupons, error)
	GetWholeCoupon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Coupons, error)
	GetMyCouponByStatus(ctx context.Context, in *MyCouponByStatusRequest, opts ...grpc.CallOption) (*Coupons, error)
	GetMyAvailableCoupon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Coupons, error)
	CollectCoupon(ctx context.Context, in *CollectCouponRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSaleExplain(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SaleExplains, error)
	GetSpuById(ctx context.Context, in *SpuByIdRequest, opts ...grpc.CallOption) (*SpuDetail, error)
	GetSpuLatest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SpuPage, error)
	GetSpuByCategory(ctx context.Context, in *SpuByCategoryRequest, opts ...grpc.CallOption) (*SpuPage, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, "/app.App/GetBannerById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetBannerByName(ctx context.Context, in *BannerByNameRequest, opts ...grpc.CallOption) (*Banner, error) {
	out := new(Banner)
	err := c.cc.Invoke(ctx, "/app.App/GetBannerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest, opts ...grpc.CallOption) (*Themes, error) {
	out := new(Themes)
	err := c.cc.Invoke(ctx, "/app.App/GetThemeByNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetThemeWithSpu(ctx context.Context, in *ThemeWithSpuRequest, opts ...grpc.CallOption) (*ThemeSpu, error) {
	out := new(ThemeSpu)
	err := c.cc.Invoke(ctx, "/app.App/GetThemeWithSpu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetActivityByName(ctx context.Context, in *ActivityByNameRequest, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/app.App/GetActivityByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetActivityWithCoupon(ctx context.Context, in *ActivityWithCouponRequest, opts ...grpc.CallOption) (*ActivityCoupon, error) {
	out := new(ActivityCoupon)
	err := c.cc.Invoke(ctx, "/app.App/GetActivityWithCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllCategory, error) {
	out := new(AllCategory)
	err := c.cc.Invoke(ctx, "/app.App/GetAllCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetGridCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GridCategories, error) {
	out := new(GridCategories)
	err := c.cc.Invoke(ctx, "/app.App/GetGridCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetTagByType(ctx context.Context, in *TagByTypeRequest, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := c.cc.Invoke(ctx, "/app.App/GetTagByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetCouponByCategory(ctx context.Context, in *CouponByCategoryRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/app.App/GetCouponByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetWholeCoupon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/app.App/GetWholeCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetMyCouponByStatus(ctx context.Context, in *MyCouponByStatusRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/app.App/GetMyCouponByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetMyAvailableCoupon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/app.App/GetMyAvailableCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) CollectCoupon(ctx context.Context, in *CollectCouponRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/app.App/CollectCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetSaleExplain(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SaleExplains, error) {
	out := new(SaleExplains)
	err := c.cc.Invoke(ctx, "/app.App/GetSaleExplain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetSpuById(ctx context.Context, in *SpuByIdRequest, opts ...grpc.CallOption) (*SpuDetail, error) {
	out := new(SpuDetail)
	err := c.cc.Invoke(ctx, "/app.App/GetSpuById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetSpuLatest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SpuPage, error) {
	out := new(SpuPage)
	err := c.cc.Invoke(ctx, "/app.App/GetSpuLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetSpuByCategory(ctx context.Context, in *SpuByCategoryRequest, opts ...grpc.CallOption) (*SpuPage, error) {
	out := new(SpuPage)
	err := c.cc.Invoke(ctx, "/app.App/GetSpuByCategory", in, out, opts...)
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
	GetActivityByName(context.Context, *ActivityByNameRequest) (*Activity, error)
	GetActivityWithCoupon(context.Context, *ActivityWithCouponRequest) (*ActivityCoupon, error)
	GetAllCategory(context.Context, *emptypb.Empty) (*AllCategory, error)
	GetGridCategory(context.Context, *emptypb.Empty) (*GridCategories, error)
	GetTagByType(context.Context, *TagByTypeRequest) (*Tags, error)
	GetCouponByCategory(context.Context, *CouponByCategoryRequest) (*Coupons, error)
	GetWholeCoupon(context.Context, *emptypb.Empty) (*Coupons, error)
	GetMyCouponByStatus(context.Context, *MyCouponByStatusRequest) (*Coupons, error)
	GetMyAvailableCoupon(context.Context, *emptypb.Empty) (*Coupons, error)
	CollectCoupon(context.Context, *CollectCouponRequest) (*emptypb.Empty, error)
	GetSaleExplain(context.Context, *emptypb.Empty) (*SaleExplains, error)
	GetSpuById(context.Context, *SpuByIdRequest) (*SpuDetail, error)
	GetSpuLatest(context.Context, *emptypb.Empty) (*SpuPage, error)
	GetSpuByCategory(context.Context, *SpuByCategoryRequest) (*SpuPage, error)
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
func (UnimplementedAppServer) GetActivityByName(context.Context, *ActivityByNameRequest) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityByName not implemented")
}
func (UnimplementedAppServer) GetActivityWithCoupon(context.Context, *ActivityWithCouponRequest) (*ActivityCoupon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityWithCoupon not implemented")
}
func (UnimplementedAppServer) GetAllCategory(context.Context, *emptypb.Empty) (*AllCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategory not implemented")
}
func (UnimplementedAppServer) GetGridCategory(context.Context, *emptypb.Empty) (*GridCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGridCategory not implemented")
}
func (UnimplementedAppServer) GetTagByType(context.Context, *TagByTypeRequest) (*Tags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagByType not implemented")
}
func (UnimplementedAppServer) GetCouponByCategory(context.Context, *CouponByCategoryRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponByCategory not implemented")
}
func (UnimplementedAppServer) GetWholeCoupon(context.Context, *emptypb.Empty) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWholeCoupon not implemented")
}
func (UnimplementedAppServer) GetMyCouponByStatus(context.Context, *MyCouponByStatusRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCouponByStatus not implemented")
}
func (UnimplementedAppServer) GetMyAvailableCoupon(context.Context, *emptypb.Empty) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyAvailableCoupon not implemented")
}
func (UnimplementedAppServer) CollectCoupon(context.Context, *CollectCouponRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectCoupon not implemented")
}
func (UnimplementedAppServer) GetSaleExplain(context.Context, *emptypb.Empty) (*SaleExplains, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaleExplain not implemented")
}
func (UnimplementedAppServer) GetSpuById(context.Context, *SpuByIdRequest) (*SpuDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpuById not implemented")
}
func (UnimplementedAppServer) GetSpuLatest(context.Context, *emptypb.Empty) (*SpuPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpuLatest not implemented")
}
func (UnimplementedAppServer) GetSpuByCategory(context.Context, *SpuByCategoryRequest) (*SpuPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpuByCategory not implemented")
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
		FullMethod: "/app.App/GetBannerById",
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
		FullMethod: "/app.App/GetBannerByName",
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
		FullMethod: "/app.App/GetThemeByNames",
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
		FullMethod: "/app.App/GetThemeWithSpu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetThemeWithSpu(ctx, req.(*ThemeWithSpuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetActivityByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetActivityByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetActivityByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetActivityByName(ctx, req.(*ActivityByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetActivityWithCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityWithCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetActivityWithCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetActivityWithCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetActivityWithCoupon(ctx, req.(*ActivityWithCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetAllCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetAllCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetAllCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetAllCategory(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetGridCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGridCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetGridCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGridCategory(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetTagByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetTagByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetTagByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetTagByType(ctx, req.(*TagByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetCouponByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CouponByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetCouponByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetCouponByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetCouponByCategory(ctx, req.(*CouponByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetWholeCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetWholeCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetWholeCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetWholeCoupon(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetMyCouponByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyCouponByStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetMyCouponByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetMyCouponByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetMyCouponByStatus(ctx, req.(*MyCouponByStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetMyAvailableCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetMyAvailableCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetMyAvailableCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetMyAvailableCoupon(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_CollectCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).CollectCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/CollectCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).CollectCoupon(ctx, req.(*CollectCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetSaleExplain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetSaleExplain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetSaleExplain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetSaleExplain(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetSpuById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpuByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetSpuById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetSpuById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetSpuById(ctx, req.(*SpuByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetSpuLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetSpuLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetSpuLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetSpuLatest(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetSpuByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpuByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetSpuByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/GetSpuByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetSpuByCategory(ctx, req.(*SpuByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
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
		{
			MethodName: "GetActivityByName",
			Handler:    _App_GetActivityByName_Handler,
		},
		{
			MethodName: "GetActivityWithCoupon",
			Handler:    _App_GetActivityWithCoupon_Handler,
		},
		{
			MethodName: "GetAllCategory",
			Handler:    _App_GetAllCategory_Handler,
		},
		{
			MethodName: "GetGridCategory",
			Handler:    _App_GetGridCategory_Handler,
		},
		{
			MethodName: "GetTagByType",
			Handler:    _App_GetTagByType_Handler,
		},
		{
			MethodName: "GetCouponByCategory",
			Handler:    _App_GetCouponByCategory_Handler,
		},
		{
			MethodName: "GetWholeCoupon",
			Handler:    _App_GetWholeCoupon_Handler,
		},
		{
			MethodName: "GetMyCouponByStatus",
			Handler:    _App_GetMyCouponByStatus_Handler,
		},
		{
			MethodName: "GetMyAvailableCoupon",
			Handler:    _App_GetMyAvailableCoupon_Handler,
		},
		{
			MethodName: "CollectCoupon",
			Handler:    _App_CollectCoupon_Handler,
		},
		{
			MethodName: "GetSaleExplain",
			Handler:    _App_GetSaleExplain_Handler,
		},
		{
			MethodName: "GetSpuById",
			Handler:    _App_GetSpuById_Handler,
		},
		{
			MethodName: "GetSpuLatest",
			Handler:    _App_GetSpuLatest_Handler,
		},
		{
			MethodName: "GetSpuByCategory",
			Handler:    _App_GetSpuByCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/app/app.proto",
}
