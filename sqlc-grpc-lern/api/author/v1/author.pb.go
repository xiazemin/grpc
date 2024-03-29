// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: author/v1/author.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAuthorResponse) Reset() {
	*x = CreateAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorResponse) ProtoMessage() {}

func (x *CreateAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{0}
}

type DeleteAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAuthorRequest) Reset() {
	*x = DeleteAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorRequest) ProtoMessage() {}

func (x *DeleteAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteAuthorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAuthorResponse) Reset() {
	*x = DeleteAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorResponse) ProtoMessage() {}

func (x *DeleteAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{2}
}

type GetAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuthorRequest) Reset() {
	*x = GetAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorRequest) ProtoMessage() {}

func (x *GetAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorResponse) Reset() {
	*x = GetAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorResponse) ProtoMessage() {}

func (x *GetAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type ListAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAuthorsRequest) Reset() {
	*x = ListAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsRequest) ProtoMessage() {}

func (x *ListAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{5}
}

type ListAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Author `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListAuthorsResponse) Reset() {
	*x = ListAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsResponse) ProtoMessage() {}

func (x *ListAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{6}
}

func (x *ListAuthorsResponse) GetList() []*Author {
	if x != nil {
		return x.List
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bio         *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	CompanyId   int32                   `protobuf:"varint,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Size        *wrapperspb.Int64Value  `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	EmptyCol    *wrapperspb.Int32Value  `protobuf:"bytes,6,opt,name=empty_col,json=emptyCol,proto3" json:"empty_col,omitempty"`
	DefaultCol  int32                   `protobuf:"varint,7,opt,name=default_col,json=defaultCol,proto3" json:"default_col,omitempty"`
	Size1       *wrapperspb.Int32Value  `protobuf:"bytes,8,opt,name=size1,proto3" json:"size1,omitempty"`
	DefaultCol1 int32                   `protobuf:"varint,9,opt,name=default_col1,json=defaultCol1,proto3" json:"default_col1,omitempty"`
	Type        string                  `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Type1       string                  `protobuf:"bytes,11,opt,name=type1,proto3" json:"type1,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{7}
}

func (x *Author) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetBio() *wrapperspb.StringValue {
	if x != nil {
		return x.Bio
	}
	return nil
}

func (x *Author) GetCompanyId() int32 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *Author) GetSize() *wrapperspb.Int64Value {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *Author) GetEmptyCol() *wrapperspb.Int32Value {
	if x != nil {
		return x.EmptyCol
	}
	return nil
}

func (x *Author) GetDefaultCol() int32 {
	if x != nil {
		return x.DefaultCol
	}
	return 0
}

func (x *Author) GetSize1() *wrapperspb.Int32Value {
	if x != nil {
		return x.Size1
	}
	return nil
}

func (x *Author) GetDefaultCol1() int32 {
	if x != nil {
		return x.DefaultCol1
	}
	return 0
}

func (x *Author) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Author) GetType1() string {
	if x != nil {
		return x.Type1
	}
	return ""
}

type AuthorsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorsType) Reset() {
	*x = AuthorsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsType) ProtoMessage() {}

func (x *AuthorsType) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsType.ProtoReflect.Descriptor instead.
func (*AuthorsType) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{8}
}

type AuthorsType1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorsType1) Reset() {
	*x = AuthorsType1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorsType1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsType1) ProtoMessage() {}

func (x *AuthorsType1) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsType1.ProtoReflect.Descriptor instead.
func (*AuthorsType1) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{9}
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bio  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
	Size *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_v1_author_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_v1_author_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_v1_author_proto_rawDescGZIP(), []int{10}
}

func (x *CreateAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAuthorRequest) GetBio() *wrapperspb.StringValue {
	if x != nil {
		return x.Bio
	}
	return nil
}

func (x *CreateAuthorRequest) GetSize() *wrapperspb.Int64Value {
	if x != nil {
		return x.Size
	}
	return nil
}

var File_author_v1_author_proto protoreflect.FileDescriptor

var file_author_v1_author_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x14, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3c, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x87, 0x03, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x2f,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x69,
	0x7a, 0x65, 0x31, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x31, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x31, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6c, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x31, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x31, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x54, 0x79, 0x70, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x54, 0x79, 0x70, 0x65, 0x31, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32, 0xa7, 0x03, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x22, 0x07, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x62, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x64, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x12, 0x08, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0xe3, 0x01, 0x5a, 0x10, 0x6d, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x2f, 0x76, 0x31, 0x92, 0x41, 0xcd, 0x01, 0x12, 0xca, 0x01, 0x0a, 0x02, 0x6d, 0x79,
	0x12, 0x83, 0x01, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x63,
	0x6f, 0x64, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x62, 0x79,
	0x20, 0x2a, 0x2a, 0x73, 0x71, 0x6c, 0x63, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2a, 0x2a, 0x2e, 0x20,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x6e, 0x20, 0x72, 0x75, 0x6e, 0x20, 0x60, 0x62, 0x75, 0x66, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x60, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x22, 0x39, 0x0a, 0x09, 0x73, 0x71, 0x6c, 0x63, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x77, 0x61,
	0x6e, 0x64, 0x65, 0x72, 0x6c, 0x65, 0x79, 0x2f, 0x73, 0x71, 0x6c, 0x63, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_author_v1_author_proto_rawDescOnce sync.Once
	file_author_v1_author_proto_rawDescData = file_author_v1_author_proto_rawDesc
)

func file_author_v1_author_proto_rawDescGZIP() []byte {
	file_author_v1_author_proto_rawDescOnce.Do(func() {
		file_author_v1_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_author_v1_author_proto_rawDescData)
	})
	return file_author_v1_author_proto_rawDescData
}

var file_author_v1_author_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_author_v1_author_proto_goTypes = []interface{}{
	(*CreateAuthorResponse)(nil),   // 0: author.v1.CreateAuthorResponse
	(*DeleteAuthorRequest)(nil),    // 1: author.v1.DeleteAuthorRequest
	(*DeleteAuthorResponse)(nil),   // 2: author.v1.DeleteAuthorResponse
	(*GetAuthorRequest)(nil),       // 3: author.v1.GetAuthorRequest
	(*GetAuthorResponse)(nil),      // 4: author.v1.GetAuthorResponse
	(*ListAuthorsRequest)(nil),     // 5: author.v1.ListAuthorsRequest
	(*ListAuthorsResponse)(nil),    // 6: author.v1.ListAuthorsResponse
	(*Author)(nil),                 // 7: author.v1.Author
	(*AuthorsType)(nil),            // 8: author.v1.AuthorsType
	(*AuthorsType1)(nil),           // 9: author.v1.AuthorsType1
	(*CreateAuthorRequest)(nil),    // 10: author.v1.CreateAuthorRequest
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 12: google.protobuf.Int64Value
	(*wrapperspb.Int32Value)(nil),  // 13: google.protobuf.Int32Value
}
var file_author_v1_author_proto_depIdxs = []int32{
	7,  // 0: author.v1.GetAuthorResponse.author:type_name -> author.v1.Author
	7,  // 1: author.v1.ListAuthorsResponse.list:type_name -> author.v1.Author
	11, // 2: author.v1.Author.bio:type_name -> google.protobuf.StringValue
	12, // 3: author.v1.Author.size:type_name -> google.protobuf.Int64Value
	13, // 4: author.v1.Author.empty_col:type_name -> google.protobuf.Int32Value
	13, // 5: author.v1.Author.size1:type_name -> google.protobuf.Int32Value
	11, // 6: author.v1.CreateAuthorRequest.bio:type_name -> google.protobuf.StringValue
	12, // 7: author.v1.CreateAuthorRequest.size:type_name -> google.protobuf.Int64Value
	10, // 8: author.v1.AuthorService.CreateAuthor:input_type -> author.v1.CreateAuthorRequest
	1,  // 9: author.v1.AuthorService.DeleteAuthor:input_type -> author.v1.DeleteAuthorRequest
	3,  // 10: author.v1.AuthorService.GetAuthor:input_type -> author.v1.GetAuthorRequest
	5,  // 11: author.v1.AuthorService.ListAuthors:input_type -> author.v1.ListAuthorsRequest
	0,  // 12: author.v1.AuthorService.CreateAuthor:output_type -> author.v1.CreateAuthorResponse
	2,  // 13: author.v1.AuthorService.DeleteAuthor:output_type -> author.v1.DeleteAuthorResponse
	4,  // 14: author.v1.AuthorService.GetAuthor:output_type -> author.v1.GetAuthorResponse
	6,  // 15: author.v1.AuthorService.ListAuthors:output_type -> author.v1.ListAuthorsResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_author_v1_author_proto_init() }
func file_author_v1_author_proto_init() {
	if File_author_v1_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_author_v1_author_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorResponse); i {
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
		file_author_v1_author_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorRequest); i {
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
		file_author_v1_author_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorResponse); i {
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
		file_author_v1_author_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorRequest); i {
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
		file_author_v1_author_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorResponse); i {
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
		file_author_v1_author_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsRequest); i {
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
		file_author_v1_author_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsResponse); i {
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
		file_author_v1_author_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_author_v1_author_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorsType); i {
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
		file_author_v1_author_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorsType1); i {
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
		file_author_v1_author_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorRequest); i {
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
			RawDescriptor: file_author_v1_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_author_v1_author_proto_goTypes,
		DependencyIndexes: file_author_v1_author_proto_depIdxs,
		MessageInfos:      file_author_v1_author_proto_msgTypes,
	}.Build()
	File_author_v1_author_proto = out.File
	file_author_v1_author_proto_rawDesc = nil
	file_author_v1_author_proto_goTypes = nil
	file_author_v1_author_proto_depIdxs = nil
}
