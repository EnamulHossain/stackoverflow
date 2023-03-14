// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: stackoverflow/gunk/v1/user/all.proto

package userpb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	IsActive  bool   `protobuf:"varint,6,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	IsAdmin   bool   `protobuf:"varint,7,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	IsActive  bool   `protobuf:"varint,6,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	IsAdmin   bool   `protobuf:"varint,7,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *RegisterRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type AdminRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	IsActive  bool   `protobuf:"varint,6,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	IsAdmin   bool   `protobuf:"varint,7,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
}

func (x *AdminRegisterRequest) Reset() {
	*x = AdminRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterRequest) ProtoMessage() {}

func (x *AdminRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterRequest.ProtoReflect.Descriptor instead.
func (*AdminRegisterRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{3}
}

func (x *AdminRegisterRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AdminRegisterRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AdminRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AdminRegisterRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *AdminRegisterRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type AdminRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *AdminRegisterResponse) Reset() {
	*x = AdminRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterResponse) ProtoMessage() {}

func (x *AdminRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterResponse.ProtoReflect.Descriptor instead.
func (*AdminRegisterResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRegisterResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{6}
}

func (x *LoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{7}
}

type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{8}
}

func (x *ListUserResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type EditUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *EditUserRequest) Reset() {
	*x = EditUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserRequest) ProtoMessage() {}

func (x *EditUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserRequest.ProtoReflect.Descriptor instead.
func (*EditUserRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{9}
}

func (x *EditUserRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type EditUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *EditUserResponse) Reset() {
	*x = EditUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditUserResponse) ProtoMessage() {}

func (x *EditUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditUserResponse.ProtoReflect.Descriptor instead.
func (*EditUserResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{10}
}

func (x *EditUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteUserRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stackoverflow_gunk_v1_user_all_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP(), []int{12}
}

var File_stackoverflow_gunk_v1_user_all_proto protoreflect.FileDescriptor

var file_stackoverflow_gunk_v1_user_all_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x22, 0xd7,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12,
	0x1d, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c,
	0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xe8, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x12, 0x1c, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1b,
	0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x22, 0x42, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xed, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12,
	0x1c, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x12, 0x1c, 0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x1b, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x47, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x52, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1c,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00,
	0x10, 0x00, 0x18, 0x00, 0x22, 0x3f, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x19, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x43, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x31, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x42, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x33, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x1c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32,
	0xe0, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x40, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90,
	0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x49, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x12, 0x49, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x4f, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88,
	0x02, 0x00, 0x42, 0x3c, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x21, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x6f,
	0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x80, 0x01, 0x00, 0x88, 0x01,
	0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stackoverflow_gunk_v1_user_all_proto_rawDescOnce sync.Once
	file_stackoverflow_gunk_v1_user_all_proto_rawDescData = file_stackoverflow_gunk_v1_user_all_proto_rawDesc
)

func file_stackoverflow_gunk_v1_user_all_proto_rawDescGZIP() []byte {
	file_stackoverflow_gunk_v1_user_all_proto_rawDescOnce.Do(func() {
		file_stackoverflow_gunk_v1_user_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_stackoverflow_gunk_v1_user_all_proto_rawDescData)
	})
	return file_stackoverflow_gunk_v1_user_all_proto_rawDescData
}

var (
	file_stackoverflow_gunk_v1_user_all_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
	file_stackoverflow_gunk_v1_user_all_proto_goTypes  = []interface{}{
		(*User)(nil),                  // 0: userpb.User
		(*RegisterRequest)(nil),       // 1: userpb.RegisterRequest
		(*RegisterResponse)(nil),      // 2: userpb.RegisterResponse
		(*AdminRegisterRequest)(nil),  // 3: userpb.AdminRegisterRequest
		(*AdminRegisterResponse)(nil), // 4: userpb.AdminRegisterResponse
		(*LoginRequest)(nil),          // 5: userpb.LoginRequest
		(*LoginResponse)(nil),         // 6: userpb.LoginResponse
		(*ListUserRequest)(nil),       // 7: userpb.ListUserRequest
		(*ListUserResponse)(nil),      // 8: userpb.ListUserResponse
		(*EditUserRequest)(nil),       // 9: userpb.EditUserRequest
		(*EditUserResponse)(nil),      // 10: userpb.EditUserResponse
		(*DeleteUserRequest)(nil),     // 11: userpb.DeleteUserRequest
		(*DeleteUserResponse)(nil),    // 12: userpb.DeleteUserResponse
	}
)

var file_stackoverflow_gunk_v1_user_all_proto_depIdxs = []int32{
	0,  // 0: userpb.RegisterResponse.User:type_name -> userpb.User
	0,  // 1: userpb.AdminRegisterResponse.User:type_name -> userpb.User
	0,  // 2: userpb.LoginResponse.User:type_name -> userpb.User
	0,  // 3: userpb.ListUserResponse.Users:type_name -> userpb.User
	0,  // 4: userpb.EditUserResponse.User:type_name -> userpb.User
	1,  // 5: userpb.UserService.Register:input_type -> userpb.RegisterRequest
	3,  // 6: userpb.UserService.AdminRegister:input_type -> userpb.AdminRegisterRequest
	5,  // 7: userpb.UserService.Login:input_type -> userpb.LoginRequest
	7,  // 8: userpb.UserService.ListUser:input_type -> userpb.ListUserRequest
	9,  // 9: userpb.UserService.EditUser:input_type -> userpb.EditUserRequest
	11, // 10: userpb.UserService.DeleteUser:input_type -> userpb.DeleteUserRequest
	2,  // 11: userpb.UserService.Register:output_type -> userpb.RegisterResponse
	4,  // 12: userpb.UserService.AdminRegister:output_type -> userpb.AdminRegisterResponse
	6,  // 13: userpb.UserService.Login:output_type -> userpb.LoginResponse
	8,  // 14: userpb.UserService.ListUser:output_type -> userpb.ListUserResponse
	10, // 15: userpb.UserService.EditUser:output_type -> userpb.EditUserResponse
	12, // 16: userpb.UserService.DeleteUser:output_type -> userpb.DeleteUserResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_stackoverflow_gunk_v1_user_all_proto_init() }
func file_stackoverflow_gunk_v1_user_all_proto_init() {
	if File_stackoverflow_gunk_v1_user_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegisterRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegisterResponse); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditUserResponse); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_stackoverflow_gunk_v1_user_all_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
			RawDescriptor: file_stackoverflow_gunk_v1_user_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stackoverflow_gunk_v1_user_all_proto_goTypes,
		DependencyIndexes: file_stackoverflow_gunk_v1_user_all_proto_depIdxs,
		MessageInfos:      file_stackoverflow_gunk_v1_user_all_proto_msgTypes,
	}.Build()
	File_stackoverflow_gunk_v1_user_all_proto = out.File
	file_stackoverflow_gunk_v1_user_all_proto_rawDesc = nil
	file_stackoverflow_gunk_v1_user_all_proto_goTypes = nil
	file_stackoverflow_gunk_v1_user_all_proto_depIdxs = nil
}
