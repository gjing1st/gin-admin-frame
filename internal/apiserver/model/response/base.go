// Path: internal/apiserver/model/response
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/28$ 17:24$

package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/global"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/pkg/errcode"
	"github.com/goccy/go-json"
	"net/http"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// Response http.code=200时统一返回格式
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func OkWithMessage(message string, c *gin.Context) {
	Result(errcode.SuccessCode, map[string]interface{}{}, message, c)
}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errcode.SuccessCode, data, message, c)
}

// Ok
// @description: 请求成功
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:03
// @success: http_code = 200
func Ok(message string, c *gin.Context) {
	Result(errcode.SuccessCode, gin.H{}, message, c)
}

// Created
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/27 17:43
// @success:
func Created(message string, c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		errcode.SuccessCode,
		gin.H{},
		message,
	})
}

// OkWithData
// @description: 请求成功并返回data数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:37
// @success:
func OkWithData(data interface{}, message string, c *gin.Context) {
	Result(errcode.SuccessCode, data, message, c)
}

// Failed
// @description: 返回错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/17 18:42
// @success: http_code = 500
func Failed(code int, message string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		gin.H{},
		message,
	})
}

// FailedNotFound
// @description: 数据没找到
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:31
// @success: http_code = 404
func FailedNotFound(code int, c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		code,
		gin.H{},
		global.DataNotFound,
	})
}

// ParamErr
// @description: 参数请求错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:11
// @success: http_code = 400
func ParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, global.RequestParamErr)
}

// OkWithLog
// @description: 返回操作成功并记录操作日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:16
// @success:
func OkWithLog(message string, username string, content string, req interface{}, c *gin.Context) {
	Result(errcode.SuccessCode, gin.H{}, message, c)
	var sysLog entity.SysLog
	sysLog.Result = dict.SysLogResultOk
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	reqJson, _ := json.Marshal(req)
	sysLog.RequestData = string(reqJson)
	mysql.SysLogMysql{}.Create(nil, &sysLog)
}

// OkWithSysLog
// @description: 返回操作成功并记录系统日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/13 17:26
// @success:
func OkWithSysLog(message string, username string, content string, c *gin.Context) {
	Result(errcode.SuccessCode, gin.H{}, message, c)
	var sysLog entity.SysLog
	sysLog.Result = dict.SysLogResultOk
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	sysLog.Category = dict.SysLogCategorySys
	mysql.SysLogMysql{}.Create(nil, &sysLog)
}

// FailWithSysLog
// @description: 返回操作失败，并记录失败日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/13 17:29
// @success:
func FailWithSysLog(code int, message string, username string, content string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		gin.H{},
		message,
	})
	var sysLog entity.SysLog
	sysLog.Result = dict.AdminLogResultFail
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	sysLog.Category = dict.SysLogCategorySys
	mysql.SysLogMysql{}.Create(nil, &sysLog)
}

// OkWithDataLog
// @description: 返回操作成功并记录操作日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:30
// @success:
func OkWithDataLog(data interface{}, message string, username string, content string, req interface{}, c *gin.Context) {
	Result(errcode.SuccessCode, data, message, c)
	var sysLog entity.SysLog
	sysLog.Result = dict.AdminLogResultOk
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	reqJson, _ := json.Marshal(req)
	sysLog.RequestData = string(reqJson)
	mysql.SysLogMysql{}.Create(nil, &sysLog)

}

// FailWithLog
// @description: 返回操作失败，并记录失败日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 16:20
// @success:
func FailWithLog(code int, message string, username string, content string, req interface{}, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		gin.H{},
		message,
	})
	var sysLog entity.SysLog
	sysLog.Result = dict.AdminLogResultFail
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	reqJson, _ := json.Marshal(req)
	sysLog.RequestData = string(reqJson)
	mysql.SysLogMysql{}.Create(nil, &sysLog)
}

// FailWithDataLog
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 17:48
// @success:
func FailWithDataLog(data interface{}, code int, message string, username string, content string, req interface{}, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		data,
		message,
	})
	var sysLog entity.SysLog
	sysLog.Result = dict.AdminLogResultFail
	sysLog.ClientIP = c.ClientIP()
	sysLog.Content = content
	if username == "" {
		userName, _ := c.Get("username")
		username = utils.String(userName)
	}
	sysLog.Username = username
	reqJson, _ := json.Marshal(req)
	sysLog.RequestData = string(reqJson)
	mysql.SysLogMysql{}.Create(nil, &sysLog)
}

// Unauthorized
// @description: 未登录的错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 18:02
// @success:
func Unauthorized(errCode int, message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		errCode,
		gin.H{},
		message,
	})
}

// Forbidden
// @description: 没有权限
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 18:05
// @success:
func Forbidden(errCode int, message string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		errCode,
		gin.H{},
		message,
	})
}
