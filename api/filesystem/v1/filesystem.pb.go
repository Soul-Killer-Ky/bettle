// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: filesystem/v1/filesystem.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type UploadType int32

const (
	UploadType_UT_File  UploadType = 0
	UploadType_UT_Image UploadType = 1
)

// Enum value maps for UploadType.
var (
	UploadType_name = map[int32]string{
		0: "UT_File",
		1: "UT_Image",
	}
	UploadType_value = map[string]int32{
		"UT_File":  0,
		"UT_Image": 1,
	}
)

func (x UploadType) Enum() *UploadType {
	p := new(UploadType)
	*p = x
	return p
}

func (x UploadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadType) Descriptor() protoreflect.EnumDescriptor {
	return file_filesystem_v1_filesystem_proto_enumTypes[0].Descriptor()
}

func (UploadType) Type() protoreflect.EnumType {
	return &file_filesystem_v1_filesystem_proto_enumTypes[0]
}

func (x UploadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadType.Descriptor instead.
func (UploadType) EnumDescriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{0}
}

type CreateStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateStorageRequest) Reset() {
	*x = CreateStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageRequest) ProtoMessage() {}

func (x *CreateStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageRequest.ProtoReflect.Descriptor instead.
func (*CreateStorageRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{0}
}

type CreateStorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateStorageReply) Reset() {
	*x = CreateStorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStorageReply) ProtoMessage() {}

func (x *CreateStorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStorageReply.ProtoReflect.Descriptor instead.
func (*CreateStorageReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{1}
}

type UpdateStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStorageRequest) Reset() {
	*x = UpdateStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStorageRequest) ProtoMessage() {}

func (x *UpdateStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStorageRequest.ProtoReflect.Descriptor instead.
func (*UpdateStorageRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{2}
}

type UpdateStorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStorageReply) Reset() {
	*x = UpdateStorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStorageReply) ProtoMessage() {}

func (x *UpdateStorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStorageReply.ProtoReflect.Descriptor instead.
func (*UpdateStorageReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{3}
}

type DeleteStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStorageRequest) Reset() {
	*x = DeleteStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageRequest) ProtoMessage() {}

func (x *DeleteStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageRequest.ProtoReflect.Descriptor instead.
func (*DeleteStorageRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{4}
}

type DeleteStorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStorageReply) Reset() {
	*x = DeleteStorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageReply) ProtoMessage() {}

func (x *DeleteStorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageReply.ProtoReflect.Descriptor instead.
func (*DeleteStorageReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{5}
}

type GetStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStorageRequest) Reset() {
	*x = GetStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageRequest) ProtoMessage() {}

func (x *GetStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageRequest.ProtoReflect.Descriptor instead.
func (*GetStorageRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{6}
}

type GetStorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStorageReply) Reset() {
	*x = GetStorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageReply) ProtoMessage() {}

func (x *GetStorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageReply.ProtoReflect.Descriptor instead.
func (*GetStorageReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{7}
}

type ListStorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStorageRequest) Reset() {
	*x = ListStorageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStorageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageRequest) ProtoMessage() {}

func (x *ListStorageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageRequest.ProtoReflect.Descriptor instead.
func (*ListStorageRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{8}
}

type ListStorageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStorageReply) Reset() {
	*x = ListStorageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStorageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageReply) ProtoMessage() {}

func (x *ListStorageReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageReply.ProtoReflect.Descriptor instead.
func (*ListStorageReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{9}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Body []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{10}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type UploadType `protobuf:"varint,1,opt,name=type,proto3,enum=api.storage.v1.UploadType" json:"type,omitempty"`
	File *File      `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{11}
}

func (x *UploadRequest) GetType() UploadType {
	if x != nil {
		return x.Type
	}
	return UploadType_UT_File
}

func (x *UploadRequest) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadReply) Reset() {
	*x = UploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filesystem_v1_filesystem_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReply) ProtoMessage() {}

func (x *UploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_filesystem_v1_filesystem_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReply.ProtoReflect.Descriptor instead.
func (*UploadReply) Descriptor() ([]byte, []int) {
	return file_filesystem_v1_filesystem_proto_rawDescGZIP(), []int{12}
}

func (x *UploadReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_filesystem_v1_filesystem_proto protoreflect.FileDescriptor

var file_filesystem_v1_filesystem_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x42, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x69, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x1f, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x2a, 0x27, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x54, 0x5f, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x54, 0x5f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x01, 0x32, 0xa1, 0x04, 0x0a,
	0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x59, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x59, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x32, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x62, 0x65, 0x65, 0x74, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filesystem_v1_filesystem_proto_rawDescOnce sync.Once
	file_filesystem_v1_filesystem_proto_rawDescData = file_filesystem_v1_filesystem_proto_rawDesc
)

func file_filesystem_v1_filesystem_proto_rawDescGZIP() []byte {
	file_filesystem_v1_filesystem_proto_rawDescOnce.Do(func() {
		file_filesystem_v1_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_filesystem_v1_filesystem_proto_rawDescData)
	})
	return file_filesystem_v1_filesystem_proto_rawDescData
}

var file_filesystem_v1_filesystem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_filesystem_v1_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_filesystem_v1_filesystem_proto_goTypes = []interface{}{
	(UploadType)(0),              // 0: api.storage.v1.UploadType
	(*CreateStorageRequest)(nil), // 1: api.storage.v1.CreateStorageRequest
	(*CreateStorageReply)(nil),   // 2: api.storage.v1.CreateStorageReply
	(*UpdateStorageRequest)(nil), // 3: api.storage.v1.UpdateStorageRequest
	(*UpdateStorageReply)(nil),   // 4: api.storage.v1.UpdateStorageReply
	(*DeleteStorageRequest)(nil), // 5: api.storage.v1.DeleteStorageRequest
	(*DeleteStorageReply)(nil),   // 6: api.storage.v1.DeleteStorageReply
	(*GetStorageRequest)(nil),    // 7: api.storage.v1.GetStorageRequest
	(*GetStorageReply)(nil),      // 8: api.storage.v1.GetStorageReply
	(*ListStorageRequest)(nil),   // 9: api.storage.v1.ListStorageRequest
	(*ListStorageReply)(nil),     // 10: api.storage.v1.ListStorageReply
	(*File)(nil),                 // 11: api.storage.v1.File
	(*UploadRequest)(nil),        // 12: api.storage.v1.UploadRequest
	(*UploadReply)(nil),          // 13: api.storage.v1.UploadReply
}
var file_filesystem_v1_filesystem_proto_depIdxs = []int32{
	0,  // 0: api.storage.v1.UploadRequest.type:type_name -> api.storage.v1.UploadType
	11, // 1: api.storage.v1.UploadRequest.file:type_name -> api.storage.v1.File
	1,  // 2: api.storage.v1.Filesystem.CreateStorage:input_type -> api.storage.v1.CreateStorageRequest
	3,  // 3: api.storage.v1.Filesystem.UpdateStorage:input_type -> api.storage.v1.UpdateStorageRequest
	5,  // 4: api.storage.v1.Filesystem.DeleteStorage:input_type -> api.storage.v1.DeleteStorageRequest
	7,  // 5: api.storage.v1.Filesystem.GetStorage:input_type -> api.storage.v1.GetStorageRequest
	9,  // 6: api.storage.v1.Filesystem.ListStorage:input_type -> api.storage.v1.ListStorageRequest
	12, // 7: api.storage.v1.Filesystem.Upload:input_type -> api.storage.v1.UploadRequest
	2,  // 8: api.storage.v1.Filesystem.CreateStorage:output_type -> api.storage.v1.CreateStorageReply
	4,  // 9: api.storage.v1.Filesystem.UpdateStorage:output_type -> api.storage.v1.UpdateStorageReply
	6,  // 10: api.storage.v1.Filesystem.DeleteStorage:output_type -> api.storage.v1.DeleteStorageReply
	8,  // 11: api.storage.v1.Filesystem.GetStorage:output_type -> api.storage.v1.GetStorageReply
	10, // 12: api.storage.v1.Filesystem.ListStorage:output_type -> api.storage.v1.ListStorageReply
	13, // 13: api.storage.v1.Filesystem.Upload:output_type -> api.storage.v1.UploadReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_filesystem_v1_filesystem_proto_init() }
func file_filesystem_v1_filesystem_proto_init() {
	if File_filesystem_v1_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filesystem_v1_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStorageRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStorageReply); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStorageRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStorageReply); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStorageReply); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStorageReply); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStorageRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStorageReply); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_filesystem_v1_filesystem_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReply); i {
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
			RawDescriptor: file_filesystem_v1_filesystem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filesystem_v1_filesystem_proto_goTypes,
		DependencyIndexes: file_filesystem_v1_filesystem_proto_depIdxs,
		EnumInfos:         file_filesystem_v1_filesystem_proto_enumTypes,
		MessageInfos:      file_filesystem_v1_filesystem_proto_msgTypes,
	}.Build()
	File_filesystem_v1_filesystem_proto = out.File
	file_filesystem_v1_filesystem_proto_rawDesc = nil
	file_filesystem_v1_filesystem_proto_goTypes = nil
	file_filesystem_v1_filesystem_proto_depIdxs = nil
}
