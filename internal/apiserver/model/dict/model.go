// $
// model.go
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/28$ 11:38$

package dict

const (
	StatusDefault   = iota //默认状态
	StatusEnable           //启用状态	1
	StatusForbidden        //禁止状态	2
)

const (
	CrontabStatusDefault  = iota
	CrontabStatusWaiting  //定时任务待执行
	CrontabStatusExecuted //定时任务已执行

)

// 文章状态
const (
	ArticleStatusDefault     = iota
	ArticleStatusWaitPublish //待发布	1
	ArticleStatusPublished   //已发布	2
)

const (
	ArticleTypeAdd        = iota + 1 //管理员添加的通知
	ArticleTypePushNotice            //推送过来的通知
)

const (
	CategoryNotification    = iota + 1 //系统通知		1
	CategoryHelp                       //帮助中心		2
	CategorySafeInformation            //安全资讯		3
	CategoryLawsRegulations            //法律法规		4
	CategoryCipherStandard             //商用密码标准规范	5
	CategoryGradePolicy                //等保政策文件		6
	CategoryEstimatePolicy             //密评政策文件		7

)

const (
	AlertMessageStatusPending  = 1 + iota //告警中-待处理
	AlertMessageStatusFinished            //已结束-告警已处理完成
)
