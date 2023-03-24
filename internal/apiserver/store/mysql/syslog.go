// Path: internal/apiserver/store/mysql
// FileName: adminlog.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/28$ 14:56$

package mysql

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type SysLogMysql struct {
}

// Create
// @description: 创建管理员日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 14:59
// @success:
func (sm SysLogMysql) Create(tx *gorm.DB, data *entity.SysLog) (errCode int) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Create(&data).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql创建管理员失败"})
		return errcode.GafSysDBCreateErr
	}
	return

}

// List
// @description: 日志列表查询
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 11:30
// @success:
func (sm SysLogMysql) List(req *request.SysLogList) (users []entity.SysLog, total int64, errCode int) {
	db := store.DB.Model(&entity.SysLog{}).Where("category = ?", req.Category)
	if req.Keyword != "" {
		db.Where("username like ? or content like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if !req.StartDate.IsZero() {
		db.Where("created_at > ?", req.StartDate)
	}

	if !req.EndDate.IsZero() {
		//db.Where("DATE_FORMAT(`created_at`,'%Y-%m-%d') <= ?", req.StartDate)
		db.Where("created_at  < ?", req.EndDate)
	}
	//总数使用缓存，避免并发CPU占用
	cacheCount, err := store.GC.Get("sysLogCount")
	if err != nil {
		total = utils.Int64(cacheCount)
	} else {
		_ = db.Count(&total).Error
		store.GC.SetWithExpire("sysLogCount", total, time.Second*2)
	}
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&users).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
		errCode = errcode.GafUserNotFoundErr
	}
	return
}

// TruncateTable
// @description: 清空日志表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/24 10:51
// @success:
func (sm SysLogMysql) TruncateTable(tx *gorm.DB) (errCode int) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Exec("TRUNCATE TABLE " + entity.SysLog{}.TableName()).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置日志失败"})
		return errcode.GafSysDBDeleteErr
	}
	return
}
