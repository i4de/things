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
	ApiCreateReq          = sys.ApiCreateReq
	ApiData               = sys.ApiData
	ApiDeleteReq          = sys.ApiDeleteReq
	ApiIndexReq           = sys.ApiIndexReq
	ApiIndexResp          = sys.ApiIndexResp
	ApiUpdateReq          = sys.ApiUpdateReq
	AuthApiIndexReq       = sys.AuthApiIndexReq
	AuthApiIndexResp      = sys.AuthApiIndexResp
	AuthApiInfo           = sys.AuthApiInfo
	AuthApiMultiUpdateReq = sys.AuthApiMultiUpdateReq
	CheckAuthReq          = sys.CheckAuthReq
	ConfigResp            = sys.ConfigResp
	DateRange             = sys.DateRange
	JwtToken              = sys.JwtToken
	LoginLogCreateReq     = sys.LoginLogCreateReq
	LoginLogIndexData     = sys.LoginLogIndexData
	LoginLogIndexReq      = sys.LoginLogIndexReq
	LoginLogIndexResp     = sys.LoginLogIndexResp
	Map                   = sys.Map
	MenuCreateReq         = sys.MenuCreateReq
	MenuData              = sys.MenuData
	MenuDeleteReq         = sys.MenuDeleteReq
	MenuIndexReq          = sys.MenuIndexReq
	MenuIndexResp         = sys.MenuIndexResp
	MenuUpdateReq         = sys.MenuUpdateReq
	OperLogCreateReq      = sys.OperLogCreateReq
	OperLogIndexData      = sys.OperLogIndexData
	OperLogIndexReq       = sys.OperLogIndexReq
	OperLogIndexResp      = sys.OperLogIndexResp
	PageInfo              = sys.PageInfo
	Response              = sys.Response
	RoleCreateReq         = sys.RoleCreateReq
	RoleDeleteReq         = sys.RoleDeleteReq
	RoleIndexData         = sys.RoleIndexData
	RoleIndexReq          = sys.RoleIndexReq
	RoleIndexResp         = sys.RoleIndexResp
	RoleMenuUpdateReq     = sys.RoleMenuUpdateReq
	RoleUpdateReq         = sys.RoleUpdateReq
	UserCheckTokenReq     = sys.UserCheckTokenReq
	UserCheckTokenResp    = sys.UserCheckTokenResp
	UserCreateReq         = sys.UserCreateReq
	UserCreateResp        = sys.UserCreateResp
	UserDeleteReq         = sys.UserDeleteReq
	UserIndexReq          = sys.UserIndexReq
	UserIndexResp         = sys.UserIndexResp
	UserInfo              = sys.UserInfo
	UserLoginReq          = sys.UserLoginReq
	UserLoginResp         = sys.UserLoginResp
	UserReadReq           = sys.UserReadReq
	UserReadResp          = sys.UserReadResp
	UserUpdateReq         = sys.UserUpdateReq

	Role interface {
		RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error)
		RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error)
		RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultRole struct {
		cli zrpc.Client
	}

	directRole struct {
		svcCtx *svc.ServiceContext
		svr    sys.RoleServer
	}
)

func NewRole(cli zrpc.Client) Role {
	return &defaultRole{
		cli: cli,
	}
}

func NewDirectRole(svcCtx *svc.ServiceContext, svr sys.RoleServer) Role {
	return &directRole{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRole) RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleCreate(ctx, in, opts...)
}

func (d *directRole) RoleCreate(ctx context.Context, in *RoleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleCreate(ctx, in)
}

func (m *defaultRole) RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleIndex(ctx, in, opts...)
}

func (d *directRole) RoleIndex(ctx context.Context, in *RoleIndexReq, opts ...grpc.CallOption) (*RoleIndexResp, error) {
	return d.svr.RoleIndex(ctx, in)
}

func (m *defaultRole) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (d *directRole) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleUpdate(ctx, in)
}

func (m *defaultRole) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (d *directRole) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleDelete(ctx, in)
}

func (m *defaultRole) RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleClient(m.cli.Conn())
	return client.RoleMenuUpdate(ctx, in, opts...)
}

func (d *directRole) RoleMenuUpdate(ctx context.Context, in *RoleMenuUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleMenuUpdate(ctx, in)
}
