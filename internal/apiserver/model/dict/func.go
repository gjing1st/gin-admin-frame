// Path: internal/pkg/dict
// FileName: func.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/30$ 16:47$

package dict

import (
	"strings"
)

func init() {
	FillPolicySendDuration()
	FillPolicyResources()
	FillAlertGrade()
}

// FillPolicySendDuration
// @description: 装填策略发送间隔对应的分钟数
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/30 16:49
// @success:
func FillPolicySendDuration() {
	PolicySendDuration[PolicySendDurationTenMinute] = 10
	PolicySendDuration[PolicySendDurationThirtyMinute] = 30
	PolicySendDuration[PolicySendDurationSixtyMinute] = 60
}

// FillPolicyResources
// @description: 填充监控对象对应的含义
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/12 19:03
// @success:
func FillPolicyResources() {
	PolicyResourcesArr[PolicyResourcesCluster] = "集群节点"
	PolicyResourcesArr[PolicyResourcesCipher] = "密码机服务器"
	PolicyResourcesArr[PolicyResourcesSign] = "电子签章服务器"
	PolicyResourcesArr[PolicyResourcesTimeStamp] = "时间戳服务器"
	PolicyResourcesArr[PolicyResourcesSignVer] = "签名验签服务器"
	PolicyResourcesArr[PolicyResourcesCA] = "ca数字证书认证系统"
	PolicyResourcesArr[PolicyResourcesCloudCipher] = "云密码机"

}

// FillAlertGrade
// @description: 填充高级等级对应含义
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/12 19:04
// @success:
func FillAlertGrade() {
	AlertGradeArr[AlertGradeHigh] = "高级"
	AlertGradeArr[AlertGradeMiddle] = "中级"
	AlertGradeArr[AlertGradeLow] = "低级"
}

// SearchResources
// @description: 模糊查询时，查询监控对象
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/12 18:50
// @success:
func SearchResources(str string) (arr []int) {
	for i := 0; i < len(PolicyResourcesArr); i++ {
		n := strings.Index(PolicyResourcesArr[i], str)
		if n > -1 {
			arr = append(arr, i)
		}
	}
	return arr
}

// SearchAlertGradeArr
// @description: 模糊查询时，查询告警级别
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/12 18:57
// @success:
func SearchAlertGradeArr(str string) (arr []int) {
	for i := 0; i < len(AlertGradeArr); i++ {
		n := strings.Index(AlertGradeArr[i], str)
		if n > -1 {
			arr = append(arr, i)
		}
	}
	return arr
}
