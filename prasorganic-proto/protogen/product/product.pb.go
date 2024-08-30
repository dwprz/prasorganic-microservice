// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/product/type/product.proto

package product

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

type ProductIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ProductIds) Reset() {
	*x = ProductIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_type_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductIds) ProtoMessage() {}

func (x *ProductIds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_type_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductIds.ProtoReflect.Descriptor instead.
func (*ProductIds) Descriptor() ([]byte, []int) {
	return file_proto_product_type_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductIds) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ProductOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId     string `protobuf:"bytes,1,opt,name=orderId,json=order_id,proto3" json:"orderId,omitempty"`
	ProductId   uint32 `protobuf:"varint,2,opt,name=productId,json=product_id,proto3" json:"productId,omitempty"`
	ProductName string `protobuf:"bytes,3,opt,name=productName,json=product_name,proto3" json:"productName,omitempty"`
	Quantity    uint32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Price       uint32 `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductOrder) Reset() {
	*x = ProductOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_type_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOrder) ProtoMessage() {}

func (x *ProductOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_type_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOrder.ProtoReflect.Descriptor instead.
func (*ProductOrder) Descriptor() ([]byte, []int) {
	return file_proto_product_type_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductOrder) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *ProductOrder) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductOrder) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductOrder) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductOrder) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ProductOrder) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ReduceStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProductOrder `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ReduceStocksReq) Reset() {
	*x = ReduceStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_type_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceStocksReq) ProtoMessage() {}

func (x *ReduceStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_type_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceStocksReq.ProtoReflect.Descriptor instead.
func (*ReduceStocksReq) Descriptor() ([]byte, []int) {
	return file_proto_product_type_product_proto_rawDescGZIP(), []int{2}
}

func (x *ReduceStocksReq) GetData() []*ProductOrder {
	if x != nil {
		return x.Data
	}
	return nil
}

type RollbackStocksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProductOrder `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RollbackStocksReq) Reset() {
	*x = RollbackStocksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_type_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackStocksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackStocksReq) ProtoMessage() {}

func (x *RollbackStocksReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_type_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackStocksReq.ProtoReflect.Descriptor instead.
func (*RollbackStocksReq) Descriptor() ([]byte, []int) {
	return file_proto_product_type_product_proto_rawDescGZIP(), []int{3}
}

func (x *RollbackStocksReq) GetData() []*ProductOrder {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_product_type_product_proto protoreflect.FileDescriptor

var file_proto_product_type_product_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0xb3, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x70,
	0x72, 0x7a, 0x2f, 0x70, 0x72, 0x61, 0x73, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_type_product_proto_rawDescOnce sync.Once
	file_proto_product_type_product_proto_rawDescData = file_proto_product_type_product_proto_rawDesc
)

func file_proto_product_type_product_proto_rawDescGZIP() []byte {
	file_proto_product_type_product_proto_rawDescOnce.Do(func() {
		file_proto_product_type_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_type_product_proto_rawDescData)
	})
	return file_proto_product_type_product_proto_rawDescData
}

var file_proto_product_type_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_product_type_product_proto_goTypes = []any{
	(*ProductIds)(nil),        // 0: product.type.ProductIds
	(*ProductOrder)(nil),      // 1: product.type.ProductOrder
	(*ReduceStocksReq)(nil),   // 2: product.type.ReduceStocksReq
	(*RollbackStocksReq)(nil), // 3: product.type.RollbackStocksReq
}
var file_proto_product_type_product_proto_depIdxs = []int32{
	1, // 0: product.type.ReduceStocksReq.data:type_name -> product.type.ProductOrder
	1, // 1: product.type.RollbackStocksReq.data:type_name -> product.type.ProductOrder
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_product_type_product_proto_init() }
func file_proto_product_type_product_proto_init() {
	if File_proto_product_type_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_type_product_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProductIds); i {
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
		file_proto_product_type_product_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProductOrder); i {
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
		file_proto_product_type_product_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReduceStocksReq); i {
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
		file_proto_product_type_product_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RollbackStocksReq); i {
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
			RawDescriptor: file_proto_product_type_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_product_type_product_proto_goTypes,
		DependencyIndexes: file_proto_product_type_product_proto_depIdxs,
		MessageInfos:      file_proto_product_type_product_proto_msgTypes,
	}.Build()
	File_proto_product_type_product_proto = out.File
	file_proto_product_type_product_proto_rawDesc = nil
	file_proto_product_type_product_proto_goTypes = nil
	file_proto_product_type_product_proto_depIdxs = nil
}
