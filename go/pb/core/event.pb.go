// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: event.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
// ルームイベントの種類を表します。
type RoomEventType int32

const (
	RoomEventType_ROOM_EVENT_ROOM_CREATED   RoomEventType = 0  // ルームが作成された。
	RoomEventType_ROOM_EVENT_ROOM_DELETED   RoomEventType = 1  // ルームが削除された。
	RoomEventType_ROOM_EVENT_DEVICE_JOINED  RoomEventType = 10 // ルームにデバイスが入室した。
	RoomEventType_ROOM_EVENT_DEVICE_LEAVED  RoomEventType = 11 // ルームからデバイスが退室した。
	RoomEventType_ROOM_EVENT_DEVICE_DELETED RoomEventType = 12 // デバイスが削除された。
	RoomEventType_ROOM_EVENT_DEVICE_OFFLINE RoomEventType = 20 // デバイスがオフラインになった。
	RoomEventType_ROOM_EVENT_DEVICE_ONLINE  RoomEventType = 21 // デバイスがオンラインになった。
)

// Enum value maps for RoomEventType.
var (
	RoomEventType_name = map[int32]string{
		0:  "ROOM_EVENT_ROOM_CREATED",
		1:  "ROOM_EVENT_ROOM_DELETED",
		10: "ROOM_EVENT_DEVICE_JOINED",
		11: "ROOM_EVENT_DEVICE_LEAVED",
		12: "ROOM_EVENT_DEVICE_DELETED",
		20: "ROOM_EVENT_DEVICE_OFFLINE",
		21: "ROOM_EVENT_DEVICE_ONLINE",
	}
	RoomEventType_value = map[string]int32{
		"ROOM_EVENT_ROOM_CREATED":   0,
		"ROOM_EVENT_ROOM_DELETED":   1,
		"ROOM_EVENT_DEVICE_JOINED":  10,
		"ROOM_EVENT_DEVICE_LEAVED":  11,
		"ROOM_EVENT_DEVICE_DELETED": 12,
		"ROOM_EVENT_DEVICE_OFFLINE": 20,
		"ROOM_EVENT_DEVICE_ONLINE":  21,
	}
)

func (x RoomEventType) Enum() *RoomEventType {
	p := new(RoomEventType)
	*p = x
	return p
}

func (x RoomEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[0].Descriptor()
}

func (RoomEventType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[0]
}

func (x RoomEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomEventType.Descriptor instead.
func (RoomEventType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

//
// オペレーションイベントの種類を表します。
type OperationEventType int32

const (
	OperationEventType_OPERATION_EVENT_CONTENT_STOP   OperationEventType = 0  // コンテンツの再生を停止する。
	OperationEventType_OPERATION_EVENT_CONTENT_PLAY   OperationEventType = 1  // コンテンツを再生する。
	OperationEventType_OPERATION_EVENT_ANIMATION_STOP OperationEventType = 10 // アニメーションの再生を停止する。
	OperationEventType_OPERATION_EVENT_ANIMATION_PLAY OperationEventType = 11 // アニメーションを再生する。
)

// Enum value maps for OperationEventType.
var (
	OperationEventType_name = map[int32]string{
		0:  "OPERATION_EVENT_CONTENT_STOP",
		1:  "OPERATION_EVENT_CONTENT_PLAY",
		10: "OPERATION_EVENT_ANIMATION_STOP",
		11: "OPERATION_EVENT_ANIMATION_PLAY",
	}
	OperationEventType_value = map[string]int32{
		"OPERATION_EVENT_CONTENT_STOP":   0,
		"OPERATION_EVENT_CONTENT_PLAY":   1,
		"OPERATION_EVENT_ANIMATION_STOP": 10,
		"OPERATION_EVENT_ANIMATION_PLAY": 11,
	}
)

func (x OperationEventType) Enum() *OperationEventType {
	p := new(OperationEventType)
	*p = x
	return p
}

func (x OperationEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[1].Descriptor()
}

func (OperationEventType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[1]
}

func (x OperationEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationEventType.Descriptor instead.
func (OperationEventType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

//
// ルームイベントを表します。
type RoomEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      RoomEventType `protobuf:"varint,1,opt,name=type,proto3,enum=at_core_service.RoomEventType" json:"type,omitempty"` // ルームイベントの種類。
	RoomState *RoomState    `protobuf:"bytes,2,opt,name=room_state,json=roomState,proto3" json:"room_state,omitempty"`          // ルームの状態。
}

func (x *RoomEvent) Reset() {
	*x = RoomEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomEvent) ProtoMessage() {}

func (x *RoomEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomEvent.ProtoReflect.Descriptor instead.
func (*RoomEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *RoomEvent) GetType() RoomEventType {
	if x != nil {
		return x.Type
	}
	return RoomEventType_ROOM_EVENT_ROOM_CREATED
}

func (x *RoomEvent) GetRoomState() *RoomState {
	if x != nil {
		return x.RoomState
	}
	return nil
}

//
// オペレーションイベントを表します。
type OperationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           OperationEventType `protobuf:"varint,1,opt,name=type,proto3,enum=at_core_service.OperationEventType" json:"type,omitempty"`    // オペレーションイベントの種類。
	ContentId      int64              `protobuf:"varint,10,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`                // コンテンツID。
	AnimationIndex int32              `protobuf:"varint,20,opt,name=animation_index,json=animationIndex,proto3" json:"animation_index,omitempty"` // アニメーションインデックス。
}

func (x *OperationEvent) Reset() {
	*x = OperationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationEvent) ProtoMessage() {}

func (x *OperationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationEvent.ProtoReflect.Descriptor instead.
func (*OperationEvent) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *OperationEvent) GetType() OperationEventType {
	if x != nil {
		return x.Type
	}
	return OperationEventType_OPERATION_EVENT_CONTENT_STOP
}

func (x *OperationEvent) GetContentId() int64 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *OperationEvent) GetAnimationIndex() int32 {
	if x != nil {
		return x.AnimationIndex
	}
	return 0
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61,
	0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x09,
	0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x74, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x6e,
	0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2a, 0xe1, 0x01, 0x0a,
	0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4f,
	0x4d, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x4f, 0x4d,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4a, 0x4f,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x56,
	0x45, 0x44, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x0c, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x14, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x15,
	0x2a, 0xa0, 0x01, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x4e, 0x49, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x0a, 0x12,
	0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x41, 0x4e, 0x49, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x10, 0x0b, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x70, 0x69, 0x72, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x64, 0x4a, 0x50, 0x2f, 0x61,
	0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_proto_goTypes = []interface{}{
	(RoomEventType)(0),      // 0: at_core_service.RoomEventType
	(OperationEventType)(0), // 1: at_core_service.OperationEventType
	(*RoomEvent)(nil),       // 2: at_core_service.RoomEvent
	(*OperationEvent)(nil),  // 3: at_core_service.OperationEvent
	(*RoomState)(nil),       // 4: at_core_service.RoomState
}
var file_event_proto_depIdxs = []int32{
	0, // 0: at_core_service.RoomEvent.type:type_name -> at_core_service.RoomEventType
	4, // 1: at_core_service.RoomEvent.room_state:type_name -> at_core_service.RoomState
	1, // 2: at_core_service.OperationEvent.type:type_name -> at_core_service.OperationEventType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomEvent); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationEvent); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		EnumInfos:         file_event_proto_enumTypes,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}