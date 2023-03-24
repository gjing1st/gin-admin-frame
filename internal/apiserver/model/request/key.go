// Path: internal/apiserver/model/request
// FileName: key.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/2$ 11:28$

package request

// CreateSm4 生成对称密钥
type CreateSm4 struct {
	Index *uint `json:"index" binding:"required"`
}

// DeleteSm4 删除对称密钥
type DeleteSm4 struct {
	Index *uint `json:"index" binding:"required"`
}

// CreateSm2 生成非对称密钥
type CreateSm2 struct {
	Index *uint  `json:"index" binding:"required"`
	Type  []int  `json:"type"`
	Pin   string `json:"pin" binding:"required"`
}

// CreateDevice 生成设备密钥
type CreateDevice struct {
	Type []int  `json:"type"`
	Pin  string `json:"pin" binding:"required"`
}

// SetSm2Pin 设置访问码
type SetSm2Pin struct {
	Index *int   `json:"index" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
}

// ExportPub 导出公钥
type ExportPub struct {
	Index *int `json:"index" form:"index" binding:"required"`
}

// DeleteSm2 删除非对称密钥
type DeleteSm2 struct {
	Index *uint `json:"index" binding:"required"`
	Type  []int `json:"type"`
}
