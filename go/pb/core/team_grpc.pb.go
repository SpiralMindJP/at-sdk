// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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
const _ = grpc.SupportPackageIsVersion7

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	// チームを取得します。
	// TeamListRequest を渡します。
	// チームのリストが設定された Teams が返ります。
	List(ctx context.Context, in *TeamListRequest, opts ...grpc.CallOption) (*Teams, error)
	// チームを取得します。
	// 取得するチームのチームIDを指定した TeamRequest を渡します。
	// チームが存在する場合、Team が返ります。
	Get(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error)
	// 新しくチームを作成します。
	// チーム名を指定した TeamCreateRequest を渡します。
	// チームの作成に成功すると、チームIDが設定さた Team が返ります。
	Create(ctx context.Context, in *TeamCreateRequest, opts ...grpc.CallOption) (*Team, error)
	// チームを更新します。
	// 更新するチームのチームIDと、新しいチーム名を指定した TeamUpdateRequest を渡します。
	// チームの作成に成功すると、Team が返ります。
	Update(ctx context.Context, in *TeamUpdateRequest, opts ...grpc.CallOption) (*Team, error)
	// チームを削除します。
	// 削除するチームのチームIDを指定した TeamRequest を渡します。
	Delete(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) List(ctx context.Context, in *TeamListRequest, opts ...grpc.CallOption) (*Teams, error) {
	out := new(Teams)
	err := c.cc.Invoke(ctx, "/at_core_service.TeamService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Get(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, "/at_core_service.TeamService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Create(ctx context.Context, in *TeamCreateRequest, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, "/at_core_service.TeamService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Update(ctx context.Context, in *TeamUpdateRequest, opts ...grpc.CallOption) (*Team, error) {
	out := new(Team)
	err := c.cc.Invoke(ctx, "/at_core_service.TeamService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Delete(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at_core_service.TeamService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	// チームを取得します。
	// TeamListRequest を渡します。
	// チームのリストが設定された Teams が返ります。
	List(context.Context, *TeamListRequest) (*Teams, error)
	// チームを取得します。
	// 取得するチームのチームIDを指定した TeamRequest を渡します。
	// チームが存在する場合、Team が返ります。
	Get(context.Context, *TeamRequest) (*Team, error)
	// 新しくチームを作成します。
	// チーム名を指定した TeamCreateRequest を渡します。
	// チームの作成に成功すると、チームIDが設定さた Team が返ります。
	Create(context.Context, *TeamCreateRequest) (*Team, error)
	// チームを更新します。
	// 更新するチームのチームIDと、新しいチーム名を指定した TeamUpdateRequest を渡します。
	// チームの作成に成功すると、Team が返ります。
	Update(context.Context, *TeamUpdateRequest) (*Team, error)
	// チームを削除します。
	// 削除するチームのチームIDを指定した TeamRequest を渡します。
	Delete(context.Context, *TeamRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) List(context.Context, *TeamListRequest) (*Teams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTeamServiceServer) Get(context.Context, *TeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTeamServiceServer) Create(context.Context, *TeamCreateRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTeamServiceServer) Update(context.Context, *TeamUpdateRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTeamServiceServer) Delete(context.Context, *TeamRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	s.RegisterService(&_TeamService_serviceDesc, srv)
}

func _TeamService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at_core_service.TeamService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).List(ctx, req.(*TeamListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at_core_service.TeamService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Get(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at_core_service.TeamService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Create(ctx, req.(*TeamCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at_core_service.TeamService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Update(ctx, req.(*TeamUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at_core_service.TeamService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Delete(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "at_core_service.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TeamService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TeamService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TeamService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TeamService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TeamService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team.proto",
}
