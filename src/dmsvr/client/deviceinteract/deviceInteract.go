// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package deviceinteract

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq               = dm.AccessAuthReq
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceRegisterReq           = dm.DeviceRegisterReq
	DeviceRegisterResp          = dm.DeviceRegisterResp
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	EventIndex                  = dm.EventIndex
	EventIndexResp              = dm.EventIndexResp
	EventLogIndexReq            = dm.EventLogIndexReq
	Firmware                    = dm.Firmware
	FirmwareFile                = dm.FirmwareFile
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GetPropertyReplyReq         = dm.GetPropertyReplyReq
	GetPropertyReplyResp        = dm.GetPropertyReplyResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq   = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	HubLogIndex                 = dm.HubLogIndex
	HubLogIndexReq              = dm.HubLogIndexReq
	HubLogIndexResp             = dm.HubLogIndexResp
	LoginAuthReq                = dm.LoginAuthReq
	MultiSendPropertyReq        = dm.MultiSendPropertyReq
	MultiSendPropertyResp       = dm.MultiSendPropertyResp
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareCreateReq        = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq        = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq         = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp        = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo             = dm.OtaFirmwareInfo
	OtaFirmwareReadReq          = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp         = dm.OtaFirmwareReadResp
	OtaFirmwareResp             = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq        = dm.OtaFirmwareUpdateReq
	OtaPageInfo                 = dm.OtaPageInfo
	OtaPromptIndexReq           = dm.OtaPromptIndexReq
	OtaPromptIndexResp          = dm.OtaPromptIndexResp
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	PropertyIndex               = dm.PropertyIndex
	PropertyIndexResp           = dm.PropertyIndexResp
	PropertyLatestIndexReq      = dm.PropertyLatestIndexReq
	PropertyLogIndexReq         = dm.PropertyLogIndexReq
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	RespActionReq               = dm.RespActionReq
	RespReadReq                 = dm.RespReadReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq
	SdkLogIndex                 = dm.SdkLogIndex
	SdkLogIndexReq              = dm.SdkLogIndexReq
	SdkLogIndexResp             = dm.SdkLogIndexResp
	SendActionReq               = dm.SendActionReq
	SendActionResp              = dm.SendActionResp
	SendMsgReq                  = dm.SendMsgReq
	SendMsgResp                 = dm.SendMsgResp
	SendOption                  = dm.SendOption
	SendPropertyMsg             = dm.SendPropertyMsg
	SendPropertyReq             = dm.SendPropertyReq
	SendPropertyResp            = dm.SendPropertyResp
	ShadowIndex                 = dm.ShadowIndex
	ShadowIndexResp             = dm.ShadowIndexResp

	DeviceInteract interface {
		// 调用设备行为
		SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error)
		// 获取异步调用设备行为的结果
		ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendActionResp, error)
		// 回复调用设备行为
		RespAction(ctx context.Context, in *RespActionReq, opts ...grpc.CallOption) (*Response, error)
		// 请求设备获取设备最新属性
		GetPropertyReply(ctx context.Context, in *GetPropertyReplyReq, opts ...grpc.CallOption) (*GetPropertyReplyResp, error)
		// 调用设备属性
		SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
		// 批量调用设备属性
		MultiSendProperty(ctx context.Context, in *MultiSendPropertyReq, opts ...grpc.CallOption) (*MultiSendPropertyResp, error)
		// 获取异步调用设备属性的结果
		PropertyRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
		// 发送消息给设备
		SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
	}

	defaultDeviceInteract struct {
		cli zrpc.Client
	}

	directDeviceInteract struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceInteractServer
	}
)

func NewDeviceInteract(cli zrpc.Client) DeviceInteract {
	return &defaultDeviceInteract{
		cli: cli,
	}
}

func NewDirectDeviceInteract(svcCtx *svc.ServiceContext, svr dm.DeviceInteractServer) DeviceInteract {
	return &directDeviceInteract{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 调用设备行为
func (m *defaultDeviceInteract) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.SendAction(ctx, in, opts...)
}

// 调用设备行为
func (d *directDeviceInteract) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	return d.svr.SendAction(ctx, in)
}

// 获取异步调用设备行为的结果
func (m *defaultDeviceInteract) ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.ActionRead(ctx, in, opts...)
}

// 获取异步调用设备行为的结果
func (d *directDeviceInteract) ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	return d.svr.ActionRead(ctx, in)
}

// 回复调用设备行为
func (m *defaultDeviceInteract) RespAction(ctx context.Context, in *RespActionReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.RespAction(ctx, in, opts...)
}

// 回复调用设备行为
func (d *directDeviceInteract) RespAction(ctx context.Context, in *RespActionReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RespAction(ctx, in)
}

// 请求设备获取设备最新属性
func (m *defaultDeviceInteract) GetPropertyReply(ctx context.Context, in *GetPropertyReplyReq, opts ...grpc.CallOption) (*GetPropertyReplyResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.GetPropertyReply(ctx, in, opts...)
}

// 请求设备获取设备最新属性
func (d *directDeviceInteract) GetPropertyReply(ctx context.Context, in *GetPropertyReplyReq, opts ...grpc.CallOption) (*GetPropertyReplyResp, error) {
	return d.svr.GetPropertyReply(ctx, in)
}

// 调用设备属性
func (m *defaultDeviceInteract) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.SendProperty(ctx, in, opts...)
}

// 调用设备属性
func (d *directDeviceInteract) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	return d.svr.SendProperty(ctx, in)
}

// 批量调用设备属性
func (m *defaultDeviceInteract) MultiSendProperty(ctx context.Context, in *MultiSendPropertyReq, opts ...grpc.CallOption) (*MultiSendPropertyResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.MultiSendProperty(ctx, in, opts...)
}

// 批量调用设备属性
func (d *directDeviceInteract) MultiSendProperty(ctx context.Context, in *MultiSendPropertyReq, opts ...grpc.CallOption) (*MultiSendPropertyResp, error) {
	return d.svr.MultiSendProperty(ctx, in)
}

// 获取异步调用设备属性的结果
func (m *defaultDeviceInteract) PropertyRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.PropertyRead(ctx, in, opts...)
}

// 获取异步调用设备属性的结果
func (d *directDeviceInteract) PropertyRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	return d.svr.PropertyRead(ctx, in)
}

// 发送消息给设备
func (m *defaultDeviceInteract) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.SendMsg(ctx, in, opts...)
}

// 发送消息给设备
func (d *directDeviceInteract) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	return d.svr.SendMsg(ctx, in)
}
