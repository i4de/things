package tenantmanagelogic

import (
	"context"
	"database/sql"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/logic"
	usermanagelogic "github.com/i-Things/things/src/syssvr/internal/logic/usermanage"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"
	"gorm.io/gorm"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantInfoCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantInfoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantInfoCreateLogic {
	return &TenantInfoCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增区域
func (l *TenantInfoCreateLogic) TenantInfoCreate(in *sys.TenantInfoCreateReq) (*sys.Response, error) {
	err := logic.IsSupperAdmin(l.ctx, def.TenantCodeDefault)
	if err != nil {
		return nil, err
	}
	userInfo := in.AdminUserInfo
	//首先校验账号格式使用正则表达式，对用户账号做格式校验：只能是大小写字母，数字和下划线，减号
	err = usermanagelogic.CheckUserName(userInfo.UserName)
	if err != nil {
		return nil, err
	}
	//校验密码强度
	err = usermanagelogic.CheckPwd(l.svcCtx, userInfo.Password)
	if err != nil {
		return nil, err
	}
	//1.生成uid
	userID := l.svcCtx.UserID.GetSnowflakeId()
	//2.对密码进行md5加密
	password := utils.MakePwd(userInfo.Password, userID, false)
	ui := relationDB.SysUserInfo{
		TenantCode: stores.TenantCode(in.Info.Code),
		UserID:     userID,
		Phone:      utils.AnyToNullString(userInfo.Phone),
		Email:      utils.AnyToNullString(userInfo.Email),
		UserName:   sql.NullString{String: userInfo.UserName, Valid: true},
		Password:   password,
		NickName:   userInfo.NickName,
		City:       userInfo.City,
		Country:    userInfo.Country,
		Province:   userInfo.Province,
		Language:   userInfo.Language,
		HeadImgUrl: userInfo.HeadImgUrl,
		Role:       userInfo.Role,
		Sex:        userInfo.Sex,
		IsAllData:  def.True,
	}
	ctxs.GetUserCtx(l.ctx).AllData = true
	defer func() {
		ctxs.GetUserCtx(l.ctx).AllData = false
	}()
	err = stores.GetCommonConn(l.ctx).Transaction(func(tx *gorm.DB) error {
		ri := relationDB.SysRoleInfo{TenantCode: stores.TenantCode(in.Info.Code), Name: "超级管理员"}
		err = relationDB.NewRoleInfoRepo(tx).Insert(l.ctx, &ri)
		ui.Role = ri.ID
		err = relationDB.NewUserInfoRepo(tx).Insert(l.ctx, &ui)
		if err != nil {
			return err
		}
		po := ToTenantInfoPo(in.Info)
		po.AdminUserID = ui.UserID
		err = relationDB.NewTenantInfoRepo(l.ctx).Insert(l.ctx, po)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &sys.Response{}, err
}