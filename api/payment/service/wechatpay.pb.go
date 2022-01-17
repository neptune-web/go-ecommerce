// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/payment/service/wechatpay.proto

package service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
}

func (x *MpRequest) Reset() {
	*x = MpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpRequest) ProtoMessage() {}

func (x *MpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpRequest.ProtoReflect.Descriptor instead.
func (*MpRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{0}
}

func (x *MpRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type MpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MpReply) Reset() {
	*x = MpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpReply) ProtoMessage() {}

func (x *MpReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpReply.ProtoReflect.Descriptor instead.
func (*MpReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{1}
}

type MiniRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
}

func (x *MiniRequest) Reset() {
	*x = MiniRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniRequest) ProtoMessage() {}

func (x *MiniRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniRequest.ProtoReflect.Descriptor instead.
func (*MiniRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{2}
}

func (x *MiniRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type MiniReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MiniReply) Reset() {
	*x = MiniReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniReply) ProtoMessage() {}

func (x *MiniReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniReply.ProtoReflect.Descriptor instead.
func (*MiniReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{3}
}

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{4}
}

type ScanReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScanReply) Reset() {
	*x = ScanReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanReply) ProtoMessage() {}

func (x *ScanReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanReply.ProtoReflect.Descriptor instead.
func (*ScanReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{5}
}

type AppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppRequest) Reset() {
	*x = AppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppRequest) ProtoMessage() {}

func (x *AppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppRequest.ProtoReflect.Descriptor instead.
func (*AppRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{6}
}

type AppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppReply) Reset() {
	*x = AppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppReply) ProtoMessage() {}

func (x *AppReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppReply.ProtoReflect.Descriptor instead.
func (*AppReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{7}
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	OutTradeNo    string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{8}
}

func (x *QueryRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *QueryRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{9}
}

type QueryRefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryRefundRequest) Reset() {
	*x = QueryRefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRefundRequest) ProtoMessage() {}

func (x *QueryRefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRefundRequest.ProtoReflect.Descriptor instead.
func (*QueryRefundRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{10}
}

type QueryRefundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryRefundReply) Reset() {
	*x = QueryRefundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRefundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRefundReply) ProtoMessage() {}

func (x *QueryRefundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRefundReply.ProtoReflect.Descriptor instead.
func (*QueryRefundReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{11}
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{12}
}

type RefundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefundReply) Reset() {
	*x = RefundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundReply) ProtoMessage() {}

func (x *RefundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundReply.ProtoReflect.Descriptor instead.
func (*RefundReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{13}
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{14}
}

type NotifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyReply) Reset() {
	*x = NotifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyReply) ProtoMessage() {}

func (x *NotifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyReply.ProtoReflect.Descriptor instead.
func (*NotifyReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{15}
}

type NotifyRefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyRefundRequest) Reset() {
	*x = NotifyRefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRefundRequest) ProtoMessage() {}

func (x *NotifyRefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRefundRequest.ProtoReflect.Descriptor instead.
func (*NotifyRefundRequest) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{16}
}

type NotifyRefundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyRefundReply) Reset() {
	*x = NotifyRefundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_payment_service_wechatpay_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRefundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRefundReply) ProtoMessage() {}

func (x *NotifyRefundReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_payment_service_wechatpay_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRefundReply.ProtoReflect.Descriptor instead.
func (*NotifyRefundReply) Descriptor() ([]byte, []int) {
	return file_api_payment_service_wechatpay_proto_rawDescGZIP(), []int{17}
}

var File_api_payment_service_wechatpay_proto protoreflect.FileDescriptor

var file_api_payment_service_wechatpay_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x70, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x09, 0x4d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x4e, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x4d, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f,
	0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0c, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22,
	0x0b, 0x0a, 0x09, 0x4d, 0x69, 0x6e, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x53,
	0x63, 0x61, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x57, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x75, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xc1, 0x06, 0x0a, 0x09, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79, 0x12, 0x4a, 0x0a,
	0x02, 0x4d, 0x70, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x08, 0x22, 0x03, 0x2f, 0x6d, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x04, 0x4d, 0x69, 0x6e,
	0x69, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22,
	0x05, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x3a, 0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x03, 0x41, 0x70, 0x70,
	0x12, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x22, 0x04, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x04, 0x53, 0x63, 0x61,
	0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x12, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x3a,
	0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x22, 0x07, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x73,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x24,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x0e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x1d, 0x5a, 0x1b, 0x6d, 0x61, 0x6c, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_payment_service_wechatpay_proto_rawDescOnce sync.Once
	file_api_payment_service_wechatpay_proto_rawDescData = file_api_payment_service_wechatpay_proto_rawDesc
)

func file_api_payment_service_wechatpay_proto_rawDescGZIP() []byte {
	file_api_payment_service_wechatpay_proto_rawDescOnce.Do(func() {
		file_api_payment_service_wechatpay_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_payment_service_wechatpay_proto_rawDescData)
	})
	return file_api_payment_service_wechatpay_proto_rawDescData
}

var file_api_payment_service_wechatpay_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_api_payment_service_wechatpay_proto_goTypes = []interface{}{
	(*MpRequest)(nil),           // 0: payment.service.MpRequest
	(*MpReply)(nil),             // 1: payment.service.MpReply
	(*MiniRequest)(nil),         // 2: payment.service.MiniRequest
	(*MiniReply)(nil),           // 3: payment.service.MiniReply
	(*ScanRequest)(nil),         // 4: payment.service.ScanRequest
	(*ScanReply)(nil),           // 5: payment.service.ScanReply
	(*AppRequest)(nil),          // 6: payment.service.AppRequest
	(*AppReply)(nil),            // 7: payment.service.AppReply
	(*QueryRequest)(nil),        // 8: payment.service.QueryRequest
	(*QueryReply)(nil),          // 9: payment.service.QueryReply
	(*QueryRefundRequest)(nil),  // 10: payment.service.QueryRefundRequest
	(*QueryRefundReply)(nil),    // 11: payment.service.QueryRefundReply
	(*RefundRequest)(nil),       // 12: payment.service.RefundRequest
	(*RefundReply)(nil),         // 13: payment.service.RefundReply
	(*NotifyRequest)(nil),       // 14: payment.service.NotifyRequest
	(*NotifyReply)(nil),         // 15: payment.service.NotifyReply
	(*NotifyRefundRequest)(nil), // 16: payment.service.NotifyRefundRequest
	(*NotifyRefundReply)(nil),   // 17: payment.service.NotifyRefundReply
}
var file_api_payment_service_wechatpay_proto_depIdxs = []int32{
	0,  // 0: payment.service.WechatPay.Mp:input_type -> payment.service.MpRequest
	0,  // 1: payment.service.WechatPay.Mini:input_type -> payment.service.MpRequest
	6,  // 2: payment.service.WechatPay.App:input_type -> payment.service.AppRequest
	4,  // 3: payment.service.WechatPay.Scan:input_type -> payment.service.ScanRequest
	8,  // 4: payment.service.WechatPay.Query:input_type -> payment.service.QueryRequest
	10, // 5: payment.service.WechatPay.QueryRefund:input_type -> payment.service.QueryRefundRequest
	12, // 6: payment.service.WechatPay.Refund:input_type -> payment.service.RefundRequest
	14, // 7: payment.service.WechatPay.Notify:input_type -> payment.service.NotifyRequest
	16, // 8: payment.service.WechatPay.NotifyRefund:input_type -> payment.service.NotifyRefundRequest
	1,  // 9: payment.service.WechatPay.Mp:output_type -> payment.service.MpReply
	1,  // 10: payment.service.WechatPay.Mini:output_type -> payment.service.MpReply
	7,  // 11: payment.service.WechatPay.App:output_type -> payment.service.AppReply
	5,  // 12: payment.service.WechatPay.Scan:output_type -> payment.service.ScanReply
	9,  // 13: payment.service.WechatPay.Query:output_type -> payment.service.QueryReply
	11, // 14: payment.service.WechatPay.QueryRefund:output_type -> payment.service.QueryRefundReply
	13, // 15: payment.service.WechatPay.Refund:output_type -> payment.service.RefundReply
	15, // 16: payment.service.WechatPay.Notify:output_type -> payment.service.NotifyReply
	17, // 17: payment.service.WechatPay.NotifyRefund:output_type -> payment.service.NotifyRefundReply
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_payment_service_wechatpay_proto_init() }
func file_api_payment_service_wechatpay_proto_init() {
	if File_api_payment_service_wechatpay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_payment_service_wechatpay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiniRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiniReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRefundReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_payment_service_wechatpay_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRefundReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_payment_service_wechatpay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_payment_service_wechatpay_proto_goTypes,
		DependencyIndexes: file_api_payment_service_wechatpay_proto_depIdxs,
		MessageInfos:      file_api_payment_service_wechatpay_proto_msgTypes,
	}.Build()
	File_api_payment_service_wechatpay_proto = out.File
	file_api_payment_service_wechatpay_proto_rawDesc = nil
	file_api_payment_service_wechatpay_proto_goTypes = nil
	file_api_payment_service_wechatpay_proto_depIdxs = nil
}
