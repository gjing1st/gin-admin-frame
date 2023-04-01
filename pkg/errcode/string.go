// Path: pkg/errcode
// FileName: string.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/4/1$ 23:03$

package errcode

func (e Err) String() string {
	switch e {
	//用户相关
	case SuccessCode:
		return "操作成功"
	case GafUserErr:
		return "用户错误"
	case GafUserLoginErr:
		return "登录错误"
	case GafUserNotFoundErr:
		return "用户不存在"
	case GafUserWithoutToken:
		return "未携带token"
	case GafUserTokenExpired:
		return "token过期"
	case GafUserRoleErr:
		return "用户的role_id错误"
	case GafUserForbiddenErr:
		return "用户没有权限"
	case GafUserHasExist:
		return "用户已存在"
	// 高级配置模块
	case GafConfErr:
		return "高级配置"
	case GafUpdateFileErr:
		return "升级包名称错误"
	case GafUpdateFileReadErr:
		return "升级包目录中文件读取错误"
	case GafUpdateFileLoadErr:
		return "升级包镜像导入错误"
	case GafUpdateAssistErr:
		return "请求助手升级失败"
	// 系统运行中的缓存错误
	case GafSysCacheErr:
		return "缓存错误"
	case GafSysCacheGetErr:
		return "缓存获取错误"
	case GafSysCacheSetErr:
		return "缓存设置错误"
	// 系统运行中数据库出现的错误
	case GafSysDBErr:
		return "数据库出错"
	case GafSysDBFindErr:
		return "数据库查询出错"
	case GafSysDBCreateErr:
		return "数据库添加数据出错"
	case GafSysDBUpdateErr:
		return "数据库更新出错"
	case GafSysDBDeleteErr:
		return "数据库删除出错"
	case GafSysCmdErr:
		return "执行宿主机cmd指令出错"
	case GafSysNetworkErr:
		return "网卡配置出错"
	// 系统运行中的错误
	case GafSysErr:
		return "系统错误"
	case GafSysJsonMarshalErr:
		return "转json失败"
	case GafSysJsonUnMarshalErr:
		return "json解析失败"
	case GafSysTimeParseErr:
		return "时间转换错误"
	case GafSysSaveFileErr:
		return "保存文件错误"
	}

	return "未知错误"
}
