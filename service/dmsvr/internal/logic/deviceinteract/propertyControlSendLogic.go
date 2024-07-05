package deviceinteractlogic

import (
	"context"
	"encoding/json"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/domain/deviceMsg"
	"gitee.com/i-Things/share/domain/deviceMsg/msgThing"
	"gitee.com/i-Things/share/domain/schema"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceLog"
	"github.com/i-Things/things/service/dmsvr/internal/domain/shadow"
	"github.com/i-Things/things/service/dmsvr/internal/domain/userShared"
	devicemanagelogic "github.com/i-Things/things/service/dmsvr/internal/logic/devicemanage"
	"github.com/i-Things/things/service/dmsvr/internal/repo/cache"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"time"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropertyControlSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	model *schema.Model
}

func NewPropertyControlSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropertyControlSendLogic {
	return &PropertyControlSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *PropertyControlSendLogic) initMsg(productID string) error {
	var err error
	l.model, err = l.svcCtx.SchemaRepo.GetData(l.ctx, productID)
	if err != nil {
		return errors.System.AddDetail(err)
	}
	return nil
}

func (l *PropertyControlSendLogic) Auth(dev devices.Core, param map[string]any) (outParam map[string]any, err error) {
	uc := ctxs.GetUserCtx(l.ctx)

	if uc != nil && !uc.IsAdmin {
		di, err := l.svcCtx.DeviceCache.GetData(l.ctx, dev)
		if err != nil {
			return nil, err
		}
		_, ok := uc.ProjectAuth[di.ProjectID]
		if !ok {
			uds, err := l.svcCtx.UserDeviceShare.GetData(l.ctx, userShared.UserShareKey{
				ProductID:    dev.ProductID,
				DeviceName:   dev.DeviceName,
				SharedUserID: uc.UserID,
			})
			if err != nil {
				return nil, err
			}
			if uds.AuthType == def.AuthAdmin {
				return param, nil
			}
			for k := range param {
				sp := uds.SchemaPerm[k]
				if sp != nil && sp.Perm != def.AuthReadWrite {
					return nil, errors.Parameter.AddMsgf("属性:%v 没有控制权限", k)
				}
			}
			return param, nil
		}
		return param, nil
	}
	return param, nil
}

// 调用设备属性
func (l *PropertyControlSendLogic) PropertyControlSend(in *dm.PropertyControlSendReq) (ret *dm.PropertyControlSendResp, err error) {
	l.Infof("%s req=%+v", utils.FuncName(), in)
	var isOnline = true
	var protocolCode string
	var dev = devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}
	if protocolCode, err = CheckIsOnline(l.ctx, l.svcCtx, dev); err != nil { //如果是不启用设备影子的模式则直接返回
		if errors.Is(err, errors.NotOnline) {
			isOnline = false
			if in.ShadowControl == shadow.ControlNo {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	err = l.initMsg(in.ProductID)
	if err != nil {
		return nil, err
	}

	param := map[string]any{}
	err = utils.Unmarshal([]byte(in.Data), &param)
	if err != nil {
		return nil, errors.Parameter.AddDetail(
			"SendProperty data not right:", in.Data)
	}
	param, err = l.Auth(dev, param)
	if err != nil {
		return nil, err
	}
	MsgToken := devices.GenMsgToken(l.ctx, l.svcCtx.NodeID)

	req := msgThing.Req{
		CommonMsg: deviceMsg.CommonMsg{
			Method:    deviceMsg.Control,
			MsgToken:  MsgToken,
			Timestamp: time.Now().UnixMilli(),
		},
		Params: param,
	}
	params, err := req.VerifyReqParam(l.model, schema.ParamProperty)
	if err != nil {
		return nil, err
	}
	if len(params) == 0 {
		l.Infof("控制的属性在设备中都不存在,req:%v", utils.Fmt(in))
		return &dm.PropertyControlSendResp{Code: errors.OK.Code, Msg: errors.OK.AddMsg("该设备无控制的属性,忽略").GetMsg()}, nil
	}
	defer func() {
		if err == nil && ret.Code == errors.OK.GetCode() && in.WithProfile != nil && in.WithProfile.Code != "" {
			_, err = devicemanagelogic.NewDeviceProfileUpdateLogic(l.ctx, l.svcCtx).DeviceProfileUpdate(&dm.DeviceProfile{
				Device: &dm.DeviceCore{
					ProductID:  in.ProductID,
					DeviceName: in.DeviceName,
				},
				Code:   in.WithProfile.Code,
				Params: in.WithProfile.Params,
			})
		}
	}()
	err = req.FmtReqParam(l.model, schema.ParamProperty)
	if err != nil {
		return nil, err
	}
	req.Params, err = msgThing.ToVal(params)
	if err != nil {
		return nil, err
	}
	if in.ShadowControl == shadow.ControlOnlyCloud {
		//插入多条设备物模型属性数据
		err = l.svcCtx.SchemaManaRepo.InsertPropertiesData(l.ctx, l.model, in.ProductID, in.DeviceName, params, time.Now())
		if err != nil {
			l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
			return nil, err
		}
		return &dm.PropertyControlSendResp{Code: errors.OK.Code, Msg: errors.OK.AddMsg("只修改云端值").GetMsg()}, nil
	}
	defer func() {
		ctxs.GoNewCtx(l.ctx, func(ctx context.Context) {
			uc := ctxs.GetUserCtx(l.ctx)
			for dataID, content := range param {
				_ = l.svcCtx.SendRepo.Insert(ctx, &deviceLog.Send{
					ProductID:  in.ProductID,
					Action:     "propertyControlSend",
					Timestamp:  time.Now(), // 操作时间
					DeviceName: in.DeviceName,
					TraceID:    utils.TraceIdFromContext(ctx),
					UserID:     uc.UserID,
					DataID:     dataID,
					Account:    uc.Account,
					Content:    utils.Fmt(content),
					ResultCode: errors.Fmt(err).GetCode(),
				})
			}
		})
	}()
	if in.ShadowControl == shadow.ControlOnlyCloudWithLog {
		//插入多条设备物模型属性数据
		err = l.svcCtx.SchemaManaRepo.InsertPropertiesData(l.ctx, l.model, in.ProductID, in.DeviceName, params, time.Now())
		if err != nil {
			l.Errorf("%s.InsertPropertyData err=%+v", utils.FuncName(), err)
			return nil, err
		}
		return &dm.PropertyControlSendResp{Code: errors.OK.Code, Msg: errors.OK.AddMsg("只修改云端值及记录操作").GetMsg()}, nil
	}
	if in.ShadowControl == shadow.ControlOnly || (!isOnline && in.ShadowControl == shadow.ControlAuto) {
		//设备影子模式
		err = shadow.CheckEnableShadow(param, l.model)
		if err != nil {
			return nil, err
		}
		err = relationDB.NewShadowRepo(l.ctx).MultiUpdate(l.ctx, shadow.NewInfo(in.ProductID, in.DeviceName, param))
		if err != nil {
			return nil, err
		}
		return &dm.PropertyControlSendResp{Code: errors.OK.Code, Msg: errors.OK.AddMsg("影子模式").GetMsg()}, nil
	}

	payload, _ := json.Marshal(req)
	reqMsg := deviceMsg.PublishMsg{
		Handle:       devices.Thing,
		Type:         msgThing.TypeProperty,
		Payload:      payload,
		Timestamp:    time.Now().UnixMilli(),
		ProductID:    in.ProductID,
		DeviceName:   in.DeviceName,
		ProtocolCode: protocolCode,
	}
	err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Cache, deviceMsg.ReqMsg, &reqMsg, req.MsgToken)
	if err != nil {
		return nil, err
	}

	if in.IsAsync { //如果是异步获取 处理结果暂不关注
		err := l.svcCtx.PubDev.PublishToDev(l.ctx, &reqMsg)
		if err != nil {
			return nil, err
		}
		return &dm.PropertyControlSendResp{
			MsgToken: req.MsgToken,
		}, nil
	}
	var resp []byte
	resp, err = l.svcCtx.PubDev.ReqToDeviceSync(l.ctx, &reqMsg, func(payload []byte) bool {
		var dresp msgThing.Resp
		err = utils.Unmarshal(payload, &dresp)
		if err != nil { //如果是没法解析的说明不是需要的包,直接跳过即可
			return false
		}
		if dresp.MsgToken != req.MsgToken { //不是该请求的回复.跳过
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	var dresp msgThing.Resp
	err = utils.Unmarshal(resp, &dresp)
	if err != nil {
		return nil, err
	}
	return &dm.PropertyControlSendResp{
		MsgToken: dresp.MsgToken,
		Msg:      dresp.Msg,
		Code:     dresp.Code,
	}, nil
}
