package deviceMsgEvent

import (
	"context"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type GatewayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatewayLogic {
	return &GatewayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GatewayLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%+v", utils.FuncName(), msg)
	return &deviceMsg.PublishMsg{
		Topic:      "",
		Payload:    nil,
		Timestamp:  time.Time{},
		ProductID:  "",
		DeviceName: "",
	}, nil
}