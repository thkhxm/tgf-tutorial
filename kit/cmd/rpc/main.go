package main

import (
	"github.com/thkhxm/tgf/tgf-tutorial/common/services"
	"github.com/thkhxm/tgf/tgf-tutorial/prop"
	"github.com/thkhxm/tgf/util"
)

//***************************************************
//@Link  https://github.com/thkhxm/tgf
//@Link  https://gitee.com/timgame/tgf
//@QQ群 7400585
//author tim.huang<thkhxm@gmail.com>
//@Description
//2023/12/10
//***************************************************

func main() {
	//设置导出的golang目录
	util.SetAutoGenerateAPICodePath("../common/api")
	//设置生成的文件后缀
	util.SetGenerateFileNameSuffix("rpc")
	//根据接口生成对应的rpc结构
	util.GeneratorRPC[services.IPropRPCService](prop.ModuleName, prop.Version, "propservice", "prop")
}
