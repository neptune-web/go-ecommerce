// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.19.4
// source: api/app/service/app.proto

package service

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAppGetActivityByName = "/app.service.App/GetActivityByName"
const OperationAppGetActivityWithCoupon = "/app.service.App/GetActivityWithCoupon"
const OperationAppGetBannerById = "/app.service.App/GetBannerById"
const OperationAppGetBannerByName = "/app.service.App/GetBannerByName"
const OperationAppGetTagByType = "/app.service.App/GetTagByType"
const OperationAppGetThemeByName = "/app.service.App/GetThemeByName"
const OperationAppGetThemeByNames = "/app.service.App/GetThemeByNames"
const OperationAppListCategory = "/app.service.App/ListCategory"
const OperationAppListGridCategory = "/app.service.App/ListGridCategory"

type AppHTTPServer interface {
	GetActivityByName(context.Context, *NameRequest) (*Activity, error)
	GetActivityWithCoupon(context.Context, *NameRequest) (*ActivityCoupon, error)
	GetBannerById(context.Context, *IdRequest) (*Banner, error)
	GetBannerByName(context.Context, *NameRequest) (*Banner, error)
	GetTagByType(context.Context, *TypeRequest) (*Tags, error)
	GetThemeByName(context.Context, *NameRequest) (*ThemeSpu, error)
	GetThemeByNames(context.Context, *ThemeByNamesRequest) (*Themes, error)
	ListCategory(context.Context, *emptypb.Empty) (*Categories, error)
	ListGridCategory(context.Context, *emptypb.Empty) (*GridCategories, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/banner/id/{id}", _App_GetBannerById0_HTTP_Handler(srv))
	r.GET("/banner/name/{name}", _App_GetBannerByName0_HTTP_Handler(srv))
	r.GET("/theme/by/names", _App_GetThemeByNames0_HTTP_Handler(srv))
	r.GET("/theme/name/{name}", _App_GetThemeByName0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}", _App_GetActivityByName0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}/with_coupon", _App_GetActivityWithCoupon0_HTTP_Handler(srv))
	r.GET("/category", _App_ListCategory0_HTTP_Handler(srv))
	r.GET("/category/grid", _App_ListGridCategory0_HTTP_Handler(srv))
	r.GET("/tag/type/{type}", _App_GetTagByType0_HTTP_Handler(srv))
}

func _App_GetBannerById0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetBannerById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerById(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Banner)
		return ctx.Result(200, reply)
	}
}

func _App_GetBannerByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetBannerByName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Banner)
		return ctx.Result(200, reply)
	}
}

func _App_GetThemeByNames0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ThemeByNamesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetThemeByNames)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetThemeByNames(ctx, req.(*ThemeByNamesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Themes)
		return ctx.Result(200, reply)
	}
}

func _App_GetThemeByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetThemeByName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetThemeByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ThemeSpu)
		return ctx.Result(200, reply)
	}
}

func _App_GetActivityByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetActivityByName)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Activity)
		return ctx.Result(200, reply)
	}
}

func _App_GetActivityWithCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetActivityWithCoupon)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityWithCoupon(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ActivityCoupon)
		return ctx.Result(200, reply)
	}
}

func _App_ListCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Categories)
		return ctx.Result(200, reply)
	}
}

func _App_ListGridCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppListGridCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGridCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GridCategories)
		return ctx.Result(200, reply)
	}
}

func _App_GetTagByType0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetTagByType)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTagByType(ctx, req.(*TypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tags)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	GetActivityByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *Activity, err error)
	GetActivityWithCoupon(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *ActivityCoupon, err error)
	GetBannerById(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetBannerByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetTagByType(ctx context.Context, req *TypeRequest, opts ...http.CallOption) (rsp *Tags, err error)
	GetThemeByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *ThemeSpu, err error)
	GetThemeByNames(ctx context.Context, req *ThemeByNamesRequest, opts ...http.CallOption) (rsp *Themes, err error)
	ListCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Categories, err error)
	ListGridCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GridCategories, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) GetActivityByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*Activity, error) {
	var out Activity
	pattern := "/activity/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetActivityByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetActivityWithCoupon(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*ActivityCoupon, error) {
	var out ActivityCoupon
	pattern := "/activity/name/{name}/with_coupon"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetActivityWithCoupon))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerById(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/id/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetBannerById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetBannerByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetTagByType(ctx context.Context, in *TypeRequest, opts ...http.CallOption) (*Tags, error) {
	var out Tags
	pattern := "/tag/type/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetTagByType))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetThemeByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*ThemeSpu, error) {
	var out ThemeSpu
	pattern := "/theme/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetThemeByName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest, opts ...http.CallOption) (*Themes, error) {
	var out Themes
	pattern := "/theme/by/names"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppGetThemeByNames))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) ListCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Categories, error) {
	var out Categories
	pattern := "/category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) ListGridCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GridCategories, error) {
	var out GridCategories
	pattern := "/category/grid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppListGridCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
