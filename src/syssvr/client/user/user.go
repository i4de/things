// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiCreateReq               = sys.ApiCreateReq
	ApiData                    = sys.ApiData
	ApiDeleteReq               = sys.ApiDeleteReq
	ApiIndexReq                = sys.ApiIndexReq
	ApiIndexResp               = sys.ApiIndexResp
	ApiUpdateReq               = sys.ApiUpdateReq
	AuthorityApiIndexReq       = sys.AuthorityApiIndexReq
	AuthorityApiIndexResp      = sys.AuthorityApiIndexResp
	AuthorityApiInfo           = sys.AuthorityApiInfo
	AuthorityApiMultiUpdateReq = sys.AuthorityApiMultiUpdateReq
	CheckAuthReq               = sys.CheckAuthReq
	CheckTokenReq              = sys.CheckTokenReq
	CheckTokenResp             = sys.CheckTokenResp
	ConfigResp                 = sys.ConfigResp
	DateRange                  = sys.DateRange
	JwtToken                   = sys.JwtToken
	LoginLogCreateReq          = sys.LoginLogCreateReq
	LoginLogIndexData          = sys.LoginLogIndexData
	LoginLogIndexReq           = sys.LoginLogIndexReq
	LoginLogIndexResp          = sys.LoginLogIndexResp
	LoginReq                   = sys.LoginReq
	LoginResp                  = sys.LoginResp
	Map                        = sys.Map
	MenuCreateReq              = sys.MenuCreateReq
	MenuData                   = sys.MenuData
	MenuDeleteReq              = sys.MenuDeleteReq
	MenuIndexReq               = sys.MenuIndexReq
	MenuIndexResp              = sys.MenuIndexResp
	MenuUpdateReq              = sys.MenuUpdateReq
	OperLogCreateReq           = sys.OperLogCreateReq
	OperLogIndexData           = sys.OperLogIndexData
	OperLogIndexReq            = sys.OperLogIndexReq
	OperLogIndexResp           = sys.OperLogIndexResp
	PageInfo                   = sys.PageInfo
	Response                   = sys.Response
	RoleCreateReq              = sys.RoleCreateReq
	RoleDeleteReq              = sys.RoleDeleteReq
	RoleIndexData              = sys.RoleIndexData
	RoleIndexReq               = sys.RoleIndexReq
	RoleIndexResp              = sys.RoleIndexResp
	RoleMenuUpdateReq          = sys.RoleMenuUpdateReq
	RoleUpdateReq              = sys.RoleUpdateReq
	UserCreateReq              = sys.UserCreateReq
	UserCreateResp             = sys.UserCreateResp
	UserDeleteReq              = sys.UserDeleteReq
	UserIndexReq               = sys.UserIndexReq
	UserIndexResp              = sys.UserIndexResp
	UserInfo                   = sys.UserInfo
	UserReadReq                = sys.UserReadReq
	UserReadResp               = sys.UserReadResp
	UserUpdateReq              = sys.UserUpdateReq

	User interface {
		Create(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error)
		Index(ctx context.Context, in *UserIndexReq, opts ...grpc.CallOption) (*UserIndexResp, error)
		Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error)
		Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadResp, error)
		Delete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*Response, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error)
		CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}

	directUser struct {
		svcCtx *svc.ServiceContext
		svr    sys.UserServer
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func NewDirectUser(svcCtx *svc.ServiceContext, svr sys.UserServer) User {
	return &directUser{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultUser) Create(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (d *directUser) Create(ctx context.Context, in *UserCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error) {
	return d.svr.Create(ctx, in)
}

func (m *defaultUser) Index(ctx context.Context, in *UserIndexReq, opts ...grpc.CallOption) (*UserIndexResp, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Index(ctx, in, opts...)
}

func (d *directUser) Index(ctx context.Context, in *UserIndexReq, opts ...grpc.CallOption) (*UserIndexResp, error) {
	return d.svr.Index(ctx, in)
}

func (m *defaultUser) Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (d *directUser) Update(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.Update(ctx, in)
}

func (m *defaultUser) Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadResp, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Read(ctx, in, opts...)
}

func (d *directUser) Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadResp, error) {
	return d.svr.Read(ctx, in)
}

func (m *defaultUser) Delete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

func (d *directUser) Delete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.Delete(ctx, in)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (d *directUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	return d.svr.Login(ctx, in)
}

func (m *defaultUser) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.CheckToken(ctx, in, opts...)
}

func (d *directUser) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error) {
	return d.svr.CheckToken(ctx, in)
}

func (m *defaultUser) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserClient(m.cli.Conn())
	return client.CheckAuth(ctx, in, opts...)
}

func (d *directUser) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.CheckAuth(ctx, in)
}
