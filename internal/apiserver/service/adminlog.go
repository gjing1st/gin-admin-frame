// Path: internal/apiserver/service
// FileName: adminlog.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 16:13$

package service

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
)

type SysLogService struct {
}

var sysLogMysql mysql.SysLogMysql

// Create
// @description: 添加管理员日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:14
// @success:
func (sls *SysLogService) Create(log *request.SysLogCreate, req interface{}) {

}

// List
// @description:日志列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 10:52
// @success:
func (sls *SysLogService) List(req *request.SysLogList) (res interface{}, total int64, errCode int) {
	res, total, errCode = sysLogMysql.List(req)
	return
}
