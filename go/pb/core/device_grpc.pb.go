// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: core/device.proto

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

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// デバイスの登録要求を行います。
	// デバイスタイプを指定した RequestDeviceRegistrationRequest を渡します。
	// デバイスの登録要求に成功すると、デバイス登録トークン及びワンタイムパスワードが設定された RequestDeviceRegistrationResponse が返ります。
	RequestDeviceRegistration(ctx context.Context, in *RequestDeviceRegistrationRequest, opts ...grpc.CallOption) (*RequestDeviceRegistrationResponse, error)
	// デバイスの登録結果を取得します。
	// デバイス登録IDを指定した GetDeviceRegistrationResultRequest を渡します。
	// 登録結果を格納した、GetDeviceRegistrationResultResponse が返ります。
	// デバイスの登録が完了すると登録完了フラグに true に設定され、デバイスIDとデバイス名が設定されます。
	GetDeviceRegistrationResult(ctx context.Context, in *GetDeviceRegistrationResultRequest, opts ...grpc.CallOption) (*GetDeviceRegistrationResultResponse, error)
	// デバイスの登録のキャンセルを行います。
	// デバイス登録IDを指定した CancelDeviceRegistrationRequest を渡します。
	CancelDeviceRegistration(ctx context.Context, in *CancelDeviceRegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// デバイス登録を受け入れます。
	// チームID、デバイス名、デバイス登録トークンまたはワンタイムパスワードを指定した AcceptDeviceRegistrationRequest を渡します。
	// デバイスの作成に成功すると、登録したデバイス情報が設定された、AcceptDeviceRegistrationResponse が返ります。
	AcceptDeviceRegistration(ctx context.Context, in *AcceptDeviceRegistrationRequest, opts ...grpc.CallOption) (*AcceptDeviceRegistrationResponse, error)
	// 新しくデバイスを登録します。
	// チームID、デバイス名、デバイスタイプを指定した、RegisterDeviceRequest を渡します。
	// デバイスの登録に成功すると、登録したデバイス情報が設定された、RegisterDeviceResponse が返ります。
	// チームID、デバイス名、ワンタイムパスワードを指定した DeviceRegistrationRequest を渡します。
	RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error)
	// デバイスのリストを取得します。
	// 取得するデバイスのチームIDを指定した DeviceListRequest を渡します。
	// 指定されたチームIDのデバイスのリストが設定された Devices が返ります。
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error)
	// デバイスを取得します。
	// 取得するデバイスのデバイスIDを指定した DeviceRequest を渡します。
	// デバイスが存在する場合、Device が返ります。
	GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	// デバイスを更新します。
	// 更新するデバイスのデバイスIDと、新しいデバイス名を指定した DeviceUpdateRequest を渡します。
	// デバイスの作成に成功すると、Device が返ります。
	UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error)
	// デバイスを削除します。
	// 削除するデバイスのデバイスIDを指定した DeviceRequest を渡します。
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) RequestDeviceRegistration(ctx context.Context, in *RequestDeviceRegistrationRequest, opts ...grpc.CallOption) (*RequestDeviceRegistrationResponse, error) {
	out := new(RequestDeviceRegistrationResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/RequestDeviceRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceRegistrationResult(ctx context.Context, in *GetDeviceRegistrationResultRequest, opts ...grpc.CallOption) (*GetDeviceRegistrationResultResponse, error) {
	out := new(GetDeviceRegistrationResultResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/GetDeviceRegistrationResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) CancelDeviceRegistration(ctx context.Context, in *CancelDeviceRegistrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/CancelDeviceRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) AcceptDeviceRegistration(ctx context.Context, in *AcceptDeviceRegistrationRequest, opts ...grpc.CallOption) (*AcceptDeviceRegistrationResponse, error) {
	out := new(AcceptDeviceRegistrationResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/AcceptDeviceRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error) {
	out := new(RegisterDeviceResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/GetDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDevice(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) UpdateDevice(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error) {
	out := new(UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.DeviceService/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	// デバイスの登録要求を行います。
	// デバイスタイプを指定した RequestDeviceRegistrationRequest を渡します。
	// デバイスの登録要求に成功すると、デバイス登録トークン及びワンタイムパスワードが設定された RequestDeviceRegistrationResponse が返ります。
	RequestDeviceRegistration(context.Context, *RequestDeviceRegistrationRequest) (*RequestDeviceRegistrationResponse, error)
	// デバイスの登録結果を取得します。
	// デバイス登録IDを指定した GetDeviceRegistrationResultRequest を渡します。
	// 登録結果を格納した、GetDeviceRegistrationResultResponse が返ります。
	// デバイスの登録が完了すると登録完了フラグに true に設定され、デバイスIDとデバイス名が設定されます。
	GetDeviceRegistrationResult(context.Context, *GetDeviceRegistrationResultRequest) (*GetDeviceRegistrationResultResponse, error)
	// デバイスの登録のキャンセルを行います。
	// デバイス登録IDを指定した CancelDeviceRegistrationRequest を渡します。
	CancelDeviceRegistration(context.Context, *CancelDeviceRegistrationRequest) (*emptypb.Empty, error)
	// デバイス登録を受け入れます。
	// チームID、デバイス名、デバイス登録トークンまたはワンタイムパスワードを指定した AcceptDeviceRegistrationRequest を渡します。
	// デバイスの作成に成功すると、登録したデバイス情報が設定された、AcceptDeviceRegistrationResponse が返ります。
	AcceptDeviceRegistration(context.Context, *AcceptDeviceRegistrationRequest) (*AcceptDeviceRegistrationResponse, error)
	// 新しくデバイスを登録します。
	// チームID、デバイス名、デバイスタイプを指定した、RegisterDeviceRequest を渡します。
	// デバイスの登録に成功すると、登録したデバイス情報が設定された、RegisterDeviceResponse が返ります。
	// チームID、デバイス名、ワンタイムパスワードを指定した DeviceRegistrationRequest を渡します。
	RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error)
	// デバイスのリストを取得します。
	// 取得するデバイスのチームIDを指定した DeviceListRequest を渡します。
	// 指定されたチームIDのデバイスのリストが設定された Devices が返ります。
	GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error)
	// デバイスを取得します。
	// 取得するデバイスのデバイスIDを指定した DeviceRequest を渡します。
	// デバイスが存在する場合、Device が返ります。
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	// デバイスを更新します。
	// 更新するデバイスのデバイスIDと、新しいデバイス名を指定した DeviceUpdateRequest を渡します。
	// デバイスの作成に成功すると、Device が返ります。
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	// デバイスを削除します。
	// 削除するデバイスのデバイスIDを指定した DeviceRequest を渡します。
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) RequestDeviceRegistration(context.Context, *RequestDeviceRegistrationRequest) (*RequestDeviceRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDeviceRegistration not implemented")
}
func (UnimplementedDeviceServiceServer) GetDeviceRegistrationResult(context.Context, *GetDeviceRegistrationResultRequest) (*GetDeviceRegistrationResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceRegistrationResult not implemented")
}
func (UnimplementedDeviceServiceServer) CancelDeviceRegistration(context.Context, *CancelDeviceRegistrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDeviceRegistration not implemented")
}
func (UnimplementedDeviceServiceServer) AcceptDeviceRegistration(context.Context, *AcceptDeviceRegistrationRequest) (*AcceptDeviceRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDeviceRegistration not implemented")
}
func (UnimplementedDeviceServiceServer) RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedDeviceServiceServer) GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedDeviceServiceServer) GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedDeviceServiceServer) UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedDeviceServiceServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_RequestDeviceRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeviceRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RequestDeviceRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/RequestDeviceRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RequestDeviceRegistration(ctx, req.(*RequestDeviceRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceRegistrationResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRegistrationResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceRegistrationResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/GetDeviceRegistrationResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceRegistrationResult(ctx, req.(*GetDeviceRegistrationResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_CancelDeviceRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelDeviceRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).CancelDeviceRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/CancelDeviceRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).CancelDeviceRegistration(ctx, req.(*CancelDeviceRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_AcceptDeviceRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptDeviceRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).AcceptDeviceRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/AcceptDeviceRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).AcceptDeviceRegistration(ctx, req.(*AcceptDeviceRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, req.(*RegisterDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDevices(ctx, req.(*GetDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDevice(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).UpdateDevice(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.DeviceService/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "at.core.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDeviceRegistration",
			Handler:    _DeviceService_RequestDeviceRegistration_Handler,
		},
		{
			MethodName: "GetDeviceRegistrationResult",
			Handler:    _DeviceService_GetDeviceRegistrationResult_Handler,
		},
		{
			MethodName: "CancelDeviceRegistration",
			Handler:    _DeviceService_CancelDeviceRegistration_Handler,
		},
		{
			MethodName: "AcceptDeviceRegistration",
			Handler:    _DeviceService_AcceptDeviceRegistration_Handler,
		},
		{
			MethodName: "RegisterDevice",
			Handler:    _DeviceService_RegisterDevice_Handler,
		},
		{
			MethodName: "GetDevices",
			Handler:    _DeviceService_GetDevices_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _DeviceService_GetDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _DeviceService_UpdateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _DeviceService_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/device.proto",
}
