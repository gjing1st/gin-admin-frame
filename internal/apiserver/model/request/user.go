// Path: internal/apiserver/model/request
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 15:49$

package request

// UserCreate 创建用户
type UserCreate struct {
	Name      string `json:"name" binding:"required"`
	Serial    string `json:"serial"`
	SignData  string `json:"sign"`
	TimeStamp string `json:"timestamp"`
	RoleId    int    `json:"role_id" form:"role_id"`
	Cert      string `json:"cert"`
	Pin       string `json:"pin"`
}

// UserLogin 用户登录请求
type UserLogin struct {
	Name      string `json:"name" binding:"required"`
	LoginType string `json:"login_type"`
	Password  string `json:"pin"  binding:"required"`
}

// UserList 用户列表
type UserList struct {
	PageInfo
	RoleId []int `json:"role_id" form:"role_id"`
}

type UKeyLogin struct {
	Name      string `json:"name" binding:"required"`
	Serial    string `json:"serial" binding:"required"`
	SignData  string `json:"sign" binding:"required"`
	TimeStamp string `json:"timestamp" binding:"required"`
}

// CreateUKeyUser 初始化时的添加后端ukey管理员
type CreateUKeyUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pin"  binding:"required"`
}

// ChangePasswd 修改密码
type ChangePasswd struct {
	OldPassword string `json:"old_pin"  binding:"required"`
	NewPassword string `json:"pin"  binding:"required"`
}

// KeyBackup 密钥备份与恢复
type KeyBackup struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pin"  binding:"required"`
}

// UserDelete 删除管理员
type UserDelete struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	KeySn    string `json:"keysn"`
}
