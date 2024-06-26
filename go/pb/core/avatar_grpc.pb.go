//
// Avatar はアバターに関する情報を操作します。

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: core/avatar.proto

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
	AvatarService_GetAvatars_FullMethodName           = "/at.core.AvatarService/GetAvatars"
	AvatarService_GetAvatar_FullMethodName            = "/at.core.AvatarService/GetAvatar"
	AvatarService_CreateAvatar_FullMethodName         = "/at.core.AvatarService/CreateAvatar"
	AvatarService_UpdateAvatar_FullMethodName         = "/at.core.AvatarService/UpdateAvatar"
	AvatarService_DeleteAvatar_FullMethodName         = "/at.core.AvatarService/DeleteAvatar"
	AvatarService_GetOperatorAvatar_FullMethodName    = "/at.core.AvatarService/GetOperatorAvatar"
	AvatarService_SetOperatorAvatar_FullMethodName    = "/at.core.AvatarService/SetOperatorAvatar"
	AvatarService_DeleteOperatorAvatar_FullMethodName = "/at.core.AvatarService/DeleteOperatorAvatar"
)

// AvatarServiceClient is the client API for AvatarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvatarServiceClient interface {
	// アバター情報のリストを取得します。
	// 取得するアバターのチームIDを指定した GetAvatarsRequest を渡します。
	// used_only に true を指定すると、Operator デバイスに設定されているアバターのみ取得します。
	// 指定されたチームIDのアバター情報のリストが設定された GetAvatarsResponse が返ります。
	GetAvatars(ctx context.Context, in *GetAvatarsRequest, opts ...grpc.CallOption) (*GetAvatarsResponse, error)
	// アバター情報を取得します。
	// 取得するアバターのアバターIDとチームIDを指定した GetAvatarRequest を渡します。
	// アバターが存在する場合、アバター情報が設定された GetAvatarResponse が返ります。
	GetAvatar(ctx context.Context, in *GetAvatarRequest, opts ...grpc.CallOption) (*GetAvatarResponse, error)
	// 新しくアバターを作成します。
	// アバター名とアバターのコンテンツID、アニメーションのコンテンツID、チームIDを指定した CreateAvatarRequest を渡します。
	// アバターの作成に成功すると、作成したアバター情報が設定された CreateAvatarResponse が返ります。
	CreateAvatar(ctx context.Context, in *CreateAvatarRequest, opts ...grpc.CallOption) (*CreateAvatarResponse, error)
	// アバター情報を更新します。
	// 更新するアバターのアバターIDとチームID、新しいアバター名、アバターのコンテンツID、アニメーションのコンテンツIDを指定した UpdateAvatarRequest を渡します。
	// アバターの更新に成功すると、更新したアバター情報が設定された UpdateAvatarResponse が返ります。
	UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...grpc.CallOption) (*UpdateAvatarResponse, error)
	// アバターを削除します。
	// 削除するアバターのアバターIDとチームIDを指定した DeleteAvatarRequest を渡します。
	DeleteAvatar(ctx context.Context, in *DeleteAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// オペレーターデバイスに設定されているアバター情報を取得します。
	// アバター情報を取得するオペレーターデバイスのデバイスIDとチームIDを指定した GetOperatorAvatarRequest を渡します。
	// デバイスにオペレーターが設定されている場合、OperatorAvatar が返ります。
	GetOperatorAvatar(ctx context.Context, in *GetOperatorAvatarRequest, opts ...grpc.CallOption) (*GetOperatorAvatarResponse, error)
	// オペレーターデバイスにアバターを設定します。
	// 設定するアバターのアバターIDと、設定するオペレーターデバイスのデバイスIDとチームIDを指定した SetOperatorAvatarRequest を渡します。
	// オペレーターデバイスへのアバターの設定に成功すると、オペレーターアバター情報を設定した SetOperatorAvatarResponse が返ります。
	SetOperatorAvatar(ctx context.Context, in *SetOperatorAvatarRequest, opts ...grpc.CallOption) (*SetOperatorAvatarResponse, error)
	// オペレーターデバイスに設定されたアバターを削除します。
	// 削除するオペレーターデバイスのデバイスIDとチームIDを指定した DeleteOperatorAvatarRequest を渡します。
	DeleteOperatorAvatar(ctx context.Context, in *DeleteOperatorAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type avatarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvatarServiceClient(cc grpc.ClientConnInterface) AvatarServiceClient {
	return &avatarServiceClient{cc}
}

func (c *avatarServiceClient) GetAvatars(ctx context.Context, in *GetAvatarsRequest, opts ...grpc.CallOption) (*GetAvatarsResponse, error) {
	out := new(GetAvatarsResponse)
	err := c.cc.Invoke(ctx, AvatarService_GetAvatars_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) GetAvatar(ctx context.Context, in *GetAvatarRequest, opts ...grpc.CallOption) (*GetAvatarResponse, error) {
	out := new(GetAvatarResponse)
	err := c.cc.Invoke(ctx, AvatarService_GetAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) CreateAvatar(ctx context.Context, in *CreateAvatarRequest, opts ...grpc.CallOption) (*CreateAvatarResponse, error) {
	out := new(CreateAvatarResponse)
	err := c.cc.Invoke(ctx, AvatarService_CreateAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...grpc.CallOption) (*UpdateAvatarResponse, error) {
	out := new(UpdateAvatarResponse)
	err := c.cc.Invoke(ctx, AvatarService_UpdateAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) DeleteAvatar(ctx context.Context, in *DeleteAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AvatarService_DeleteAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) GetOperatorAvatar(ctx context.Context, in *GetOperatorAvatarRequest, opts ...grpc.CallOption) (*GetOperatorAvatarResponse, error) {
	out := new(GetOperatorAvatarResponse)
	err := c.cc.Invoke(ctx, AvatarService_GetOperatorAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) SetOperatorAvatar(ctx context.Context, in *SetOperatorAvatarRequest, opts ...grpc.CallOption) (*SetOperatorAvatarResponse, error) {
	out := new(SetOperatorAvatarResponse)
	err := c.cc.Invoke(ctx, AvatarService_SetOperatorAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *avatarServiceClient) DeleteOperatorAvatar(ctx context.Context, in *DeleteOperatorAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AvatarService_DeleteOperatorAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvatarServiceServer is the server API for AvatarService service.
// All implementations must embed UnimplementedAvatarServiceServer
// for forward compatibility
type AvatarServiceServer interface {
	// アバター情報のリストを取得します。
	// 取得するアバターのチームIDを指定した GetAvatarsRequest を渡します。
	// used_only に true を指定すると、Operator デバイスに設定されているアバターのみ取得します。
	// 指定されたチームIDのアバター情報のリストが設定された GetAvatarsResponse が返ります。
	GetAvatars(context.Context, *GetAvatarsRequest) (*GetAvatarsResponse, error)
	// アバター情報を取得します。
	// 取得するアバターのアバターIDとチームIDを指定した GetAvatarRequest を渡します。
	// アバターが存在する場合、アバター情報が設定された GetAvatarResponse が返ります。
	GetAvatar(context.Context, *GetAvatarRequest) (*GetAvatarResponse, error)
	// 新しくアバターを作成します。
	// アバター名とアバターのコンテンツID、アニメーションのコンテンツID、チームIDを指定した CreateAvatarRequest を渡します。
	// アバターの作成に成功すると、作成したアバター情報が設定された CreateAvatarResponse が返ります。
	CreateAvatar(context.Context, *CreateAvatarRequest) (*CreateAvatarResponse, error)
	// アバター情報を更新します。
	// 更新するアバターのアバターIDとチームID、新しいアバター名、アバターのコンテンツID、アニメーションのコンテンツIDを指定した UpdateAvatarRequest を渡します。
	// アバターの更新に成功すると、更新したアバター情報が設定された UpdateAvatarResponse が返ります。
	UpdateAvatar(context.Context, *UpdateAvatarRequest) (*UpdateAvatarResponse, error)
	// アバターを削除します。
	// 削除するアバターのアバターIDとチームIDを指定した DeleteAvatarRequest を渡します。
	DeleteAvatar(context.Context, *DeleteAvatarRequest) (*emptypb.Empty, error)
	// オペレーターデバイスに設定されているアバター情報を取得します。
	// アバター情報を取得するオペレーターデバイスのデバイスIDとチームIDを指定した GetOperatorAvatarRequest を渡します。
	// デバイスにオペレーターが設定されている場合、OperatorAvatar が返ります。
	GetOperatorAvatar(context.Context, *GetOperatorAvatarRequest) (*GetOperatorAvatarResponse, error)
	// オペレーターデバイスにアバターを設定します。
	// 設定するアバターのアバターIDと、設定するオペレーターデバイスのデバイスIDとチームIDを指定した SetOperatorAvatarRequest を渡します。
	// オペレーターデバイスへのアバターの設定に成功すると、オペレーターアバター情報を設定した SetOperatorAvatarResponse が返ります。
	SetOperatorAvatar(context.Context, *SetOperatorAvatarRequest) (*SetOperatorAvatarResponse, error)
	// オペレーターデバイスに設定されたアバターを削除します。
	// 削除するオペレーターデバイスのデバイスIDとチームIDを指定した DeleteOperatorAvatarRequest を渡します。
	DeleteOperatorAvatar(context.Context, *DeleteOperatorAvatarRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAvatarServiceServer()
}

// UnimplementedAvatarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvatarServiceServer struct {
}

func (UnimplementedAvatarServiceServer) GetAvatars(context.Context, *GetAvatarsRequest) (*GetAvatarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatars not implemented")
}
func (UnimplementedAvatarServiceServer) GetAvatar(context.Context, *GetAvatarRequest) (*GetAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) CreateAvatar(context.Context, *CreateAvatarRequest) (*CreateAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) UpdateAvatar(context.Context, *UpdateAvatarRequest) (*UpdateAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) DeleteAvatar(context.Context, *DeleteAvatarRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) GetOperatorAvatar(context.Context, *GetOperatorAvatarRequest) (*GetOperatorAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) SetOperatorAvatar(context.Context, *SetOperatorAvatarRequest) (*SetOperatorAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOperatorAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) DeleteOperatorAvatar(context.Context, *DeleteOperatorAvatarRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperatorAvatar not implemented")
}
func (UnimplementedAvatarServiceServer) mustEmbedUnimplementedAvatarServiceServer() {}

// UnsafeAvatarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvatarServiceServer will
// result in compilation errors.
type UnsafeAvatarServiceServer interface {
	mustEmbedUnimplementedAvatarServiceServer()
}

func RegisterAvatarServiceServer(s grpc.ServiceRegistrar, srv AvatarServiceServer) {
	s.RegisterService(&AvatarService_ServiceDesc, srv)
}

func _AvatarService_GetAvatars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).GetAvatars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_GetAvatars_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).GetAvatars(ctx, req.(*GetAvatarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_GetAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).GetAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_GetAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).GetAvatar(ctx, req.(*GetAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_CreateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).CreateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_CreateAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).CreateAvatar(ctx, req.(*CreateAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_UpdateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).UpdateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_UpdateAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).UpdateAvatar(ctx, req.(*UpdateAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_DeleteAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).DeleteAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_DeleteAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).DeleteAvatar(ctx, req.(*DeleteAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_GetOperatorAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).GetOperatorAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_GetOperatorAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).GetOperatorAvatar(ctx, req.(*GetOperatorAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_SetOperatorAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOperatorAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).SetOperatorAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_SetOperatorAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).SetOperatorAvatar(ctx, req.(*SetOperatorAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AvatarService_DeleteOperatorAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperatorAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvatarServiceServer).DeleteOperatorAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AvatarService_DeleteOperatorAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvatarServiceServer).DeleteOperatorAvatar(ctx, req.(*DeleteOperatorAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AvatarService_ServiceDesc is the grpc.ServiceDesc for AvatarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvatarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "at.core.AvatarService",
	HandlerType: (*AvatarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvatars",
			Handler:    _AvatarService_GetAvatars_Handler,
		},
		{
			MethodName: "GetAvatar",
			Handler:    _AvatarService_GetAvatar_Handler,
		},
		{
			MethodName: "CreateAvatar",
			Handler:    _AvatarService_CreateAvatar_Handler,
		},
		{
			MethodName: "UpdateAvatar",
			Handler:    _AvatarService_UpdateAvatar_Handler,
		},
		{
			MethodName: "DeleteAvatar",
			Handler:    _AvatarService_DeleteAvatar_Handler,
		},
		{
			MethodName: "GetOperatorAvatar",
			Handler:    _AvatarService_GetOperatorAvatar_Handler,
		},
		{
			MethodName: "SetOperatorAvatar",
			Handler:    _AvatarService_SetOperatorAvatar_Handler,
		},
		{
			MethodName: "DeleteOperatorAvatar",
			Handler:    _AvatarService_DeleteOperatorAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/avatar.proto",
}
