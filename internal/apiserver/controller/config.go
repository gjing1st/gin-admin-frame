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
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type ConfigController struct {
}

var configService service.ConfigService

// GetInit
// @description: 获取初始化状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 14:42
// @success:
func (cc ConfigController) GetInit(c *gin.Context) {
	v, errCode := configService.GetValueStr(dict.ConfigInitKey)
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	var res response.Init
	res.Initialized = utils.Bool(v)
	response.OkWithData(res, global.QuerySuccess, c)
}

// GetInitStep
// @description: 获取初始化步骤
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 9:52
// @success:
func (cc ConfigController) GetInitStep(c *gin.Context) {
	res, errCode := configService.GetInitStep()
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	response.OkWithData(res, global.QuerySuccess, c)
}

// LoginType
// @description: 获取登录方式
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 18:23
// @success:
func (cc ConfigController) LoginType(c *gin.Context) {
	v, errCode := configService.GetValueStr(dict.ConfigLoginType)
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	var res response.LoginTypeRes
	res.LoginType = utils.Int(v)
	response.OkWithData(res, global.QuerySuccess, c)
}

// InitNetwork
// @description: 初始化网络
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 11:18
// @success:
func (cc ConfigController) InitNetwork(c *gin.Context) {
	//判断是否已初始化完成
	res, errCode := configService.GetInitStep()
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	//已初始化完成
	if res.User == dict.InitStepValueDown && res.Network == dict.InitStepValueDown {
		response.Forbidden(errcode2.GafUserForbiddenErr, global.InitDown, c)
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
		response.FailWithLog(errCode, global.CreatedFailed, "", content, nil, c)
		return
	}
	//记录初始化操作步骤
	errCode = configService.SetInitStep(dict.InitStepNetwork)
	if errCode != 0 {
		response.FailWithLog(errCode, global.CreatedFailed, "", content, nil, c)
		return
	}
	response.OkWithLog(global.OperateSuccess, "", content, nil, c)
}

// InitReset
// @description: 初始化中重置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 18:23
// @success:
func (cc ConfigController) InitReset(c *gin.Context) {
	//判断是否已初始化完成
	res, errCode := configService.GetInitStep()
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	//已初始化完成
	if res.User == dict.InitStepValueDown && res.Network == dict.InitStepValueDown {
		response.Forbidden(errcode2.GafUserForbiddenErr, global.InitDown, c)
		return
	}
	errCode = configService.InitReset()
	content := "初始化重置"
	if errCode != 0 {
		response.FailWithLog(errCode, global.CreatedFailed, "", content, nil, c)
		return
	}
	response.OkWithLog(global.OperateSuccess, "", content, nil, c)
}

// VersionInfo
// @description: 软件版本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/1/4 9:25
// @success:
func (cc ConfigController) VersionInfo(c *gin.Context) {
	res := configService.VersionInfo()
	response.OkWithData(res, global.QuerySuccess, c)
}
