// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package otamodulemanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                      = dm.ActionRespReq
	ActionSendReq                      = dm.ActionSendReq
	ActionSendResp                     = dm.ActionSendResp
	CommonSchemaCreateReq              = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq               = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp              = dm.CommonSchemaIndexResp
	CommonSchemaInfo                   = dm.CommonSchemaInfo
	CommonSchemaUpdateReq              = dm.CommonSchemaUpdateReq
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceCountInfo                    = dm.DeviceCountInfo
	DeviceCountReq                     = dm.DeviceCountReq
	DeviceCountResp                    = dm.DeviceCountResp
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCount                    = dm.DeviceInfoCount
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq           = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                  = dm.DeviceInfoReadReq
	DeviceTypeCountReq                 = dm.DeviceTypeCountReq
	DeviceTypeCountResp                = dm.DeviceTypeCountResp
	DynamicUpgradeJobReq               = dm.DynamicUpgradeJobReq
	Empty                              = dm.Empty
	EventLogIndexReq                   = dm.EventLogIndexReq
	EventLogIndexResp                  = dm.EventLogIndexResp
	EventLogInfo                       = dm.EventLogInfo
	Firmware                           = dm.Firmware
	FirmwareFile                       = dm.FirmwareFile
	FirmwareInfo                       = dm.FirmwareInfo
	FirmwareInfoDeleteReq              = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp             = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq               = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp              = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp               = dm.FirmwareInfoReadResp
	FirmwareResp                       = dm.FirmwareResp
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq            = dm.GroupDeviceMultiSaveReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	HubLogInfo                         = dm.HubLogInfo
	JobReq                             = dm.JobReq
	OTAModuleDeleteReq                 = dm.OTAModuleDeleteReq
	OTAModuleDetail                    = dm.OTAModuleDetail
	OTAModuleIndexReq                  = dm.OTAModuleIndexReq
	OTAModuleIndexResp                 = dm.OTAModuleIndexResp
	OTAModuleReq                       = dm.OTAModuleReq
	OTAModuleVersionsIndexResp         = dm.OTAModuleVersionsIndexResp
	OTATaskByDeviceCancelReq           = dm.OTATaskByDeviceCancelReq
	OTATaskByDeviceNameReq             = dm.OTATaskByDeviceNameReq
	OTATaskByJobCancelReq              = dm.OTATaskByJobCancelReq
	OTATaskByJobIndexReq               = dm.OTATaskByJobIndexReq
	OTATaskConfirmReq                  = dm.OTATaskConfirmReq
	OTATaskReUpgradeReq                = dm.OTATaskReUpgradeReq
	OTAUnfinishedTaskByDeviceIndexReq  = dm.OTAUnfinishedTaskByDeviceIndexReq
	OTAUnfinishedTaskByDeviceIndexResp = dm.OTAUnfinishedTaskByDeviceIndexResp
	OtaCommonResp                      = dm.OtaCommonResp
	OtaFirmwareCreateReq               = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq               = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq           = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp          = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile                    = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq            = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp           = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                 = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq                = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp               = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo                    = dm.OtaFirmwareInfo
	OtaFirmwareReadReq                 = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp                = dm.OtaFirmwareReadResp
	OtaFirmwareResp                    = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq               = dm.OtaFirmwareUpdateReq
	OtaFirmwareVerifyReq               = dm.OtaFirmwareVerifyReq
	OtaJobByDeviceIndexReq             = dm.OtaJobByDeviceIndexReq
	OtaJobByFirmwareIndexReq           = dm.OtaJobByFirmwareIndexReq
	OtaJobInfo                         = dm.OtaJobInfo
	OtaJobInfoIndexResp                = dm.OtaJobInfoIndexResp
	OtaModuleInfo                      = dm.OtaModuleInfo
	OtaPageInfo                        = dm.OtaPageInfo
	OtaPromptIndexReq                  = dm.OtaPromptIndexReq
	OtaPromptIndexResp                 = dm.OtaPromptIndexResp
	OtaTaskBatchReq                    = dm.OtaTaskBatchReq
	OtaTaskBatchResp                   = dm.OtaTaskBatchResp
	OtaTaskByJobIndexResp              = dm.OtaTaskByJobIndexResp
	OtaTaskCancleReq                   = dm.OtaTaskCancleReq
	OtaTaskCreatResp                   = dm.OtaTaskCreatResp
	OtaTaskCreateReq                   = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq             = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq              = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp             = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo                  = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq            = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq               = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq                    = dm.OtaTaskIndexReq
	OtaTaskIndexResp                   = dm.OtaTaskIndexResp
	OtaTaskInfo                        = dm.OtaTaskInfo
	OtaTaskReadReq                     = dm.OtaTaskReadReq
	OtaTaskReadResp                    = dm.OtaTaskReadResp
	OtaUpTaskInfo                      = dm.OtaUpTaskInfo
	PageInfo                           = dm.PageInfo
	PageInfo_OrderBy                   = dm.PageInfo_OrderBy
	Point                              = dm.Point
	ProductCategory                    = dm.ProductCategory
	ProductCategoryIndexReq            = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp           = dm.ProductCategoryIndexResp
	ProductCategoryReadReq             = dm.ProductCategoryReadReq
	ProductCustom                      = dm.ProductCustom
	ProductCustomReadReq               = dm.ProductCustomReadReq
	ProductInfo                        = dm.ProductInfo
	ProductInfoDeleteReq               = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                = dm.ProductInfoIndexReq
	ProductInfoIndexResp               = dm.ProductInfoIndexResp
	ProductInfoReadReq                 = dm.ProductInfoReadReq
	ProductInitReq                     = dm.ProductInitReq
	ProductRemoteConfig                = dm.ProductRemoteConfig
	ProductSchemaCreateReq             = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq             = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq              = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp             = dm.ProductSchemaIndexResp
	ProductSchemaInfo                  = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq        = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq          = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq            = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp           = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq             = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq        = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp       = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg             = dm.PropertyControlSendMsg
	PropertyControlSendReq             = dm.PropertyControlSendReq
	PropertyControlSendResp            = dm.PropertyControlSendResp
	PropertyGetReportSendReq           = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp          = dm.PropertyGetReportSendResp
	PropertyLogIndexReq                = dm.PropertyLogIndexReq
	PropertyLogIndexResp               = dm.PropertyLogIndexResp
	PropertyLogInfo                    = dm.PropertyLogInfo
	PropertyLogLatestIndexReq          = dm.PropertyLogLatestIndexReq
	ProtocolConfigField                = dm.ProtocolConfigField
	ProtocolConfigInfo                 = dm.ProtocolConfigInfo
	ProtocolInfo                       = dm.ProtocolInfo
	ProtocolInfoIndexReq               = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp              = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq              = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq               = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp              = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq            = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp           = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq             = dm.RemoteConfigPushAllReq
	RespReadReq                        = dm.RespReadReq
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SdkLogInfo                         = dm.SdkLogInfo
	SendLogIndexReq                    = dm.SendLogIndexReq
	SendLogIndexResp                   = dm.SendLogIndexResp
	SendLogInfo                        = dm.SendLogInfo
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	StatusLogIndexReq                  = dm.StatusLogIndexReq
	StatusLogIndexResp                 = dm.StatusLogIndexResp
	StatusLogInfo                      = dm.StatusLogInfo
	Tag                                = dm.Tag
	TimeRange                          = dm.TimeRange
	UpgradeJobResp                     = dm.UpgradeJobResp
	UserDeviceCollectSave              = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq            = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp           = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo                = dm.UserDeviceShareInfo
	UserDeviceShareReadReq             = dm.UserDeviceShareReadReq
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq
	WithID                             = dm.WithID
	WithIDCode                         = dm.WithIDCode

	OTAModuleManage interface {
		// 创建产品的OTA模块
		OtaModuleCreate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error)
		// 修改OTA模块别名、描述
		OtaModuleUpdate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error)
		// 删除自定义OTA模块
		OtaModuleDelete(ctx context.Context, in *OTAModuleDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 查询产品下的OTA模块列表
		OtaModuleByProductIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleIndexResp, error)
		// 查询设备上报过的OTA模块版本列表
		OtaModuleVersionsByDeviceIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleVersionsIndexResp, error)
	}

	defaultOTAModuleManage struct {
		cli zrpc.Client
	}

	directOTAModuleManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.OTAModuleManageServer
	}
)

func NewOTAModuleManage(cli zrpc.Client) OTAModuleManage {
	return &defaultOTAModuleManage{
		cli: cli,
	}
}

func NewDirectOTAModuleManage(svcCtx *svc.ServiceContext, svr dm.OTAModuleManageServer) OTAModuleManage {
	return &directOTAModuleManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 创建产品的OTA模块
func (m *defaultOTAModuleManage) OtaModuleCreate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOTAModuleManageClient(m.cli.Conn())
	return client.OtaModuleCreate(ctx, in, opts...)
}

// 创建产品的OTA模块
func (d *directOTAModuleManage) OtaModuleCreate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaModuleCreate(ctx, in)
}

// 修改OTA模块别名、描述
func (m *defaultOTAModuleManage) OtaModuleUpdate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOTAModuleManageClient(m.cli.Conn())
	return client.OtaModuleUpdate(ctx, in, opts...)
}

// 修改OTA模块别名、描述
func (d *directOTAModuleManage) OtaModuleUpdate(ctx context.Context, in *OTAModuleReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaModuleUpdate(ctx, in)
}

// 删除自定义OTA模块
func (m *defaultOTAModuleManage) OtaModuleDelete(ctx context.Context, in *OTAModuleDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOTAModuleManageClient(m.cli.Conn())
	return client.OtaModuleDelete(ctx, in, opts...)
}

// 删除自定义OTA模块
func (d *directOTAModuleManage) OtaModuleDelete(ctx context.Context, in *OTAModuleDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.OtaModuleDelete(ctx, in)
}

// 查询产品下的OTA模块列表
func (m *defaultOTAModuleManage) OtaModuleByProductIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleIndexResp, error) {
	client := dm.NewOTAModuleManageClient(m.cli.Conn())
	return client.OtaModuleByProductIndex(ctx, in, opts...)
}

// 查询产品下的OTA模块列表
func (d *directOTAModuleManage) OtaModuleByProductIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleIndexResp, error) {
	return d.svr.OtaModuleByProductIndex(ctx, in)
}

// 查询设备上报过的OTA模块版本列表
func (m *defaultOTAModuleManage) OtaModuleVersionsByDeviceIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleVersionsIndexResp, error) {
	client := dm.NewOTAModuleManageClient(m.cli.Conn())
	return client.OtaModuleVersionsByDeviceIndex(ctx, in, opts...)
}

// 查询设备上报过的OTA模块版本列表
func (d *directOTAModuleManage) OtaModuleVersionsByDeviceIndex(ctx context.Context, in *OTAModuleIndexReq, opts ...grpc.CallOption) (*OTAModuleVersionsIndexResp, error) {
	return d.svr.OtaModuleVersionsByDeviceIndex(ctx, in)
}
