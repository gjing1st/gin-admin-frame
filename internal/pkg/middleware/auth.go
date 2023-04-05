// Path: internal/pkg/middleware
// FileName: auth.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 10:00$

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// LoginRequired
// @description: token认证中间件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 10:37
// @success:
func LoginRequired(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		//请求头没有传Authorization参数，直接返回401未授权
		response.Unauthorized(errcode.GafUserWithoutToken, ctx)
		ctx.Abort()
		return
	}
	var bearer = "Bearer "
	i := strings.Index(token, bearer)
	if i == -1 {
		functions.AddWarnLog(log.Fields{"msg": "token非法", "token": token})
		response.Unauthorized(errcode.GafUserWithoutToken, ctx)
		ctx.Abort()
		return
	}
	token = strings.Replace(token, bearer, "", 1)
	var tokenService service.TokenService
	userInfo, err := tokenService.GetInfo(token)
	if err != 0 || userInfo.Id == 0 {
		//token错误或token过期，返回401
		functions.AddWarnLog(log.Fields{"err": err, "msg": "用户token解析错误", "token": token})
		response.Unauthorized(errcode.GafUserTokenExpired, ctx)
		ctx.Abort()
		return
	}

	ctx.Set("userId", userInfo.Id)
	ctx.Set("username", userInfo.Name)
	ctx.Set("roleId", userInfo.RoleId)
	ctx.Next()
}

// AdminRequired
// @description: 管理员权限认证
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:06
// @success:
func AdminRequired(c *gin.Context) {
	role, b := c.Get("roleId")
	if !b {
		response.Forbidden(errcode.GafUserRoleErr, c)
		c.Abort()
		return
	}
	roleId := utils.Int(role)
	if roleId != dict.RoleIdAdmin && roleId != dict.RoleIdSuperAdmin {
		//不是管理员或者超管角色，返回403状态码
		response.Forbidden(errcode.GafUserForbiddenErr, c)
		c.Abort()
		return
	}
	c.Next()
}

// OperatorRequired
// @description: 需要操作员权限
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:09
// @success:
func OperatorRequired(c *gin.Context) {
	role, b := c.Get("roleId")
	if !b {
		c.JSON(http.StatusForbidden, global.Unauthorized)
		c.Abort()
		return
	}
	roleId := utils.Int(role)
	if roleId != dict.RoleIdOperator && roleId != dict.RoleIdSuperAdmin {
		c.JSON(http.StatusForbidden, global.AuthForbidden)
		c.Abort()
		return

	}
	c.Next()
}
