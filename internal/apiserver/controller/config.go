// Path: internal/apiserver/controller
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 10:55$

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	errcode "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type ConfigController struct {
}

var configService service.ConfigService

// GetInit godoc
// @Summary 获取初始化状态
// @Description 获取初始化状态
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.Init "操作成功"
// @Failure 500 {object} string
// @Router /init [get]
func (cc ConfigController) GetInit(c *gin.Context) {
	v, errCode := configService.GetValueStr(dict.ConfigInitKey)
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	var res response.Init
	res.Initialized = utils.Bool(v)
	response.OkWithData(res, c)
}

// GetInitStep godoc
// @Summary 获取初始化步骤
// @Description 获取初始化步骤
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.InitStepValue "操作成功"
// @Failure 500 {object} string
// @Router /init/step [get]
func (cc ConfigController) GetInitStep(c *gin.Context) {
	res, errCode := configService.GetInitStep()
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	response.OkWithData(res, c)
}

// LoginType godoc
// @Summary 获取登录方式
// @Description 获取登录方式
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.LoginTypeRes "操作成功"
// @Failure 500 {object} string
// @Router /login-type [get]
func (cc ConfigController) LoginType(c *gin.Context) {
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	var res response.LoginTypeRes
	res.LoginType = utils.Int(v)
	response.OkWithData(res, c)
}

// InitNetwork godoc
// @Summary 网络配置
// @Description 网络配置
// @Param data body request.SetNetwork true "网络配置"
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /init/network [post]
func (cc ConfigController) InitNetwork(c *gin.Context) {
	//判断是否已初始化完成
	res, errCode := configService.GetInitStep()
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	//已初始化完成
	if res.User == dict.InitStepValueDown && res.Network == dict.InitStepValueDown {
		response.Forbidden(errcode.GafUserForbiddenErr, c)
		return
	}
	var req request.SetNetwork
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	errCode = configService.SetNetwork(&req)
	content := "初始化网络"
	if errCode != 0 {
		response.FailWithLog(errCode, "", content, nil, c)
		return
	}
	//记录初始化操作步骤
	errCode = configService.SetInitStep(dict.InitStepNetwork)
	if errCode != 0 {
		response.FailWithLog(errCode, "", content, nil, c)
		return
	}
	response.OkWithLog("", content, nil, c)
}

// InitReset godoc
// @Summary 初始化中重置
// @Description 初始化中重置
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} string "操作成功"
// @Failure 500 {object} string
// @Router /init/network [delete]
func (cc ConfigController) InitReset(c *gin.Context) {
	//判断是否已初始化完成
	res, errCode := configService.GetInitStep()
	if errCode != errcode.SuccessCode {
		response.Failed(errCode, c)
		return
	}
	//已初始化完成
	if res.User == dict.InitStepValueDown && res.Network == dict.InitStepValueDown {
		response.Forbidden(errcode.GafUserForbiddenErr, c)
		return
	}
	errCode = configService.InitReset()
	content := "初始化重置"
	if errCode != 0 {
		response.FailWithLog(errCode, "", content, nil, c)
		return
	}
	response.OkWithLog("", content, nil, c)
}

// VersionInfo godoc
// @Summary 软件版本信息
// @Description 软件版本信息
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @Accept application/json
// @Success 200 {object} response.VersionInfo "操作成功"
// @Failure 500 {object} string
// @Router /sys/version/info [get]
func (cc ConfigController) VersionInfo(c *gin.Context) {
	res := configService.VersionInfo()
	response.OkWithData(res, c)
}
