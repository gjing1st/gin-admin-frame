// Path: internal/apiserver/router/v1
// FileName: sys.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/30$ 13:07$

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/controller"
)

func initSysRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("sys")
	sysController := controller.SysController{}
	api.DELETE("reset", sysController.SysReset)   //恢复出厂设置
	api.POST("reboot", sysController.Reboot)      //重启服务器
	api.GET("run", sysController.SysRunDate)      //系统运行时长
	api.GET("status", sysController.ServerStatus) //系统运行状态
	api.GET("network", sysController.GetNetwork)  //当前网络配置
	api.PUT("network", sysController.SetNetwork)  //网络配置
	api.GET("version", sysController.VersionInfo) //关于
	//升级
	api = api.Group("update")
	api.GET("version/info", sysController.UpdateVersionInfo) //版本信息
	api.POST("", sysController.Update)                       //升级
	api.POST("upload", sysController.UploadFileV1)           //上传升级包
	api.GET("auto", sysController.GetAuto)                   //获取自动更新配置
	api.POST("auto", sysController.SetAuto)                  //设置自动更新配置
}
