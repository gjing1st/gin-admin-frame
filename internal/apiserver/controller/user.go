// Path: internal/apiserver/controller
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 15:48$

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
}

var userService = service.UserService{}

// InitAdmin
// @description: 添加管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:07
// @success:
func (uc *UserController) InitAdmin(c *gin.Context) {
	var req request.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	req.RoleId = dict.RoleIdAdmin
	errCode := userService.Create(&req)
	//记录操作日志
	content := "初始化添加管理员"
	if errCode != 0 {
		response.FailWithLog(errCode, global.CreatedFailed, "", content, &req, c)
		return
	}
	//记录初始化操作步骤
	errCode = configService.SetInitStep(dict.InitStepUser)
	if errCode != 0 {
		response.FailWithLog(errCode, global.CreatedFailed, "", content, &req, c)
		return
	}
	response.OkWithLog(global.CreatedSuccess, "", content, &req, c)

}

// Login
// @description: 用户登录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:04
// @success:
func (uc *UserController) Login(c *gin.Context) {

	//参数接收
	var req request.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	//判断登录方式
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}
	loginType := utils.Int(v)

	var res interface{}
	//该接口只允许用户名密码登录或者后端key登录
	if loginType == dict.LoginTypePasswd {
		res, errCode = userService.Login(&req)
	} else if loginType == dict.LoginTypeBackendUKey {

	} else {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}

	content := "登录"
	if errCode != 0 {
		if errCode == errcode2.GafUserNotFoundErr {
			response.FailWithLog(errCode, global.UserNotFound, req.Name, content, nil, c)
			return
		}
		//response.FailWithLog(errCode, global.LoginFail, req.Name, content, nil, c)
		response.FailWithDataLog(res, errCode, global.LoginFail, req.Name, content, nil, c)
	} else {
		response.OkWithDataLog(res, global.LoginSuccess, req.Name, content, &req, c)
	}
}

// List
// @description: 用户列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:17
// @success:
func (uc *UserController) List(c *gin.Context) {
	var req request.UserList
	_ = c.ShouldBindQuery(&req)
	if req.PageSize == 0 {
		req.PageSize = global.PageSizeDefault
	}
	var (
		list    interface{}
		total   int64
		errCode int
	)

	//判断用户角色
	roleId, _ := c.Get("roleId")
	roleIdInt := utils.Int(roleId)
	userName, _ := c.Get("username")
	username := utils.String(userName)
	if roleIdInt == dict.RoleIdAdmin {
		list, total, errCode = userService.List(&req)
	} else if roleIdInt == dict.RoleIdOperator {
		list, total, errCode = userService.InfoByName(username)
	}

	if errCode != 0 {
		response.Failed(errCode, global.QueryFailed, c)

	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, global.QuerySuccess, c)
	}
}

// Create
// @description: 添加用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 14:27
// @success:
func (uc *UserController) Create(c *gin.Context) {
	var req request.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	errCode := userService.Create(&req)
	//记录操作日志
	content := "添加管理员"
	if errCode != 0 {
		if errCode == errcode2.GafUserHasExist {
			response.FailWithLog(errCode, global.UKeyHasAddAdmin, "", content, &req, c)
			return
		}
		response.FailWithLog(errCode, global.UserHasAddFailed, "", content, &req, c)
		return
	}
	response.OkWithLog(global.CreatedSuccess, "", content, &req, c)
}

// UKeyLogin
// @description: 远端(前端)UKey登录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 15:25
// @success:
func (uc *UserController) UKeyLogin(c *gin.Context) {
	//判断当前登录方式
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}
	loginType := utils.Int(v)
	if loginType != dict.LoginTypeFrontUKey {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}
	var req request.UKeyLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err, "msg": "ukey登录参数错误"})
		return
	}
	user, errCode := userService.UKeyLogin(&req)
	content := "UKey登录"
	if errCode != 0 {
		if errCode == errcode2.GafUserNotFoundErr {
			response.FailWithLog(errCode, global.UserNotFound, req.Name, content, nil, c)
			return
		}
		response.FailWithLog(errCode, global.LoginFail, req.Name, content, nil, c)
	} else {
		response.OkWithDataLog(user, global.LoginSuccess, req.Name, content, &req, c)
	}
}

// Delete
// @description: 删除管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 11:06
// @success:
func (uc *UserController) Delete(c *gin.Context) {
	userid := c.Param("userid")
	userId := utils.Int(userid)
	errCode := userService.DeleteById(userId)
	content := "删除管理员"
	if errCode != 0 {
		response.FailWithLog(errCode, global.DeleteUserFail, "", content, nil, c)
	} else {
		response.OkWithLog(global.DeleteUserSuccess, "", content, userId, c)
	}
}

// UserDelete
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/17 15:12
// @success:
func (uc *UserController) UserDelete(c *gin.Context) {
	var req request.UserDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	errCode := userService.DeleteUser(&req)
	content := "删除管理员"
	if errCode != 0 {
		response.FailWithLog(errCode, global.DeleteUserFail, "", content, nil, c)
	} else {
		response.OkWithLog(global.DeleteUserSuccess, "", content, req, c)
	}
}

// ChangePasswd
// @description: 修改UKey密码
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 10:59
// @success:
func (uc *UserController) ChangePasswd(c *gin.Context) {
	var req request.ChangePasswd
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err, "msg": "修改密码参数错误"})
		return
	}
	//判断登录方式
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	//该接口只允许用户名密码登录或者后端key登录
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}
	loginType := utils.Int(v)
	var res interface{}
	if loginType == dict.LoginTypePasswd {
	} else if loginType == dict.LoginTypeBackendUKey {
		//后端key
	} else {
		response.Failed(errCode, global.LoginTypeErr, c)
		return
	}
	content := "修改密码"
	if errCode != 0 {
		response.FailWithDataLog(res, errCode, global.OperateFailed, "", content, nil, c)
		return
	}
	response.OkWithDataLog(res, global.OperateSuccess, "", content, nil, c)
}
