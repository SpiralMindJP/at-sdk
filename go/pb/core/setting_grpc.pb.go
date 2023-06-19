//
// Setting は設定に関する情報を操作します。

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: core/setting.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SettingService_GetSettings_FullMethodName       = "/at.core.SettingService/GetSettings"
	SettingService_GetSetting_FullMethodName        = "/at.core.SettingService/GetSetting"
	SettingService_CreateSetting_FullMethodName     = "/at.core.SettingService/CreateSetting"
	SettingService_UpdateSetting_FullMethodName     = "/at.core.SettingService/UpdateSetting"
	SettingService_DeleteSetting_FullMethodName     = "/at.core.SettingService/DeleteSetting"
	SettingService_GetTeamSetting_FullMethodName    = "/at.core.SettingService/GetTeamSetting"
	SettingService_SetTeamSetting_FullMethodName    = "/at.core.SettingService/SetTeamSetting"
	SettingService_DeleteTeamSetting_FullMethodName = "/at.core.SettingService/DeleteTeamSetting"
)

// SettingServiceClient is the client API for SettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingServiceClient interface {
	// 設定のリストを取得します。
	// チームIDを指定した、GetSettingsRequest を渡します。
	// 取得した設定のリストがセットされた GetSettingsResponse が帰ります。
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	// 設定を取得します。
	// 取得する設定のチームIDと設定IDを指定した、GetSettingRequest を渡します。
	// 取得した設定がセットされた GetSettingResponse が帰ります。
	GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error)
	// 新しく設定を作成します。
	// 設定名と属性、チームIDを指定した CreateSettingRequest を渡します。
	// 設定の作成に成功すると、作成された設定情報が設定された CreateSettingResponse が返ります。
	CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error)
	// 設定を更新します。
	// 更新する設定のチームID、設定ID、設定名、属性を指定した UpdateSettingRequest を渡します。
	// 設定の更新に成功すると、更新された設定情報が設定された UpdateSettingResponse が返ります。
	UpdateSetting(ctx context.Context, in *UpdateSettingRequest, opts ...grpc.CallOption) (*UpdateSettingResponse, error)
	// 設定を削除します。
	// 削除する設定のチームIDと設定IDを指定した DeleteSettingRequest を渡します。
	DeleteSetting(ctx context.Context, in *DeleteSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// チーム設定を取得します。
	// 取得するチーム設定のチームIDを指定した、GetTeamSettingRequest を渡します。
	// チーム設定がセットされた GetTeamSettingResponse が帰ります。
	GetTeamSetting(ctx context.Context, in *GetTeamSettingRequest, opts ...grpc.CallOption) (*GetTeamSettingResponse, error)
	// チーム設定を設定します。
	// チーム設定を指定した SetTeamSettingRequest を渡します。
	SetTeamSetting(ctx context.Context, in *SetTeamSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// チーム設定を削除します。
	// 削除するチーム設定のチームIDを指定した、DeleteTeamSettingRequest を渡します。
	DeleteTeamSetting(ctx context.Context, in *DeleteTeamSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type settingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingServiceClient(cc grpc.ClientConnInterface) SettingServiceClient {
	return &settingServiceClient{cc}
}

func (c *settingServiceClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := c.cc.Invoke(ctx, SettingService_GetSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) GetSetting(ctx context.Context, in *GetSettingRequest, opts ...grpc.CallOption) (*GetSettingResponse, error) {
	out := new(GetSettingResponse)
	err := c.cc.Invoke(ctx, SettingService_GetSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error) {
	out := new(CreateSettingResponse)
	err := c.cc.Invoke(ctx, SettingService_CreateSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) UpdateSetting(ctx context.Context, in *UpdateSettingRequest, opts ...grpc.CallOption) (*UpdateSettingResponse, error) {
	out := new(UpdateSettingResponse)
	err := c.cc.Invoke(ctx, SettingService_UpdateSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) DeleteSetting(ctx context.Context, in *DeleteSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SettingService_DeleteSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) GetTeamSetting(ctx context.Context, in *GetTeamSettingRequest, opts ...grpc.CallOption) (*GetTeamSettingResponse, error) {
	out := new(GetTeamSettingResponse)
	err := c.cc.Invoke(ctx, SettingService_GetTeamSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) SetTeamSetting(ctx context.Context, in *SetTeamSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SettingService_SetTeamSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingServiceClient) DeleteTeamSetting(ctx context.Context, in *DeleteTeamSettingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SettingService_DeleteTeamSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingServiceServer is the server API for SettingService service.
// All implementations must embed UnimplementedSettingServiceServer
// for forward compatibility
type SettingServiceServer interface {
	// 設定のリストを取得します。
	// チームIDを指定した、GetSettingsRequest を渡します。
	// 取得した設定のリストがセットされた GetSettingsResponse が帰ります。
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	// 設定を取得します。
	// 取得する設定のチームIDと設定IDを指定した、GetSettingRequest を渡します。
	// 取得した設定がセットされた GetSettingResponse が帰ります。
	GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error)
	// 新しく設定を作成します。
	// 設定名と属性、チームIDを指定した CreateSettingRequest を渡します。
	// 設定の作成に成功すると、作成された設定情報が設定された CreateSettingResponse が返ります。
	CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error)
	// 設定を更新します。
	// 更新する設定のチームID、設定ID、設定名、属性を指定した UpdateSettingRequest を渡します。
	// 設定の更新に成功すると、更新された設定情報が設定された UpdateSettingResponse が返ります。
	UpdateSetting(context.Context, *UpdateSettingRequest) (*UpdateSettingResponse, error)
	// 設定を削除します。
	// 削除する設定のチームIDと設定IDを指定した DeleteSettingRequest を渡します。
	DeleteSetting(context.Context, *DeleteSettingRequest) (*emptypb.Empty, error)
	// チーム設定を取得します。
	// 取得するチーム設定のチームIDを指定した、GetTeamSettingRequest を渡します。
	// チーム設定がセットされた GetTeamSettingResponse が帰ります。
	GetTeamSetting(context.Context, *GetTeamSettingRequest) (*GetTeamSettingResponse, error)
	// チーム設定を設定します。
	// チーム設定を指定した SetTeamSettingRequest を渡します。
	SetTeamSetting(context.Context, *SetTeamSettingRequest) (*emptypb.Empty, error)
	// チーム設定を削除します。
	// 削除するチーム設定のチームIDを指定した、DeleteTeamSettingRequest を渡します。
	DeleteTeamSetting(context.Context, *DeleteTeamSettingRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSettingServiceServer()
}

// UnimplementedSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSettingServiceServer struct {
}

func (UnimplementedSettingServiceServer) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (UnimplementedSettingServiceServer) GetSetting(context.Context, *GetSettingRequest) (*GetSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetting not implemented")
}
func (UnimplementedSettingServiceServer) CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}
func (UnimplementedSettingServiceServer) UpdateSetting(context.Context, *UpdateSettingRequest) (*UpdateSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSetting not implemented")
}
func (UnimplementedSettingServiceServer) DeleteSetting(context.Context, *DeleteSettingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSetting not implemented")
}
func (UnimplementedSettingServiceServer) GetTeamSetting(context.Context, *GetTeamSettingRequest) (*GetTeamSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamSetting not implemented")
}
func (UnimplementedSettingServiceServer) SetTeamSetting(context.Context, *SetTeamSettingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTeamSetting not implemented")
}
func (UnimplementedSettingServiceServer) DeleteTeamSetting(context.Context, *DeleteTeamSettingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeamSetting not implemented")
}
func (UnimplementedSettingServiceServer) mustEmbedUnimplementedSettingServiceServer() {}

// UnsafeSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingServiceServer will
// result in compilation errors.
type UnsafeSettingServiceServer interface {
	mustEmbedUnimplementedSettingServiceServer()
}

func RegisterSettingServiceServer(s grpc.ServiceRegistrar, srv SettingServiceServer) {
	s.RegisterService(&SettingService_ServiceDesc, srv)
}

func _SettingService_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_GetSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).GetSettings(ctx, req.(*GetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_GetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).GetSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_GetSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).GetSetting(ctx, req.(*GetSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_CreateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).CreateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_CreateSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).CreateSetting(ctx, req.(*CreateSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_UpdateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).UpdateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_UpdateSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).UpdateSetting(ctx, req.(*UpdateSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_DeleteSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).DeleteSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_DeleteSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).DeleteSetting(ctx, req.(*DeleteSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_GetTeamSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).GetTeamSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_GetTeamSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).GetTeamSetting(ctx, req.(*GetTeamSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_SetTeamSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTeamSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).SetTeamSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_SetTeamSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).SetTeamSetting(ctx, req.(*SetTeamSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingService_DeleteTeamSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServiceServer).DeleteTeamSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingService_DeleteTeamSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServiceServer).DeleteTeamSetting(ctx, req.(*DeleteTeamSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingService_ServiceDesc is the grpc.ServiceDesc for SettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "at.core.SettingService",
	HandlerType: (*SettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler:    _SettingService_GetSettings_Handler,
		},
		{
			MethodName: "GetSetting",
			Handler:    _SettingService_GetSetting_Handler,
		},
		{
			MethodName: "CreateSetting",
			Handler:    _SettingService_CreateSetting_Handler,
		},
		{
			MethodName: "UpdateSetting",
			Handler:    _SettingService_UpdateSetting_Handler,
		},
		{
			MethodName: "DeleteSetting",
			Handler:    _SettingService_DeleteSetting_Handler,
		},
		{
			MethodName: "GetTeamSetting",
			Handler:    _SettingService_GetTeamSetting_Handler,
		},
		{
			MethodName: "SetTeamSetting",
			Handler:    _SettingService_SetTeamSetting_Handler,
		},
		{
			MethodName: "DeleteTeamSetting",
			Handler:    _SettingService_DeleteTeamSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/setting.proto",
}