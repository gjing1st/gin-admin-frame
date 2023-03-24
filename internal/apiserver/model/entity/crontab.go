// Path: internal/apiserver/model
// FileName: crontab.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 17:52$

package entity

import (
	"time"
)

type Crontab struct {
	BaseModel
	CrontabTime time.Time `json:"crontab_time" gorm:"column:crontab_time;comment:定时发布时间;"`
	//Status      int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:1;"`
	ArticleId uint `json:"-" gorm:"column:article_id;comment:文章id;size:32;"`
	NoticeId  uint `json:"-" gorm:"column:notice_id;comment:通知id;size:32;"`
}
