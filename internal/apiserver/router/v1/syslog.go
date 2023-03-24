// Path: internal/apiserver/router/v1
// FileName: log.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/9$ 10:36$

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/controller"
)

func initLogRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("logs")
	var logController controller.SysLogController
	api.GET("", logController.List) //查看设备密钥

}
