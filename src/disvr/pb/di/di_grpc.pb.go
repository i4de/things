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

// DeviceLogClient is the client API for DeviceLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceLogClient interface {
	//获取设备sdk调试日志
	DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error)
	//获取设备数据信息
	DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
	//获取设备数据信息
	DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
}

type deviceLogClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceLogClient(cc grpc.ClientConnInterface) DeviceLogClient {
	return &deviceLogClient{cc}
}

func (c *deviceLogClient) DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error) {
	out := new(DataSdkLogIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceLog/dataSdkLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceLogClient) DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error) {
	out := new(DataHubLogIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceLog/dataHubLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceLogClient) DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	out := new(DataSchemaIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceLog/dataSchemaLatestIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceLogClient) DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	out := new(DataSchemaIndexResp)
	err := c.cc.Invoke(ctx, "/di.DeviceLog/dataSchemaLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceLogServer is the server API for DeviceLog service.
// All implementations must embed UnimplementedDeviceLogServer
// for forward compatibility
type DeviceLogServer interface {
	//获取设备sdk调试日志
	DataSdkLogIndex(context.Context, *DataSdkLogIndexReq) (*DataSdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	DataHubLogIndex(context.Context, *DataHubLogIndexReq) (*DataHubLogIndexResp, error)
	//获取设备数据信息
	DataSchemaLatestIndex(context.Context, *DataSchemaLatestIndexReq) (*DataSchemaIndexResp, error)
	//获取设备数据信息
	DataSchemaLogIndex(context.Context, *DataSchemaLogIndexReq) (*DataSchemaIndexResp, error)
	mustEmbedUnimplementedDeviceLogServer()
}

// UnimplementedDeviceLogServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceLogServer struct {
}

func (UnimplementedDeviceLogServer) DataSdkLogIndex(context.Context, *DataSdkLogIndexReq) (*DataSdkLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSdkLogIndex not implemented")
}
func (UnimplementedDeviceLogServer) DataHubLogIndex(context.Context, *DataHubLogIndexReq) (*DataHubLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataHubLogIndex not implemented")
}
func (UnimplementedDeviceLogServer) DataSchemaLatestIndex(context.Context, *DataSchemaLatestIndexReq) (*DataSchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSchemaLatestIndex not implemented")
}
func (UnimplementedDeviceLogServer) DataSchemaLogIndex(context.Context, *DataSchemaLogIndexReq) (*DataSchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSchemaLogIndex not implemented")
}
func (UnimplementedDeviceLogServer) mustEmbedUnimplementedDeviceLogServer() {}

// UnsafeDeviceLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceLogServer will
// result in compilation errors.
type UnsafeDeviceLogServer interface {
	mustEmbedUnimplementedDeviceLogServer()
}

func RegisterDeviceLogServer(s grpc.ServiceRegistrar, srv DeviceLogServer) {
	s.RegisterService(&DeviceLog_ServiceDesc, srv)
}

func _DeviceLog_DataSdkLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSdkLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceLogServer).DataSdkLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceLog/dataSdkLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceLogServer).DataSdkLogIndex(ctx, req.(*DataSdkLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceLog_DataHubLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataHubLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceLogServer).DataHubLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceLog/dataHubLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceLogServer).DataHubLogIndex(ctx, req.(*DataHubLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceLog_DataSchemaLatestIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSchemaLatestIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceLogServer).DataSchemaLatestIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceLog/dataSchemaLatestIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceLogServer).DataSchemaLatestIndex(ctx, req.(*DataSchemaLatestIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceLog_DataSchemaLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSchemaLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceLogServer).DataSchemaLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di.DeviceLog/dataSchemaLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceLogServer).DataSchemaLogIndex(ctx, req.(*DataSchemaLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceLog_ServiceDesc is the grpc.ServiceDesc for DeviceLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "di.DeviceLog",
	HandlerType: (*DeviceLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "dataSdkLogIndex",
			Handler:    _DeviceLog_DataSdkLogIndex_Handler,
		},
		{
			MethodName: "dataHubLogIndex",
			Handler:    _DeviceLog_DataHubLogIndex_Handler,
		},
		{
			MethodName: "dataSchemaLatestIndex",
			Handler:    _DeviceLog_DataSchemaLatestIndex_Handler,
		},
		{
			MethodName: "dataSchemaLogIndex",
			Handler:    _DeviceLog_DataSchemaLogIndex_Handler,
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
