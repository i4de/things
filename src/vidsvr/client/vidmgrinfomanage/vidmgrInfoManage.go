// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package vidmgrinfomanage

import (
	"context"

	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PageInfo              = vid.PageInfo
	PageInfo_OrderBy      = vid.PageInfo_OrderBy
	Response              = vid.Response
	StreamTrack           = vid.StreamTrack
	VidPlanInfo           = vid.VidPlanInfo
	VidmgrConfig          = vid.VidmgrConfig
	VidmgrConfigDeleteReq = vid.VidmgrConfigDeleteReq
	VidmgrConfigIndexReq  = vid.VidmgrConfigIndexReq
	VidmgrConfigIndexResp = vid.VidmgrConfigIndexResp
	VidmgrConfigReadReq   = vid.VidmgrConfigReadReq
	VidmgrInfo            = vid.VidmgrInfo
	VidmgrInfoActiveReq   = vid.VidmgrInfoActiveReq
	VidmgrInfoCountReq    = vid.VidmgrInfoCountReq
	VidmgrInfoCountResp   = vid.VidmgrInfoCountResp
	VidmgrInfoDeleteReq   = vid.VidmgrInfoDeleteReq
	VidmgrInfoIndexReq    = vid.VidmgrInfoIndexReq
	VidmgrInfoIndexResp   = vid.VidmgrInfoIndexResp
	VidmgrInfoReadReq     = vid.VidmgrInfoReadReq
	VidmgrStream          = vid.VidmgrStream
	VidmgrStreamCountReq  = vid.VidmgrStreamCountReq
	VidmgrStreamCountResp = vid.VidmgrStreamCountResp
	VidmgrStreamCreateReq = vid.VidmgrStreamCreateReq
	VidmgrStreamDeleteReq = vid.VidmgrStreamDeleteReq
	VidmgrStreamIndexReq  = vid.VidmgrStreamIndexReq
	VidmgrStreamIndexResp = vid.VidmgrStreamIndexResp
	VidmgrStreamReadReq   = vid.VidmgrStreamReadReq
	VidrecordfileInfo     = vid.VidrecordfileInfo

	VidmgrInfoManage interface {
		// 新建服务
		VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
		// 激活服务
		VidmgrInfoActive(ctx context.Context, in *VidmgrInfoActiveReq, opts ...grpc.CallOption) (*Response, error)
		// 更新服务
		VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除服务
		VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取服务列表
		VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error)
		// 获取服务信息详情
		VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error)
		// 获取服务统计  在线，离线，未激活
		VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error)
	}

	defaultVidmgrInfoManage struct {
		cli zrpc.Client
	}

	directVidmgrInfoManage struct {
		svcCtx *svc.ServiceContext
		svr    vid.VidmgrInfoManageServer
	}
)

func NewVidmgrInfoManage(cli zrpc.Client) VidmgrInfoManage {
	return &defaultVidmgrInfoManage{
		cli: cli,
	}
}

func NewDirectVidmgrInfoManage(svcCtx *svc.ServiceContext, svr vid.VidmgrInfoManageServer) VidmgrInfoManage {
	return &directVidmgrInfoManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新建服务
func (m *defaultVidmgrInfoManage) VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoCreate(ctx, in, opts...)
}

// 新建服务
func (d *directVidmgrInfoManage) VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrInfoCreate(ctx, in)
}

// 激活服务
func (m *defaultVidmgrInfoManage) VidmgrInfoActive(ctx context.Context, in *VidmgrInfoActiveReq, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoActive(ctx, in, opts...)
}

// 激活服务
func (d *directVidmgrInfoManage) VidmgrInfoActive(ctx context.Context, in *VidmgrInfoActiveReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrInfoActive(ctx, in)
}

// 更新服务
func (m *defaultVidmgrInfoManage) VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoUpdate(ctx, in, opts...)
}

// 更新服务
func (d *directVidmgrInfoManage) VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrInfoUpdate(ctx, in)
}

// 删除服务
func (m *defaultVidmgrInfoManage) VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoDelete(ctx, in, opts...)
}

// 删除服务
func (d *directVidmgrInfoManage) VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrInfoDelete(ctx, in)
}

// 获取服务列表
func (m *defaultVidmgrInfoManage) VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoIndex(ctx, in, opts...)
}

// 获取服务列表
func (d *directVidmgrInfoManage) VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error) {
	return d.svr.VidmgrInfoIndex(ctx, in)
}

// 获取服务信息详情
func (m *defaultVidmgrInfoManage) VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoRead(ctx, in, opts...)
}

// 获取服务信息详情
func (d *directVidmgrInfoManage) VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error) {
	return d.svr.VidmgrInfoRead(ctx, in)
}

// 获取服务统计  在线，离线，未激活
func (m *defaultVidmgrInfoManage) VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error) {
	client := vid.NewVidmgrInfoManageClient(m.cli.Conn())
	return client.VidmgrInfoCount(ctx, in, opts...)
}

// 获取服务统计  在线，离线，未激活
func (d *directVidmgrInfoManage) VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error) {
	return d.svr.VidmgrInfoCount(ctx, in)
}
