package hall

import (
	"context"
	"fmt"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/tgf-tutorial/common/pb"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/9
//***************************************************

type service struct {
}

func (s *service) HelloWorld(ctx context.Context, args *rpc.Args[*pb.HelloWorldRequest], reply *rpc.Reply[*pb.HelloWorldResponse]) (err error) {
	// args.GetData() -> *pb.HelloWorldRequest
	msg := fmt.Sprintf("hello world %s", args.GetData().GetName())
	reply.SetData(&pb.HelloWorldResponse{Message: msg})
	return
}

func (s *service) GetName() string {
	return "hall"
}

func (s *service) GetVersion() string {
	return "1.0.0"
}

func (s *service) Startup() (bool, error) {
	log.DebugTag("hall", "hall startup")
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {
	log.DebugTag("hall", "hall destroy")
}

func NewService() rpc.IService {
	return &service{}
}
