// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package service

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OrderHTTPServer interface {
	CreateUser(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.POST("/order", _Order_CreateUser0_HTTP_Handler(srv))
}

func _Order_CreateUser0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/order.service.Order/CreateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderReply)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) CreateUser(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderReply, error) {
	var out CreateOrderReply
	pattern := "/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/order.service.Order/CreateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}