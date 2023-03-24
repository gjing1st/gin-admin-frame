// Path: internal/apiserver/store/cache
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 11:09$

package cache

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
	"time"
)

type Cache struct {
}

// GetValueStr
// @description: 从缓存中获取value
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:20
// @success:
func (cc *Cache) GetValueStr(k interface{}) (v string, errCode int) {
	value, err := store.GC.Get(k)
	if err != nil {
		errCode = errcode2.GafSysCacheGetErr
		functions.AddWarnLog(log.Fields{"err": err, "msg": "cache获取value错误", "key": k})
		if err == errcode2.ErrKeyNotFound {
			errCode = errcode2.SuccessCode
		}
	}
	v = utils.String(value)
	return
}

// RemoveSet
// @description: 删除旧缓存设置新缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 11:24
// @success:
func (cc *Cache) RemoveSet(k, v interface{}) (errCode int) {
	store.GC.Remove(k)
	err := store.GC.SetWithExpire(k, v, time.Hour*8)
	if err != nil {
		functions.AddErrLog(log.Fields{"err": err, "msg": "cache缓存key,value错误", "key": k, "value": v})
		errCode = errcode2.GafSysCacheSetErr
	}
	return
}

// Get
// @description: 获取缓存中k对应的值
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:00
// @success:
func (cc Cache) Get(k interface{}) (v interface{}, errCode int) {
	value, err := store.GC.Get(k)
	if err != nil {
		errCode = errcode2.GafSysCacheGetErr
		functions.AddWarnLog(log.Fields{"err": err, "msg": "cache获取value错误", "key": k})
		if err == errcode2.ErrKeyNotFound {
			errCode = 0
		}
	}
	v = value
	return
}

// Remove
// @description: 删除缓存
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:10
// @success:
func (cc Cache) Remove(k interface{}) (errCode int) {
	store.GC.Remove(k)
	return
}
