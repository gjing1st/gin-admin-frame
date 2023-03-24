// Path: internal/apiserver/controller
// FileName: sys.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/9$ 18:43$

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type SysController struct {
}

var sysService service.SysService

// ServerStatus
// @description: 设备运行状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 21:30
// @success:
func (slc *SysController) ServerStatus(c *gin.Context) {
	res, _ := sysService.ServerStatus()
	response.OkWithData(res, global.OperateSuccess, c)

}

// SysReset
// @description: 恢复出厂设置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 13:48
// @success:
func (slc *SysController) SysReset(c *gin.Context) {
	var req request.UserLogin
	//参数接收
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	errCode := configService.SysReset(&req)
	content := "恢复出厂设置"
	if errCode != 0 {
		response.FailWithSysLog(errCode, global.CreatedFailed, "", content, c)
		return
	}
	response.OkWithSysLog(global.OperateSuccess, "", content, c)
}

// Reboot
// @description: 重启服务器
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 13:53
// @success:
func (slc *SysController) Reboot(c *gin.Context) {
	content := "重启服务器"
	go func() {
		err := sysService.Reboot()
		if err != 0 {
			response.FailWithSysLog(errcode2.GafSysCmdErr, global.OperateFailed, "", content, c)
			return
		}
	}()
	//先返回请求成功，记录日志
	response.OkWithSysLog(global.OperateSuccess, "", content, c)

}

// SysRunDate
// @description: 设备运行时间
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 11:19
// @success:
func (slc *SysController) SysRunDate(c *gin.Context) {
	res, errCode := configService.GetRunDate()
	if errCode != errcode2.SuccessCode {
		response.Failed(errCode, global.QueryFailed, c)
		return
	}
	response.OkWithData(res, global.QuerySuccess, c)

}

// GetNetwork
// @description: 获取当前网络配置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 17:20
// @success:
func (slc *SysController) GetNetwork(c *gin.Context) {
	res, errCode := sysService.GetNetwork()
	if errCode != 0 {
		response.Failed(errCode, global.OperateFailed, c)
		return
	}
	response.OkWithData(res, global.OperateSuccess, c)
}

// SetNetwork
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 18:38
// @success:
func (slc *SysController) SetNetwork(c *gin.Context) {
	var req request.SetNetwork
	//参数接收
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err})
		return
	}
	errCode := sysService.SetNetwork(&req)
	content := "网络配置"
	if errCode != 0 {
		response.FailWithSysLog(errCode, global.CreatedFailed, "", content, c)
		return
	}
	response.OkWithSysLog(global.OperateSuccess, "", content, c)
	//sysService.RestartNetwork()
}

// VersionInfo
// @description: 系统配置-关于
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/14 15:25
// @success:
func (slc *SysController) VersionInfo(c *gin.Context) {
	res := sysService.VersionInfo()
	response.OkWithData(res, global.OperateSuccess, c)
}

// Update
// @description: 升级
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/15 18:27
// @success:
func (slc *SysController) Update(c *gin.Context) {
	errCode, version := sysService.Update()
	content := "升级至" + version + "版本"
	if errCode != 0 {
		response.FailWithSysLog(errCode, global.OperateFailed, "", content, c)
		return
	}
	response.OkWithSysLog(global.OperateSuccess, "", content, c)
}

// UpdateVersionInfo
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 9:40
// @success:
func (slc *SysController) UpdateVersionInfo(c *gin.Context) {
	res, errCode := sysService.UpdateVersionInfo()
	if errCode != 0 {
		response.Failed(errCode, global.OperateFailed, c)
	} else {
		response.OkWithData(res, global.OperateSuccess, c)
	}
}

// UploadFile
// @description: 升级升级包
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 10:23
// @success:
func (slc *SysController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err, "msg": "上传升级包失败"})
		return
	}
	//文件名为:项目名_版本号格式 ex:hss_V2.7.0.zip
	fileExt := path.Ext(file.Filename)
	if fileExt != ".zip" {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"fileExt": fileExt, "msg": "上传升级包失败"})
		return
	}
	//日期存放路径
	dirName := "/" + time.Now().Format("2006-01-02_15_04") + "/"
	content := "上传更新包-" + file.Filename
	//创建存放目录
	err = os.MkdirAll(config.Config.UploadPath+dirName, 0777)
	if err != nil {
		response.FailWithLog(errcode2.GafSysSaveFileErr, global.UploadFailed, "", content, nil, c)
		return
	}
	//完整路径文件名
	fullName := config.Config.UploadPath + dirName + file.Filename
	//存放文件
	if err = c.SaveUploadedFile(file, fullName); err != nil {
		response.FailWithLog(errcode2.GafSysSaveFileErr, global.UploadFailed, "", content, nil, c)
		return
	}
	errCode := sysService.DealFile(file.Filename, config.Config.UploadPath+dirName)
	if errCode != 0 {
		response.FailWithLog(errCode, global.UploadFailed, "", content, nil, c)
	} else {
		response.OkWithLog(global.OperateSuccess, "", content, nil, c)
	}

}

// UploadFileV1
// @description: 上传升级包V1版本
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/13 17:31
// @success:
func (slc *SysController) UploadFileV1(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err, "msg": "上传升级包失败"})
		return
	}
	//文件名为:项目名_版本号格式 ex:hss_2.7.0.zip
	fileExt := path.Ext(file.Filename)
	if fileExt != ".zip" {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"fileExt": fileExt, "msg": "上传升级包失败"})
		return
	}
	////日期存放路径
	//dirName := "/" + time.Now().Format("2006-01-02_15_04") + "/"
	content := "上传更新包-" + file.Filename
	//创建存放目录
	err = os.MkdirAll(config.Config.UploadPath, 0777)
	if err != nil {
		response.FailWithLog(errcode2.GafSysSaveFileErr, global.UploadFailed, "", content, nil, c)
		return
	}
	//完整路径文件名
	fullName := config.Config.UploadPath + file.Filename
	//存放文件
	if err = c.SaveUploadedFile(file, fullName); err != nil {
		response.FailWithLog(errcode2.GafSysSaveFileErr, global.UploadFailed, "", content, nil, c)
		return
	}
	errCode := sysService.DealFileV1(file.Filename, config.Config.UploadPath)
	if errCode != 0 {
		response.FailWithLog(errCode, global.UploadFailed, "", content, nil, c)
	} else {
		response.OkWithLog(global.OperateSuccess, "", content, nil, c)
	}

}

// GetAuto
// @description: 获取自动更新策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 15:54
// @success:
func (slc *SysController) GetAuto(c *gin.Context) {
	res, errCode := sysService.GetAuto()
	if errCode != 0 {
		response.Failed(errCode, global.OperateFailed, c)
		return
	}
	response.OkWithData(res, global.OperateSuccess, c)
}

// SetAuto
// @description: 设置自动更新
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 16:25
// @success:
func (slc *SysController) SetAuto(c *gin.Context) {
	var req request.AutoUpdateConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddErrLog(log.Fields{"err": err, "msg": "设置自动更新参数错误"})
		return
	}
	content := "配置自动更新"
	errCode := sysService.SetAuto(&req)
	if errCode != 0 {
		response.FailWithLog(errCode, global.OperateFailed, "", content, req, c)
		return
	}
	response.OkWithLog(global.OperateSuccess, "", content, req, c)

}
