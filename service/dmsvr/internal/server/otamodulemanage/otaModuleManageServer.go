// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/logic/otamodulemanage"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

type OTAModuleManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedOTAModuleManageServer
}

func NewOTAModuleManageServer(svcCtx *svc.ServiceContext) *OTAModuleManageServer {
	return &OTAModuleManageServer{
		svcCtx: svcCtx,
	}
}

// 创建产品的OTA模块
func (s *OTAModuleManageServer) OtaModuleCreate(ctx context.Context, in *dm.OTAModuleReq) (*dm.Empty, error) {
	l := otamodulemanagelogic.NewOtaModuleCreateLogic(ctx, s.svcCtx)
	return l.OtaModuleCreate(in)
}

// 修改OTA模块别名、描述
func (s *OTAModuleManageServer) OtaModuleUpdate(ctx context.Context, in *dm.OTAModuleReq) (*dm.Empty, error) {
	l := otamodulemanagelogic.NewOtaModuleUpdateLogic(ctx, s.svcCtx)
	return l.OtaModuleUpdate(in)
}

// 删除自定义OTA模块
func (s *OTAModuleManageServer) OtaModuleDelete(ctx context.Context, in *dm.OTAModuleDeleteReq) (*dm.Empty, error) {
	l := otamodulemanagelogic.NewOtaModuleDeleteLogic(ctx, s.svcCtx)
	return l.OtaModuleDelete(in)
}

// 查询产品下的OTA模块列表
func (s *OTAModuleManageServer) OtaModuleByProductIndex(ctx context.Context, in *dm.OTAModuleIndexReq) (*dm.OTAModuleIndexResp, error) {
	l := otamodulemanagelogic.NewOtaModuleByProductIndexLogic(ctx, s.svcCtx)
	return l.OtaModuleByProductIndex(in)
}

// 查询设备上报过的OTA模块版本列表
func (s *OTAModuleManageServer) OtaModuleVersionsByDeviceIndex(ctx context.Context, in *dm.OTAModuleIndexReq) (*dm.OTAModuleVersionsIndexResp, error) {
	l := otamodulemanagelogic.NewOtaModuleVersionsByDeviceIndexLogic(ctx, s.svcCtx)
	return l.OtaModuleVersionsByDeviceIndex(in)
}
