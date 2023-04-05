// Path: internal/apiserver/model/dict
// FileName: config.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 9:47$

package dict

const (
	ConfigInitKey           = "initialized"      //系统是否已初始化
	ConfigSysFirstStartDate = "first_start_date" //系统首次运行的时间
	ConfigSysBreakDate      = "sys_break_date"   //系统故障时间
	ConfigLoginType         = "login_type"       //登录方式
	ConfigInitStep          = "init_step"        //初始化步骤
	ConfigGuideStep         = "guide_step"       //向导步骤
	ConfigVersion           = "version"          //当前版本信息
	ConfigLatestVersion     = "latest_version"   //最新版本信息
	ConfigBackupTime        = "backup_time"      //备份时间
	ConfigRestoreTime       = "restore_time"     //恢复时间
	ConfigAutoUpdate        = "auto_update"      //是否自动更新
	ConfigUpdateRange       = "update_range"     //自动更新周期
	ConfigUpdateTime        = "update_time"      //自动更新时间
)

const (
	LoginTypePasswd      = 1 //用户名口令
	LoginTypeFrontUKey   = 2 //前端UKey登录
	LoginTypeBackendUKey = 3 //后端UKey登录
)

// 是否已初始化完成
const (
	InitStepValueNot  = 0 //未配置
	InitStepValueDown = 1 //已完成配置
)

// 初始化步骤
const (
	InitStepUser    = 1 //步骤添加管理
	InitStepNetwork = 2 //步骤配置网络
	InitStepReset   = 3 //初始化重置
)

// GuideStepValue 向导步骤
type GuideStepValue struct {
	DeviceKey int `json:"device_key"`
	Sm4Key    int `json:"sm4_key"`
	Sm2Key    int `json:"sm2_key"`
	BakKey    int `json:"bak_key"`
	NetWork   int `json:"network"`
}
