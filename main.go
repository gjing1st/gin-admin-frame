package main

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/database"
)

// @title GAFAdmin
// @version 1.0
// @description 基于Gin的后台管理框架.
// @contact.name 这里写联系人信息
// @contact.url http://tna.cn
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9681
// @BasePath /gaf/v1
//
//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
func main() {
	//加载配置文件
	config.InitConfig()
	//加载数据库驱动并初始化数据
	store.DB = database.GetDB()
	if store.DB != nil {
		db, _ := store.DB.DB()
		// 程序结束前关闭数据库链接
		defer db.Close()
	}
	//开启定时任务
	service.AddCron()
	//go func() {
	//	//开启升级检查
	//	var ss service.SysService
	//	ss.UpdateStateVersion()
	//}()
	//开启http服务
	apiserver.HttpStart()

}
