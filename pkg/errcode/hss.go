// Path: internal/pkg/errcode
// FileName: hss.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/26$ 19:18$

package errcode

// 模块代码
const (
	GafSysCode      = 1 * ModuleCode //系统中产生的错误代码
	GafInitCode     = 2 * ModuleCode //初始化模块
	GafUserCode     = 3 * ModuleCode //用户信息模块
	GafKeyCode      = 4 * ModuleCode //密钥管理
	GafSysConfCode  = 5 * ModuleCode //系统配置模块
	GafLogCode      = 6 * ModuleCode //日志模块
	GafSelfTestCode = 8 * ModuleCode //自检程序
)

// 系统运行中的错误
const (
	GafSysErr              = GafServer + GafSysCode + iota + 1
	GafSysJsonMarshalErr   //转json失败
	GafSysJsonUnMarshalErr //json解析失败
	GafSysTimeParseErr     //时间转换错误
	GafSysSaveFileErr      //保存文件错误
)

// 系统运行中数据库出现的错误
const (
	GafSysDBErr       = GafServer + GafSysCode + iota + 20 //数据库出错
	GafSysDBFindErr                                        //数据库查询出错
	GafSysDBCreateErr                                      //数据库添加数据出错
	GafSysDBUpdateErr                                      //数据库更新出错
	GafSysDBDeleteErr                                      //数据库删除出错
	GafSysCmdErr                                           //执行宿主机cmd指令出错
	GafSysNetworkErr                                       //网卡配置出错
)

// 系统运行中的缓存错误
const (
	GafSysCacheErr    = GafServer + GafSysCode + iota + 30 //缓存错误
	GafSysCacheGetErr                                      //缓存获取错误
	GafSysCacheSetErr                                      //缓存获取错误
)

// 高级配置模块
const (
	GafConfErr           = GafServer + GafSysConfCode + iota + 1 //高级配置
	GafUpdateFileErr                                             //升级包名称错误
	GafUpdateFileReadErr                                         //升级包目录中文件读取错误
	GafUpdateFileLoadErr                                         //升级包镜像导入错误
	GafUpdateAssistErr                                           // 请求助手升级失败
)
