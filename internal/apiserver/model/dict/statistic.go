// Path: internal/pkg/dict
// FileName: statistic.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/9$ 18:26$

package dict

const (
	StatisticYesterdayThresholdTotal     = "yesterday_threshold_total"     //昨日阈值告警总数
	StatisticYesterdayMalfunctionTotal   = "yesterday_malfunction_total"   //昨日故障告警总数
	StatisticYesterdayThresholdPending   = "yesterday_threshold_pending"   //昨日阈值告警待处理数量
	StatisticYesterdayMalfunctionPending = "yesterday_malfunction_pending" //昨日故障告警待处理数量
)

// StatisticYesterdayTotal 昨日统计数据
type StatisticYesterdayTotal struct {
	StatisticYesterdayThresholdTotal     int64
	StatisticYesterdayMalfunctionTotal   int64
	StatisticYesterdayThresholdPending   int64
	StatisticYesterdayMalfunctionPending int64
}
