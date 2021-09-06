// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ova-joke-api/ova-joke-api.proto

package ova_joke_api

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

type Joke struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId uint64 `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *Joke) Reset() {
	*x = Joke{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Joke) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Joke) ProtoMessage() {}

func (x *Joke) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Joke.ProtoReflect.Descriptor instead.
func (*Joke) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{0}
}

func (x *Joke) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Joke) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Joke) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type CreateJokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId uint64 `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *CreateJokeRequest) Reset() {
	*x = CreateJokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJokeRequest) ProtoMessage() {}

func (x *CreateJokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJokeRequest.ProtoReflect.Descriptor instead.
func (*CreateJokeRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJokeRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateJokeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateJokeRequest) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type CreateJokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateJokeResponse) Reset() {
	*x = CreateJokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJokeResponse) ProtoMessage() {}

func (x *CreateJokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJokeResponse.ProtoReflect.Descriptor instead.
func (*CreateJokeResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{2}
}

type DescribeJokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeJokeRequest) Reset() {
	*x = DescribeJokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeJokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJokeRequest) ProtoMessage() {}

func (x *DescribeJokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJokeRequest.ProtoReflect.Descriptor instead.
func (*DescribeJokeRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeJokeRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeJokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId uint64 `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *DescribeJokeResponse) Reset() {
	*x = DescribeJokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeJokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJokeResponse) ProtoMessage() {}

func (x *DescribeJokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJokeResponse.ProtoReflect.Descriptor instead.
func (*DescribeJokeResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeJokeResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DescribeJokeResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *DescribeJokeResponse) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type ListJokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListJokeRequest) Reset() {
	*x = ListJokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJokeRequest) ProtoMessage() {}

func (x *ListJokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJokeRequest.ProtoReflect.Descriptor instead.
func (*ListJokeRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListJokeRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListJokeRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListJokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jokes []*Joke `protobuf:"bytes,1,rep,name=jokes,proto3" json:"jokes,omitempty"`
}

func (x *ListJokeResponse) Reset() {
	*x = ListJokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJokeResponse) ProtoMessage() {}

func (x *ListJokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJokeResponse.ProtoReflect.Descriptor instead.
func (*ListJokeResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListJokeResponse) GetJokes() []*Joke {
	if x != nil {
		return x.Jokes
	}
	return nil
}

type RemoveJokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveJokeRequest) Reset() {
	*x = RemoveJokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveJokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveJokeRequest) ProtoMessage() {}

func (x *RemoveJokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveJokeRequest.ProtoReflect.Descriptor instead.
func (*RemoveJokeRequest) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveJokeRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveJokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveJokeResponse) Reset() {
	*x = RemoveJokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveJokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveJokeResponse) ProtoMessage() {}

func (x *RemoveJokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_joke_api_ova_joke_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveJokeResponse.ProtoReflect.Descriptor instead.
func (*RemoveJokeResponse) Descriptor() ([]byte, []int) {
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP(), []int{8}
}

var File_api_ova_joke_api_ova_joke_api_proto protoreflect.FileDescriptor

var file_api_ova_joke_api_ova_joke_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6a, 0x6f, 0x6b, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6a, 0x6f, 0x6b, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76,
	0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x47, 0x0a, 0x04, 0x4a, 0x6f,
	0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22,
	0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x43, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6a, 0x6f, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61,
	0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x05,
	0x6a, 0x6f, 0x6b, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4a,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x91, 0x03, 0x0a, 0x0b, 0x4a, 0x6f, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x12, 0x26,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e,
	0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x6b,
	0x65, 0x12, 0x28, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a,
	0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x4a, 0x6f, 0x6b, 0x65, 0x12, 0x24, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76,
	0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a,
	0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x6b,
	0x65, 0x12, 0x26, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a,
	0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4a, 0x6f,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x76, 0x61, 0x2e, 0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6a, 0x6f,
	0x6b, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x6a,
	0x6f, 0x6b, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x6a, 0x6f, 0x6b, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ova_joke_api_ova_joke_api_proto_rawDescOnce sync.Once
	file_api_ova_joke_api_ova_joke_api_proto_rawDescData = file_api_ova_joke_api_ova_joke_api_proto_rawDesc
)

func file_api_ova_joke_api_ova_joke_api_proto_rawDescGZIP() []byte {
	file_api_ova_joke_api_ova_joke_api_proto_rawDescOnce.Do(func() {
		file_api_ova_joke_api_ova_joke_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ova_joke_api_ova_joke_api_proto_rawDescData)
	})
	return file_api_ova_joke_api_ova_joke_api_proto_rawDescData
}

var file_api_ova_joke_api_ova_joke_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ova_joke_api_ova_joke_api_proto_goTypes = []interface{}{
	(*Joke)(nil),                 // 0: ozonva.ova_joke_api.Joke
	(*CreateJokeRequest)(nil),    // 1: ozonva.ova_joke_api.CreateJokeRequest
	(*CreateJokeResponse)(nil),   // 2: ozonva.ova_joke_api.CreateJokeResponse
	(*DescribeJokeRequest)(nil),  // 3: ozonva.ova_joke_api.DescribeJokeRequest
	(*DescribeJokeResponse)(nil), // 4: ozonva.ova_joke_api.DescribeJokeResponse
	(*ListJokeRequest)(nil),      // 5: ozonva.ova_joke_api.ListJokeRequest
	(*ListJokeResponse)(nil),     // 6: ozonva.ova_joke_api.ListJokeResponse
	(*RemoveJokeRequest)(nil),    // 7: ozonva.ova_joke_api.RemoveJokeRequest
	(*RemoveJokeResponse)(nil),   // 8: ozonva.ova_joke_api.RemoveJokeResponse
}
var file_api_ova_joke_api_ova_joke_api_proto_depIdxs = []int32{
	0, // 0: ozonva.ova_joke_api.ListJokeResponse.jokes:type_name -> ozonva.ova_joke_api.Joke
	1, // 1: ozonva.ova_joke_api.JokeService.CreateJoke:input_type -> ozonva.ova_joke_api.CreateJokeRequest
	3, // 2: ozonva.ova_joke_api.JokeService.DescribeJoke:input_type -> ozonva.ova_joke_api.DescribeJokeRequest
	5, // 3: ozonva.ova_joke_api.JokeService.ListJoke:input_type -> ozonva.ova_joke_api.ListJokeRequest
	7, // 4: ozonva.ova_joke_api.JokeService.RemoveJoke:input_type -> ozonva.ova_joke_api.RemoveJokeRequest
	2, // 5: ozonva.ova_joke_api.JokeService.CreateJoke:output_type -> ozonva.ova_joke_api.CreateJokeResponse
	4, // 6: ozonva.ova_joke_api.JokeService.DescribeJoke:output_type -> ozonva.ova_joke_api.DescribeJokeResponse
	6, // 7: ozonva.ova_joke_api.JokeService.ListJoke:output_type -> ozonva.ova_joke_api.ListJokeResponse
	8, // 8: ozonva.ova_joke_api.JokeService.RemoveJoke:output_type -> ozonva.ova_joke_api.RemoveJokeResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ova_joke_api_ova_joke_api_proto_init() }
func file_api_ova_joke_api_ova_joke_api_proto_init() {
	if File_api_ova_joke_api_ova_joke_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Joke); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJokeRequest); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJokeResponse); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeJokeRequest); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeJokeResponse); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJokeRequest); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJokeResponse); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveJokeRequest); i {
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
		file_api_ova_joke_api_ova_joke_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveJokeResponse); i {
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
			RawDescriptor: file_api_ova_joke_api_ova_joke_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ova_joke_api_ova_joke_api_proto_goTypes,
		DependencyIndexes: file_api_ova_joke_api_ova_joke_api_proto_depIdxs,
		MessageInfos:      file_api_ova_joke_api_ova_joke_api_proto_msgTypes,
	}.Build()
	File_api_ova_joke_api_ova_joke_api_proto = out.File
	file_api_ova_joke_api_ova_joke_api_proto_rawDesc = nil
	file_api_ova_joke_api_ova_joke_api_proto_goTypes = nil
	file_api_ova_joke_api_ova_joke_api_proto_depIdxs = nil
}
