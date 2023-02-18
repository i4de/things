// Code generated by goctl. DO NOT EDIT.
// Source: rule.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/rulesvr/internal/svc"
	"github.com/i-Things/things/src/rulesvr/pb/rule"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FlowInfo           = rule.FlowInfo
	FlowInfoDeleteReq  = rule.FlowInfoDeleteReq
	FlowInfoIndexReq   = rule.FlowInfoIndexReq
	FlowInfoIndexResp  = rule.FlowInfoIndexResp
	FlowInfoReadReq    = rule.FlowInfoReadReq
	PageInfo           = rule.PageInfo
	Response           = rule.Response
	SceneInfo          = rule.SceneInfo
	SceneInfoDeleteReq = rule.SceneInfoDeleteReq
	SceneInfoIndexReq  = rule.SceneInfoIndexReq
	SceneInfoIndexResp = rule.SceneInfoIndexResp
	SceneInfoReadReq   = rule.SceneInfoReadReq

	RuleEngine interface {
		FlowInfoCreate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error)
		FlowInfoUpdate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error)
		FlowInfoDelete(ctx context.Context, in *FlowInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		FlowInfoIndex(ctx context.Context, in *FlowInfoIndexReq, opts ...grpc.CallOption) (*FlowInfoIndexResp, error)
		FlowInfoRead(ctx context.Context, in *FlowInfoReadReq, opts ...grpc.CallOption) (*FlowInfo, error)
	}

	defaultRuleEngine struct {
		cli zrpc.Client
	}

	directRuleEngine struct {
		svcCtx *svc.ServiceContext
		svr    rule.RuleEngineServer
	}
)

func NewRuleEngine(cli zrpc.Client) RuleEngine {
	return &defaultRuleEngine{
		cli: cli,
	}
}

func NewDirectRuleEngine(svcCtx *svc.ServiceContext, svr rule.RuleEngineServer) RuleEngine {
	return &directRuleEngine{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRuleEngine) FlowInfoCreate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewRuleEngineClient(m.cli.Conn())
	return client.FlowInfoCreate(ctx, in, opts...)
}

func (d *directRuleEngine) FlowInfoCreate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.FlowInfoCreate(ctx, in)
}

func (m *defaultRuleEngine) FlowInfoUpdate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewRuleEngineClient(m.cli.Conn())
	return client.FlowInfoUpdate(ctx, in, opts...)
}

func (d *directRuleEngine) FlowInfoUpdate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.FlowInfoUpdate(ctx, in)
}

func (m *defaultRuleEngine) FlowInfoDelete(ctx context.Context, in *FlowInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := rule.NewRuleEngineClient(m.cli.Conn())
	return client.FlowInfoDelete(ctx, in, opts...)
}

func (d *directRuleEngine) FlowInfoDelete(ctx context.Context, in *FlowInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.FlowInfoDelete(ctx, in)
}

func (m *defaultRuleEngine) FlowInfoIndex(ctx context.Context, in *FlowInfoIndexReq, opts ...grpc.CallOption) (*FlowInfoIndexResp, error) {
	client := rule.NewRuleEngineClient(m.cli.Conn())
	return client.FlowInfoIndex(ctx, in, opts...)
}

func (d *directRuleEngine) FlowInfoIndex(ctx context.Context, in *FlowInfoIndexReq, opts ...grpc.CallOption) (*FlowInfoIndexResp, error) {
	return d.svr.FlowInfoIndex(ctx, in)
}

func (m *defaultRuleEngine) FlowInfoRead(ctx context.Context, in *FlowInfoReadReq, opts ...grpc.CallOption) (*FlowInfo, error) {
	client := rule.NewRuleEngineClient(m.cli.Conn())
	return client.FlowInfoRead(ctx, in, opts...)
}

func (d *directRuleEngine) FlowInfoRead(ctx context.Context, in *FlowInfoReadReq, opts ...grpc.CallOption) (*FlowInfo, error) {
	return d.svr.FlowInfoRead(ctx, in)
}
