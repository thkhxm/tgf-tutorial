package main

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/db"
	"github.com/thkhxm/tgf/util"
)

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2023/4/18
// ***************************************************
func main() {
	//关闭框架缓存
	db.WithCacheModule(tgf.CacheModuleClose)
	//设置excel路径
	util.SetExcelPath("./excel")
	//设置excel导出的go文件路径
	util.SetExcelToGoPath("../common/conf")
	//设置excel导出的json文件路径
	util.SetExcelToJsonPath("../common/conf/js")
	//给前端路径也导出一份
	//util.SetExcelToJsonPath("E:\\unity\\project\\t2\\Assets\\Config\\js")

	//设置excel导出的unity文件路径
	//util.SetExcelToUnityPath("E:\\unity\\project\\t2\\Assets\\HotFix\\Config")
	//设置excel导出的unity命名空间
	//util.SetExcelToUnityNamespace("HotFix.Config")

	//开始导出excel
	util.ExcelExport()

	//设置配置源数据路径
	//component.WithConfPath("../common/conf/js")
	//初始化游戏配置到内存中
	//component.InitGameConfToMem()
	//获取配置数据
	//codes := component.GetAllGameConf[*conf.ErrorCodeConf]()
	////初始化结构化kv数据源
	//data := make([]util.TemplateKeyValueData, len(codes), len(codes))
	//for i, code := range codes {
	//	data[i] = util.TemplateKeyValueData{
	//		FieldName: code.FieldName,
	//		Values:    code.Code,
	//	}
	//}
	////将数据源写入文件 生成kv结构文件
	//util.JsonToKeyValueGoFile("errorcodes", "error_code", "../common/define/errorcode", "int32", data)
	//
	//获取配置数据
	//c := component.GetAllGameConf[*conf.ConstantConf]()
	////初始化结构化kv数据源
	//d := make([]util.TemplateKeyValueData, len(c), len(c))
	//for i, code := range c {
	//	d[i] = util.TemplateKeyValueData{
	//		FieldName: code.Key,
	//		Values:    code.Key,
	//	}
	//}
	//将数据源写入文件 生成kv结构文件
	//util.JsonToKeyValueGoFile("constant", "constant_key", "../common/define/constant", "string", d)

}
