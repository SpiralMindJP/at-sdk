//
// Setting は設定に関する情報を操作します。

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.26.1
// source: core/setting.proto

package core

import (
	common "github.com/SpiralMindJP/at-sdk/go/pb/common"
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

// 設定リスト取得クエストを表します。
type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{0}
}

func (x *GetSettingsRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

// 設定リスト取得レスポンスを表します。
type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings []*common.Setting `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty"` // 設定情報のリスト。
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{1}
}

func (x *GetSettingsResponse) GetSettings() []*common.Setting {
	if x != nil {
		return x.Settings
	}
	return nil
}

// 設定取得リクエストを表します。
type GetSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`          // チームID。
	SettingId int64 `protobuf:"varint,2,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"` // 設定ID。
}

func (x *GetSettingRequest) Reset() {
	*x = GetSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingRequest) ProtoMessage() {}

func (x *GetSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingRequest.ProtoReflect.Descriptor instead.
func (*GetSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{2}
}

func (x *GetSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *GetSettingRequest) GetSettingId() int64 {
	if x != nil {
		return x.SettingId
	}
	return 0
}

// 設定取得レスポンスを表します。
type GetSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *common.Setting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"` // 設定情報。
}

func (x *GetSettingResponse) Reset() {
	*x = GetSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingResponse) ProtoMessage() {}

func (x *GetSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingResponse.ProtoReflect.Descriptor instead.
func (*GetSettingResponse) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{3}
}

func (x *GetSettingResponse) GetSetting() *common.Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

// 設定作成リクエストを表します。
type CreateSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                    // 設定名。
	Attributes string `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`        // 属性。
}

func (x *CreateSettingRequest) Reset() {
	*x = CreateSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingRequest) ProtoMessage() {}

func (x *CreateSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingRequest.ProtoReflect.Descriptor instead.
func (*CreateSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *CreateSettingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSettingRequest) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

// 設定作成レスポンスを表します。
type CreateSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *common.Setting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"` // 設定情報。
}

func (x *CreateSettingResponse) Reset() {
	*x = CreateSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSettingResponse) ProtoMessage() {}

func (x *CreateSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSettingResponse.ProtoReflect.Descriptor instead.
func (*CreateSettingResponse) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSettingResponse) GetSetting() *common.Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

// 設定更新リクエストを表します。
type UpdateSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`          // チームID。
	SettingId  int64  `protobuf:"varint,2,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"` // 設定ID。
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                             // 設定名。
	Attributes string `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`                 // 属性。
}

func (x *UpdateSettingRequest) Reset() {
	*x = UpdateSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingRequest) ProtoMessage() {}

func (x *UpdateSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingRequest.ProtoReflect.Descriptor instead.
func (*UpdateSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *UpdateSettingRequest) GetSettingId() int64 {
	if x != nil {
		return x.SettingId
	}
	return 0
}

func (x *UpdateSettingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSettingRequest) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

// 設定更新レスポンスを表します。
type UpdateSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *common.Setting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"` // 設定情報。
}

func (x *UpdateSettingResponse) Reset() {
	*x = UpdateSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSettingResponse) ProtoMessage() {}

func (x *UpdateSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSettingResponse.ProtoReflect.Descriptor instead.
func (*UpdateSettingResponse) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSettingResponse) GetSetting() *common.Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

// 設定削除リクエストを表します。
type DeleteSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`          // チームID。
	SettingId int64 `protobuf:"varint,2,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"` // 設定ID。
}

func (x *DeleteSettingRequest) Reset() {
	*x = DeleteSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSettingRequest) ProtoMessage() {}

func (x *DeleteSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSettingRequest.ProtoReflect.Descriptor instead.
func (*DeleteSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *DeleteSettingRequest) GetSettingId() int64 {
	if x != nil {
		return x.SettingId
	}
	return 0
}

// チーム設定取得リクエストを表します。
type GetTeamSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
}

func (x *GetTeamSettingRequest) Reset() {
	*x = GetTeamSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamSettingRequest) ProtoMessage() {}

func (x *GetTeamSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamSettingRequest.ProtoReflect.Descriptor instead.
func (*GetTeamSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{9}
}

func (x *GetTeamSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

// チーム設定取得レスポンスを表します。
type GetTeamSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *common.TeamSetting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"` // チーム設定。
}

func (x *GetTeamSettingResponse) Reset() {
	*x = GetTeamSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamSettingResponse) ProtoMessage() {}

func (x *GetTeamSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamSettingResponse.ProtoReflect.Descriptor instead.
func (*GetTeamSettingResponse) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{10}
}

func (x *GetTeamSettingResponse) GetSetting() *common.TeamSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

// チーム設定の設定リクエストを表します。
type SetTeamSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *common.TeamSetting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"` // チーム設定。
}

func (x *SetTeamSettingRequest) Reset() {
	*x = SetTeamSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTeamSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTeamSettingRequest) ProtoMessage() {}

func (x *SetTeamSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTeamSettingRequest.ProtoReflect.Descriptor instead.
func (*SetTeamSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{11}
}

func (x *SetTeamSettingRequest) GetSetting() *common.TeamSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

// チーム設定削除リクエストを表します。
type DeleteTeamSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"` // チームID。
}

func (x *DeleteTeamSettingRequest) Reset() {
	*x = DeleteTeamSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_setting_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamSettingRequest) ProtoMessage() {}

func (x *DeleteTeamSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_setting_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamSettingRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamSettingRequest) Descriptor() ([]byte, []int) {
	return file_core_setting_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteTeamSettingRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

var File_core_setting_proto protoreflect.FileDescriptor

var file_core_setting_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x63, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4e,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x30,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x49, 0x0a, 0x15,
	0x53, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x32, 0xf6, 0x04, 0x0a,
	0x0e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b,
	0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x2e, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x61, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e,
	0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x69, 0x72, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x64, 0x4a, 0x50,
	0x2f, 0x61, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_setting_proto_rawDescOnce sync.Once
	file_core_setting_proto_rawDescData = file_core_setting_proto_rawDesc
)

func file_core_setting_proto_rawDescGZIP() []byte {
	file_core_setting_proto_rawDescOnce.Do(func() {
		file_core_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_setting_proto_rawDescData)
	})
	return file_core_setting_proto_rawDescData
}

var file_core_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_core_setting_proto_goTypes = []interface{}{
	(*GetSettingsRequest)(nil),       // 0: at.core.GetSettingsRequest
	(*GetSettingsResponse)(nil),      // 1: at.core.GetSettingsResponse
	(*GetSettingRequest)(nil),        // 2: at.core.GetSettingRequest
	(*GetSettingResponse)(nil),       // 3: at.core.GetSettingResponse
	(*CreateSettingRequest)(nil),     // 4: at.core.CreateSettingRequest
	(*CreateSettingResponse)(nil),    // 5: at.core.CreateSettingResponse
	(*UpdateSettingRequest)(nil),     // 6: at.core.UpdateSettingRequest
	(*UpdateSettingResponse)(nil),    // 7: at.core.UpdateSettingResponse
	(*DeleteSettingRequest)(nil),     // 8: at.core.DeleteSettingRequest
	(*GetTeamSettingRequest)(nil),    // 9: at.core.GetTeamSettingRequest
	(*GetTeamSettingResponse)(nil),   // 10: at.core.GetTeamSettingResponse
	(*SetTeamSettingRequest)(nil),    // 11: at.core.SetTeamSettingRequest
	(*DeleteTeamSettingRequest)(nil), // 12: at.core.DeleteTeamSettingRequest
	(*common.Setting)(nil),           // 13: at.common.Setting
	(*common.TeamSetting)(nil),       // 14: at.common.TeamSetting
	(*emptypb.Empty)(nil),            // 15: google.protobuf.Empty
}
var file_core_setting_proto_depIdxs = []int32{
	13, // 0: at.core.GetSettingsResponse.settings:type_name -> at.common.Setting
	13, // 1: at.core.GetSettingResponse.setting:type_name -> at.common.Setting
	13, // 2: at.core.CreateSettingResponse.setting:type_name -> at.common.Setting
	13, // 3: at.core.UpdateSettingResponse.setting:type_name -> at.common.Setting
	14, // 4: at.core.GetTeamSettingResponse.setting:type_name -> at.common.TeamSetting
	14, // 5: at.core.SetTeamSettingRequest.setting:type_name -> at.common.TeamSetting
	0,  // 6: at.core.SettingService.GetSettings:input_type -> at.core.GetSettingsRequest
	2,  // 7: at.core.SettingService.GetSetting:input_type -> at.core.GetSettingRequest
	4,  // 8: at.core.SettingService.CreateSetting:input_type -> at.core.CreateSettingRequest
	6,  // 9: at.core.SettingService.UpdateSetting:input_type -> at.core.UpdateSettingRequest
	8,  // 10: at.core.SettingService.DeleteSetting:input_type -> at.core.DeleteSettingRequest
	9,  // 11: at.core.SettingService.GetTeamSetting:input_type -> at.core.GetTeamSettingRequest
	11, // 12: at.core.SettingService.SetTeamSetting:input_type -> at.core.SetTeamSettingRequest
	12, // 13: at.core.SettingService.DeleteTeamSetting:input_type -> at.core.DeleteTeamSettingRequest
	1,  // 14: at.core.SettingService.GetSettings:output_type -> at.core.GetSettingsResponse
	3,  // 15: at.core.SettingService.GetSetting:output_type -> at.core.GetSettingResponse
	5,  // 16: at.core.SettingService.CreateSetting:output_type -> at.core.CreateSettingResponse
	7,  // 17: at.core.SettingService.UpdateSetting:output_type -> at.core.UpdateSettingResponse
	15, // 18: at.core.SettingService.DeleteSetting:output_type -> google.protobuf.Empty
	10, // 19: at.core.SettingService.GetTeamSetting:output_type -> at.core.GetTeamSettingResponse
	15, // 20: at.core.SettingService.SetTeamSetting:output_type -> google.protobuf.Empty
	15, // 21: at.core.SettingService.DeleteTeamSetting:output_type -> google.protobuf.Empty
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_core_setting_proto_init() }
func file_core_setting_proto_init() {
	if File_core_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
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
		file_core_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
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
		file_core_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingRequest); i {
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
		file_core_setting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingResponse); i {
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
		file_core_setting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSettingRequest); i {
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
		file_core_setting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSettingResponse); i {
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
		file_core_setting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingRequest); i {
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
		file_core_setting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSettingResponse); i {
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
		file_core_setting_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSettingRequest); i {
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
		file_core_setting_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamSettingRequest); i {
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
		file_core_setting_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamSettingResponse); i {
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
		file_core_setting_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTeamSettingRequest); i {
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
		file_core_setting_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeamSettingRequest); i {
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
			RawDescriptor: file_core_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_setting_proto_goTypes,
		DependencyIndexes: file_core_setting_proto_depIdxs,
		MessageInfos:      file_core_setting_proto_msgTypes,
	}.Build()
	File_core_setting_proto = out.File
	file_core_setting_proto_rawDesc = nil
	file_core_setting_proto_goTypes = nil
	file_core_setting_proto_depIdxs = nil
}
