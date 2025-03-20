// language: proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/model/v1/product_grid.proto

package model_v1

import (
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

type ProductGrid struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               *string                `protobuf:"bytes,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Sku              *string                `protobuf:"bytes,20,opt,name=sku,proto3,oneof" json:"sku,omitempty"`
	Brand            *string                `protobuf:"bytes,30,opt,name=brand,proto3,oneof" json:"brand,omitempty"`
	Name             *string                `protobuf:"bytes,40,opt,name=name,proto3,oneof" json:"name,omitempty"`
	ShortDescription *string                `protobuf:"bytes,50,opt,name=short_description,json=shortDescription,proto3,oneof" json:"short_description,omitempty"`
	Url              *string                `protobuf:"bytes,60,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Status           *string                `protobuf:"bytes,70,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Price            *float32               `protobuf:"fixed32,80,opt,name=price,proto3,oneof" json:"price,omitempty"`
	SalePrice        *float32               `protobuf:"fixed32,90,opt,name=sale_price,json=salePrice,proto3,oneof" json:"sale_price,omitempty"`
	Currency         *string                `protobuf:"bytes,100,opt,name=currency,proto3,oneof" json:"currency,omitempty"`
	Quantity         *int32                 `protobuf:"varint,110,opt,name=quantity,proto3,oneof" json:"quantity,omitempty"`
	IsTaxable        *bool                  `protobuf:"varint,120,opt,name=is_taxable,json=isTaxable,proto3,oneof" json:"is_taxable,omitempty"`
	SeoTitle         *string                `protobuf:"bytes,130,opt,name=seo_title,json=seoTitle,proto3,oneof" json:"seo_title,omitempty"`
	SeoDescription   *string                `protobuf:"bytes,140,opt,name=seo_description,json=seoDescription,proto3,oneof" json:"seo_description,omitempty"`
	Keywords         []string               `protobuf:"bytes,150,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Categories       []string               `protobuf:"bytes,160,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags             []*Tag                 `protobuf:"bytes,170,rep,name=tags,proto3" json:"tags,omitempty"`
	MainImage        *Image                 `protobuf:"bytes,180,opt,name=main_image,json=mainImage,proto3,oneof" json:"main_image,omitempty"`
	HoverImage       *Image                 `protobuf:"bytes,190,opt,name=hover_image,json=hoverImage,proto3,oneof" json:"hover_image,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ProductGrid) Reset() {
	*x = ProductGrid{}
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductGrid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductGrid) ProtoMessage() {}

func (x *ProductGrid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductGrid.ProtoReflect.Descriptor instead.
func (*ProductGrid) Descriptor() ([]byte, []int) {
	return file_proto_model_v1_product_grid_proto_rawDescGZIP(), []int{0}
}

func (x *ProductGrid) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ProductGrid) GetSku() string {
	if x != nil && x.Sku != nil {
		return *x.Sku
	}
	return ""
}

func (x *ProductGrid) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *ProductGrid) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ProductGrid) GetShortDescription() string {
	if x != nil && x.ShortDescription != nil {
		return *x.ShortDescription
	}
	return ""
}

func (x *ProductGrid) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *ProductGrid) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *ProductGrid) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *ProductGrid) GetSalePrice() float32 {
	if x != nil && x.SalePrice != nil {
		return *x.SalePrice
	}
	return 0
}

func (x *ProductGrid) GetCurrency() string {
	if x != nil && x.Currency != nil {
		return *x.Currency
	}
	return ""
}

func (x *ProductGrid) GetQuantity() int32 {
	if x != nil && x.Quantity != nil {
		return *x.Quantity
	}
	return 0
}

func (x *ProductGrid) GetIsTaxable() bool {
	if x != nil && x.IsTaxable != nil {
		return *x.IsTaxable
	}
	return false
}

func (x *ProductGrid) GetSeoTitle() string {
	if x != nil && x.SeoTitle != nil {
		return *x.SeoTitle
	}
	return ""
}

func (x *ProductGrid) GetSeoDescription() string {
	if x != nil && x.SeoDescription != nil {
		return *x.SeoDescription
	}
	return ""
}

func (x *ProductGrid) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *ProductGrid) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ProductGrid) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ProductGrid) GetMainImage() *Image {
	if x != nil {
		return x.MainImage
	}
	return nil
}

func (x *ProductGrid) GetHoverImage() *Image {
	if x != nil {
		return x.HoverImage
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           *string                `protobuf:"bytes,10,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Name          *string                `protobuf:"bytes,20,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Value         *string                `protobuf:"bytes,30,opt,name=value,proto3,oneof" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tag) Reset() {
	*x = Tag{}
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_proto_model_v1_product_grid_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      *string                `protobuf:"bytes,10,opt,name=filename,proto3,oneof" json:"filename,omitempty"`
	Extension     *string                `protobuf:"bytes,20,opt,name=extension,proto3,oneof" json:"extension,omitempty"`
	Title         *string                `protobuf:"bytes,30,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Alt           *string                `protobuf:"bytes,40,opt,name=alt,proto3,oneof" json:"alt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_v1_product_grid_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_model_v1_product_grid_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *Image) GetExtension() string {
	if x != nil && x.Extension != nil {
		return *x.Extension
	}
	return ""
}

func (x *Image) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Image) GetAlt() string {
	if x != nil && x.Alt != nil {
		return *x.Alt
	}
	return ""
}

var File_proto_model_v1_product_grid_proto protoreflect.FileDescriptor

var file_proto_model_v1_product_grid_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x22, 0xfa, 0x06, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x05, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x02, 0x48, 0x07, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x08, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0b,
	0x52, 0x09, 0x69, 0x73, 0x54, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x73, 0x65, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x6f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x73, 0x65, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d, 0x52, 0x0e, 0x73,
	0x65, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x96, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1f, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0xa0, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0xaa, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x48, 0x0e, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0b, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0xbe, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x48, 0x0f, 0x52, 0x0a, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x6b,
	0x75, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75,
	0x72, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x61, 0x6c, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x73, 0x65, 0x6f, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x6b, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xaa, 0x01,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x6c, 0x74, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x52, 0x75, 0x73, 0x61, 0x6b,
	0x6f, 0x76, 0x2f, 0x74, 0x6f, 0x6e, 0x6f, 0x63, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_model_v1_product_grid_proto_rawDescOnce sync.Once
	file_proto_model_v1_product_grid_proto_rawDescData []byte
)

func file_proto_model_v1_product_grid_proto_rawDescGZIP() []byte {
	file_proto_model_v1_product_grid_proto_rawDescOnce.Do(func() {
		file_proto_model_v1_product_grid_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_model_v1_product_grid_proto_rawDesc), len(file_proto_model_v1_product_grid_proto_rawDesc)))
	})
	return file_proto_model_v1_product_grid_proto_rawDescData
}

var file_proto_model_v1_product_grid_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_model_v1_product_grid_proto_goTypes = []any{
	(*ProductGrid)(nil), // 0: proto.model.v1.ProductGrid
	(*Tag)(nil),         // 1: proto.model.v1.Tag
	(*Image)(nil),       // 2: proto.model.v1.Image
}
var file_proto_model_v1_product_grid_proto_depIdxs = []int32{
	1, // 0: proto.model.v1.ProductGrid.tags:type_name -> proto.model.v1.Tag
	2, // 1: proto.model.v1.ProductGrid.main_image:type_name -> proto.model.v1.Image
	2, // 2: proto.model.v1.ProductGrid.hover_image:type_name -> proto.model.v1.Image
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_model_v1_product_grid_proto_init() }
func file_proto_model_v1_product_grid_proto_init() {
	if File_proto_model_v1_product_grid_proto != nil {
		return
	}
	file_proto_model_v1_product_grid_proto_msgTypes[0].OneofWrappers = []any{}
	file_proto_model_v1_product_grid_proto_msgTypes[1].OneofWrappers = []any{}
	file_proto_model_v1_product_grid_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_model_v1_product_grid_proto_rawDesc), len(file_proto_model_v1_product_grid_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_model_v1_product_grid_proto_goTypes,
		DependencyIndexes: file_proto_model_v1_product_grid_proto_depIdxs,
		MessageInfos:      file_proto_model_v1_product_grid_proto_msgTypes,
	}.Build()
	File_proto_model_v1_product_grid_proto = out.File
	file_proto_model_v1_product_grid_proto_goTypes = nil
	file_proto_model_v1_product_grid_proto_depIdxs = nil
}
