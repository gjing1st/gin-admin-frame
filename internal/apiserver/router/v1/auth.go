// Path: internal/apiserver/router/v1
// FileName: auth.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/28$ 17:10$

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/middleware"
)

// initLoginRouter
// @description: 需要登录权限的路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:14
// @success:
func initLoginRouter(apiV1 *gin.RouterGroup) {
	//需要登录权限路由组
	apiV1.Use(middleware.LoginRequired)
	initLogRouter(apiV1)
	operatorApi := apiV1
	initAuthAdminRouter(apiV1)          //需要管理员权限
	initAuthOperatorRouter(operatorApi) //需要操作员权限

}

// initAuthAdminRouter
// @description: 需要管理员权限的路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:14
// @success:
func initAuthAdminRouter(apiV1 *gin.RouterGroup) {
	apiV1.Use(middleware.AdminRequired)
	initUserRouter(apiV1)
	initSysRouter(apiV1)
}

// initAuthOperatorRouter
// @description: 需要操作员权限的路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:14
// @success:
func initAuthOperatorRouter(apiV1 *gin.RouterGroup) {
	apiV1.Use(middleware.OperatorRequired)
}
