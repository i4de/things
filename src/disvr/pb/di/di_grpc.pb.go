// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/di.proto

package di

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeviceMsgClient is the client API for DeviceMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceMsgClient interface {
	//获取设备sdk调试日志
	SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error)
	//获取设备数据信息
	SchemaLatestIndex(ctx context.Context, in *SchemaLatestIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error)
	//获取设备数据信息
	SchemaLogIndex(ctx context.Context, in *SchemaLogIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error)
}

type deviceMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceMsgClient(cc grpc.ClientConnInterface) DeviceMsgClient {
	return &deviceMsgClient{cc}
}

func (c *deviceMsgClient) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	out := new(SdkLogIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceMsg/sdkLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMsgClient) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	out := new(HubLogIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceMsg/hubLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMsgClient) SchemaLatestIndex(ctx context.Context, in *SchemaLatestIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	out := new(SchemaIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceMsg/schemaLatestIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceMsgClient) SchemaLogIndex(ctx context.Context, in *SchemaLogIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	out := new(SchemaIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceMsg/schemaLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceMsgServer is the server API for DeviceMsg service.
// All implementations must embed UnimplementedDeviceMsgServer
// for forward compatibility
type DeviceMsgServer interface {
	//获取设备sdk调试日志
	SdkLogIndex(context.Context, *SdkLogIndexReq) (*SdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	HubLogIndex(context.Context, *HubLogIndexReq) (*HubLogIndexResp, error)
	//获取设备数据信息
	SchemaLatestIndex(context.Context, *SchemaLatestIndexReq) (*SchemaIndexResp, error)
	//获取设备数据信息
	SchemaLogIndex(context.Context, *SchemaLogIndexReq) (*SchemaIndexResp, error)
	mustEmbedUnimplementedDeviceMsgServer()
}

// UnimplementedDeviceMsgServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceMsgServer struct {
}

func (UnimplementedDeviceMsgServer) SdkLogIndex(context.Context, *SdkLogIndexReq) (*SdkLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SdkLogIndex not implemented")
}
func (UnimplementedDeviceMsgServer) HubLogIndex(context.Context, *HubLogIndexReq) (*HubLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HubLogIndex not implemented")
}
func (UnimplementedDeviceMsgServer) SchemaLatestIndex(context.Context, *SchemaLatestIndexReq) (*SchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaLatestIndex not implemented")
}
func (UnimplementedDeviceMsgServer) SchemaLogIndex(context.Context, *SchemaLogIndexReq) (*SchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchemaLogIndex not implemented")
}
func (UnimplementedDeviceMsgServer) mustEmbedUnimplementedDeviceMsgServer() {}

// UnsafeDeviceMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceMsgServer will
// result in compilation errors.
type UnsafeDeviceMsgServer interface {
	mustEmbedUnimplementedDeviceMsgServer()
}

func RegisterDeviceMsgServer(s grpc.ServiceRegistrar, srv DeviceMsgServer) {
	s.RegisterService(&DeviceMsg_ServiceDesc, srv)
}

func _DeviceMsg_SdkLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SdkLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMsgServer).SdkLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceMsg/sdkLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMsgServer).SdkLogIndex(ctx, req.(*SdkLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMsg_HubLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMsgServer).HubLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceMsg/hubLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMsgServer).HubLogIndex(ctx, req.(*HubLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMsg_SchemaLatestIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaLatestIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMsgServer).SchemaLatestIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceMsg/schemaLatestIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMsgServer).SchemaLatestIndex(ctx, req.(*SchemaLatestIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceMsg_SchemaLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceMsgServer).SchemaLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceMsg/schemaLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceMsgServer).SchemaLogIndex(ctx, req.(*SchemaLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceMsg_ServiceDesc is the grpc.ServiceDesc for DeviceMsg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceMsg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "di.DeviceMsg",
	HandlerType: (*DeviceMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sdkLogIndex",
			Handler:    _DeviceMsg_SdkLogIndex_Handler,
		},
		{
			MethodName: "hubLogIndex",
			Handler:    _DeviceMsg_HubLogIndex_Handler,
		},
		{
			MethodName: "schemaLatestIndex",
			Handler:    _DeviceMsg_SchemaLatestIndex_Handler,
		},
		{
			MethodName: "schemaLogIndex",
			Handler:    _DeviceMsg_SchemaLogIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/di.proto",
}

// DeviceInteractClient is the client API for DeviceInteract service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceInteractClient interface {
	//同步调用设备行为
	SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error)
	//同步调用设备属性
	SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
}

type deviceInteractClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceInteractClient(cc grpc.ClientConnInterface) DeviceInteractClient {
	return &deviceInteractClient{cc}
}

func (c *deviceInteractClient) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	out := new(SendActionResp)
	err := c.cc.Invoke(ctx, "/di.DeviceInteract/sendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceInteractClient) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	out := new(SendPropertyResp)
	err := c.cc.Invoke(ctx, "/di.DeviceInteract/sendProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceInteractServer is the server API for DeviceInteract service.
// All implementations must embed UnimplementedDeviceInteractServer
// for forward compatibility
type DeviceInteractServer interface {
	//同步调用设备行为
	SendAction(context.Context, *SendActionReq) (*SendActionResp, error)
	//同步调用设备属性
	SendProperty(context.Context, *SendPropertyReq) (*SendPropertyResp, error)
	mustEmbedUnimplementedDeviceInteractServer()
}

// UnimplementedDeviceInteractServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceInteractServer struct {
}

func (UnimplementedDeviceInteractServer) SendAction(context.Context, *SendActionReq) (*SendActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAction not implemented")
}
func (UnimplementedDeviceInteractServer) SendProperty(context.Context, *SendPropertyReq) (*SendPropertyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProperty not implemented")
}
func (UnimplementedDeviceInteractServer) mustEmbedUnimplementedDeviceInteractServer() {}

// UnsafeDeviceInteractServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceInteractServer will
// result in compilation errors.
type UnsafeDeviceInteractServer interface {
	mustEmbedUnimplementedDeviceInteractServer()
}

func RegisterDeviceInteractServer(s grpc.ServiceRegistrar, srv DeviceInteractServer) {
	s.RegisterService(&DeviceInteract_ServiceDesc, srv)
}

func _DeviceInteract_SendAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceInteractServer).SendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceInteract/sendAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceInteractServer).SendAction(ctx, req.(*SendActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceInteract_SendProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceInteractServer).SendProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceInteract/sendProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceInteractServer).SendProperty(ctx, req.(*SendPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceInteract_ServiceDesc is the grpc.ServiceDesc for DeviceInteract service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceInteract_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "di.DeviceInteract",
	HandlerType: (*DeviceInteractServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendAction",
			Handler:    _DeviceInteract_SendAction_Handler,
		},
		{
			MethodName: "sendProperty",
			Handler:    _DeviceInteract_SendProperty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/di.proto",
}
