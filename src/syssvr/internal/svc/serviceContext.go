package svc

import (
	"context"
	"github.com/casbin/casbin/v2"
	cas "github.com/i-Things/things/shared/casbin"
	"github.com/i-Things/things/shared/clients"
	"github.com/i-Things/things/shared/oss"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/config"
	"github.com/i-Things/things/src/syssvr/internal/repo/cache"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"os"
)

type ServiceContext struct {
	Config        config.Config
	ProjectID     *utils.SnowFlake
	AreaID        *utils.SnowFlake
	WxMiniProgram *clients.MiniProgram
	UserID        *utils.SnowFlake
	Casbin        *casbin.Enforcer
	OssClient     *oss.Client
	Store         kv.Store
	PwdCheck      *cache.PwdCheck
	Captcha       *cache.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	//conn := sqlx.NewMysql(c.Database.DSN)
	stores.InitConn(c.Database)
	err := relationDB.Migrate(c.Database)
	if err != nil {
		logx.Error("syssvr 数据库初始化失败 err", err)
		os.Exit(-1)
	}
	// 自动迁移数据库
	db := stores.GetCommonConn(context.Background())
	nodeID := utils.GetNodeID(c.CacheRedis, c.Name)
	ProjectID := utils.NewSnowFlake(nodeID)
	AreaID := utils.NewSnowFlake(nodeID)
	WxMiniProgram := clients.NewWxMiniProgram(context.Background(), c.WxMiniProgram, c.CacheRedis)
	nodeId := utils.GetNodeID(c.CacheRedis, c.Name)
	UserID := utils.NewSnowFlake(nodeId)
	dbRaw, err := db.DB()
	if err != nil {
		logx.Error("sys failed to  database err: %v", err)
	}
	ca := cas.NewCasbinWithRedisWatcher(dbRaw, c.Database.DBType, c.CacheRedis[0].RedisConf)
	store := kv.NewStore(c.CacheRedis)
	ossClient, err := oss.NewOssClient(c.OssConf)
	if err != nil {
		logx.Errorf("NewOss err err:%v", err)
		os.Exit(-1)
	}
	return &ServiceContext{
		Captcha:       cache.NewCaptcha(store),
		PwdCheck:      cache.NewPwdCheck(store),
		Config:        c,
		ProjectID:     ProjectID,
		OssClient:     ossClient,
		AreaID:        AreaID,
		WxMiniProgram: WxMiniProgram,
		UserID:        UserID,
		Casbin:        ca,
		Store:         store,
	}
}
