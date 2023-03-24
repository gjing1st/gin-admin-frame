// Path: pkg/utils/global
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/24$ 14:24$

package dict

// 告警策略类型
const (
	AlertTypeMalfunction = 1 //故障告警
	AlertTypeThreshold   = 2 //阈值告警
	AlertTypeAbnormal    = 3 //异常告警
)

// 告警级别
const (
	AlertGradeHigh   = 1 + iota //告警级别高
	AlertGradeMiddle            //告警级别中
	AlertGradeLow               //告警级别低
)

// 监控内容
const (
	PolicyContentCPU       = iota + 1 //CPU
	PolicyContentMemory               //内存
	PolicyContentDisk                 //硬盘
	PolicyContentServerTPS            //服务能力

)

// 数据时段-持续时间
const (
	PolicyDurationNow       = iota + 1 //即时
	PolicyContentFiveMinute            //5分钟
	PolicyContentOneHour               //1小时
	PolicyContentOneDay                //1天
)

// 发送间隔
const (
	PolicySendDurationTenMinute    = iota + 1 //10分钟
	PolicySendDurationThirtyMinute            //30分钟
	PolicySendDurationSixtyMinute             //60分钟
)

var PolicySendDuration = make([]int, 3+1)

// 通知方式
const (
	PolicySendTypeEmail = iota + 1 //email
	PolicySendTypeSms              //短信
)

const (
	KSPolicySeverityCritical = "critical" //危险告警
	KSPolicySeverityError    = "error"    //重要告警
	KSPolicySummaryCPU       = "CPU使用率已达到阈值"
	KSPolicySummaryMemory    = "内存使用率已达到阈值"
	KSPolicySummaryDisk      = "硬盘使用率已达到阈值"
)

// AlertGradeArr 以上接口类型数量+1
var AlertGradeArr = [4]string{}

// 监控对象
const (
	PolicyResourcesCluster     = iota + 1 //ks集群节点
	PolicyResourcesCipher                 //密码机服务器
	PolicyResourcesSign                   //电子签章服务器
	PolicyResourcesTimeStamp              //时间戳服务器
	PolicyResourcesSignVer                //签名验签服务器
	PolicyResourcesCA                     //CA
	PolicyResourcesCloudCipher            //云密码机
)

// PolicyResourcesArr 以上接口类型数量+1
var PolicyResourcesArr = [8]string{}

// 策略启用禁用状态
const (
	PolicyStatusDisabled = 0 //禁用
	PolicyStatusEnable   = 1 //启用
)
