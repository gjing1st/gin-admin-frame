// Path: internal/apiserver/store/mysql
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 16:00$

package mysql

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserMysql struct {
}

// Create
// @description: mysql存储
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:00
// @success:
func (um *UserMysql) Create(tx *gorm.DB, user *entity.User) (id uint, errCode int) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Create(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql创建管理员失败"})
		return id, errcode2.GafSysDBCreateErr
	}
	return user.ID, 0
}

// GetByName
// @description: 通过用户名查询用户信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:10
// @success:
func (um UserMysql) GetByName(name string) (user *entity.User, errCode int) {
	err := store.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql创建管理员失败"})
		if err == errcode2.ErrRecordNotFound {
			return user, errcode2.GafUserNotFoundErr
		}
		return user, errcode2.GafSysDBFindErr
	}
	return
}

// UpdateToken
// @description: 修改数据库用户token
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:14
// @success:
func (um UserMysql) UpdateToken(id uint, token string) (errCode int) {
	err := store.DB.Model(&entity.User{}).Where("id = ?", id).Update("token", token).Error
	if err != nil {
		errCode = errcode2.GafSysDBUpdateErr
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql修改token失败", "id": id, "token": token})
	}
	return
}

// List
// @description: 获取用户列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:31
// @success:
func (um UserMysql) List(req *request.UserList) (users []entity.User, total int64, errCode int) {
	db := store.DB.Model(&entity.User{})
	if req.Keyword != "" {
		db.Where("name like ?", "%"+req.Keyword+"%")
	}
	if len(req.RoleId) > 0 {
		db.Where("role_id in ?", req.RoleId)
	}
	err := db.Count(&total).Error
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&users).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
		errCode = errcode2.GafUserNotFoundErr
	}
	return
}

// GetByNameAndSerialNum
// @description: 查询ukey管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 15:36
// @success:
func (um UserMysql) GetByNameAndSerialNum(name, serialNum string) (user *entity.User, errCode int) {
	err := store.DB.Where("name = ? and user_serial_num = ?", name, serialNum).First(&user).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询ukey管理员失败"})
		if err == errcode2.ErrRecordNotFound {
			return user, errcode2.GafUserNotFoundErr
		}
		return user, errcode2.GafSysDBFindErr
	}
	return
}

// ResetUser
// @description: 删除管理员意外的其他用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 18:38
// @success:
func (um UserMysql) ResetUser(tx *gorm.DB) (errCode int) {
	if tx == nil {
		tx = store.DB
	}
	//永久删除
	err := tx.Unscoped().Where("name != ?", "superadmin").Delete(&entity.User{}).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置管理员失败"})
		return errcode2.GafSysDBDeleteErr
	}
	return
}

// DeleteById
// @description: 通过id删除用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 11:28
// @success:
func (um UserMysql) DeleteById(tx *gorm.DB, userid int) (errCode int) {
	if tx == nil {
		tx = store.DB
	}
	err := tx.Where("id = ?", userid).Delete(&entity.User{}).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql重置管理员失败"})
		return errcode2.GafSysDBDeleteErr
	}
	return
}
