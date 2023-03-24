// Path: internal/apiserver/model/entity
// FileName: role.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 9:08$

package entity

// Role 角色表
type Role struct {
	BaseModel
	Name   string `json:"name" gorm:"column:name;comment:名称;type:varchar(64);NOT NULL;"`
	Status int    `json:"status" form:"status" gorm:"column:status;default:1;comment:状态;type:tinyint(3);"`
}

func (Role) TableName() string {
	return "role"
}
