package prop

import (
	"context"
	"github.com/thkhxm/tgf/component"
	"github.com/thkhxm/tgf/log"
	"github.com/thkhxm/tgf/rpc"
	"github.com/thkhxm/tgf/tgf-tutorial/common/conf"
	"github.com/thkhxm/tgf/tgf-tutorial/common/model"
	"github.com/thkhxm/tgf/util"
	"math/rand"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/10
//***************************************************

var (
	ModuleName = "prop"
	Version    = "v1.0.0"
)

type service struct {
	rpc.Module
	propCountCache map[string]int32
}

func (s *service) GetUserPropCount(ctx context.Context, args *model.GetUserPropArgs, reply *model.GetUserPropReply) (err error) {
	userId := rpc.GetUserId(ctx)
	reply.Count = s.propCountCache[args.PropId]
	if propConfig, h := component.GetGameConf[*conf.PropConf](args.PropId); h {
		reply.Name = propConfig.Name
	}
	log.DebugTag("prop", "get %s user %s prop %s count %d ", userId, args.PropId, reply.Name, reply.Count)

	return
}

func (s *service) GetName() string {
	return ModuleName
}

func (s *service) GetVersion() string {
	return Version
}

func (s *service) Startup() (bool, error) {
	s.propCountCache = make(map[string]int32)
	r := rand.New(rand.NewSource(12345))
	for i := 0; i < 10; i++ {
		propId, _ := util.AnyToStr(i)
		s.propCountCache[propId] = r.Int31n(100)
	}
	return true, nil
}

func (s *service) Destroy(sub rpc.IService) {

}

func NewService() rpc.IService {
	return &service{}
}
