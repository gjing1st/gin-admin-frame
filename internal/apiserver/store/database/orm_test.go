// Path: internal/apiserver/store/mysql
// FileName: m_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/22$ 11:23$

package database

import (
	"fmt"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"testing"
	"time"
)

var dbi DBI

func Init() {
	config.InitConfig()
	//加载数据库驱动并初始化数据
	//dbi = GetOrm("gorm")
	dbi = GetOrm("zorm")

}
func TestQuery(t *testing.T) {
	Init()
	var cc, c1 entity.Config
	var list []entity.Config

	go func() {
		dbi.GetDB().Model(entity.Config{}).Where("status = ?", 1).Limit(3).Offset(3).Find(&list)
		fmt.Println("==========")
		fmt.Println("=", list)
		fmt.Println("==========")
	}()
	time.Sleep(time.Millisecond * 5)
	//store.DBI.GetDB().SetTableName("config").Where("status =?", 1).Where("value = ?", "false").First(&cc)
	go func() {
		dbi.GetDB().Where("id =?", 20).First(&c1)
		fmt.Println(c1)
	}()
	time.Sleep(time.Millisecond * 5)

	dbi.GetDB().Where("status =?", 1).First(&cc)
	fmt.Println(utils.String(cc.Value))
	time.Sleep(time.Millisecond * 100)

}

func TestUpdate(t *testing.T) {
	Init()

	dbi.GetDB().Model(&entity.Config{}).Where("id = ?", 8).Update("status", 2)
}

func TestDelete(t *testing.T) {
	Init()
	dbi.GetDB().Debug().Where("id = ?", 7).Where("status=?", 2).Delete(&entity.Config{})
}

func TestSelect(t *testing.T) {
	Init()
	dbi.GetDB().Select("id,name,age", "sex")
}
func TestJoin(t *testing.T) {
	Init()
	var list []entity.Config
	dbi.GetDB().Model(&entity.Config{}).Select("config.id,config.status", "user.name").
		Joins("left join user on user.id = config.id").Find(&list)
	for _, v := range list {
		fmt.Println(v.Name)

	}
}
