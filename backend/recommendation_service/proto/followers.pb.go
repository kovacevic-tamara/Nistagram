// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
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

	UserId                string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FollowerId            string `protobuf:"bytes,2,opt,name=FollowerId,proto3" json:"FollowerId,omitempty"`
	IsMuted               bool   `protobuf:"varint,3,opt,name=IsMuted,proto3" json:"IsMuted,omitempty"`
	IsCloseFriends        bool   `protobuf:"varint,4,opt,name=IsCloseFriends,proto3" json:"IsCloseFriends,omitempty"`
	IsApprovedRequest     bool   `protobuf:"varint,5,opt,name=IsApprovedRequest,proto3" json:"IsApprovedRequest,omitempty"`
	IsNotificationEnabled bool   `protobuf:"varint,6,opt,name=IsNotificationEnabled,proto3" json:"IsNotificationEnabled,omitempty"`
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

type CreateFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follower *Follower `protobuf:"bytes,1,opt,name=follower,proto3" json:"follower,omitempty"`
}

func (x *CreateFollowerRequest) Reset() {
	*x = CreateFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFollowerRequest) ProtoMessage() {}

func (x *CreateFollowerRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateFollowerRequest.ProtoReflect.Descriptor instead.
func (*CreateFollowerRequest) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFollowerRequest) GetFollower() *Follower {
	if x != nil {
		return x.Follower
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_followers_proto_rawDescGZIP(), []int{2}
}

var File_followers_proto protoreflect.FileDescriptor

var file_followers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x49,
	0x73, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73,
	0x4d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49,
	0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x49, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x49, 0x73, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x49,
	0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x49, 0x73, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x44, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x75, 0x0a, 0x09, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x68, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x08, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x77,
	0x73, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_followers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_followers_proto_goTypes = []interface{}{
	(*Follower)(nil),              // 0: proto.Follower
	(*CreateFollowerRequest)(nil), // 1: proto.CreateFollowerRequest
	(*EmptyResponse)(nil),         // 2: proto.EmptyResponse
}
var file_followers_proto_depIdxs = []int32{
	0, // 0: proto.CreateFollowerRequest.follower:type_name -> proto.Follower
	1, // 1: proto.Followers.CreateUserConnection:input_type -> proto.CreateFollowerRequest
	2, // 2: proto.Followers.CreateUserConnection:output_type -> proto.EmptyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_followers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			NumMessages:   3,
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