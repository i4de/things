// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package otamanage

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AbnormalLogIndexReq               = dm.AbnormalLogIndexReq
	AbnormalLogIndexResp              = dm.AbnormalLogIndexResp
	AbnormalLogInfo                   = dm.AbnormalLogInfo
	ActionRespReq                     = dm.ActionRespReq
	ActionSendReq                     = dm.ActionSendReq
	ActionSendResp                    = dm.ActionSendResp
	CommonSchemaCreateReq             = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq              = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp             = dm.CommonSchemaIndexResp
	CommonSchemaInfo                  = dm.CommonSchemaInfo
	CommonSchemaUpdateReq             = dm.CommonSchemaUpdateReq
	CompareInt64                      = dm.CompareInt64
	CompareString                     = dm.CompareString
	CustomTopic                       = dm.CustomTopic
	DeviceCore                        = dm.DeviceCore
	DeviceCountInfo                   = dm.DeviceCountInfo
	DeviceCountReq                    = dm.DeviceCountReq
	DeviceCountResp                   = dm.DeviceCountResp
	DeviceError                       = dm.DeviceError
	DeviceGatewayBindDevice           = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq             = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp            = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq       = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiSaveReq         = dm.DeviceGatewayMultiSaveReq
	DeviceGatewaySign                 = dm.DeviceGatewaySign
	DeviceGroupMultiSaveReq           = dm.DeviceGroupMultiSaveReq
	DeviceInfo                        = dm.DeviceInfo
	DeviceInfoBindReq                 = dm.DeviceInfoBindReq
	DeviceInfoCanBindReq              = dm.DeviceInfoCanBindReq
	DeviceInfoCount                   = dm.DeviceInfoCount
	DeviceInfoCountReq                = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq               = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp               = dm.DeviceInfoIndexResp
	DeviceInfoMultiBindReq            = dm.DeviceInfoMultiBindReq
	DeviceInfoMultiBindResp           = dm.DeviceInfoMultiBindResp
	DeviceInfoMultiUpdateReq          = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                 = dm.DeviceInfoReadReq
	DeviceModuleVersion               = dm.DeviceModuleVersion
	DeviceModuleVersionIndexReq       = dm.DeviceModuleVersionIndexReq
	DeviceModuleVersionIndexResp      = dm.DeviceModuleVersionIndexResp
	DeviceModuleVersionReadReq        = dm.DeviceModuleVersionReadReq
	DeviceMoveReq                     = dm.DeviceMoveReq
	DeviceOnlineMultiFix              = dm.DeviceOnlineMultiFix
	DeviceOnlineMultiFixReq           = dm.DeviceOnlineMultiFixReq
	DeviceProfile                     = dm.DeviceProfile
	DeviceProfileIndexReq             = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp            = dm.DeviceProfileIndexResp
	DeviceProfileReadReq              = dm.DeviceProfileReadReq
	DeviceResetReq                    = dm.DeviceResetReq
	DeviceSchema                      = dm.DeviceSchema
	DeviceSchemaIndexReq              = dm.DeviceSchemaIndexReq
	DeviceSchemaIndexResp             = dm.DeviceSchemaIndexResp
	DeviceSchemaMultiCreateReq        = dm.DeviceSchemaMultiCreateReq
	DeviceSchemaMultiDeleteReq        = dm.DeviceSchemaMultiDeleteReq
	DeviceSchemaTslReadReq            = dm.DeviceSchemaTslReadReq
	DeviceSchemaTslReadResp           = dm.DeviceSchemaTslReadResp
	DeviceShareInfo                   = dm.DeviceShareInfo
	DeviceTransferReq                 = dm.DeviceTransferReq
	DeviceTypeCountReq                = dm.DeviceTypeCountReq
	DeviceTypeCountResp               = dm.DeviceTypeCountResp
	EdgeSendReq                       = dm.EdgeSendReq
	EdgeSendResp                      = dm.EdgeSendResp
	Empty                             = dm.Empty
	EventLogIndexReq                  = dm.EventLogIndexReq
	EventLogIndexResp                 = dm.EventLogIndexResp
	EventLogInfo                      = dm.EventLogInfo
	Firmware                          = dm.Firmware
	FirmwareFile                      = dm.FirmwareFile
	FirmwareInfo                      = dm.FirmwareInfo
	FirmwareInfoDeleteReq             = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp            = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq              = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp             = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq               = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp              = dm.FirmwareInfoReadResp
	FirmwareResp                      = dm.FirmwareResp
	GatewayCanBindIndexReq            = dm.GatewayCanBindIndexReq
	GatewayCanBindIndexResp           = dm.GatewayCanBindIndexResp
	GatewayGetFoundReq                = dm.GatewayGetFoundReq
	GatewayNotifyBindSendReq          = dm.GatewayNotifyBindSendReq
	GroupCore                         = dm.GroupCore
	GroupDeviceMultiDeleteReq         = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq           = dm.GroupDeviceMultiSaveReq
	GroupInfo                         = dm.GroupInfo
	GroupInfoCreateReq                = dm.GroupInfoCreateReq
	GroupInfoIndexReq                 = dm.GroupInfoIndexReq
	GroupInfoIndexResp                = dm.GroupInfoIndexResp
	GroupInfoMultiCreateReq           = dm.GroupInfoMultiCreateReq
	GroupInfoReadReq                  = dm.GroupInfoReadReq
	GroupInfoUpdateReq                = dm.GroupInfoUpdateReq
	HubLogIndexReq                    = dm.HubLogIndexReq
	HubLogIndexResp                   = dm.HubLogIndexResp
	HubLogInfo                        = dm.HubLogInfo
	IDPath                            = dm.IDPath
	IDPathWithUpdate                  = dm.IDPathWithUpdate
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceConfirmReq       = dm.OtaFirmwareDeviceConfirmReq
	OtaFirmwareDeviceIndexReq         = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp        = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo             = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq         = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                   = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq           = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp          = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo               = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp               = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                   = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq          = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq           = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp          = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq          = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq            = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp           = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq            = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                 = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                  = dm.OtaJobStaticInfo
	OtaModuleInfo                     = dm.OtaModuleInfo
	OtaModuleInfoIndexReq             = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp            = dm.OtaModuleInfoIndexResp
	PageInfo                          = dm.PageInfo
	PageInfo_OrderBy                  = dm.PageInfo_OrderBy
	Point                             = dm.Point
	ProductCategory                   = dm.ProductCategory
	ProductCategoryIndexReq           = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp          = dm.ProductCategoryIndexResp
	ProductCategorySchemaIndexReq     = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp    = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiSaveReq = dm.ProductCategorySchemaMultiSaveReq
	ProductCustom                     = dm.ProductCustom
	ProductCustomReadReq              = dm.ProductCustomReadReq
	ProductCustomUi                   = dm.ProductCustomUi
	ProductInfo                       = dm.ProductInfo
	ProductInfoDeleteReq              = dm.ProductInfoDeleteReq
	ProductInfoIndexReq               = dm.ProductInfoIndexReq
	ProductInfoIndexResp              = dm.ProductInfoIndexResp
	ProductInfoReadReq                = dm.ProductInfoReadReq
	ProductInitReq                    = dm.ProductInitReq
	ProductRemoteConfig               = dm.ProductRemoteConfig
	ProductSchemaCreateReq            = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq            = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq             = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp            = dm.ProductSchemaIndexResp
	ProductSchemaInfo                 = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq       = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq         = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq           = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp          = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq            = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq       = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp      = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg            = dm.PropertyControlSendMsg
	PropertyControlSendReq            = dm.PropertyControlSendReq
	PropertyControlSendResp           = dm.PropertyControlSendResp
	PropertyGetReportSendReq          = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp         = dm.PropertyGetReportSendResp
	PropertyLogIndexReq               = dm.PropertyLogIndexReq
	PropertyLogIndexResp              = dm.PropertyLogIndexResp
	PropertyLogInfo                   = dm.PropertyLogInfo
	PropertyLogLatestIndexReq         = dm.PropertyLogLatestIndexReq
	ProtocolConfigField               = dm.ProtocolConfigField
	ProtocolConfigInfo                = dm.ProtocolConfigInfo
	ProtocolInfo                      = dm.ProtocolInfo
	ProtocolInfoIndexReq              = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp             = dm.ProtocolInfoIndexResp
	ProtocolService                   = dm.ProtocolService
	ProtocolServiceIndexReq           = dm.ProtocolServiceIndexReq
	ProtocolServiceIndexResp          = dm.ProtocolServiceIndexResp
	RemoteConfigCreateReq             = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq              = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp             = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq           = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp          = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq            = dm.RemoteConfigPushAllReq
	RespReadReq                       = dm.RespReadReq
	RootCheckReq                      = dm.RootCheckReq
	SdkLogIndexReq                    = dm.SdkLogIndexReq
	SdkLogIndexResp                   = dm.SdkLogIndexResp
	SdkLogInfo                        = dm.SdkLogInfo
	SendLogIndexReq                   = dm.SendLogIndexReq
	SendLogIndexResp                  = dm.SendLogIndexResp
	SendLogInfo                       = dm.SendLogInfo
	SendMsgReq                        = dm.SendMsgReq
	SendMsgResp                       = dm.SendMsgResp
	SendOption                        = dm.SendOption
	ShadowIndex                       = dm.ShadowIndex
	ShadowIndexResp                   = dm.ShadowIndexResp
	SharePerm                         = dm.SharePerm
	StatusLogIndexReq                 = dm.StatusLogIndexReq
	StatusLogIndexResp                = dm.StatusLogIndexResp
	StatusLogInfo                     = dm.StatusLogInfo
	TimeRange                         = dm.TimeRange
	UserDeviceCollectSave             = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq           = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp          = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo               = dm.UserDeviceShareInfo
	UserDeviceShareMultiAcceptReq     = dm.UserDeviceShareMultiAcceptReq
	UserDeviceShareMultiDeleteReq     = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareMultiInfo          = dm.UserDeviceShareMultiInfo
	UserDeviceShareMultiToken         = dm.UserDeviceShareMultiToken
	UserDeviceShareReadReq            = dm.UserDeviceShareReadReq
	WithID                            = dm.WithID
	WithIDChildren                    = dm.WithIDChildren
	WithIDCode                        = dm.WithIDCode
	WithProfile                       = dm.WithProfile

	OtaManage interface {
		// 添加升级包
		OtaFirmwareInfoCreate(ctx context.Context, in *OtaFirmwareInfoCreateReq, opts ...grpc.CallOption) (*WithID, error)
		// 修改升级包
		OtaFirmwareInfoUpdate(ctx context.Context, in *OtaFirmwareInfoUpdateReq, opts ...grpc.CallOption) (*WithID, error)
		// 删除升级包
		OtaFirmwareInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 升级包列表
		OtaFirmwareInfoIndex(ctx context.Context, in *OtaFirmwareInfoIndexReq, opts ...grpc.CallOption) (*OtaFirmwareInfoIndexResp, error)
		// 查询升级包
		OtaFirmwareInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareInfo, error)
		// 创建静态升级批次
		OtaFirmwareJobCreate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*WithID, error)
		// 获取升级包下的升级任务批次列表
		OtaFirmwareJobIndex(ctx context.Context, in *OtaFirmwareJobIndexReq, opts ...grpc.CallOption) (*OtaFirmwareJobIndexResp, error)
		// 查询指定升级批次的详情
		OtaFirmwareJobRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareJobInfo, error)
		// 取消动态升级策略
		OtaFirmwareJobUpdate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*Empty, error)
		// 查询指定升级批次下的设备升级作业列表
		OtaFirmwareDeviceIndex(ctx context.Context, in *OtaFirmwareDeviceIndexReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceIndexResp, error)
		// 取消指定批次下的设备升级作业
		OtaFirmwareDeviceCancel(ctx context.Context, in *OtaFirmwareDeviceCancelReq, opts ...grpc.CallOption) (*Empty, error)
		// 重新升级指定批次下升级失败或升级取消的设备升级作业
		OtaFirmwareDeviceRetry(ctx context.Context, in *OtaFirmwareDeviceRetryReq, opts ...grpc.CallOption) (*Empty, error)
		// app确认设备升级
		OtaFirmwareDeviceConfirm(ctx context.Context, in *OtaFirmwareDeviceConfirmReq, opts ...grpc.CallOption) (*Empty, error)
		OtaModuleInfoCreate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*WithID, error)
		OtaModuleInfoUpdate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*Empty, error)
		OtaModuleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		OtaModuleInfoIndex(ctx context.Context, in *OtaModuleInfoIndexReq, opts ...grpc.CallOption) (*OtaModuleInfoIndexResp, error)
		OtaModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*OtaModuleInfo, error)
	}

	defaultOtaManage struct {
		cli zrpc.Client
	}

	directOtaManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.OtaManageServer
	}
)

func NewOtaManage(cli zrpc.Client) OtaManage {
	return &defaultOtaManage{
		cli: cli,
	}
}

func NewDirectOtaManage(svcCtx *svc.ServiceContext, svr dm.OtaManageServer) OtaManage {
	return &directOtaManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 添加升级包
func (m *defaultOtaManage) OtaFirmwareInfoCreate(ctx context.Context, in *OtaFirmwareInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareInfoCreate(ctx, in, opts...)
}

// 添加升级包
func (d *directOtaManage) OtaFirmwareInfoCreate(ctx context.Context, in *OtaFirmwareInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.OtaFirmwareInfoCreate(ctx, in)
}

// 修改升级包
func (m *defaultOtaManage) OtaFirmwareInfoUpdate(ctx context.Context, in *OtaFirmwareInfoUpdateReq, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareInfoUpdate(ctx, in, opts...)
}

// 修改升级包
func (d *directOtaManage) OtaFirmwareInfoUpdate(ctx context.Context, in *OtaFirmwareInfoUpdateReq, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.OtaFirmwareInfoUpdate(ctx, in)
}

// 删除升级包
func (m *defaultOtaManage) OtaFirmwareInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareInfoDelete(ctx, in, opts...)
}

// 删除升级包
func (d *directOtaManage) OtaFirmwareInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaFirmwareInfoDelete(ctx, in)
}

// 升级包列表
func (m *defaultOtaManage) OtaFirmwareInfoIndex(ctx context.Context, in *OtaFirmwareInfoIndexReq, opts ...grpc.CallOption) (*OtaFirmwareInfoIndexResp, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareInfoIndex(ctx, in, opts...)
}

// 升级包列表
func (d *directOtaManage) OtaFirmwareInfoIndex(ctx context.Context, in *OtaFirmwareInfoIndexReq, opts ...grpc.CallOption) (*OtaFirmwareInfoIndexResp, error) {
	return d.svr.OtaFirmwareInfoIndex(ctx, in)
}

// 查询升级包
func (m *defaultOtaManage) OtaFirmwareInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareInfo, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareInfoRead(ctx, in, opts...)
}

// 查询升级包
func (d *directOtaManage) OtaFirmwareInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareInfo, error) {
	return d.svr.OtaFirmwareInfoRead(ctx, in)
}

// 创建静态升级批次
func (m *defaultOtaManage) OtaFirmwareJobCreate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareJobCreate(ctx, in, opts...)
}

// 创建静态升级批次
func (d *directOtaManage) OtaFirmwareJobCreate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.OtaFirmwareJobCreate(ctx, in)
}

// 获取升级包下的升级任务批次列表
func (m *defaultOtaManage) OtaFirmwareJobIndex(ctx context.Context, in *OtaFirmwareJobIndexReq, opts ...grpc.CallOption) (*OtaFirmwareJobIndexResp, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareJobIndex(ctx, in, opts...)
}

// 获取升级包下的升级任务批次列表
func (d *directOtaManage) OtaFirmwareJobIndex(ctx context.Context, in *OtaFirmwareJobIndexReq, opts ...grpc.CallOption) (*OtaFirmwareJobIndexResp, error) {
	return d.svr.OtaFirmwareJobIndex(ctx, in)
}

// 查询指定升级批次的详情
func (m *defaultOtaManage) OtaFirmwareJobRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareJobInfo, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareJobRead(ctx, in, opts...)
}

// 查询指定升级批次的详情
func (d *directOtaManage) OtaFirmwareJobRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*OtaFirmwareJobInfo, error) {
	return d.svr.OtaFirmwareJobRead(ctx, in)
}

// 取消动态升级策略
func (m *defaultOtaManage) OtaFirmwareJobUpdate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareJobUpdate(ctx, in, opts...)
}

// 取消动态升级策略
func (d *directOtaManage) OtaFirmwareJobUpdate(ctx context.Context, in *OtaFirmwareJobInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaFirmwareJobUpdate(ctx, in)
}

// 查询指定升级批次下的设备升级作业列表
func (m *defaultOtaManage) OtaFirmwareDeviceIndex(ctx context.Context, in *OtaFirmwareDeviceIndexReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceIndexResp, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceIndex(ctx, in, opts...)
}

// 查询指定升级批次下的设备升级作业列表
func (d *directOtaManage) OtaFirmwareDeviceIndex(ctx context.Context, in *OtaFirmwareDeviceIndexReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceIndexResp, error) {
	return d.svr.OtaFirmwareDeviceIndex(ctx, in)
}

// 取消指定批次下的设备升级作业
func (m *defaultOtaManage) OtaFirmwareDeviceCancel(ctx context.Context, in *OtaFirmwareDeviceCancelReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceCancel(ctx, in, opts...)
}

// 取消指定批次下的设备升级作业
func (d *directOtaManage) OtaFirmwareDeviceCancel(ctx context.Context, in *OtaFirmwareDeviceCancelReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaFirmwareDeviceCancel(ctx, in)
}

// 重新升级指定批次下升级失败或升级取消的设备升级作业
func (m *defaultOtaManage) OtaFirmwareDeviceRetry(ctx context.Context, in *OtaFirmwareDeviceRetryReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceRetry(ctx, in, opts...)
}

// 重新升级指定批次下升级失败或升级取消的设备升级作业
func (d *directOtaManage) OtaFirmwareDeviceRetry(ctx context.Context, in *OtaFirmwareDeviceRetryReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaFirmwareDeviceRetry(ctx, in)
}

// app确认设备升级
func (m *defaultOtaManage) OtaFirmwareDeviceConfirm(ctx context.Context, in *OtaFirmwareDeviceConfirmReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceConfirm(ctx, in, opts...)
}

// app确认设备升级
func (d *directOtaManage) OtaFirmwareDeviceConfirm(ctx context.Context, in *OtaFirmwareDeviceConfirmReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaFirmwareDeviceConfirm(ctx, in)
}

func (m *defaultOtaManage) OtaModuleInfoCreate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaModuleInfoCreate(ctx, in, opts...)
}

func (d *directOtaManage) OtaModuleInfoCreate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.OtaModuleInfoCreate(ctx, in)
}

func (m *defaultOtaManage) OtaModuleInfoUpdate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaModuleInfoUpdate(ctx, in, opts...)
}

func (d *directOtaManage) OtaModuleInfoUpdate(ctx context.Context, in *OtaModuleInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaModuleInfoUpdate(ctx, in)
}

func (m *defaultOtaManage) OtaModuleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaModuleInfoDelete(ctx, in, opts...)
}

func (d *directOtaManage) OtaModuleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaModuleInfoDelete(ctx, in)
}

func (m *defaultOtaManage) OtaModuleInfoIndex(ctx context.Context, in *OtaModuleInfoIndexReq, opts ...grpc.CallOption) (*OtaModuleInfoIndexResp, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaModuleInfoIndex(ctx, in, opts...)
}

func (d *directOtaManage) OtaModuleInfoIndex(ctx context.Context, in *OtaModuleInfoIndexReq, opts ...grpc.CallOption) (*OtaModuleInfoIndexResp, error) {
	return d.svr.OtaModuleInfoIndex(ctx, in)
}

func (m *defaultOtaManage) OtaModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*OtaModuleInfo, error) {
	client := dm.NewOtaManageClient(m.cli.Conn())
	return client.OtaModuleInfoRead(ctx, in, opts...)
}

func (d *directOtaManage) OtaModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*OtaModuleInfo, error) {
	return d.svr.OtaModuleInfoRead(ctx, in)
}
