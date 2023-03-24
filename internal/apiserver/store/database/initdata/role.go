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
)

type InitRole struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitRole) DataInserted(db *gorm.DB) bool {
	if errors.Is(db.Where(entity.Role{}.TableName()+".id = ?", dict.RoleIdAdmin).First(&entity.Role{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
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
func (i InitRole) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	entities := []entity.Role{
		{Name: "超管", BaseModel: entity.BaseModel{ID: dict.RoleIdSuperAdmin}},
		{Name: "管理员", BaseModel: entity.BaseModel{ID: dict.RoleIdAdmin}},
		{Name: "操作员", BaseModel: entity.BaseModel{ID: dict.RoleIdOperator}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}
