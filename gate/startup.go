package gate

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/rpc"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/9
//***************************************************

func Startup() {
	r := rpc.NewRPCServer().
		WithGatewayWS("8443", "/tgf", nil).
		WithCache(tgf.CacheModuleClose).
		Run()

	<-r
}
