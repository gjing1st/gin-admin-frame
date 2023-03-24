// Path: internal/pkg/functions
// FileName: log_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 16:03$

package functions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/third/kubesphere/monitor"
	"testing"
)

func TestAddErrLog(t *testing.T) {
	AddErrLog(log.Fields{"test": "1111111"})

}

func TestGetToken(t *testing.T) {
	config.InitConfig()
	s, err := monitor.GetToken()
	fmt.Println("s=", s, "err=", err)
}
