// Path: internal/apiserver/router/v1
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 10:17$

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/controller"
)

// initWithoutTokenRouter
// @description: 不需要权限的接口，开放权限
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:19
// @success:
func initWithoutTokenRouter(apiV1 *gin.RouterGroup) {
	initWithoutConfigRouter(apiV1)
	initWithoutAdminRouter(apiV1)
}

// 初始化
func initWithoutConfigRouter(apiV1 *gin.RouterGroup) {
	configController := controller.ConfigController{}
	apiV1.GET("init", configController.GetInit)                 //暂无使用
	apiV1.GET("init/step", configController.GetInitStep)        //初始化状态步骤
	apiV1.POST("init/network", configController.InitNetwork)    //初始化网络
	apiV1.DELETE("init/reset", configController.InitReset)      //初始化时重置
	apiV1.GET("login-type", configController.LoginType)         //登录方式
	apiV1.GET("sys/version/info", configController.VersionInfo) //软件信息

}
