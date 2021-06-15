// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: followers.proto

package proto

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

type Follower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FollowerId            string `protobuf:"bytes,2,opt,name=followerId,proto3" json:"followerId,omitempty"`
	IsMuted               bool   `protobuf:"varint,3,opt,name=isMuted,proto3" json:"isMuted,omitempty"`
	IsCloseFriends        bool   `protobuf:"varint,4,opt,name=isCloseFriends,proto3" json:"isCloseFriends,omitempty"`
	IsApprovedRequest     bool   `protobuf:"varint,5,opt,name=isApprovedRequest,proto3" json:"isApprovedRequest,omitempty"`
	IsNotificationEnabled bool   `protobuf:"varint,6,opt,name=isNotificationEnabled,proto3" json:"isNotificationEnabled,omitempty"`
}

func (x *Follower) Reset() {
	*x = Follower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follower) ProtoMessage() {}

func (x *Follower) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follower.ProtoReflect.Descriptor instead.
func (*Follower) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{0}
}

func (x *Follower) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Follower) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *Follower) GetIsMuted() bool {
	if x != nil {
		return x.IsMuted
	}
	return false
}

func (x *Follower) GetIsCloseFriends() bool {
	if x != nil {
		return x.IsCloseFriends
	}
	return false
}

func (x *Follower) GetIsApprovedRequest() bool {
	if x != nil {
		return x.IsApprovedRequest
	}
	return false
}

func (x *Follower) GetIsNotificationEnabled() bool {
	if x != nil {
		return x.IsNotificationEnabled
	}
	return false
}

type UserFollowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *UserFollowers) Reset() {
	*x = UserFollowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFollowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFollowers) ProtoMessage() {}

func (x *UserFollowers) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFollowers.ProtoReflect.Descriptor instead.
func (*UserFollowers) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{1}
}

func (x *UserFollowers) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateUserRequestFollowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserFollowers `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRequestFollowers) Reset() {
	*x = CreateUserRequestFollowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequestFollowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequestFollowers) ProtoMessage() {}

func (x *CreateUserRequestFollowers) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequestFollowers.ProtoReflect.Descriptor instead.
func (*CreateUserRequestFollowers) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequestFollowers) GetUser() *UserFollowers {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserFollowers `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResponse) GetUsers() []*UserFollowers {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateFollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follower *Follower `protobuf:"bytes,1,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *CreateFollowerResponse) Reset() {
	*x = CreateFollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowerResponse) ProtoMessage() {}

func (x *CreateFollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowerResponse.ProtoReflect.Descriptor instead.
func (*CreateFollowerResponse) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFollowerResponse) GetFollower() *Follower {
	if x != nil {
		return x.Follower
	}
	return nil
}

type CreateFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follower *Follower `protobuf:"bytes,1,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *CreateFollowerRequest) Reset() {
	*x = CreateFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowerRequest) ProtoMessage() {}

func (x *CreateFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFollowerRequest.ProtoReflect.Descriptor instead.
func (*CreateFollowerRequest) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{5}
}

func (x *CreateFollowerRequest) GetFollower() *Follower {
	if x != nil {
		return x.Follower
	}
	return nil
}

type RequestIdFollowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestIdFollowers) Reset() {
	*x = RequestIdFollowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestIdFollowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIdFollowers) ProtoMessage() {}

func (x *RequestIdFollowers) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestIdFollowers.ProtoReflect.Descriptor instead.
func (*RequestIdFollowers) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{6}
}

func (x *RequestIdFollowers) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyResponseFollowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponseFollowers) Reset() {
	*x = EmptyResponseFollowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponseFollowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponseFollowers) ProtoMessage() {}

func (x *EmptyResponseFollowers) ProtoReflect() protoreflect.Message {
	mi := &file_followers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponseFollowers.ProtoReflect.Descriptor instead.
func (*EmptyResponseFollowers) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{7}
}

var File_followers_proto protoreflect.FileDescriptor

var file_followers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x4d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69,
	0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x69, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x27, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x1a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x45, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x32, 0xef, 0x09, 0x0a, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12,
	0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x89, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x3a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x8d, 0x01, 0x0a,
	0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x3a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x78, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22,
	0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x89, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6a, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x3a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x60, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x77, 0x73, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_followers_proto_rawDescOnce sync.Once
	file_followers_proto_rawDescData = file_followers_proto_rawDesc
)

func file_followers_proto_rawDescGZIP() []byte {
	file_followers_proto_rawDescOnce.Do(func() {
		file_followers_proto_rawDescData = protoimpl.X.CompressGZIP(file_followers_proto_rawDescData)
	})
	return file_followers_proto_rawDescData
}

var file_followers_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_followers_proto_goTypes = []interface{}{
	(*Follower)(nil),                   // 0: proto.Follower
	(*UserFollowers)(nil),              // 1: proto.UserFollowers
	(*CreateUserRequestFollowers)(nil), // 2: proto.CreateUserRequestFollowers
	(*CreateUserResponse)(nil),         // 3: proto.CreateUserResponse
	(*CreateFollowerResponse)(nil),     // 4: proto.CreateFollowerResponse
	(*CreateFollowerRequest)(nil),      // 5: proto.CreateFollowerRequest
	(*RequestIdFollowers)(nil),         // 6: proto.RequestIdFollowers
	(*EmptyResponseFollowers)(nil),     // 7: proto.EmptyResponseFollowers
}
var file_followers_proto_depIdxs = []int32{
	1,  // 0: proto.CreateUserRequestFollowers.user:type_name -> proto.UserFollowers
	1,  // 1: proto.CreateUserResponse.users:type_name -> proto.UserFollowers
	0,  // 2: proto.CreateFollowerResponse.follower:type_name -> proto.Follower
	0,  // 3: proto.CreateFollowerRequest.follower:type_name -> proto.Follower
	5,  // 4: proto.Followers.CreateUserConnection:input_type -> proto.CreateFollowerRequest
	2,  // 5: proto.Followers.CreateUser:input_type -> proto.CreateUserRequestFollowers
	5,  // 6: proto.Followers.DeleteDirectedConnection:input_type -> proto.CreateFollowerRequest
	5,  // 7: proto.Followers.DeleteBiDirectedConnection:input_type -> proto.CreateFollowerRequest
	2,  // 8: proto.Followers.GetAllFollowers:input_type -> proto.CreateUserRequestFollowers
	2,  // 9: proto.Followers.GetAllFollowing:input_type -> proto.CreateUserRequestFollowers
	2,  // 10: proto.Followers.GetAllFollowingsForHomepage:input_type -> proto.CreateUserRequestFollowers
	6,  // 11: proto.Followers.GetCloseFriends:input_type -> proto.RequestIdFollowers
	5,  // 12: proto.Followers.UpdateUserConnection:input_type -> proto.CreateFollowerRequest
	0,  // 13: proto.Followers.GetFollowersConnection:input_type -> proto.Follower
	7,  // 14: proto.Followers.CreateUserConnection:output_type -> proto.EmptyResponseFollowers
	7,  // 15: proto.Followers.CreateUser:output_type -> proto.EmptyResponseFollowers
	7,  // 16: proto.Followers.DeleteDirectedConnection:output_type -> proto.EmptyResponseFollowers
	7,  // 17: proto.Followers.DeleteBiDirectedConnection:output_type -> proto.EmptyResponseFollowers
	3,  // 18: proto.Followers.GetAllFollowers:output_type -> proto.CreateUserResponse
	3,  // 19: proto.Followers.GetAllFollowing:output_type -> proto.CreateUserResponse
	3,  // 20: proto.Followers.GetAllFollowingsForHomepage:output_type -> proto.CreateUserResponse
	3,  // 21: proto.Followers.GetCloseFriends:output_type -> proto.CreateUserResponse
	4,  // 22: proto.Followers.UpdateUserConnection:output_type -> proto.CreateFollowerResponse
	0,  // 23: proto.Followers.GetFollowersConnection:output_type -> proto.Follower
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_followers_proto_init() }
func file_followers_proto_init() {
	if File_followers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_followers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follower); i {
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
		file_followers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFollowers); i {
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
		file_followers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequestFollowers); i {
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
		file_followers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_followers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowerResponse); i {
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
		file_followers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFollowerRequest); i {
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
		file_followers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestIdFollowers); i {
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
		file_followers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponseFollowers); i {
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
			RawDescriptor: file_followers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_followers_proto_goTypes,
		DependencyIndexes: file_followers_proto_depIdxs,
		MessageInfos:      file_followers_proto_msgTypes,
	}.Build()
	File_followers_proto = out.File
	file_followers_proto_rawDesc = nil
	file_followers_proto_goTypes = nil
	file_followers_proto_depIdxs = nil
}
