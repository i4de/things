package productmanagelogic

import (
	"context"
	"gitee.com/i-Things/share/def"
	"github.com/i-Things/things/service/dmsvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryReadLogic {
	return &ProductCategoryReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取产品信息详情
func (l *ProductCategoryReadLogic) ProductCategoryRead(in *dm.ProductCategoryReadReq) (*dm.ProductCategory, error) {
	var (
		po  *relationDB.DmProductCategory
		err error
	)
	switch in.Id {
	case def.RootNode, 0:
		po = &relationDB.DmProductCategory{
			ID:   def.RootNode,
			Name: "全部产品品类",
		}
	default:
		po, err = relationDB.NewProductCategoryRepo(l.ctx).FindOne(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
	}
	if !in.WithChildren {
		return logic.ToProductCategoryPb(l.ctx, l.svcCtx, po, nil), nil
	}
	children, err := relationDB.NewProductCategoryRepo(l.ctx).FindByFilter(l.ctx,
		relationDB.ProductCategoryFilter{IDPath: po.IDPath}, nil)
	if err != nil {
		return nil, err
	}
	return logic.ToProductCategoryPb(l.ctx, l.svcCtx, po, children), err
}
