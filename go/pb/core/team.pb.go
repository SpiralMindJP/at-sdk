//
// Team はチームに関する情報を操作します。

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: team.proto

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
// チームを表します。
type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                    // チーム名。
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// チームの作成要求を表します。
type TeamCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // チーム名。
}

func (x *TeamCreateRequest) Reset() {
	*x = TeamCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCreateRequest) ProtoMessage() {}

func (x *TeamCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCreateRequest.ProtoReflect.Descriptor instead.
func (*TeamCreateRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// チームの更新要求を表します。
type TeamUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                    // チーム名。
}

func (x *TeamUpdateRequest) Reset() {
	*x = TeamUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamUpdateRequest) ProtoMessage() {}

func (x *TeamUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamUpdateRequest.ProtoReflect.Descriptor instead.
func (*TeamUpdateRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2}
}

func (x *TeamUpdateRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *TeamUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// チームを取得する際のリクエストを表します。
type TeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
}

func (x *TeamRequest) Reset() {
	*x = TeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRequest) ProtoMessage() {}

func (x *TeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRequest.ProtoReflect.Descriptor instead.
func (*TeamRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{3}
}

func (x *TeamRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

//
// チームのリストを取得する際のリクエストを表します。
type TeamListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TeamListRequest) Reset() {
	*x = TeamListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamListRequest) ProtoMessage() {}

func (x *TeamListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamListRequest.ProtoReflect.Descriptor instead.
func (*TeamListRequest) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{4}
}

//
// チームのリストを表します。
type Teams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"` // チームのリスト。
}

func (x *Teams) Reset() {
	*x = Teams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teams) ProtoMessage() {}

func (x *Teams) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teams.ProtoReflect.Descriptor instead.
func (*Teams) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{5}
}

func (x *Teams) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

var File_team_proto protoreflect.FileDescriptor

var file_team_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x74,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x04, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0b, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x05, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x2b,
	0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x32, 0xd5, 0x02, 0x0a, 0x0b,
	0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x3a, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x43,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x74, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x70, 0x69, 0x72, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x64, 0x4a, 0x50, 0x2f, 0x61,
	0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_team_proto_rawDescOnce sync.Once
	file_team_proto_rawDescData = file_team_proto_rawDesc
)

func file_team_proto_rawDescGZIP() []byte {
	file_team_proto_rawDescOnce.Do(func() {
		file_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_proto_rawDescData)
	})
	return file_team_proto_rawDescData
}

var file_team_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_team_proto_goTypes = []interface{}{
	(*Team)(nil),              // 0: at_core_service.Team
	(*TeamCreateRequest)(nil), // 1: at_core_service.TeamCreateRequest
	(*TeamUpdateRequest)(nil), // 2: at_core_service.TeamUpdateRequest
	(*TeamRequest)(nil),       // 3: at_core_service.TeamRequest
	(*TeamListRequest)(nil),   // 4: at_core_service.TeamListRequest
	(*Teams)(nil),             // 5: at_core_service.Teams
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_team_proto_depIdxs = []int32{
	0, // 0: at_core_service.Teams.teams:type_name -> at_core_service.Team
	4, // 1: at_core_service.TeamService.List:input_type -> at_core_service.TeamListRequest
	3, // 2: at_core_service.TeamService.Get:input_type -> at_core_service.TeamRequest
	1, // 3: at_core_service.TeamService.Create:input_type -> at_core_service.TeamCreateRequest
	2, // 4: at_core_service.TeamService.Update:input_type -> at_core_service.TeamUpdateRequest
	3, // 5: at_core_service.TeamService.Delete:input_type -> at_core_service.TeamRequest
	5, // 6: at_core_service.TeamService.List:output_type -> at_core_service.Teams
	0, // 7: at_core_service.TeamService.Get:output_type -> at_core_service.Team
	0, // 8: at_core_service.TeamService.Create:output_type -> at_core_service.Team
	0, // 9: at_core_service.TeamService.Update:output_type -> at_core_service.Team
	6, // 10: at_core_service.TeamService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_team_proto_init() }
func file_team_proto_init() {
	if File_team_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCreateRequest); i {
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
		file_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamUpdateRequest); i {
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
		file_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRequest); i {
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
		file_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamListRequest); i {
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
		file_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teams); i {
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
			RawDescriptor: file_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_team_proto_goTypes,
		DependencyIndexes: file_team_proto_depIdxs,
		MessageInfos:      file_team_proto_msgTypes,
	}.Build()
	File_team_proto = out.File
	file_team_proto_rawDesc = nil
	file_team_proto_goTypes = nil
	file_team_proto_depIdxs = nil
}
