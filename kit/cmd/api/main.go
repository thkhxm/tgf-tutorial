package main

import (
	"github.com/thkhxm/tgf/tgf-tutorial/common/services"
	"github.com/thkhxm/tgf/tgf-tutorial/hall"
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
	//设置导出的C#目录(用于Unity项目)
	//util.SetAutoGenerateAPICSCode("E:\\unity\\project\\t2\\Assets\\HotFix\\Code", "HotFix.Code")
	//根据接口生成对应的api结构
	util.GeneratorAPI[services.IHallService](hall.ModuleName, hall.Version, "api")
	//生成Cs代码(用于Unity)
	//util.GenerateCSApiService()
}
