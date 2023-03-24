// Path: internal/apiserver/store/mysql
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 10:59$

package mysql

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type ConfigMysql struct {
}

// GetValue
// @description: 根据k获取v
// @param: k string
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:06
// @success:
func (cs *ConfigMysql) GetValue(k string) (v interface{}, errCode int) {
	var conf entity.Config
	err := store.DB.Where("name = ?", k).Where("status = ?", dict.StatusEnable).First(&conf).Error
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "mysql查询value错误", "key": k})
		return "", errcode2.GafSysDBFindErr
	}
	return conf.Value, errcode2.SuccessCode
}

// SetValue
// @description: 修改值
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 10:57
// @success:
func (cs ConfigMysql) SetValue(k string, v interface{}) (errCode int) {
	conf := &entity.Config{
		Name:  k,
		Value: v,
	}
	err := store.DB.Model(&entity.Config{}).Where("name = ?", k).Update("value", v).Error
	if err != nil {
		if err != nil {
			functions.AddErrLog(log.Fields{"err": err, "msg": "mysql修改config表错误", "key": k})
			return errcode2.GafSysDBUpdateErr
		}
	}
	go func() {
		//此处防止k不存在表中
		_, errCode1 := cs.GetValue(k)
		if errCode1 != 0 {
			store.DB.Model(&entity.Config{}).Where("name = ?", k).Save(conf)
		}
	}()
	return
}
