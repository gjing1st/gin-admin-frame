package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	v1 "github.com/gjing1st/gin-admin-frame/internal/apiserver/router/v1"
)

// HttpStart
// @description: 开始http服务
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/12 18:56
// @success:
func HttpStart() {
	run()
}

// @description: 启动http服务
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/12 18:56
// @success:
func run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	{
		v1.InitApi(router) //v1版本相关接口
	}

	//启动gin路由服务
	log.Println("端口号：", config.Config.Web.Port)
	err := router.Run(fmt.Sprintf(":%s", config.Config.Web.Port))
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("http服务启动失败")
	}
}
