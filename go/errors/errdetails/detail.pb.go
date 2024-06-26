// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.26.1
// source: errdetails/detail.proto

package errdetails

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

// エラー情報を表します。
type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"` // エラーが発生した理由。
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// フィールドの違反を表します。
type BadRequestFieldViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field       string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`             // フィールドの名前。
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // 違反の内容。
}

func (x *BadRequestFieldViolation) Reset() {
	*x = BadRequestFieldViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadRequestFieldViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadRequestFieldViolation) ProtoMessage() {}

func (x *BadRequestFieldViolation) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadRequestFieldViolation.ProtoReflect.Descriptor instead.
func (*BadRequestFieldViolation) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{1}
}

func (x *BadRequestFieldViolation) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *BadRequestFieldViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 不正なリクエストの詳細を表します。
type BadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldViolations []*BadRequestFieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations,proto3" json:"field_violations,omitempty"` // フィールド違反のリスト。
}

func (x *BadRequest) Reset() {
	*x = BadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadRequest) ProtoMessage() {}

func (x *BadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadRequest.ProtoReflect.Descriptor instead.
func (*BadRequest) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{2}
}

func (x *BadRequest) GetFieldViolations() []*BadRequestFieldViolation {
	if x != nil {
		return x.FieldViolations
	}
	return nil
}

// 必要条件の違反を表します。
type PreconditionFailureViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`               // 必要条件のタイプ。
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`         // タイプに対する相対的な対象。
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // 違反の内容。
}

func (x *PreconditionFailureViolation) Reset() {
	*x = PreconditionFailureViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailureViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailureViolation) ProtoMessage() {}

func (x *PreconditionFailureViolation) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailureViolation.ProtoReflect.Descriptor instead.
func (*PreconditionFailureViolation) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{3}
}

func (x *PreconditionFailureViolation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PreconditionFailureViolation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PreconditionFailureViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 必要条件の不足を表わします。
type PreconditionFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Violations []*PreconditionFailureViolation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"` // 必要条件の違反のリスト。
}

func (x *PreconditionFailure) Reset() {
	*x = PreconditionFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailure) ProtoMessage() {}

func (x *PreconditionFailure) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailure.ProtoReflect.Descriptor instead.
func (*PreconditionFailure) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{4}
}

func (x *PreconditionFailure) GetViolations() []*PreconditionFailureViolation {
	if x != nil {
		return x.Violations
	}
	return nil
}

// 割り当ての違反を表します。
type QuotaFailureViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject     string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`         // 違反の対象。
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // 違反の内容。
}

func (x *QuotaFailureViolation) Reset() {
	*x = QuotaFailureViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailureViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailureViolation) ProtoMessage() {}

func (x *QuotaFailureViolation) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailureViolation.ProtoReflect.Descriptor instead.
func (*QuotaFailureViolation) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{5}
}

func (x *QuotaFailureViolation) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *QuotaFailureViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 割り当ての失敗を表します。
type QuotaFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Violations []*QuotaFailureViolation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"` // 割り当ての違反のリスト。
}

func (x *QuotaFailure) Reset() {
	*x = QuotaFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errdetails_detail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaFailure) ProtoMessage() {}

func (x *QuotaFailure) ProtoReflect() protoreflect.Message {
	mi := &file_errdetails_detail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaFailure.ProtoReflect.Descriptor instead.
func (*QuotaFailure) Descriptor() ([]byte, []int) {
	return file_errdetails_detail_proto_rawDescGZIP(), []int{6}
}

func (x *QuotaFailure) GetViolations() []*QuotaFailureViolation {
	if x != nil {
		return x.Violations
	}
	return nil
}

var File_errdetails_detail_proto protoreflect.FileDescriptor

var file_errdetails_detail_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x74, 0x2e, 0x65, 0x72,
	0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x23, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x52, 0x0a,
	0x18, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x60, 0x0a, 0x0a, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x52, 0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x74, 0x2e, 0x65,
	0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x6e, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x76, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x61, 0x74, 0x2e, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x53, 0x0a, 0x15, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0c,
	0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x44, 0x0a, 0x0a,
	0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x61, 0x74, 0x2e, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x56, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_errdetails_detail_proto_rawDescOnce sync.Once
	file_errdetails_detail_proto_rawDescData = file_errdetails_detail_proto_rawDesc
)

func file_errdetails_detail_proto_rawDescGZIP() []byte {
	file_errdetails_detail_proto_rawDescOnce.Do(func() {
		file_errdetails_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_errdetails_detail_proto_rawDescData)
	})
	return file_errdetails_detail_proto_rawDescData
}

var file_errdetails_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_errdetails_detail_proto_goTypes = []interface{}{
	(*ErrorInfo)(nil),                    // 0: at.errdetails.ErrorInfo
	(*BadRequestFieldViolation)(nil),     // 1: at.errdetails.BadRequestFieldViolation
	(*BadRequest)(nil),                   // 2: at.errdetails.BadRequest
	(*PreconditionFailureViolation)(nil), // 3: at.errdetails.PreconditionFailureViolation
	(*PreconditionFailure)(nil),          // 4: at.errdetails.PreconditionFailure
	(*QuotaFailureViolation)(nil),        // 5: at.errdetails.QuotaFailureViolation
	(*QuotaFailure)(nil),                 // 6: at.errdetails.QuotaFailure
}
var file_errdetails_detail_proto_depIdxs = []int32{
	1, // 0: at.errdetails.BadRequest.field_violations:type_name -> at.errdetails.BadRequestFieldViolation
	3, // 1: at.errdetails.PreconditionFailure.violations:type_name -> at.errdetails.PreconditionFailureViolation
	5, // 2: at.errdetails.QuotaFailure.violations:type_name -> at.errdetails.QuotaFailureViolation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_errdetails_detail_proto_init() }
func file_errdetails_detail_proto_init() {
	if File_errdetails_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errdetails_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
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
		file_errdetails_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadRequestFieldViolation); i {
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
		file_errdetails_detail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadRequest); i {
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
		file_errdetails_detail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailureViolation); i {
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
		file_errdetails_detail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailure); i {
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
		file_errdetails_detail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailureViolation); i {
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
		file_errdetails_detail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaFailure); i {
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
			RawDescriptor: file_errdetails_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errdetails_detail_proto_goTypes,
		DependencyIndexes: file_errdetails_detail_proto_depIdxs,
		MessageInfos:      file_errdetails_detail_proto_msgTypes,
	}.Build()
	File_errdetails_detail_proto = out.File
	file_errdetails_detail_proto_rawDesc = nil
	file_errdetails_detail_proto_goTypes = nil
	file_errdetails_detail_proto_depIdxs = nil
}
