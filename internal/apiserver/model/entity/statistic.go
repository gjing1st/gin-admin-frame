// Path: internal/apiserver/model/entity
// FileName: statistic.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/9$ 18:14$

package entity

// Statistic 结构体
type Statistic struct {
	BaseModel
	Key   string `json:"key" form:"key" gorm:"column:key;comment:键;type:varchar(32);index:key_name;NOT NULL;"`
	Value string `json:"value" form:"value" gorm:"column:value;comment:值;type:int(11);"`
	//Date  string `json:"date" form:"date" gorm:"column:date;comment:日期;type:varchar(32)"`
}

// TableName Category 表名
func (Statistic) TableName() string {
	return "statistic"
}
