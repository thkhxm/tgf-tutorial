package prop

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
//2023/12/10
//***************************************************

func Startup() {
	r := rpc.NewRPCServer().
		WithRandomServicePort(8021, 8030).
		WithCache(tgf.CacheModuleClose).
		WithService(NewService()).
		Run()

	<-r
}
