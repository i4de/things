package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/dmsvr/internal/config"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceLog"
	"github.com/i-Things/things/src/dmsvr/internal/event/eventDevSub"
	"github.com/i-Things/things/src/dmsvr/internal/repo/event/innerLink"
	"github.com/i-Things/things/src/dmsvr/internal/server"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "net/http/pprof"
	"time"
)

var configFile = flag.String("f", "etc/dm.yaml", "the config file")

func main() {
	flag.Parse()
	//go device.NewDevice()
	//device.TestMongo()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	Test(svcCtx.DeviceLogRepo)
	svcCtx.InnerLink.Subscribe(func(ctx context.Context) innerLink.InnerSubHandle {
		return eventDevSub.NewDeviceMsgHandle(ctx, svcCtx)
	})
	//grpc服务初始化

	srv := server.NewDmServer(svcCtx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		dm.RegisterDmServer(grpcServer, srv)
		reflection.Register(grpcServer)
	})
	defer s.Stop()
	s.AddUnaryInterceptors(errors.ErrorInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func Test(DeviceDataRepo deviceLog.DeviceLogRepo) {
	var (
		err error
	)
	if true {
		err = DeviceDataRepo.InitProduct(context.Background(),
			"23FIPSIJPsk")
		fmt.Println(err)
	}
	if false {
		err = DeviceDataRepo.DropProduct(context.Background(),
			"23FIPSIJPsk")
		fmt.Println(err)
	}
	if false {
		err = DeviceDataRepo.DropDevice(context.Background(),
			"23FIPSIJPsk", "test5")
		fmt.Println(err)
	}
	if true {
		err = DeviceDataRepo.Insert(context.Background(),
			&deviceLog.DeviceLog{
				ProductID:  "23FIPSIJPsk",
				DeviceName: "test5",
				Content:    "ade",
				Topic:      "dfae",
				Action:     "dsfas",
				Timestamp:  time.Now(),
				RequestID:  "dfgasdf",
				TranceID:   "dgasdf",
				ResultType: 44,
			})
		fmt.Println(err)
	}
	if true {
		log, err := DeviceDataRepo.GetDeviceLog(context.Background(),
			"23FIPSIJPsk", "test5", def.PageInfo2{})
		fmt.Println(log, err)
	}
	return
}
