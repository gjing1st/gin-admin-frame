// Path: third/assist
// FileName: update.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/17$ 16:22$

package assist

import (
	"errors"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	urlAssistUpdate      = "/engine-assist/v1/update"       //POST 升级
	urlAssistUpdateState = "/engine-assist/v1/update/state" //GET 升级状态
)

// Update
// @description: 请求助手进行升级
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/17 16:25
// @success:
func Update(path string) error {
	var req struct {
		Path string `json:"path"`
	}
	req.Path = path
	reqUrl := config.Config.AssistAddr + urlAssistUpdate
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&req).
		Post(reqUrl)
	if err != nil || resp.StatusCode() != http.StatusOK {
		functions.AddErrLog(log.Fields{"err": err, "resp": resp, "msg": "请求助手升级失败"})
		return errors.New("请求助手升级失败")
	}
	return nil
}

type UpdateStateResponse struct {
	State int `json:"state"`
}

const (
	UpdatePending = 0 //升级中
	UpdateSuccess = 1 //升级完成
	UpdateFailed  = 2 //升级失败
)

// UpdateState
// @description: 升级状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/17 17:14
// @success:
func UpdateState() (res UpdateStateResponse, err error) {
	reqUrl := config.Config.AssistAddr + urlAssistUpdateState
	client := resty.New().SetTimeout(time.Second * 3)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&res).
		Get(reqUrl)
	if err != nil || resp.StatusCode() != http.StatusOK {
		functions.AddErrLog(log.Fields{"err": err, "resp": resp, "msg": "请求助手升级失败"})
		return res, errors.New("请求助手升级失败")
	}
	return
}
