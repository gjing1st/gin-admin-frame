// Path: internal/apiserver/service
// FileName: cron.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 22:59$

package service

import (
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils/crontab"
	"strconv"
	"time"
)

type CrontabService struct {
}

var crontabStore mysql.CrontabStore

// CrontabPublish
// @description: 定时发布
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 10:53
// @success:
func CrontabPublish() {
	now := time.Now()
	cronList, err := crontabStore.SearchWaitingStatus(now)
	if err != nil {
		return
	}
	_ = store.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range cronList {

			if err != nil {
				return err
			}
			//删除待执行的定时任务
			crontabId := v.ID
			err = crontabStore.Delete(tx, crontabId)
		}

		return err
	})

}

// AddCron
// @description: 添加定时计划
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 13:52
// @success:
func AddCron() {
	//seconds := config.Config.CrontabTime
	//fmt.Println("定时任务设置时间seconds=", seconds)
	//每个整点开始定时任务
	//_, err := crontab.GetCron().AddFunc("00 00 * * * *", CrontabPublish)
	//if err != nil {
	//	functions.AddErrLog(log.Fields{"定时任务添加失败": err})
	//}
}

var cronId cron.EntryID

// AutoCron
// @description: 添加自动升级定时任务
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 18:13
// @success:
func AutoCron(updateRange, hour string) {
	crontab.GetCron().Remove(cronId)
	weekDay := time.Now().Weekday()
	day := time.Now().Day()
	var spec string
	switch updateRange {
	case "day":
		spec = "0 0 " + hour + " 00 * * *"
	case "week":
		spec = "0 0 " + hour + " * * " + strconv.Itoa(int(weekDay))
	case "month":
		spec = "0 0 " + hour + " " + strconv.Itoa(day) + " * *"
	}
	var ss SysService
	cronId, _ = crontab.GetCron().AddFunc(spec, ss.CronUpdate)
}
