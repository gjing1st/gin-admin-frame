// Path: internal/apiserver/model/entity
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/26$ 19:35$

package entity

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// User 用户表
type User struct {
	ID            uint                  `gorm:"primaryKey" json:"id"`
	Name          string                `json:"name" gorm:"column:name;comment:用户名;uniqueIndex:user_name;type:varchar(64);NOT NULL;"`
	NickName      string                `json:"nick_name" gorm:"column:nick_name;comment:昵称;type:varchar(64);NOT NULL;"`
	RoleId        int                   `json:"role_id" gorm:"column:role_id;comment:角色类型;type:tinyint(3);" `
	Password      string                `json:"-"  gorm:"column:password;comment:密码;type:varchar(64);"`
	Token         string                `json:"token"  gorm:"column:token;comment:令牌;type:varchar(255);"`
	UserSerialNum string                `json:"user_serial_num"  gorm:"column:user_serial_num;index:user_serial_num;comment:ukey序列号;type:varchar(255);"`
	UserPub       string                `json:"-"  gorm:"column:user_pub;comment:ukey签名公钥,与签名证书里的公钥一致;type:varchar(255);"`
	CertSerialNum string                `json:"-"  gorm:"column:cert_serial_num;comment:签名证书序列号;type:varchar(255);"`
	Status        int                   `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;type:tinyint(3);"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `gorm:"uniqueIndex:user_name" json:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserTokenInfo struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	RoleId int    `json:"role_id"`
}
