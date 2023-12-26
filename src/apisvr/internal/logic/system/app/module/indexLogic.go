package module

import (
	"context"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.AppModuleIndexReq) (resp *types.AppModuleIndexResp, err error) {
	ret, err := l.svcCtx.AppRpc.AppModuleIndex(l.ctx, &sys.AppModuleIndexReq{Code: req.Code})
	if err != nil {
		return nil, err
	}
	return &types.AppModuleIndexResp{
		ModuleCodes: ret.ModuleCodes,
	}, nil
}
