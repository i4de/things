// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/logic/otamanage"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

type OtaManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedOtaManageServer
}

func NewOtaManageServer(svcCtx *svc.ServiceContext) *OtaManageServer {
	return &OtaManageServer{
		svcCtx: svcCtx,
	}
}

// 添加升级包
func (s *OtaManageServer) OtaFirmwareInfoCreate(ctx context.Context, in *dm.OtaFirmwareInfoCreateReq) (*dm.WithID, error) {
	l := otamanagelogic.NewOtaFirmwareInfoCreateLogic(ctx, s.svcCtx)
	return l.OtaFirmwareInfoCreate(in)
}

// 修改升级包
func (s *OtaManageServer) OtaFirmwareInfoUpdate(ctx context.Context, in *dm.OtaFirmwareInfoUpdateReq) (*dm.WithID, error) {
	l := otamanagelogic.NewOtaFirmwareInfoUpdateLogic(ctx, s.svcCtx)
	return l.OtaFirmwareInfoUpdate(in)
}

// 删除升级包
func (s *OtaManageServer) OtaFirmwareInfoDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaFirmwareInfoDeleteLogic(ctx, s.svcCtx)
	return l.OtaFirmwareInfoDelete(in)
}

// 升级包列表
func (s *OtaManageServer) OtaFirmwareInfoIndex(ctx context.Context, in *dm.OtaFirmwareInfoIndexReq) (*dm.OtaFirmwareInfoIndexResp, error) {
	l := otamanagelogic.NewOtaFirmwareInfoIndexLogic(ctx, s.svcCtx)
	return l.OtaFirmwareInfoIndex(in)
}

// 查询升级包
func (s *OtaManageServer) OtaFirmwareInfoRead(ctx context.Context, in *dm.WithID) (*dm.OtaFirmwareInfo, error) {
	l := otamanagelogic.NewOtaFirmwareInfoReadLogic(ctx, s.svcCtx)
	return l.OtaFirmwareInfoRead(in)
}

// 创建静态升级批次
func (s *OtaManageServer) OtaFirmwareJobCreate(ctx context.Context, in *dm.OtaFirmwareJobInfo) (*dm.WithID, error) {
	l := otamanagelogic.NewOtaFirmwareJobCreateLogic(ctx, s.svcCtx)
	return l.OtaFirmwareJobCreate(in)
}

// 获取升级包下的升级任务批次列表
func (s *OtaManageServer) OtaFirmwareJobIndex(ctx context.Context, in *dm.OtaFirmwareJobIndexReq) (*dm.OtaFirmwareJobIndexResp, error) {
	l := otamanagelogic.NewOtaFirmwareJobIndexLogic(ctx, s.svcCtx)
	return l.OtaFirmwareJobIndex(in)
}

// 查询指定升级批次的详情
func (s *OtaManageServer) OtaFirmwareJobRead(ctx context.Context, in *dm.WithID) (*dm.OtaFirmwareJobInfo, error) {
	l := otamanagelogic.NewOtaFirmwareJobReadLogic(ctx, s.svcCtx)
	return l.OtaFirmwareJobRead(in)
}

// 取消动态升级策略
func (s *OtaManageServer) OtaFirmwareJobUpdate(ctx context.Context, in *dm.OtaFirmwareJobInfo) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaFirmwareJobUpdateLogic(ctx, s.svcCtx)
	return l.OtaFirmwareJobUpdate(in)
}

// 查询指定升级批次下的设备升级作业列表
func (s *OtaManageServer) OtaFirmwareDeviceIndex(ctx context.Context, in *dm.OtaFirmwareDeviceIndexReq) (*dm.OtaFirmwareDeviceIndexResp, error) {
	l := otamanagelogic.NewOtaFirmwareDeviceIndexLogic(ctx, s.svcCtx)
	return l.OtaFirmwareDeviceIndex(in)
}

// 取消指定批次下的设备升级作业
func (s *OtaManageServer) OtaFirmwareDeviceCancel(ctx context.Context, in *dm.OtaFirmwareDeviceCancelReq) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaFirmwareDeviceCancelLogic(ctx, s.svcCtx)
	return l.OtaFirmwareDeviceCancel(in)
}

// 重新升级指定批次下升级失败或升级取消的设备升级作业
func (s *OtaManageServer) OtaFirmwareDeviceRetry(ctx context.Context, in *dm.OtaFirmwareDeviceRetryReq) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaFirmwareDeviceRetryLogic(ctx, s.svcCtx)
	return l.OtaFirmwareDeviceRetry(in)
}

// app确认设备升级
func (s *OtaManageServer) OtaFirmwareDeviceConfirm(ctx context.Context, in *dm.OtaFirmwareDeviceConfirmReq) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaFirmwareDeviceConfirmLogic(ctx, s.svcCtx)
	return l.OtaFirmwareDeviceConfirm(in)
}

func (s *OtaManageServer) OtaModuleInfoCreate(ctx context.Context, in *dm.OtaModuleInfo) (*dm.WithID, error) {
	l := otamanagelogic.NewOtaModuleInfoCreateLogic(ctx, s.svcCtx)
	return l.OtaModuleInfoCreate(in)
}

func (s *OtaManageServer) OtaModuleInfoUpdate(ctx context.Context, in *dm.OtaModuleInfo) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaModuleInfoUpdateLogic(ctx, s.svcCtx)
	return l.OtaModuleInfoUpdate(in)
}

func (s *OtaManageServer) OtaModuleInfoDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := otamanagelogic.NewOtaModuleInfoDeleteLogic(ctx, s.svcCtx)
	return l.OtaModuleInfoDelete(in)
}

func (s *OtaManageServer) OtaModuleInfoIndex(ctx context.Context, in *dm.OtaModuleInfoIndexReq) (*dm.OtaModuleInfoIndexResp, error) {
	l := otamanagelogic.NewOtaModuleInfoIndexLogic(ctx, s.svcCtx)
	return l.OtaModuleInfoIndex(in)
}

func (s *OtaManageServer) OtaModuleInfoRead(ctx context.Context, in *dm.WithIDCode) (*dm.OtaModuleInfo, error) {
	l := otamanagelogic.NewOtaModuleInfoReadLogic(ctx, s.svcCtx)
	return l.OtaModuleInfoRead(in)
}
