// Path: pkg/utils
// FileName: errcode.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/1$ 20:06$

package errcode

// type ErrCode int
// 错误代码100101，其中 10 代表gaf管理平台；中间的 01 代表系统运行中产生的错误；最后的 01 代表模块下的错误码序号，每个模块可以注册 100 个错误
// 0代表成功
const (
	SuccessCode = 0 //成功返回错误码
	ErrCode     = 1
	ModuleCode  = 100 * ErrCode
	ServerCode  = ModuleCode * 100
)

const (
	GafServer = 10 * ServerCode // gaf管理平台
)
