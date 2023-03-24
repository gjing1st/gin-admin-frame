// Path: internal/apiserver/model/request
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/28$ 16:09$

package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`           // 页码
	PageSize int    `json:"page_size" form:"page_size"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`     //关键字
	Sort     string `json:"sort" form:"sort"`           //排序参数
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
