// Path: internal/apiserver/controller
// FileName: syslog.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/9$ 10:44$

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/service"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
)

type SysLogController struct {
}

var sysLogService service.SysLogService

// List
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 10:48
// @success:
func (slc *SysLogController) List(c *gin.Context) {
	var req request.SysLogList
	_ = c.ShouldBindQuery(&req)
	if req.PageSize == 0 {
		req.PageSize = global.PageSizeDefault
	}
	if req.Category == 0 {
		req.Category = dict.SysLogCategoryOperation
	}

	list, total, errCode := sysLogService.List(&req)
	if errCode != 0 {
		response.Failed(errCode, global.QueryFailed, c)

	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, global.QuerySuccess, c)
	}
}
