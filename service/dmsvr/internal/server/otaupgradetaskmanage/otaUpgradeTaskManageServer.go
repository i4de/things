// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/logic/otaupgradetaskmanage"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

type OTAUpgradeTaskManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedOTAUpgradeTaskManageServer
}

func NewOTAUpgradeTaskManageServer(svcCtx *svc.ServiceContext) *OTAUpgradeTaskManageServer {
	return &OTAUpgradeTaskManageServer{
		svcCtx: svcCtx,
	}
}

// 查询指定升级批次下的设备升级作业列表
func (s *OTAUpgradeTaskManageServer) OtaTaskByJobIndex(ctx context.Context, in *dm.OTATaskByJobIndexReq) (*dm.OtaTaskByJobIndexResp, error) {
	l := otaupgradetaskmanagelogic.NewOtaTaskByJobIndexLogic(ctx, s.svcCtx)
	return l.OtaTaskByJobIndex(in)
}

// 取消指定批次下的设备升级作业
func (s *OTAUpgradeTaskManageServer) OtaTaskByJobCancel(ctx context.Context, in *dm.OTATaskByJobCancelReq) (*dm.Empty, error) {
	l := otaupgradetaskmanagelogic.NewOtaTaskByJobCancelLogic(ctx, s.svcCtx)
	return l.OtaTaskByJobCancel(in)
}

// 取消指定OTA升级包下状态为待确认、待推送、已推送、升级中状态的设备升级作业
func (s *OTAUpgradeTaskManageServer) OtaTaskByDeviceCancel(ctx context.Context, in *dm.OTATaskByDeviceCancelReq) (*dm.Empty, error) {
	l := otaupgradetaskmanagelogic.NewOtaTaskByDeviceCancelLogic(ctx, s.svcCtx)
	return l.OtaTaskByDeviceCancel(in)
}

// 批量确认，处于待确认状态的设备升级作业
func (s *OTAUpgradeTaskManageServer) OtaTaskConfirm(ctx context.Context, in *dm.OTATaskConfirmReq) (*dm.Empty, error) {
	l := otaupgradetaskmanagelogic.NewOtaTaskConfirmLogic(ctx, s.svcCtx)
	return l.OtaTaskConfirm(in)
}

// 查询指定设备下，未完成状态的设备升级作业
func (s *OTAUpgradeTaskManageServer) OtaUnfinishedTaskByDeviceIndex(ctx context.Context, in *dm.OTAUnfinishedTaskByDeviceIndexReq) (*dm.OTAUnfinishedTaskByDeviceIndexResp, error) {
	l := otaupgradetaskmanagelogic.NewOtaUnfinishedTaskByDeviceIndexLogic(ctx, s.svcCtx)
	return l.OtaUnfinishedTaskByDeviceIndex(in)
}

// 重新升级指定批次下升级失败或升级取消的设备升级作业
func (s *OTAUpgradeTaskManageServer) OtaTaskReUpgrade(ctx context.Context, in *dm.OTATaskReUpgradeReq) (*dm.Empty, error) {
	l := otaupgradetaskmanagelogic.NewOtaTaskReUpgradeLogic(ctx, s.svcCtx)
	return l.OtaTaskReUpgrade(in)
}
