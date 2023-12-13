package services

import (
	"context"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/tgf-tutorial/common/pb"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/10
//***************************************************

type IHallService interface {
	LoadUserData(ctx context.Context, args *rpc.Args[*pb.LoadUserDataRequest], reply *rpc.Reply[*pb.LoadUserDataResponse]) (err error)
	Login(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error)
	HelloWorld(ctx context.Context, args *rpc.Args[*pb.HelloWorldRequest], reply *rpc.Reply[*pb.HelloWorldResponse]) (err error)
	UpdatePassword(ctx context.Context, args *rpc.Args[*pb.LoginRequest], reply *rpc.Reply[*pb.LoginResponse]) (err error)
}

type IHallRPCService interface {
}
