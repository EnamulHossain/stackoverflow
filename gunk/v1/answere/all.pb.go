// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: stackoverflow/gunk/v1/answere/all.proto

package answerepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Answere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId     int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	QuestionId int32  `protobuf:"varint,3,opt,name=QuestionId,proto3" json:"QuestionId,omitempty"`
	Answere    string `protobuf:"bytes,4,opt,name=Answere,proto3" json:"Answere,omitempty"`
	IsCorrect  bool   `protobuf:"varint,5,opt,name=IsCorrect,proto3" json:"IsCorrect,omitempty"`
}

func (x *Answere) Reset() {
	*x = Answere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answere) ProtoMessage() {}

func (x *Answere) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answere.ProtoReflect.Descriptor instead.
func (*Answere) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{0}
}

func (x *Answere) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Answere) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Answere) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answere) GetAnswere() string {
	if x != nil {
		return x.Answere
	}
	return ""
}

func (x *Answere) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

type AnswereCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	QuestionId int32  `protobuf:"varint,2,opt,name=QuestionId,proto3" json:"QuestionId,omitempty"`
	Answere    string `protobuf:"bytes,3,opt,name=Answere,proto3" json:"Answere,omitempty"`
}

func (x *AnswereCreateRequest) Reset() {
	*x = AnswereCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereCreateRequest) ProtoMessage() {}

func (x *AnswereCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereCreateRequest.ProtoReflect.Descriptor instead.
func (*AnswereCreateRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{1}
}

func (x *AnswereCreateRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AnswereCreateRequest) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *AnswereCreateRequest) GetAnswere() string {
	if x != nil {
		return x.Answere
	}
	return ""
}

type AnswereCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answere *Answere `protobuf:"bytes,1,opt,name=Answere,proto3" json:"Answere,omitempty"`
}

func (x *AnswereCreateResponse) Reset() {
	*x = AnswereCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereCreateResponse) ProtoMessage() {}

func (x *AnswereCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereCreateResponse.ProtoReflect.Descriptor instead.
func (*AnswereCreateResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{2}
}

func (x *AnswereCreateResponse) GetAnswere() *Answere {
	if x != nil {
		return x.Answere
	}
	return nil
}

type AnswereListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnswereListRequest) Reset() {
	*x = AnswereListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereListRequest) ProtoMessage() {}

func (x *AnswereListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereListRequest.ProtoReflect.Descriptor instead.
func (*AnswereListRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{3}
}

type AnswereListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answere []*Answere `protobuf:"bytes,1,rep,name=Answere,proto3" json:"Answere,omitempty"`
}

func (x *AnswereListResponse) Reset() {
	*x = AnswereListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereListResponse) ProtoMessage() {}

func (x *AnswereListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereListResponse.ProtoReflect.Descriptor instead.
func (*AnswereListResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{4}
}

func (x *AnswereListResponse) GetAnswere() []*Answere {
	if x != nil {
		return x.Answere
	}
	return nil
}

type AnswereDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AnswereDeleteRequest) Reset() {
	*x = AnswereDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereDeleteRequest) ProtoMessage() {}

func (x *AnswereDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereDeleteRequest.ProtoReflect.Descriptor instead.
func (*AnswereDeleteRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{5}
}

func (x *AnswereDeleteRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type AnswereDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnswereDeleteResponse) Reset() {
	*x = AnswereDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereDeleteResponse) ProtoMessage() {}

func (x *AnswereDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereDeleteResponse.ProtoReflect.Descriptor instead.
func (*AnswereDeleteResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{6}
}

type AnswereEditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AnswereEditRequest) Reset() {
	*x = AnswereEditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereEditRequest) ProtoMessage() {}

func (x *AnswereEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereEditRequest.ProtoReflect.Descriptor instead.
func (*AnswereEditRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{7}
}

func (x *AnswereEditRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type AnswereEditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answere *Answere `protobuf:"bytes,1,opt,name=Answere,proto3" json:"Answere,omitempty"`
}

func (x *AnswereEditResponse) Reset() {
	*x = AnswereEditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereEditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereEditResponse) ProtoMessage() {}

func (x *AnswereEditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereEditResponse.ProtoReflect.Descriptor instead.
func (*AnswereEditResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{8}
}

func (x *AnswereEditResponse) GetAnswere() *Answere {
	if x != nil {
		return x.Answere
	}
	return nil
}

type AnswereUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Answere string `protobuf:"bytes,2,opt,name=Answere,proto3" json:"Answere,omitempty"`
}

func (x *AnswereUpdateRequest) Reset() {
	*x = AnswereUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereUpdateRequest) ProtoMessage() {}

func (x *AnswereUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereUpdateRequest.ProtoReflect.Descriptor instead.
func (*AnswereUpdateRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{9}
}

func (x *AnswereUpdateRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnswereUpdateRequest) GetAnswere() string {
	if x != nil {
		return x.Answere
	}
	return ""
}

type AnswereUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnswereUpdateResponse) Reset() {
	*x = AnswereUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswereUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswereUpdateResponse) ProtoMessage() {}

func (x *AnswereUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswereUpdateResponse.ProtoReflect.Descriptor instead.
func (*AnswereUpdateResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{10}
}

type CorrectAnswereRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IsCorrect bool  `protobuf:"varint,2,opt,name=IsCorrect,proto3" json:"IsCorrect,omitempty"`
}

func (x *CorrectAnswereRequest) Reset() {
	*x = CorrectAnswereRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrectAnswereRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrectAnswereRequest) ProtoMessage() {}

func (x *CorrectAnswereRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrectAnswereRequest.ProtoReflect.Descriptor instead.
func (*CorrectAnswereRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{11}
}

func (x *CorrectAnswereRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CorrectAnswereRequest) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

type CorrectAnswereResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CorrectAnswereResponse) Reset() {
	*x = CorrectAnswereResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorrectAnswereResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorrectAnswereResponse) ProtoMessage() {}

func (x *CorrectAnswereResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorrectAnswereResponse.ProtoReflect.Descriptor instead.
func (*CorrectAnswereResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP(), []int{12}
}

var File_stackoverflow_gunk_v1_answere_all_proto protoreflect.FileDescriptor

var file_stackoverflow_gunk_v1_answere_all_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x2f,
	0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x70, 0x62, 0x22, 0xa1, 0x01, 0x0a, 0x07, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x12, 0x1d, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x77, 0x0a, 0x14, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x50, 0x0a, 0x15, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x22, 0x1c, 0x0a, 0x12, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x4e, 0x0a, 0x13, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x36, 0x0a, 0x14, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x1f, 0x0a, 0x15, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x34, 0x0a, 0x12, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x4e, 0x0a, 0x13, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x53, 0x0a, 0x14, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x12, 0x1b, 0x0a, 0x07, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x1f, 0x0a, 0x15, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x56, 0x0a, 0x15, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1d, 0x0a, 0x09, 0x49, 0x73, 0x43, 0x6f, 0x72,
	0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x20,
	0x0a, 0x16, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x32, 0xcc, 0x04, 0x0a, 0x0e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62,
	0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70,
	0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x5e, 0x0a,
	0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x58, 0x0a,
	0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x45,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00,
	0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x5e, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00,
	0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x61, 0x0a, 0x0e, 0x43, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42,
	0x42, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x27, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6f, 0x76, 0x65, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x65, 0x3b, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x70, 0x62, 0x80, 0x01,
	0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01,
	0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stackoverflow_gunk_v1_answere_all_proto_rawDescOnce sync.Once
	file_stackoverflow_gunk_v1_answere_all_proto_rawDescData = file_stackoverflow_gunk_v1_answere_all_proto_rawDesc
)

func file_stackoverflow_gunk_v1_answere_all_proto_rawDescGZIP() []byte {
	file_stackoverflow_gunk_v1_answere_all_proto_rawDescOnce.Do(func() {
		file_stackoverflow_gunk_v1_answere_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_stackoverflow_gunk_v1_answere_all_proto_rawDescData)
	})
	return file_stackoverflow_gunk_v1_answere_all_proto_rawDescData
}

var (
	file_stackoverflow_gunk_v1_answere_all_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
	file_stackoverflow_gunk_v1_answere_all_proto_goTypes  = []interface{}{
		(*Answere)(nil),                // 0: answerepb.Answere
		(*AnswereCreateRequest)(nil),   // 1: answerepb.AnswereCreateRequest
		(*AnswereCreateResponse)(nil),  // 2: answerepb.AnswereCreateResponse
		(*AnswereListRequest)(nil),     // 3: answerepb.AnswereListRequest
		(*AnswereListResponse)(nil),    // 4: answerepb.AnswereListResponse
		(*AnswereDeleteRequest)(nil),   // 5: answerepb.AnswereDeleteRequest
		(*AnswereDeleteResponse)(nil),  // 6: answerepb.AnswereDeleteResponse
		(*AnswereEditRequest)(nil),     // 7: answerepb.AnswereEditRequest
		(*AnswereEditResponse)(nil),    // 8: answerepb.AnswereEditResponse
		(*AnswereUpdateRequest)(nil),   // 9: answerepb.AnswereUpdateRequest
		(*AnswereUpdateResponse)(nil),  // 10: answerepb.AnswereUpdateResponse
		(*CorrectAnswereRequest)(nil),  // 11: answerepb.CorrectAnswereRequest
		(*CorrectAnswereResponse)(nil), // 12: answerepb.CorrectAnswereResponse
	}
)

var file_stackoverflow_gunk_v1_answere_all_proto_depIdxs = []int32{
	0,  // 0: answerepb.AnswereCreateResponse.Answere:type_name -> answerepb.Answere
	0,  // 1: answerepb.AnswereListResponse.Answere:type_name -> answerepb.Answere
	0,  // 2: answerepb.AnswereEditResponse.Answere:type_name -> answerepb.Answere
	1,  // 3: answerepb.AnswereService.AnswereCreate:input_type -> answerepb.AnswereCreateRequest
	3,  // 4: answerepb.AnswereService.AnswereList:input_type -> answerepb.AnswereListRequest
	5,  // 5: answerepb.AnswereService.AnswereDelete:input_type -> answerepb.AnswereDeleteRequest
	7,  // 6: answerepb.AnswereService.AnswereEdit:input_type -> answerepb.AnswereEditRequest
	9,  // 7: answerepb.AnswereService.AnswereUpdate:input_type -> answerepb.AnswereUpdateRequest
	11, // 8: answerepb.AnswereService.CorrectAnswere:input_type -> answerepb.CorrectAnswereRequest
	2,  // 9: answerepb.AnswereService.AnswereCreate:output_type -> answerepb.AnswereCreateResponse
	4,  // 10: answerepb.AnswereService.AnswereList:output_type -> answerepb.AnswereListResponse
	6,  // 11: answerepb.AnswereService.AnswereDelete:output_type -> answerepb.AnswereDeleteResponse
	8,  // 12: answerepb.AnswereService.AnswereEdit:output_type -> answerepb.AnswereEditResponse
	10, // 13: answerepb.AnswereService.AnswereUpdate:output_type -> answerepb.AnswereUpdateResponse
	12, // 14: answerepb.AnswereService.CorrectAnswere:output_type -> answerepb.CorrectAnswereResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_stackoverflow_gunk_v1_answere_all_proto_init() }
func file_stackoverflow_gunk_v1_answere_all_proto_init() {
	if File_stackoverflow_gunk_v1_answere_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answere); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereCreateRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereCreateResponse); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereListRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereListResponse); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereDeleteRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereDeleteResponse); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereEditRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereEditResponse); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereUpdateRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswereUpdateResponse); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorrectAnswereRequest); i {
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
		file_stackoverflow_gunk_v1_answere_all_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorrectAnswereResponse); i {
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
			RawDescriptor: file_stackoverflow_gunk_v1_answere_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stackoverflow_gunk_v1_answere_all_proto_goTypes,
		DependencyIndexes: file_stackoverflow_gunk_v1_answere_all_proto_depIdxs,
		MessageInfos:      file_stackoverflow_gunk_v1_answere_all_proto_msgTypes,
	}.Build()
	File_stackoverflow_gunk_v1_answere_all_proto = out.File
	file_stackoverflow_gunk_v1_answere_all_proto_rawDesc = nil
	file_stackoverflow_gunk_v1_answere_all_proto_goTypes = nil
	file_stackoverflow_gunk_v1_answere_all_proto_depIdxs = nil
}
