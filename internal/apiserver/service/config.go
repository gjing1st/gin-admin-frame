// Path: internal/apiserver/service
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 10:57$

package service

import (
	"encoding/json"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/cache"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"time"
)

type ConfigService struct {
}

var (
	gCache      = cache.Cache{}
	configMysql = mysql.ConfigMysql{}
)

// GetValueStr
// @description: 根据k获取v
// @param: k string
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:08
// @success:
func (cs *ConfigService) GetValueStr(k string) (v string, errCode errcode.Err) {
	v, errCode = gCache.GetValueStr(k)
	if v == "" {
		vi, errCode1 := configMysql.GetValue(k)
		if errCode1 != 0 {
			errCode = errCode1
		}
		v = utils.String(vi)
		//存入缓存
		cs.SetCacheValue(k, v)
	}
	return
}

// GetInitStep
// @description: 获取初始化步骤
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 10:01
// @success:
func (cs *ConfigService) GetInitStep() (res response.InitStepValue, errCode errcode.Err) {
	v, errCode1 := cs.GetValueStr(dict.ConfigInitStep)
	if errCode1 != 0 {
		errCode = errCode1
		return
	}
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "初始化步骤值转json错误", "err": err, "data": v, "v": res})
		errCode = errcode.GafSysJsonUnMarshalErr
		return
	}
	return
}

func (cs *ConfigService) Get() {

}

// SetCacheValue
// @description: 设置缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 15:27
// @success:
func (cs *ConfigService) SetCacheValue(k, v interface{}) (errCode errcode.Err) {
	errCode = gCache.RemoveSet(k, v)
	return
}

// SetValue
// @description: 设置k-v值，先持久化再更新缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/15 14:01
// @success:
func (cs *ConfigService) SetValue(k string, v interface{}) (errCode errcode.Err) {
	errCode = configMysql.SetValue(k, v)
	if errCode == 0 {
		cs.SetCacheValue(k, v)
	}
	return
}

// GetRunDate
// @description: 获取系统运行时长
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 11:36
// @success:
func (cs *ConfigService) GetRunDate() (res response.SysRunDate, errCode errcode.Err) {
	v, errCode1 := cs.GetValueStr(dict.ConfigSysBreakDate)
	if errCode1 != 0 {
		errCode = errCode1
		return
	}
	//转换为time类型
	breakTime, err := time.ParseInLocation(global.TimeFormat, v, time.Local)
	if err != nil {
		errCode = errcode.GafSysTimeParseErr
		functions.AddErrLog(log.Fields{"msg": "breakTime时间转换出错", "err": err, "time": v})
		return
	}
	//当前时间与故障时间对比
	d := time.Now().Sub(breakTime)
	//day
	h := d / time.Hour //相差的小时数
	res.Minute = int(d % time.Hour / time.Minute)
	res.Hour = int(h % 24)
	res.Day = int(h / 24)
	return
}

// SetInitStep
// @description: 初始化状态修改
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 10:55
// @success:
func (cs *ConfigService) SetInitStep(s int) (errCode errcode.Err) {
	var step response.InitStepValue
	if s == dict.InitStepUser {
		step.User = dict.InitStepValueDown
	} else if s == dict.InitStepNetwork {
		step.User = dict.InitStepValueDown
		step.Network = dict.InitStepValueDown
	} else if s == dict.InitStepReset {
		//重置全部置零
	}
	stepStr, err := json.Marshal(step)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "初始化步骤转换"})
		errCode = errcode.GafSysJsonMarshalErr
		return
	}
	errCode = configMysql.SetValue(dict.ConfigInitStep, string(stepStr))
	if errCode == 0 {
		cs.SetCacheValue(dict.ConfigInitStep, stepStr)
	}
	return

}

// InitReset
// @description: 初始化时重置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 18:51
// @success:
func (cs *ConfigService) InitReset() (errCode errcode.Err) {
	errCode = cs.SetInitStep(dict.InitStepReset)
	if errCode != 0 {
		return
	}
	//删除管理员
	errCode = userMysql.ResetUser(nil)
	return
}

// SysReset
// @description: 恢复出厂设置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 13:48
// @success:
func (cs *ConfigService) SysReset(req *request.UserLogin) (errCode errcode.Err) {
	//初始化状态
	errCode = cs.SetInitStep(dict.InitStepReset)
	if errCode != 0 {
		return
	}
	//删除管理员
	errCode = userMysql.ResetUser(nil)
	//清空日志记录
	errCode = sysLogMysql.TruncateTable(nil)
	//TODO 删除其他数据
	return
}

// VersionInfo
// @description: 获取当前版本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/1/4 9:35
// @success:
func (cs *ConfigService) VersionInfo() (res response.VersionInfo) {
	res.Version, _ = cs.GetValueStr(dict.ConfigVersion)
	res.Manufacturer = config.Config.VersionInfo.Manufacturer
	res.Serial = config.Config.VersionInfo.Serial
	res.DeviceModel = config.Config.VersionInfo.Serial
	return
}

// SetNetwork
// @description: 初始化配置网络
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 20:29
// @success:
func (cs *ConfigService) SetNetwork(req *request.SetNetwork) (errCode errcode.Err) {
	var sysService SysService
	err := sysService.setNetwork(req.Admin.Addr, req.Admin.Gateway, req.Admin.Netmask, config.Config.Adapter.AdminPath)
	if err != nil {
		errCode = errcode.GafSysNetworkErr
		return
	}
	return
}
