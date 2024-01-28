package schemamanagelogic

import (
	"context"
	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonSchemaIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PsDB *relationDB.CommonSchemaRepo
}

func NewCommonSchemaIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonSchemaIndexLogic {
	return &CommonSchemaIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PsDB:   relationDB.NewCommonSchemaRepo(ctx),
	}
}

// 获取产品信息列表
func (l *CommonSchemaIndexLogic) CommonSchemaIndex(in *dm.CommonSchemaIndexReq) (*dm.CommonSchemaIndexResp, error) {
	filter := relationDB.CommonSchemaFilter{
		Type:        in.Type,
		Identifiers: in.Identifiers,
	}
	if len(in.ProductIDs) != 0 {
		rst, err := relationDB.NewProductSchemaRepo(l.ctx).FindByFilter(l.ctx, relationDB.ProductSchemaFilter{ProductIDs: in.ProductIDs, Tag: 2}, nil)
		if err != nil {
			return nil, err
		}
		var identifyMap = map[string]int{}
		for _, v := range rst {
			identifyMap[v.Identifier]++
		}
		for k, v := range identifyMap {
			if v == len(in.ProductIDs) { //每个产品都有的物模型
				filter.Identifiers = append(filter.Identifiers, k)
			}
		}
	}

	schemas, err := l.PsDB.FindByFilter(l.ctx, filter, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	total, err := l.PsDB.CountByFilter(l.ctx, filter)
	if err != nil {
		return nil, err
	}
	list := make([]*dm.CommonSchemaInfo, 0, len(schemas))
	for _, s := range schemas {
		list = append(list, ToCommonSchemaRpc(s))
	}
	return &dm.CommonSchemaIndexResp{List: list, Total: total}, nil
}
