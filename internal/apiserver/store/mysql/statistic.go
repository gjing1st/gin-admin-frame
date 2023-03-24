// Path: internal/apiserver/store/mysql
// FileName: statistic.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/13$ 22:42$

package mysql

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
)

type StatisticStore struct {
}

// Save
// @description: 保存统计数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/13 22:44
// @success:
func (ss StatisticStore) Save(tx *gorm.DB, s entity.Statistic) (err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Model(&entity.Statistic{}).Where(entity.Statistic{Key: s.Key}).Update("value", s.Value).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql保存统计数据失败", "Statistic": s})
	}
	return
}

// GetByKey
// @description: 通过Key查询统计数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/13 22:58
// @success:
func (ss StatisticStore) GetByKey(key string) (value string, err error) {
	var s entity.Statistic
	err = store.DB.Where(entity.Statistic{}.TableName()+".key = ?", key).First(&s).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询统计数据失败", "key": key})
		return
	}
	value = s.Value
	return
}

// GetAll
// @description: 获取所有统计数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/13 23:07
// @success:
func (ss StatisticStore) GetAll() (arr []entity.Statistic, err error) {
	err = store.DB.Find(&arr).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询统计数据失败"})
		return
	}
	return
}
