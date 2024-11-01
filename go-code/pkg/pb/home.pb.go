// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: home.proto

package pb

import (
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

type HomeIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *HomeIndexRequest) Reset() {
	*x = HomeIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeIndexRequest) ProtoMessage() {}

func (x *HomeIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeIndexRequest.ProtoReflect.Descriptor instead.
func (*HomeIndexRequest) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{0}
}

func (x *HomeIndexRequest) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

type HomeIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories         []string                `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	TodaysSpecialOffer *HomeTodaysSpecialOffer `protobuf:"bytes,2,opt,name=todays_special_offer,json=todaysSpecialOffer,proto3" json:"todays_special_offer,omitempty"`
	PopularNow         []*HomePopularNow       `protobuf:"bytes,3,rep,name=popular_now,json=popularNow,proto3" json:"popular_now,omitempty"`
}

func (x *HomeIndexResponse) Reset() {
	*x = HomeIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeIndexResponse) ProtoMessage() {}

func (x *HomeIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeIndexResponse.ProtoReflect.Descriptor instead.
func (*HomeIndexResponse) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{1}
}

func (x *HomeIndexResponse) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *HomeIndexResponse) GetTodaysSpecialOffer() *HomeTodaysSpecialOffer {
	if x != nil {
		return x.TodaysSpecialOffer
	}
	return nil
}

func (x *HomeIndexResponse) GetPopularNow() []*HomePopularNow {
	if x != nil {
		return x.PopularNow
	}
	return nil
}

type HomeTodaysSpecialOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Img             string     `protobuf:"bytes,2,opt,name=img,proto3" json:"img,omitempty"`
	Description     string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price           *HomeMoney `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Discount        float32    `protobuf:"fixed32,5,opt,name=discount,proto3" json:"discount,omitempty"`
	PriceDiscounted *HomeMoney `protobuf:"bytes,6,opt,name=price_discounted,json=priceDiscounted,proto3" json:"price_discounted,omitempty"`
}

func (x *HomeTodaysSpecialOffer) Reset() {
	*x = HomeTodaysSpecialOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeTodaysSpecialOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeTodaysSpecialOffer) ProtoMessage() {}

func (x *HomeTodaysSpecialOffer) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeTodaysSpecialOffer.ProtoReflect.Descriptor instead.
func (*HomeTodaysSpecialOffer) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{2}
}

func (x *HomeTodaysSpecialOffer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HomeTodaysSpecialOffer) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *HomeTodaysSpecialOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HomeTodaysSpecialOffer) GetPrice() *HomeMoney {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *HomeTodaysSpecialOffer) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *HomeTodaysSpecialOffer) GetPriceDiscounted() *HomeMoney {
	if x != nil {
		return x.PriceDiscounted
	}
	return nil
}

type HomePopularNow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Img        string     `protobuf:"bytes,2,opt,name=img,proto3" json:"img,omitempty"`
	Price      *HomeMoney `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	IsFavorite bool       `protobuf:"varint,4,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
}

func (x *HomePopularNow) Reset() {
	*x = HomePopularNow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomePopularNow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomePopularNow) ProtoMessage() {}

func (x *HomePopularNow) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomePopularNow.ProtoReflect.Descriptor instead.
func (*HomePopularNow) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{3}
}

func (x *HomePopularNow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HomePopularNow) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *HomePopularNow) GetPrice() *HomeMoney {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *HomePopularNow) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type HomeMoney struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Decimals uint32 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *HomeMoney) Reset() {
	*x = HomeMoney{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeMoney) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeMoney) ProtoMessage() {}

func (x *HomeMoney) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeMoney.ProtoReflect.Descriptor instead.
func (*HomeMoney) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{4}
}

func (x *HomeMoney) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *HomeMoney) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *HomeMoney) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_home_proto protoreflect.FileDescriptor

var file_home_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x79, 0x75,
	0x6d, 0x6d, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x10, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x14, 0x74,
	0x6f, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x75, 0x6d, 0x6d,
	0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x12, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x0b, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x6e, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x79, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f,
	0x6d, 0x65, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x77, 0x52, 0x0a, 0x70, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x77, 0x22, 0xe5, 0x01, 0x0a, 0x16, 0x48, 0x6f, 0x6d,
	0x65, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x75, 0x6d,
	0x6d, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x75,
	0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52,
	0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64,
	0x22, 0x81, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72,
	0x4e, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x75, 0x6d, 0x6d, 0x69,
	0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x22, 0x5b, 0x0a, 0x09, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x32, 0x4d, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x2e, 0x79, 0x75, 0x6d, 0x6d,
	0x69, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x79, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x6f, 0x73, 0x65, 0x6c, 0x75, 0x69, 0x73, 0x38, 0x39, 0x30, 0x36, 0x2f, 0x79, 0x75, 0x6d, 0x6d,
	0x69, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_proto_rawDescOnce sync.Once
	file_home_proto_rawDescData = file_home_proto_rawDesc
)

func file_home_proto_rawDescGZIP() []byte {
	file_home_proto_rawDescOnce.Do(func() {
		file_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_proto_rawDescData)
	})
	return file_home_proto_rawDescData
}

var file_home_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_home_proto_goTypes = []interface{}{
	(*HomeIndexRequest)(nil),       // 0: yummies.HomeIndexRequest
	(*HomeIndexResponse)(nil),      // 1: yummies.HomeIndexResponse
	(*HomeTodaysSpecialOffer)(nil), // 2: yummies.HomeTodaysSpecialOffer
	(*HomePopularNow)(nil),         // 3: yummies.HomePopularNow
	(*HomeMoney)(nil),              // 4: yummies.HomeMoney
}
var file_home_proto_depIdxs = []int32{
	2, // 0: yummies.HomeIndexResponse.todays_special_offer:type_name -> yummies.HomeTodaysSpecialOffer
	3, // 1: yummies.HomeIndexResponse.popular_now:type_name -> yummies.HomePopularNow
	4, // 2: yummies.HomeTodaysSpecialOffer.price:type_name -> yummies.HomeMoney
	4, // 3: yummies.HomeTodaysSpecialOffer.price_discounted:type_name -> yummies.HomeMoney
	4, // 4: yummies.HomePopularNow.price:type_name -> yummies.HomeMoney
	0, // 5: yummies.HomeService.Index:input_type -> yummies.HomeIndexRequest
	1, // 6: yummies.HomeService.Index:output_type -> yummies.HomeIndexResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_home_proto_init() }
func file_home_proto_init() {
	if File_home_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeIndexRequest); i {
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
		file_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeIndexResponse); i {
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
		file_home_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeTodaysSpecialOffer); i {
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
		file_home_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomePopularNow); i {
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
		file_home_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeMoney); i {
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
			RawDescriptor: file_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_home_proto_goTypes,
		DependencyIndexes: file_home_proto_depIdxs,
		MessageInfos:      file_home_proto_msgTypes,
	}.Build()
	File_home_proto = out.File
	file_home_proto_rawDesc = nil
	file_home_proto_goTypes = nil
	file_home_proto_depIdxs = nil
}
