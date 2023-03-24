// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"errors"
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils/gm"
)

type InitUser struct {
}

var (
	adminName     = "admin"
	adminPassword = "12345678"
)

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitUser) DataInserted(db *gorm.DB) bool {
	if errors.Is(db.Where(entity.User{}.TableName()+".name = ?", adminName).First(&entity.User{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// InitializeData
// @description: 初始化数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:26
// @success:
func (i InitUser) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	adminPasswd := gm.EncryptPasswd(adminName, adminPassword)
	entities := []entity.User{
		{Name: adminName, RoleId: dict.RoleIdAdmin, Password: adminPasswd},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}
