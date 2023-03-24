// Path: internal/apiserver/store/mysql
// FileName: crontab.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 17:57$

package mysql

import (
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"time"
)

type CrontabStore struct {
}

// Create
// @description: 创建定时任务数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:58
// @success:
func (cs CrontabStore) Create(tx *gorm.DB, cron entity.Crontab) (id uint, err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Create(&cron).Error

	return cron.ID, err
}

// Update
// @description: 修改定时时间
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:52
// @success:
func (cs CrontabStore) Update(tx *gorm.DB, articleId uint, cronTime time.Time) (err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Model(&entity.Crontab{}).Where("article_id = ?", articleId).Update("CrontabTime", cronTime).Error
	return
}

// SearchWaitingStatus
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 10:54
// @success:
func (cs CrontabStore) SearchWaitingStatus(cronTime time.Time) (cronList []entity.Crontab, err error) {
	err = store.DB.Where("crontab_time <= ?", cronTime).Find(&cronList).Error
	return
}

// UpdateStatus
// @description: 修改状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 11:05
// @success:
func (cs CrontabStore) UpdateStatus(tx *gorm.DB, status int, id uint) (err error) {
	//if tx != nil {
	//	err = tx.Model(&model.Crontab{}).Where("id = ?", id).Update("status", status).Error
	//} else {
	//	err = store.DB.Model(&model.Crontab{}).Where("id = ?", id).Update("status", status).Error
	//}
	return
}

// Delete
// @description: 删除任务
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 11:11
// @success:
func (cs CrontabStore) Delete(tx *gorm.DB, id uint) (err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Delete(&entity.Crontab{}, id).Error
	return
}
