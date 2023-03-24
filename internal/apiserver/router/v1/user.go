// Path: internal/apiserver/router/v1
// FileName: admin.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/28$ 17:15$

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/controller"
)

// initWithoutAdminRouter
// @description: 无需权限，用户相关接口
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:19
// @success:
func initWithoutAdminRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("user")
	userController := controller.UserController{}
	api.POST("init", userController.InitAdmin)
	api.POST("login", userController.Login)
	api.POST("ukey/login", userController.UKeyLogin)
}

// initUserRouter
// @description: 需要权限，用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:20
// @success:
func initUserRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("users")
	userController := controller.UserController{}
	api.GET("", userController.List)               //登录即可访问
	api.POST("", userController.Create)            //添加管理员
	api.DELETE("/:userid", userController.Delete)  //删除管理员
	api.PUT("passwd", userController.ChangePasswd) //修改密码
}
