package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/middleware"
	"net/http"
)

// InitApi
// @description: 初始化路由
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/28 15:27
// @success:
func InitApi(router *gin.Engine) {
	//是否跨域
	if config.Config.Web.Cors {
		router.Use(middleware.CORS)
	}
	apiV1 := router.Group("gaf/v1")
	//ping服务检测接口
	apiV1.GET("ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	{
		//无需登录即可访问的相关接口
		initWithoutTokenRouter(apiV1)
	}

	{
		//需要登录才可访问的接口
		initLoginRouter(apiV1)
	}

}
