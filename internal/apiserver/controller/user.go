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
	errcode "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
}

var userService = service.UserService{}

// InitAdmin godoc
// @Summary 添加管理员
// @Description 添加管理员
// @Param data body request.UserCreate true "管理员信息"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /user/init [post]
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
		response.FailWithLog(errCode, "", content, &req, c)
		return
	}
	//记录初始化操作步骤
	errCode = configService.SetInitStep(dict.InitStepUser)
	if errCode != 0 {
		response.FailWithLog(errCode, "", content, &req, c)
		return
	}
	response.OkWithLog("", content, &req, c)

}

// Login godoc
// @Summary 用户登录
// @Description 用户登录
// @Param data body request.UserLogin true "用户名密码"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} entity.User "操作成功"
// @Failure 500 {object} string
// @Router /user/login [post]
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
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	loginType := utils.Int(v)

	var res interface{}
	//该接口只允许用户名密码登录或者后端key登录
	if loginType == dict.LoginTypePasswd {
		res, errCode = userService.Login(&req)
	} else if loginType == dict.LoginTypeBackendUKey {

	} else {
		response.Failed(errCode, c)
		return
	}

	content := "登录"
	if errCode != 0 {
		if errCode == errcode.GafUserNotFoundErr {
			response.FailWithLog(errCode, req.Name, content, nil, c)
			return
		}
		//response.FailWithLog(errCode, global.LoginFail, req.Name, content, nil, c)
		response.FailWithDataLog(res, errCode, req.Name, content, nil, c)
	} else {
		response.OkWithDataLog(res, req.Name, content, &req, c)
	}
}

// List godoc
// @Summary 用户列表
// @Description 用户列表
// @Param data query request.UserList false "分页搜索"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /users [get]
func (uc *UserController) List(c *gin.Context) {
	var req request.UserList
	_ = c.ShouldBindQuery(&req)
	if req.PageSize == 0 {
		req.PageSize = global.PageSizeDefault
	}
	var (
		list    interface{}
		total   int64
		errCode errcode.Err
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
		response.Failed(errCode, c)

	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, c)
	}
}

// Create godoc
// @Summary 添加管理员
// @Description 添加管理员
// @Param data body request.UserCreate true "管理员信息"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /users [post]
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
		if errCode == errcode.GafUserHasExist {
			response.FailWithLog(errCode, "", content, &req, c)
			return
		}
		response.FailWithLog(errCode, "", content, &req, c)
		return
	}
	response.OkWithLog("", content, &req, c)
}

// UKeyLogin godoc
// @Summary 远端(前端)UKey登录
// @Description 远端(前端)UKey登录
// @Param data body request.UKeyLogin true "管理员信息"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} entity.User "操作成功"
// @Failure 500 {object} string
// @Router /ukey/login [post]
func (uc *UserController) UKeyLogin(c *gin.Context) {
	//判断当前登录方式
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	loginType := utils.Int(v)
	if loginType != dict.LoginTypeFrontUKey {
		response.Failed(errCode, c)
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
		if errCode == errcode.GafUserNotFoundErr {
			response.FailWithLog(errCode, req.Name, content, nil, c)
			return
		}
		response.FailWithLog(errCode, req.Name, content, nil, c)
	} else {
		response.OkWithDataLog(user, req.Name, content, &req, c)
	}
}

// Delete godoc
// @Summary 删除管理员
// @Description 删除管理员
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /users/:userid [delete]
func (uc *UserController) Delete(c *gin.Context) {
	userid := c.Param("userid")
	userId := utils.Int(userid)
	errCode := userService.DeleteById(userId)
	content := "删除管理员"
	if errCode != 0 {
		response.FailWithLog(errCode, "", content, nil, c)
	} else {
		response.OkWithLog("", content, userId, c)
	}
}

// UserDelete godoc
// @Summary 删除管理员
// @Description 删除管理员
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /users/:userid [delete]
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
		response.FailWithLog(errCode, "", content, nil, c)
	} else {
		response.OkWithLog("", content, req, c)
	}
}

// ChangePasswd godoc
// @Summary 修改UKey密码
// @Description 修改UKey密码
// @Param data body request.ChangePasswd true "用户信息"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /users/passwd [put]
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
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	loginType := utils.Int(v)
	var res interface{}
	if loginType == dict.LoginTypePasswd {
	} else if loginType == dict.LoginTypeBackendUKey {
		//后端key
	} else {
		response.Failed(errCode, c)
		return
	}
	content := "修改密码"
	if errCode != 0 {
		response.FailWithDataLog(res, errCode, "", content, nil, c)
		return
	}
	response.OkWithDataLog(res, "", content, nil, c)
}
