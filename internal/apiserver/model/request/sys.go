// Path: internal/apiserver/model/request
// FileName: sys.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/13$ 19:08$

package request

type Network struct {
	Addr    string `json:"addr"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

// SetNetwork 网络配置
type SetNetwork struct {
	Admin Network `json:"admin" binding:"required"`
	SDF   Network `json:"sdf"`
}

// AutoUpdateConfig 自动升级配置
type AutoUpdateConfig struct {
	AutoUpdate  bool   `json:"auto_update"`
	UpdateRange string `json:"update_range"`
	Time        string `json:"time"`
}
