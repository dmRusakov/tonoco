// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/service/v1/product.proto

package service_v1

import (
	v11 "github.com/dmRusakov/tonoco-grpc/gen/go/proto/model/v1"
	v1 "github.com/dmRusakov/tonoco/gen/go/proto/model/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetProductGrid
type GetProductGridRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductGridRequest) Reset() {
	*x = GetProductGridRequest{}
	mi := &file_proto_service_v1_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductGridRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductGridRequest) ProtoMessage() {}

func (x *GetProductGridRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_v1_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductGridRequest.ProtoReflect.Descriptor instead.
func (*GetProductGridRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductGridRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductGridResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductGrid   *v1.ProductGrid        `protobuf:"bytes,10,opt,name=productGrid,proto3" json:"productGrid,omitempty"`
	Error         *v11.Error             `protobuf:"bytes,9999,opt,name=error,proto3,oneof" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductGridResponse) Reset() {
	*x = GetProductGridResponse{}
	mi := &file_proto_service_v1_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductGridResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductGridResponse) ProtoMessage() {}

func (x *GetProductGridResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_v1_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductGridResponse.ProtoReflect.Descriptor instead.
func (*GetProductGridResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_v1_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductGridResponse) GetProductGrid() *v1.ProductGrid {
	if x != nil {
		return x.ProductGrid
	}
	return nil
}

func (x *GetProductGridResponse) GetError() *v11.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// GetStatus
type GetStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	mi := &file_proto_service_v1_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_v1_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_v1_product_proto_rawDescGZIP(), []int{2}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *v11.Status            `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	mi := &file_proto_service_v1_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_v1_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_v1_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatusResponse) GetStatus() *v11.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_proto_service_v1_product_proto protoreflect.FileDescriptor

var file_proto_service_v1_product_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x69, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x8f, 0x4e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xff, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x69, 0x64, 0x12, 0x27,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x69, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x68, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x52, 0x75, 0x73, 0x61, 0x6b,
	0x6f, 0x76, 0x2f, 0x74, 0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_service_v1_product_proto_rawDescOnce sync.Once
	file_proto_service_v1_product_proto_rawDescData []byte
)

func file_proto_service_v1_product_proto_rawDescGZIP() []byte {
	file_proto_service_v1_product_proto_rawDescOnce.Do(func() {
		file_proto_service_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_service_v1_product_proto_rawDesc), len(file_proto_service_v1_product_proto_rawDesc)))
	})
	return file_proto_service_v1_product_proto_rawDescData
}

var file_proto_service_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_service_v1_product_proto_goTypes = []any{
	(*GetProductGridRequest)(nil),  // 0: proto.service.v1.GetProductGridRequest
	(*GetProductGridResponse)(nil), // 1: proto.service.v1.GetProductGridResponse
	(*GetStatusRequest)(nil),       // 2: proto.service.v1.GetStatusRequest
	(*GetStatusResponse)(nil),      // 3: proto.service.v1.GetStatusResponse
	(*v1.ProductGrid)(nil),         // 4: proto.model.v1.ProductGrid
	(*v11.Error)(nil),              // 5: proto.model.v1.Error
	(*v11.Status)(nil),             // 6: proto.model.v1.Status
}
var file_proto_service_v1_product_proto_depIdxs = []int32{
	4, // 0: proto.service.v1.GetProductGridResponse.productGrid:type_name -> proto.model.v1.ProductGrid
	5, // 1: proto.service.v1.GetProductGridResponse.error:type_name -> proto.model.v1.Error
	6, // 2: proto.service.v1.GetStatusResponse.status:type_name -> proto.model.v1.Status
	0, // 3: proto.service.v1.ProductService.GetProductGrid:input_type -> proto.service.v1.GetProductGridRequest
	2, // 4: proto.service.v1.ProductService.GetStatus:input_type -> proto.service.v1.GetStatusRequest
	1, // 5: proto.service.v1.ProductService.GetProductGrid:output_type -> proto.service.v1.GetProductGridResponse
	3, // 6: proto.service.v1.ProductService.GetStatus:output_type -> proto.service.v1.GetStatusResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_service_v1_product_proto_init() }
func file_proto_service_v1_product_proto_init() {
	if File_proto_service_v1_product_proto != nil {
		return
	}
	file_proto_service_v1_product_proto_msgTypes[1].OneofWrappers = []any{}
	file_proto_service_v1_product_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_v1_product_proto_rawDesc), len(file_proto_service_v1_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_v1_product_proto_goTypes,
		DependencyIndexes: file_proto_service_v1_product_proto_depIdxs,
		MessageInfos:      file_proto_service_v1_product_proto_msgTypes,
	}.Build()
	File_proto_service_v1_product_proto = out.File
	file_proto_service_v1_product_proto_goTypes = nil
	file_proto_service_v1_product_proto_depIdxs = nil
}
