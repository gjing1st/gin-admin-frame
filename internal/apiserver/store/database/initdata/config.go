// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"strconv"
	"time"
)

type InitConfig struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitConfig) DataInserted(db *gorm.DB) bool {
	if errors.Is(db.Where(entity.Config{}.TableName()+".name = ?", dict.ConfigSysFirstStartDate).First(&entity.Config{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
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
func (i InitConfig) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	//初始化步骤
	var step1 dict.InitStepValue
	initStep, _ := json.Marshal(step1)
	//向导步骤
	//var step2 dict.GuideStepValue
	//guideStep, _ := json.Marshal(step2)
	entities := []entity.Config{
		{Name: dict.ConfigInitKey, Value: "false"},
		{Name: dict.ConfigSysFirstStartDate, Value: time.Now().Format(global.TimeFormat)},
		{Name: dict.ConfigSysBreakDate, Value: time.Now().Format(global.TimeFormat)},
		{Name: dict.ConfigLoginType, Value: strconv.Itoa(dict.LoginTypeBackendUKey)},
		{Name: dict.ConfigInitStep, Value: string(initStep)},
		//{Name: dict.ConfigGuideStep, Value: string(guideStep)},
		{Name: dict.ConfigVersion, Value: global.Version},
		{Name: dict.ConfigLatestVersion},
		{Name: dict.ConfigBackupTime},
		{Name: dict.ConfigRestoreTime},
		{Name: dict.ConfigAutoUpdate},
		{Name: dict.ConfigUpdateRange},
		{Name: dict.ConfigUpdateTime},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}

// Update
// @description: 更新数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/10 16:50
// @success:
func (i InitConfig) Update(db *gorm.DB) error {
	return db.Model(&entity.Config{}).Where("name = ?", dict.ConfigSysBreakDate).Update("value", time.Now().Format(global.TimeFormat)).Error
}
