// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: core/room.proto

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

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// 全てのルームグループのリストを取得します。
	// 取得するルームグループのチームIDを指定した GetAllRoomGroupsRequest を渡します。
	// 取得したルームグループ情報のリストが設定された GetAllRoomGroupsResponse が返ります。
	GetAllRoomGroups(ctx context.Context, in *GetAllRoomGroupsRequest, opts ...grpc.CallOption) (*GetAllRoomGroupsResponse, error)
	// ルームグループのリストを取得します。
	// 取得するルームのチームIDを指定した GetRoomGroupsRequest を渡します。
	// 親ルームグループIDを指定した場合は、そのルームグループに所属するルームグループのリストを取得します。
	// 親ルームグループIDに 0 を指定した場合は、ルームグループに所属していないルームグループのリストを取得します。
	// 取得したルームグループのリストが設定された GetRoomGroupsResponse が返ります。
	GetRoomGroups(ctx context.Context, in *GetRoomGroupsRequest, opts ...grpc.CallOption) (*GetRoomGroupsResponse, error)
	// ルームグループを取得します。
	// 取得するルームグループのチームIDおよびルームグループIDを指定した GetRoomGroupRequest を渡します。
	// ルームグループが存在する場合、ルームグループ情報が設定された GetRoomGroupResponse が返ります。
	GetRoomGroup(ctx context.Context, in *GetRoomGroupRequest, opts ...grpc.CallOption) (*GetRoomGroupResponse, error)
	// 新しくルームグループを作成します。
	// ルームグループ名とチームIDを指定した CreateRoomGroupRequest を渡します。
	// ルームグループをルームグループに所属させる場合は、親ルームグループIDを指定します。
	// ルームグループをルームグループに所属させない場合は、親ルームグループIDに 0 を指定します。
	// ルームグループの作成に成功すると、作成されたルームグループ情報が設定された CreateRoomGroupResponse が返ります。
	CreateRoomGroup(ctx context.Context, in *CreateRoomGroupRequest, opts ...grpc.CallOption) (*CreateRoomGroupResponse, error)
	// ルームグループを更新します。
	// 更新するルームグループのルームグループIDと、新しいルームグループ名を指定した UpdateRoomGroupRequest を渡します。
	// ルームグループをルームグループに所属させる場合は、親ルームグループIDを指定します。
	// ルームグループをルームグループに所属させない場合は、親ルームグループIDに 0 を指定します。
	// ルームグループの更新に成功すると、更新されたルームグループ情報が設定された UpdateRoomGroupResponse が返ります。
	UpdateRoomGroup(ctx context.Context, in *UpdateRoomGroupRequest, opts ...grpc.CallOption) (*UpdateRoomGroupResponse, error)
	// ルームグループを削除します。
	// 削除するルームグループのルームグループIDを指定した DeleteRoomGroupRequest を渡します。
	DeleteRoomGroup(ctx context.Context, in *DeleteRoomGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 全てのルームのリストを取得します。
	// 取得するルームのチームIDを指定した GetAllRoomsRequest を渡します。
	// 取得したルーム情報のリストが設定された GetAllRoomsResponse が返ります。
	GetAllRooms(ctx context.Context, in *GetAllRoomsRequest, opts ...grpc.CallOption) (*GetAllRoomsResponse, error)
	// ルームのリストを取得します。
	// 取得するルームのチームIDを指定した GetRoomsRequest を渡します。
	// ルームグループIDを指定した場合は、そのルームグループに所属するルームのリストを取得します。
	// ルームグループIDに 0 を指定した場合は、ルームグループに所属していないルームのリストを取得します。
	// 取得したルーム情報のリストが設定された GetRoomsResponse が返ります。
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	// ルームを取得します。
	// 取得するルームのチームIDおよびルームIDを指定した GetRoomRequest を渡します。
	// ルームが存在する場合、ルーム情報が設定された GetRoomResponse が返ります。
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	// 新しくルームを作成します。
	// ルーム名とチームIDを指定した CreateRoomRequest を渡します。
	// ルームをルームグループに所属させる場合は、ルームグループIDを指定します。
	// ルームをルームグループに所属させない場合は、ルームグループIDに 0 を指定します。
	// ルームの作成に成功すると、作成されたルーム情報が設定された CreateRoomResponse が返ります。
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	// ルームを更新します。
	// 更新するルームのルームIDと、新しいルーム名を指定した UpdateRoomRequest を渡します。
	// ルームをルームグループに所属させる場合は、ルームグループIDを指定します。
	// ルームをルームグループに所属させない場合は、ルームグループIDに 0 を指定します。
	// ルームの更新に成功すると、更新されたルーム情報が設定された UpdateRoomResponse が返ります。
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error)
	// ルームを削除します。
	// 削除するルームのルームIDを指定した DeleteRoomRequest を渡します。
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ルームへのカスタマーデバイスの入室を行います。
	// 入室先のルームのルームIDと、入室を行うカスタマーデバイスのデバイスIDを指定した JoinCustomerDeviceRequest を渡します。
	// 入室先のルームに既にカスタマーデバイスが入室済の場合は、入室を行いません。
	// ただし、force に true をした場合は入室済のカスタマーデバイスを強制的に退室させ、入室を行います。
	// カスタマーデバイスの入室に成功すると、入室したルームのルーム情報を設定した JoinCustomerDeviceResponse が返ります。
	JoinCustomerDevice(ctx context.Context, in *JoinCustomerDeviceRequest, opts ...grpc.CallOption) (*JoinCustomerDeviceResponse, error)
	// ルームからのカスタマーデバイスの退室を行います。
	// カスタマーデバイスを退室させるルームのルームIDを指定した LeaveCustomerDeviceRequest を渡します。
	// カスタマーデバイスの退室に成功すると、退室したルームのルーム情報を設定した LeaveCustomerDeviceResponse が返ります。
	LeaveCustomerDevice(ctx context.Context, in *LeaveCustomerDeviceRequest, opts ...grpc.CallOption) (*LeaveCustomerDeviceResponse, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) GetAllRoomGroups(ctx context.Context, in *GetAllRoomGroupsRequest, opts ...grpc.CallOption) (*GetAllRoomGroupsResponse, error) {
	out := new(GetAllRoomGroupsResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetAllRoomGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRoomGroups(ctx context.Context, in *GetRoomGroupsRequest, opts ...grpc.CallOption) (*GetRoomGroupsResponse, error) {
	out := new(GetRoomGroupsResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetRoomGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRoomGroup(ctx context.Context, in *GetRoomGroupRequest, opts ...grpc.CallOption) (*GetRoomGroupResponse, error) {
	out := new(GetRoomGroupResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetRoomGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) CreateRoomGroup(ctx context.Context, in *CreateRoomGroupRequest, opts ...grpc.CallOption) (*CreateRoomGroupResponse, error) {
	out := new(CreateRoomGroupResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/CreateRoomGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdateRoomGroup(ctx context.Context, in *UpdateRoomGroupRequest, opts ...grpc.CallOption) (*UpdateRoomGroupResponse, error) {
	out := new(UpdateRoomGroupResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/UpdateRoomGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) DeleteRoomGroup(ctx context.Context, in *DeleteRoomGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/DeleteRoomGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetAllRooms(ctx context.Context, in *GetAllRoomsRequest, opts ...grpc.CallOption) (*GetAllRoomsResponse, error) {
	out := new(GetAllRoomsResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetAllRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error) {
	out := new(UpdateRoomResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/UpdateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) JoinCustomerDevice(ctx context.Context, in *JoinCustomerDeviceRequest, opts ...grpc.CallOption) (*JoinCustomerDeviceResponse, error) {
	out := new(JoinCustomerDeviceResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/JoinCustomerDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) LeaveCustomerDevice(ctx context.Context, in *LeaveCustomerDeviceRequest, opts ...grpc.CallOption) (*LeaveCustomerDeviceResponse, error) {
	out := new(LeaveCustomerDeviceResponse)
	err := c.cc.Invoke(ctx, "/at.core.RoomService/LeaveCustomerDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	// 全てのルームグループのリストを取得します。
	// 取得するルームグループのチームIDを指定した GetAllRoomGroupsRequest を渡します。
	// 取得したルームグループ情報のリストが設定された GetAllRoomGroupsResponse が返ります。
	GetAllRoomGroups(context.Context, *GetAllRoomGroupsRequest) (*GetAllRoomGroupsResponse, error)
	// ルームグループのリストを取得します。
	// 取得するルームのチームIDを指定した GetRoomGroupsRequest を渡します。
	// 親ルームグループIDを指定した場合は、そのルームグループに所属するルームグループのリストを取得します。
	// 親ルームグループIDに 0 を指定した場合は、ルームグループに所属していないルームグループのリストを取得します。
	// 取得したルームグループのリストが設定された GetRoomGroupsResponse が返ります。
	GetRoomGroups(context.Context, *GetRoomGroupsRequest) (*GetRoomGroupsResponse, error)
	// ルームグループを取得します。
	// 取得するルームグループのチームIDおよびルームグループIDを指定した GetRoomGroupRequest を渡します。
	// ルームグループが存在する場合、ルームグループ情報が設定された GetRoomGroupResponse が返ります。
	GetRoomGroup(context.Context, *GetRoomGroupRequest) (*GetRoomGroupResponse, error)
	// 新しくルームグループを作成します。
	// ルームグループ名とチームIDを指定した CreateRoomGroupRequest を渡します。
	// ルームグループをルームグループに所属させる場合は、親ルームグループIDを指定します。
	// ルームグループをルームグループに所属させない場合は、親ルームグループIDに 0 を指定します。
	// ルームグループの作成に成功すると、作成されたルームグループ情報が設定された CreateRoomGroupResponse が返ります。
	CreateRoomGroup(context.Context, *CreateRoomGroupRequest) (*CreateRoomGroupResponse, error)
	// ルームグループを更新します。
	// 更新するルームグループのルームグループIDと、新しいルームグループ名を指定した UpdateRoomGroupRequest を渡します。
	// ルームグループをルームグループに所属させる場合は、親ルームグループIDを指定します。
	// ルームグループをルームグループに所属させない場合は、親ルームグループIDに 0 を指定します。
	// ルームグループの更新に成功すると、更新されたルームグループ情報が設定された UpdateRoomGroupResponse が返ります。
	UpdateRoomGroup(context.Context, *UpdateRoomGroupRequest) (*UpdateRoomGroupResponse, error)
	// ルームグループを削除します。
	// 削除するルームグループのルームグループIDを指定した DeleteRoomGroupRequest を渡します。
	DeleteRoomGroup(context.Context, *DeleteRoomGroupRequest) (*emptypb.Empty, error)
	// 全てのルームのリストを取得します。
	// 取得するルームのチームIDを指定した GetAllRoomsRequest を渡します。
	// 取得したルーム情報のリストが設定された GetAllRoomsResponse が返ります。
	GetAllRooms(context.Context, *GetAllRoomsRequest) (*GetAllRoomsResponse, error)
	// ルームのリストを取得します。
	// 取得するルームのチームIDを指定した GetRoomsRequest を渡します。
	// ルームグループIDを指定した場合は、そのルームグループに所属するルームのリストを取得します。
	// ルームグループIDに 0 を指定した場合は、ルームグループに所属していないルームのリストを取得します。
	// 取得したルーム情報のリストが設定された GetRoomsResponse が返ります。
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	// ルームを取得します。
	// 取得するルームのチームIDおよびルームIDを指定した GetRoomRequest を渡します。
	// ルームが存在する場合、ルーム情報が設定された GetRoomResponse が返ります。
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	// 新しくルームを作成します。
	// ルーム名とチームIDを指定した CreateRoomRequest を渡します。
	// ルームをルームグループに所属させる場合は、ルームグループIDを指定します。
	// ルームをルームグループに所属させない場合は、ルームグループIDに 0 を指定します。
	// ルームの作成に成功すると、作成されたルーム情報が設定された CreateRoomResponse が返ります。
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	// ルームを更新します。
	// 更新するルームのルームIDと、新しいルーム名を指定した UpdateRoomRequest を渡します。
	// ルームをルームグループに所属させる場合は、ルームグループIDを指定します。
	// ルームをルームグループに所属させない場合は、ルームグループIDに 0 を指定します。
	// ルームの更新に成功すると、更新されたルーム情報が設定された UpdateRoomResponse が返ります。
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error)
	// ルームを削除します。
	// 削除するルームのルームIDを指定した DeleteRoomRequest を渡します。
	DeleteRoom(context.Context, *DeleteRoomRequest) (*emptypb.Empty, error)
	// ルームへのカスタマーデバイスの入室を行います。
	// 入室先のルームのルームIDと、入室を行うカスタマーデバイスのデバイスIDを指定した JoinCustomerDeviceRequest を渡します。
	// 入室先のルームに既にカスタマーデバイスが入室済の場合は、入室を行いません。
	// ただし、force に true をした場合は入室済のカスタマーデバイスを強制的に退室させ、入室を行います。
	// カスタマーデバイスの入室に成功すると、入室したルームのルーム情報を設定した JoinCustomerDeviceResponse が返ります。
	JoinCustomerDevice(context.Context, *JoinCustomerDeviceRequest) (*JoinCustomerDeviceResponse, error)
	// ルームからのカスタマーデバイスの退室を行います。
	// カスタマーデバイスを退室させるルームのルームIDを指定した LeaveCustomerDeviceRequest を渡します。
	// カスタマーデバイスの退室に成功すると、退室したルームのルーム情報を設定した LeaveCustomerDeviceResponse が返ります。
	LeaveCustomerDevice(context.Context, *LeaveCustomerDeviceRequest) (*LeaveCustomerDeviceResponse, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) GetAllRoomGroups(context.Context, *GetAllRoomGroupsRequest) (*GetAllRoomGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoomGroups not implemented")
}
func (UnimplementedRoomServiceServer) GetRoomGroups(context.Context, *GetRoomGroupsRequest) (*GetRoomGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomGroups not implemented")
}
func (UnimplementedRoomServiceServer) GetRoomGroup(context.Context, *GetRoomGroupRequest) (*GetRoomGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomGroup not implemented")
}
func (UnimplementedRoomServiceServer) CreateRoomGroup(context.Context, *CreateRoomGroupRequest) (*CreateRoomGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoomGroup not implemented")
}
func (UnimplementedRoomServiceServer) UpdateRoomGroup(context.Context, *UpdateRoomGroupRequest) (*UpdateRoomGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomGroup not implemented")
}
func (UnimplementedRoomServiceServer) DeleteRoomGroup(context.Context, *DeleteRoomGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoomGroup not implemented")
}
func (UnimplementedRoomServiceServer) GetAllRooms(context.Context, *GetAllRoomsRequest) (*GetAllRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRooms not implemented")
}
func (UnimplementedRoomServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedRoomServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedRoomServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedRoomServiceServer) JoinCustomerDevice(context.Context, *JoinCustomerDeviceRequest) (*JoinCustomerDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCustomerDevice not implemented")
}
func (UnimplementedRoomServiceServer) LeaveCustomerDevice(context.Context, *LeaveCustomerDeviceRequest) (*LeaveCustomerDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCustomerDevice not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_GetAllRoomGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRoomGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetAllRoomGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetAllRoomGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetAllRoomGroups(ctx, req.(*GetAllRoomGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRoomGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoomGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetRoomGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoomGroups(ctx, req.(*GetRoomGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRoomGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoomGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetRoomGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoomGroup(ctx, req.(*GetRoomGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_CreateRoomGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoomGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/CreateRoomGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoomGroup(ctx, req.(*CreateRoomGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdateRoomGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdateRoomGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/UpdateRoomGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdateRoomGroup(ctx, req.(*UpdateRoomGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_DeleteRoomGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).DeleteRoomGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/DeleteRoomGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).DeleteRoomGroup(ctx, req.(*DeleteRoomGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetAllRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetAllRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetAllRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetAllRooms(ctx, req.(*GetAllRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/UpdateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_JoinCustomerDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinCustomerDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).JoinCustomerDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/JoinCustomerDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).JoinCustomerDevice(ctx, req.(*JoinCustomerDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_LeaveCustomerDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveCustomerDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).LeaveCustomerDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.RoomService/LeaveCustomerDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).LeaveCustomerDevice(ctx, req.(*LeaveCustomerDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "at.core.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRoomGroups",
			Handler:    _RoomService_GetAllRoomGroups_Handler,
		},
		{
			MethodName: "GetRoomGroups",
			Handler:    _RoomService_GetRoomGroups_Handler,
		},
		{
			MethodName: "GetRoomGroup",
			Handler:    _RoomService_GetRoomGroup_Handler,
		},
		{
			MethodName: "CreateRoomGroup",
			Handler:    _RoomService_CreateRoomGroup_Handler,
		},
		{
			MethodName: "UpdateRoomGroup",
			Handler:    _RoomService_UpdateRoomGroup_Handler,
		},
		{
			MethodName: "DeleteRoomGroup",
			Handler:    _RoomService_DeleteRoomGroup_Handler,
		},
		{
			MethodName: "GetAllRooms",
			Handler:    _RoomService_GetAllRooms_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _RoomService_GetRooms_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _RoomService_GetRoom_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _RoomService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _RoomService_DeleteRoom_Handler,
		},
		{
			MethodName: "JoinCustomerDevice",
			Handler:    _RoomService_JoinCustomerDevice_Handler,
		},
		{
			MethodName: "LeaveCustomerDevice",
			Handler:    _RoomService_LeaveCustomerDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/room.proto",
}
