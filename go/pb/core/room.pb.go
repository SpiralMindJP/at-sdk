//
// Room はルームに関する情報を操作します。

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: room.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
// ルームを表します。
type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`             // チームID。
	RoomId     int64  `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`             // ルームID。
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                                // ルーム名。
	DeviceId   int64  `protobuf:"varint,11,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`      // デバイスID。
	DeviceName string `protobuf:"bytes,12,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"` // デバイス名。
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *Room) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *Room) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

//
// ルームの作成要求を表します。
type RoomCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                    // ルーム名。
}

func (x *RoomCreateRequest) Reset() {
	*x = RoomCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateRequest) ProtoMessage() {}

func (x *RoomCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateRequest.ProtoReflect.Descriptor instead.
func (*RoomCreateRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *RoomCreateRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *RoomCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// ルームの更新要求を表します。
type RoomUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	RoomId int64  `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // ルームID。
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                    // ルーム名。
}

func (x *RoomUpdateRequest) Reset() {
	*x = RoomUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomUpdateRequest) ProtoMessage() {}

func (x *RoomUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoomUpdateRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *RoomUpdateRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *RoomUpdateRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RoomUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// ルームへのデバイス設定要求を表します。
type RoomDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`       // チームID。
	RoomId   int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`       // ルームID。
	DeviceId int64 `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"` // デバイスID。
	Force    bool  `protobuf:"varint,4,opt,name=force,proto3" json:"force,omitempty"`                       // 強制的にデバイスを設定するかどうか。True の場合、他のルームに既に設定されているデバイスであっても、登録します。(optional)
}

func (x *RoomDeviceRequest) Reset() {
	*x = RoomDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDeviceRequest) ProtoMessage() {}

func (x *RoomDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDeviceRequest.ProtoReflect.Descriptor instead.
func (*RoomDeviceRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *RoomDeviceRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *RoomDeviceRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RoomDeviceRequest) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *RoomDeviceRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

//
// ルームを取得する際のリクエストを表します。
type RoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	RoomId int64 `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // ルームID。
}

func (x *RoomRequest) Reset() {
	*x = RoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRequest) ProtoMessage() {}

func (x *RoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRequest.ProtoReflect.Descriptor instead.
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *RoomRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *RoomRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

//
// ルームのリストを取得する際のリクエストを表します。
type RoomListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
}

func (x *RoomListRequest) Reset() {
	*x = RoomListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListRequest) ProtoMessage() {}

func (x *RoomListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListRequest.ProtoReflect.Descriptor instead.
func (*RoomListRequest) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

func (x *RoomListRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

//
// ルームのリストを表します。
type Rooms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"` // ルームのリスト。
}

func (x *Rooms) Reset() {
	*x = Rooms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rooms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rooms) ProtoMessage() {}

func (x *Rooms) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rooms.ProtoReflect.Descriptor instead.
func (*Rooms) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{6}
}

func (x *Rooms) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x74,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x52, 0x6f, 0x6f,
	0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x3f,
	0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22,
	0x2a, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x05, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x32, 0xe2, 0x03, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x74, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x74,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x74, 0x5f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x74, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x5f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x43, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x69, 0x72, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x64, 0x4a,
	0x50, 0x2f, 0x61, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_room_proto_goTypes = []interface{}{
	(*Room)(nil),              // 0: at_core_service.Room
	(*RoomCreateRequest)(nil), // 1: at_core_service.RoomCreateRequest
	(*RoomUpdateRequest)(nil), // 2: at_core_service.RoomUpdateRequest
	(*RoomDeviceRequest)(nil), // 3: at_core_service.RoomDeviceRequest
	(*RoomRequest)(nil),       // 4: at_core_service.RoomRequest
	(*RoomListRequest)(nil),   // 5: at_core_service.RoomListRequest
	(*Rooms)(nil),             // 6: at_core_service.Rooms
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_room_proto_depIdxs = []int32{
	0, // 0: at_core_service.Rooms.rooms:type_name -> at_core_service.Room
	5, // 1: at_core_service.RoomService.List:input_type -> at_core_service.RoomListRequest
	4, // 2: at_core_service.RoomService.Get:input_type -> at_core_service.RoomRequest
	1, // 3: at_core_service.RoomService.Create:input_type -> at_core_service.RoomCreateRequest
	2, // 4: at_core_service.RoomService.Update:input_type -> at_core_service.RoomUpdateRequest
	3, // 5: at_core_service.RoomService.SetDevice:input_type -> at_core_service.RoomDeviceRequest
	4, // 6: at_core_service.RoomService.DeleteDevice:input_type -> at_core_service.RoomRequest
	4, // 7: at_core_service.RoomService.Delete:input_type -> at_core_service.RoomRequest
	6, // 8: at_core_service.RoomService.List:output_type -> at_core_service.Rooms
	0, // 9: at_core_service.RoomService.Get:output_type -> at_core_service.Room
	0, // 10: at_core_service.RoomService.Create:output_type -> at_core_service.Room
	0, // 11: at_core_service.RoomService.Update:output_type -> at_core_service.Room
	0, // 12: at_core_service.RoomService.SetDevice:output_type -> at_core_service.Room
	0, // 13: at_core_service.RoomService.DeleteDevice:output_type -> at_core_service.Room
	7, // 14: at_core_service.RoomService.Delete:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreateRequest); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomUpdateRequest); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomDeviceRequest); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRequest); i {
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
		file_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListRequest); i {
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
		file_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rooms); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
