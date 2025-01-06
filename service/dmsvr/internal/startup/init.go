package startup

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/unitedrhino/core/service/syssvr/pb/sys"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceGroup"
	"time"

	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/unitedrhino/share/caches"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/application"
	"gitee.com/unitedrhino/share/domain/deviceMsg"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/share/events/topics"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceStatus"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/userShared"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/event/deviceMsgEvent"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/event/otaEvent"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/event/serverEvent"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/event/staticEvent"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	devicemanagelogic "gitee.com/unitedrhino/things/service/dmsvr/internal/logic/devicemanage"
	userdevicelogic "gitee.com/unitedrhino/things/service/dmsvr/internal/logic/userdevice"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/event/subscribe/server"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

func Init(svcCtx *svc.ServiceContext) {
	logic.Init(svcCtx)
	VersionUpdate(svcCtx)
	InitCache(svcCtx)
	TimerInit(svcCtx)
	InitSubscribe(svcCtx)
	InitEventBus(svcCtx)
	DictInit(svcCtx)
}

func init() {
	//var (
	//	TagsTypes []*types.Tag
	//	TagMap    map[string]string
	//)
	//utils.AddConverter(
	//	utils.TypeConverter{SrcType: TagsTypes, DstType: TagMap, Fn: func(src interface{}) (dst interface{}, err error) {
	//		return logic.ToTagsMap(src.([]*types.Tag)), nil
	//	}},
	//	utils.TypeConverter{SrcType: TagMap, DstType: TagsTypes, Fn: func(src interface{}) (dst interface{}, err error) {
	//		return logic.ToTagsType(src.(map[string]string)), nil
	//	}},
	//)

}

const Version = "v1.1.0"

func VersionUpdate(svcCtx *svc.ServiceContext) {
	err := svcCtx.AbnormalRepo.InitProduct(context.Background(), "")
	logx.Must(err)
}

func InitSubscribe(svcCtx *svc.ServiceContext) {
	//{
	//	cli, err := subDev.NewSubDev(svcCtx.Config.Event, svcCtx.NodeID)
	//	logx.Must(err)
	//	err = cli.Subscribe(func(ctx context.Context) subDev.InnerSubEvent {
	//		return deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx)
	//	})
	//	logx.Must(err)
	//}
	{
		cli, err := server.NewServer(svcCtx.Config.Event, svcCtx.NodeID)
		logx.Must(err)
		err = cli.Subscribe(func(ctx context.Context) server.ServerHandle {
			return serverEvent.NewServerHandle(ctx, svcCtx)
		})
		logx.Must(err)
	}
}

func InitCache(svcCtx *svc.ServiceContext) {
	{
		userDeviceShare, err := caches.NewCache(caches.CacheConfig[dm.UserDeviceShareInfo, userShared.UserShareKey]{
			KeyType:   eventBus.ServerCacheKeyDmUserShareDevice,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key userShared.UserShareKey) (*dm.UserDeviceShareInfo, error) {
				db := relationDB.NewUserDeviceShareRepo(ctx)
				f := relationDB.UserDeviceShareFilter{
					DeviceName:   key.DeviceName,
					ProductID:    key.ProductID,
					SharedUserID: key.SharedUserID,
				}
				uds, err := db.FindOneByFilter(ctx, f)
				if err != nil {
					return nil, err
				}
				pb := userdevicelogic.ToUserDeviceSharePb(uds)
				return pb, err
			},
			ExpireTime: 3 * time.Minute,
		})
		logx.Must(err)
		svcCtx.UserDeviceShare = userDeviceShare
	}
	{
		userMultiDeviceShare, err := caches.NewCache(caches.CacheConfig[dm.UserDeviceShareMultiInfo, string]{
			KeyType:   eventBus.ServerCacheKeyDmMultiDevicesShare,
			FastEvent: svcCtx.FastEvent,
			GetData: func(ctx context.Context, key string) (*dm.UserDeviceShareMultiInfo, error) {
				return &dm.UserDeviceShareMultiInfo{}, errors.Failure.WithMsg("分享已过期")
			},
			ExpireTime: 24 * time.Hour,
		})
		logx.Must(err)
		svcCtx.UserMultiDeviceShare = userMultiDeviceShare
	}

	productCache, err := caches.NewCache(caches.CacheConfig[dm.ProductInfo, string]{
		KeyType:   eventBus.ServerCacheKeyDmProduct,
		FastEvent: svcCtx.FastEvent,
		GetData: func(ctx context.Context, key string) (*dm.ProductInfo, error) {
			db := relationDB.NewProductInfoRepo(ctx)
			pi, err := db.FindOneByFilter(ctx, relationDB.ProductFilter{
				ProductIDs: []string{key}, WithProtocol: true, WithCategory: true})
			if err != nil {
				return nil, err
			}
			pb := logic.ToProductInfo(ctx, svcCtx, pi)
			return pb, err
		},
		ExpireTime: 3 * time.Minute,
	})
	logx.Must(err)
	svcCtx.ProductCache = productCache
	deviceCache, err := caches.NewCache(caches.CacheConfig[dm.DeviceInfo, devices.Core]{
		KeyType:   eventBus.ServerCacheKeyDmDevice,
		FastEvent: svcCtx.FastEvent,
		GetData: func(ctx context.Context, key devices.Core) (*dm.DeviceInfo, error) {
			ctx = ctxs.WithRoot(ctx)
			db := relationDB.NewDeviceInfoRepo(ctx)
			di, err := db.FindOneByFilter(ctx, relationDB.DeviceFilter{
				ProductID: key.ProductID, DeviceNames: []string{key.DeviceName}})
			if err != nil {
				return nil, err
			}
			pb := logic.ToDeviceInfo(ctx, svcCtx, di)
			return pb, err
		},
		ExpireTime: 3 * time.Minute,
	})
	logx.Must(err)
	svcCtx.DeviceCache = deviceCache

	relationDB.ClearDeviceInfo = func(ctx context.Context, dev devices.Core) error {
		return svcCtx.DeviceCache.SetData(ctx, dev, nil)
	}
}

func InitEventBus(svcCtx *svc.ServiceContext) {
	{ //设备数据订阅
		f := func(ctx context.Context, msg []byte, ff func(ctx context.Context, msg *deviceMsg.PublishMsg) error) error {
			ctx = ctxs.WithRoot(ctx)
			defer utils.Recover(ctx)
			ele, err := deviceMsg.GetDevPublish(ctx, msg)
			if err != nil {
				logx.WithContext(ctx).Errorf("%s.GetDevPublish failure err:%v", utils.FuncName(), err)
				return err
			}
			return ff(ctx, ele)
		}
		err := svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpThingAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Thing(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpOtaAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Ota(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpExtAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Ext(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpConfigAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Config(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpSDKLogAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).SDKLog(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpShadowAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Shadow(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpGatewayAll,
			func(ctx context.Context, t time.Time, body []byte) error {
				return f(ctx, body, func(ctx context.Context, msg *deviceMsg.PublishMsg) error {
					err := deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Gateway(msg)
					if err != nil {
						logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
						return err
					}
					return nil
				})
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpStatusConnected,
			func(ctx context.Context, t time.Time, body []byte) error {
				ctx = ctxs.WithRoot(ctx)
				ele, err := deviceStatus.GetDevConnMsg(ctx, body)
				if err != nil {
					logx.WithContext(ctx).Errorf("%s.GetDevConnMsg failure err:%v", utils.FuncName(), err)
					return err
				}
				err = deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Connected(ele)
				if err != nil {
					logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
					return err
				}
				return nil
			})
		logx.Must(err)
		err = svcCtx.FastEvent.QueueSubscribe(topics.DeviceUpStatusDisconnected,
			func(ctx context.Context, t time.Time, body []byte) error {
				ctx = ctxs.WithRoot(ctx)
				ele, err := deviceStatus.GetDevConnMsg(ctx, body)
				if err != nil {
					logx.WithContext(ctx).Errorf("%s.GetDevConnMsg failure err:%v", utils.FuncName(), err)
					return err
				}
				err = deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx).Disconnected(ele)
				if err != nil {
					logx.WithContext(ctx).Errorf("%s.Thing failure err:%v", utils.FuncName(), err)
					return err
				}
				return nil
			})
		logx.Must(err)
	}

	err := svcCtx.FastEvent.QueueSubscribe(eventBus.CoreUserDelete, func(ctx context.Context, t time.Time, body []byte) error {
		var value application.IDs
		err := json.Unmarshal(body, &value)
		if err != nil {
			return err
		}
		logx.WithContext(ctx).Infof("CoreUserDelete value:%v err:%v", utils.Fmt(value), err)
		ctx = ctxs.WithRoot(ctx)
		dis, err := relationDB.NewDeviceInfoRepo(ctx).FindByFilter(ctx, relationDB.DeviceFilter{UserIDs: value.IDs}, nil)
		for _, v := range dis {
			_, err := devicemanagelogic.NewDeviceInfoUnbindLogic(ctx, svcCtx).DeviceInfoUnbind(&dm.DeviceCore{
				ProductID:  v.ProductID,
				DeviceName: v.DeviceName,
			})
			if err != nil {
				logx.WithContext(ctx).Errorf("DeviceInfoUnbind dev:%v err:%v", utils.Fmt(v), err)
			}
		}
		err = relationDB.NewUserDeviceShareRepo(ctx).DeleteByFilter(ctx, relationDB.UserDeviceShareFilter{SharedUserIDs: value.IDs})
		if err != nil {
			logx.WithContext(ctx).Errorf("NewUserDeviceShareRepo.DeleteByFilter err:%v", err)
		}
		return nil
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.CoreProjectInfoDelete, func(ctx context.Context, t time.Time, body []byte) error {
		pi := cast.ToInt64(string(body))
		logx.WithContext(ctx).Infof("CoreProjectInfoDelete value:%v err:%v", string(body), err)
		if pi == 0 {
			return nil
		}
		ctx = ctxs.WithRoot(ctx)
		dis, err := relationDB.NewDeviceInfoRepo(ctx).FindByFilter(ctx, relationDB.DeviceFilter{ProjectIDs: []int64{pi}}, nil)
		if err != nil {
			logx.WithContext(ctx).Error(err)
			return err
		}
		for _, v := range dis {
			_, err := devicemanagelogic.NewDeviceInfoUnbindLogic(ctx, svcCtx).DeviceInfoUnbind(&dm.DeviceCore{
				ProductID:  v.ProductID,
				DeviceName: v.DeviceName,
			})
			if err != nil {
				logx.WithContext(ctx).Errorf("DeviceInfoUnbind dev:%v err:%v", utils.Fmt(v), err)
			}
		}
		return nil
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.CoreAreaInfoDelete, func(ctx context.Context, t time.Time, body []byte) error {
		var value application.IDs
		err := json.Unmarshal(body, &value)
		if err != nil {
			return err
		}
		logx.WithContext(ctx).Infof("CoreAreaInfoDelete value:%v err:%v", utils.Fmt(value), err)
		if len(value.IDs) == 0 {
			return nil
		}
		ctx = ctxs.WithRoot(ctx)
		dis, err := relationDB.NewDeviceInfoRepo(ctx).FindByFilter(ctx, relationDB.DeviceFilter{AreaIDs: value.IDs}, nil)
		if err != nil {
			logx.WithContext(ctx).Error(err)
			return err
		}
		var devs []*dm.DeviceCore
		for _, v := range dis {
			devs = append(devs, &dm.DeviceCore{
				ProductID:  v.ProductID,
				DeviceName: v.DeviceName,
			})
		}
		_, err = devicemanagelogic.NewDeviceInfoMultiUpdateLogic(ctx, svcCtx).DeviceInfoMultiUpdate(&dm.DeviceInfoMultiUpdateReq{
			Devices: devs,
			AreaID:  def.NotClassified,
		})
		if err != nil {
			logx.WithContext(ctx).Errorf("DeviceInfoMultiUpdate dev:%v err:%v", utils.Fmt(devs), err)
		}
		return nil
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmOtaJobDelayRun, func(ctx context.Context, t time.Time, body []byte) error {
		return otaEvent.NewOtaEvent(svcCtx, ctxs.WithRoot(ctx)).JobDelayRun(cast.ToInt64(string(body)))
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmOtaDeviceUpgradePush, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		return otaEvent.NewOtaEvent(svcCtx, ctxs.WithRoot(ctx)).DeviceUpgradePush()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmDeviceOnlineStatusChange, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		return serverEvent.NewServerHandle(ctxs.WithRoot(ctx), svcCtx).OnlineStatusHandle()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmDeviceStaticOneHour, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		return staticEvent.NewOneHourHandle(ctxs.WithRoot(ctx), svcCtx).Handle()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmDeviceStaticHalfHour, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		return staticEvent.NewHalfHourHandle(ctxs.WithRoot(ctx), svcCtx).Handle()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.QueueSubscribe(eventBus.DmDeviceStaticOneMinute, func(ctx context.Context, t time.Time, body []byte) error {
		if t.Before(time.Now().Add(-time.Second * 2)) { //2秒之前的跳过
			return nil
		}
		return staticEvent.NewOneMinuteHandle(ctxs.WithRoot(ctx), svcCtx).Handle()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.Start()
	logx.Must(err)
}

func TimerInit(svcCtx *svc.ServiceContext) {
	ctx := context.Background()
	_, err := svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedUnitedRhinoQueueGroupCode,                                          //组编码
		Type:      1,                                                                           //任务类型 1 定时任务 2 延时任务
		Name:      "联犀 ota升级定时任务",                                                      // 任务名称
		Code:      "iThingsOtaDeviceUpgradePush",                                               //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmOtaDeviceUpgradePush), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 5s",                                                                 // cron执行表达式
		Status:    def.StatusWaitRun,                                                           // 状态
		Priority:  3,                                                                           //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	if err != nil && !errors.Cmp(errors.Fmt(err), errors.Duplicate) {
		logx.Must(err)
	}
	_, err = svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedUnitedRhinoQueueGroupCode,                                              //组编码
		Type:      1,                                                                               //任务类型 1 定时任务 2 延时任务
		Name:      "联犀 设备在线状态改变处理",                                                     // 任务名称
		Code:      "dmDeviceOnlineStatusChange",                                                    //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmDeviceOnlineStatusChange), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 1s",                                                                     // cron执行表达式
		Status:    def.StatusWaitRun,                                                               // 状态
		Priority:  3,                                                                               //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	if err != nil && !errors.Cmp(errors.Fmt(err), errors.Duplicate) {
		logx.Must(err)
	}
	_, err = svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedUnitedRhinoQueueGroupCode,                                          //组编码
		Type:      1,                                                                           //任务类型 1 定时任务 2 延时任务
		Name:      "联犀 设备半小时统计",                                                       // 任务名称
		Code:      "dmDeviceStaticHalfHour",                                                    //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmDeviceStaticHalfHour), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 30m",                                                                // cron执行表达式
		Status:    def.StatusWaitRun,                                                           // 状态
		Priority:  3,                                                                           //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	_, err = svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedUnitedRhinoQueueGroupCode,                                         //组编码
		Type:      1,                                                                          //任务类型 1 定时任务 2 延时任务
		Name:      "联犀 设备1小时统计",                                                       // 任务名称
		Code:      "dmDeviceStaticOneHour",                                                    //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmDeviceStaticOneHour), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 60m",                                                               // cron执行表达式
		Status:    def.StatusWaitRun,                                                          // 状态
		Priority:  3,                                                                          //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	_, err = svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedUnitedRhinoQueueGroupCode,                                           //组编码
		Type:      1,                                                                            //任务类型 1 定时任务 2 延时任务
		Name:      "联犀 设备1分钟统计",                                                         // 任务名称
		Code:      "dmDeviceStaticOneMinute",                                                    //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmDeviceStaticOneMinute), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 1m",                                                                  // cron执行表达式
		Status:    def.StatusWaitRun,                                                            // 状态
		Priority:  3,                                                                            //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	if err != nil && !errors.Cmp(errors.Fmt(err), errors.Duplicate) {
		logx.Must(err)
	}
}

// 用到的字典初始化
func DictInit(svcCtx *svc.ServiceContext) {
	ctx := ctxs.WithRoot(context.Background())
	svcCtx.DictM.DictInfoCreate(ctx, &sys.DictInfo{
		Group: def.DictGroupThings,
		Name:  "设备分组用途",
		Code:  deviceGroup.DictCode,
	})
	svcCtx.DictM.DictDetailCreate(ctx, &sys.DictDetail{
		DictCode: deviceGroup.DictCode,
		Label:    "默认",
		Value:    deviceGroup.DictDefault,
		Sort:     1,
	})
	svcCtx.DictM.DictInfoCreate(ctx, &sys.DictInfo{
		Group: def.DictGroupThings,
		Name:  "设备异常类型",
		Code:  "deviceAbnormal",
	})
	svcCtx.DictM.DictDetailCreate(ctx, &sys.DictDetail{
		DictCode: "deviceAbnormal",
		Label:    "上下线异常",
		Value:    "online",
		Sort:     1,
	})
}
