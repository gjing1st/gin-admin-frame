// Path: internal/apiserver/service
// FileName: sys.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/2/9$ 19:55$

package service

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/response"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/pkg/errcode"
	"github.com/gjing1st/gin-admin-frame/third/assist"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"net"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SysService struct {
}

// ServerStatus
// @description: 设备运行状态
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/9 20:29
// @success:
func (ss *SysService) ServerStatus() (res response.ServerStatus, errCode int) {
	res.ServiceStatus = dict.ServiceStatusInit
	res.RunStatus = dict.RunStatusAbnormal
	var wg sync.WaitGroup
	//wg.Add(4)
	wg.Add(4)
	//serviceStatus := make(chan bool, 3)
	serviceStatus := make(chan bool, 1)
	runStatus := make(chan bool, 3)
	var err error
	go func() {
		//1. 是否存在设备密钥
		defer wg.Done()
		//TODO
		if err != nil {
			serviceStatus <- false
			return
		}
		serviceStatus <- true
	}()

	go func() {
		//运行状态
		defer wg.Done()
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	}()
	go func() {
		//运行状态
		defer wg.Done()
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	}()
	go func() {
		//运行状态
		defer wg.Done()
		//TODO
		if err != nil {
			runStatus <- false
			return
		}
		runStatus <- true
	}()

	wg.Wait()
	//运行状态
	//runStatusRes := <-runStatus
	//if runStatusRes {
	//	res.RunStatus = dict.RunStatusNormal
	//}
	//服务状态
	serviceStatusRes := <-serviceStatus
	if serviceStatusRes {
		res.ServiceStatus = dict.ServiceStatusReady
	}
	//运行状态
	for i := 0; i < 3; i++ {
		status := <-runStatus

		if status == false {
			//有未完成的
			return
		}
	}
	res.RunStatus = dict.RunStatusNormal
	return
}

// Reboot
// @description: 服务器重启
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/11 16:30
// @success:
func (ss *SysService) Reboot() (errCode int) {
	err := utils.DockerRunCommand("reboot")
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "服务器重启指令执行失败", "err": err})
		errCode = errcode.GafSysCmdErr
	}
	return
}

// RestartNetwork
// @description: 重启网卡
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/11 16:32
// @success:
func (ss *SysService) RestartNetwork() (errCode int) {
	err := utils.DockerRunCommand("systemctl restart network")
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "服务器重启网卡指令执行失败", "err": err})
		errCode = errcode.GafSysCmdErr
	}
	return
}

// GetNetwork
// @description: 获取当前网卡信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 17:25
// @success:
func (ss *SysService) GetNetwork() (res *response.GetNetwork, errCode int) {
	//管理网口
	adminNetwork, errCode1 := ss.getNetwork(config.Config.Adapter.AdminPath)
	if errCode1 != 0 {
		return
	}
	//密码服务网口
	sdfNetwork, errCode2 := ss.getNetwork(config.Config.Adapter.CipherPath)
	if errCode2 != 0 {
		//return
	}
	res = &response.GetNetwork{
		Admin: adminNetwork,
		SDF:   sdfNetwork,
	}
	return
}
func getNetmask(addr string) string {
	_, ipNet, _ := net.ParseCIDR(addr)
	val := make([]byte, len(ipNet.Mask))
	copy(val, ipNet.Mask)

	var s []string
	for _, i := range val[:] {
		s = append(s, strconv.Itoa(int(i)))
	}
	return strings.Join(s, ".")
}

// getNetwork
// @description: 获取网卡的网络配置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 19:03
// @success:
func (ss *SysService) getNetwork(filePath string) (res response.Network, errCode int) {
	conf, err := ini.Load(filePath)
	if err != nil {
		log.Println("ini conf newkey error:", err)
		functions.AddErrLog(log.Fields{"adapterPath": filePath, "err": err})
		errCode = errcode.GafSysNetworkErr
		return
	}
	addr := conf.Section("").Key("IPADDR").String()
	prefix := conf.Section("").Key("PREFIX").String()
	netmask := getNetmask(addr + "/" + prefix)

	res = response.Network{
		Addr:    addr,
		Gateway: conf.Section("").Key("GATEWAY").String(),
		Netmask: netmask,
	}
	return
}

// SetNetwork
// @description: 配置网络
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/13 20:21
// @success:
func (ss *SysService) SetNetwork(req *request.SetNetwork) (errCode int) {
	if req.Admin.Addr != "" {
		err := ss.setNetwork(req.Admin.Addr, req.Admin.Gateway, req.Admin.Netmask, config.Config.Adapter.AdminPath)
		if err != nil {
			errCode = errcode.GafSysNetworkErr
			return
		}
	}
	if req.SDF.Addr != "" {
		err := ss.setNetwork(req.SDF.Addr, req.SDF.Gateway, req.SDF.Netmask, config.Config.Adapter.CipherPath)
		if err != nil {
			errCode = errcode.GafSysNetworkErr
			return
		}
	}

	errCode = ss.RestartNetwork()
	return
}

// 配置网卡信息
func (ss *SysService) setNetwork(addr, gateway, netmask, filePath string) error {
	netmaskArray := strings.Split(netmask, ".")
	//netmaskArray := convert(netmaskSlice)
	prefix, _ := net.IPv4Mask(byte(utils.Int(netmaskArray[0])), byte(utils.Int(netmaskArray[1])), byte(utils.Int(netmaskArray[2])), byte(utils.Int(netmaskArray[3]))).Size()
	conf, err := ini.Load(filePath)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "配置网络错误", "addr": addr, "gateway": gateway, "netmask": netmask, "err": err})
		return err
	}

	if conf.Section("").Key("IPADDR").String() == "" {
		_, err = conf.Section("").NewKey("IPADDR", addr)
	} else {
		conf.Section("").Key("IPADDR").SetValue(addr)
	}

	if conf.Section("").Key("PREFIX").String() == "" {
		_, err = conf.Section("").NewKey("PREFIX", strconv.Itoa(prefix))
	} else {
		conf.Section("").Key("PREFIX").SetValue(strconv.Itoa(prefix))
	}

	if conf.Section("").Key("GATEWAY").String() == "" {
		_, err = conf.Section("").NewKey("GATEWAY", gateway)
	} else {
		conf.Section("").Key("GATEWAY").SetValue(gateway)
	}

	if conf.Section("").Key("ONBOOT").String() == "" {
		_, err = conf.Section("").NewKey("ONBOOT", "yes")
	} else {
		conf.Section("").Key("ONBOOT").SetValue("yes")
	}

	if conf.Section("").Key("BOOTPROTO").String() == "" {
		_, err = conf.Section("").NewKey("BOOTPROTO", "none")
	} else {
		conf.Section("").Key("BOOTPROTO").SetValue("none")
	}

	if conf.Section("").Key("ZONE").String() == "" {
		_, err = conf.Section("").NewKey("ZONE", "public")
	} else {
		conf.Section("").Key("ZONE").SetValue("public")
	}

	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "配置网络错误", "addr": addr, "gateway": gateway, "netmask": netmask, "err": err})
		return err
	}

	err = conf.SaveTo(filePath)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "配置网络错误", "addr": addr, "gateway": gateway, "netmask": netmask, "err": err})
		return err
	}

	b, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "配置网络错误", "addr": addr, "gateway": gateway, "netmask": netmask, "err": err})
		return err
	}

	str := string(b) // convert content to a 'string'
	str2 := strings.Replace(str, " ", "", -1)

	err = ioutil.WriteFile(filePath, []byte(str2), 0666)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "配置网络错误", "addr": addr, "gateway": gateway, "netmask": netmask, "err": err})
		return err
	}

	// reboot()
	return nil
}

// VersionInfo
// @description: 关于
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/14 15:36
// @success:
func (ss *SysService) VersionInfo() (res response.VersionInfo) {
	res.Version = config.Config.VersionInfo.Version
	res.Manufacturer = config.Config.VersionInfo.Manufacturer
	res.Serial = config.Config.VersionInfo.Serial
	res.DeviceModel = config.Config.VersionInfo.DeviceModel
	return
}

// Update
// @description: 更新
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/15 18:29
// @success:
func (ss *SysService) Update() (errCode int, version string) {
	//err := utils.DockerRunCommand("/home/app/hss/up.sh")
	//if err != nil {
	//	functions.AddErrLog(log.Fields{"msg": "执行升级失败", "err": err})
	//	errCode = errcode.GafSysCmdErr
	//}
	var cs ConfigService
	lastVersion, errCode1 := cs.GetValueStr(dict.ConfigLatestVersion)
	if errCode1 != 0 {
		return errCode1, ""
	}
	go func() {
		err := assist.Update(config.Config.UploadPath + "hss_" + lastVersion)
		if err != nil {
			functions.AddErrLog(log.Fields{"msg": "请求助手升级失败", "err": err})
		}
		ss.UpdateStateVersion()
	}()

	return
}

// CronUpdate
// @description: 定时升级
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/13 17:22
// @success:
func (ss *SysService) CronUpdate() {
	err, version := ss.Update()
	//记录日志
	content := "升级至" + version + "版本"
	var sysLog entity.SysLog
	sysLog.Content = content
	if err != 0 {
		functions.AddErrLog(log.Fields{"msg": "执行升级失败", "err": err})
		sysLog.Result = dict.AdminLogResultFail
	} else {
		sysLog.Result = dict.SysLogResultOk
	}

	mysql.SysLogMysql{}.Create(nil, &sysLog)

	return
}

// UpdateVersionInfo
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 9:44
// @success:
func (ss *SysService) UpdateVersionInfo() (res response.UpdateVersionInfo, errCode int) {
	var cs ConfigService
	res.CurrentVersion, errCode = cs.GetValueStr(dict.ConfigVersion)
	res.LatestVersion, errCode = cs.GetValueStr(dict.ConfigLatestVersion)
	res.CanUpdate = utils.VersionCompare(res.CurrentVersion, res.LatestVersion)
	return
}

// DealFile
// @description: 处理升级包文件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 14:27
// @success:
func (ss *SysService) DealFile(fileName, filePath string) (errCode int) {
	fullName := filePath + fileName

	//解压缩升级包
	utils.UnzipDir(fullName, filePath)
	files := strings.Split(fileName, "_")
	if len(files) < 2 {
		return
	}
	var cs ConfigService
	currentVersion, errCode := cs.GetValueStr(dict.ConfigVersion)
	version := utils.UnExt(files[1])
	if version != currentVersion {
		if !utils.VersionCompare(currentVersion, version) {
			errCode = errcode.GafUpdateFileErr
			return
		}
	}
	//记录最新版本
	errCode = cs.SetValue(dict.ConfigLatestVersion, version)
	if errCode != 0 {
		return
	}
	//解压后的路径
	dirPath := filePath + utils.UnExt(fileName)
	//解压后处理解压后的文件
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		functions.AddErrLog(log.Fields{"dirPath": dirPath, "msg": "读取升级包目录中的文件错误", "err": err})
		errCode = errcode.GafUpdateFileReadErr
		return
	}
	//遍历解压后目录中的所有文件
	for _, fileInfo := range fileInfos {
		filename := fileInfo.Name()
		fileExt := path.Ext(filename)
		if fileExt == ".tar" {
			//导入镜像
			err = utils.RunCommand("docker", "load", "-i", dirPath+"/"+filename)
			if err != nil {
				errCode = errcode.GafUpdateFileLoadErr
				return
			}
		} else if fileExt == ".tgz" {

		}

	}
	return
}

func (ss *SysService) DealFileV1(fileName, filePath string) (errCode int) {
	fullName := filePath + fileName

	//解压缩升级包
	utils.UnzipDir(fullName, filePath)
	files := strings.Split(fileName, "_")
	if len(files) < 2 {
		return
	}
	var cs ConfigService
	currentVersion, errCode := cs.GetValueStr(dict.ConfigVersion)
	version := utils.UnExt(files[1])
	if version != currentVersion {
		if !utils.VersionCompare(currentVersion, version) {
			errCode = errcode.GafUpdateFileErr
			return
		}
	}
	//记录最新版本
	errCode = cs.SetValue(dict.ConfigLatestVersion, version)
	if errCode != 0 {
		return
	}
	//升级助手的授予install.sh权限有问题，改到这里实现
	err := utils.RunCommand("chmod", "+x", filePath+utils.UnExt(fileName)+"/install.sh")
	if err != nil {
		errCode = errcode.GafUpdateFileLoadErr
		return
	}
	//解压后的路径
	//dirPath := filePath + utils.UnExt(fileName)
	////解压后处理解压后的文件
	//fileInfos, err := ioutil.ReadDir(dirPath)
	//if err != nil {
	//	functions.AddErrLog(log.Fields{"dirPath": dirPath, "msg": "读取升级包目录中的文件错误", "err": err})
	//	errCode = errcode.GafUpdateFileReadErr
	//	return
	//}
	////遍历解压后目录中的所有文件
	//for _, fileInfo := range fileInfos {
	//	filename := fileInfo.Name()
	//	fileExt := path.Ext(filename)
	//	if fileExt == ".tar" {
	//		//导入镜像
	//		err = utils.RunCommand("docker", "load", "-i", dirPath+"/"+filename)
	//		if err != nil {
	//			errCode = errcode.GafUpdateFileLoadErr
	//			return
	//		}
	//	} else if fileExt == ".tgz" {
	//
	//	}
	//
	//}
	return
}

// GetAuto
// @description: 当前自动更新配置
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 16:22
// @success:
func (ss *SysService) GetAuto() (res response.AutoUpdateConfig, errCode int) {
	var cs ConfigService
	auto, errCode1 := cs.GetValueStr(dict.ConfigAutoUpdate)
	if errCode1 != 0 {
		errCode = errCode1
		return
	}
	res.AutoUpdate = utils.Bool(auto)
	res.UpdateRange, errCode = cs.GetValueStr(dict.ConfigUpdateRange)
	if errCode != 0 {
		return
	}
	res.Time, errCode = cs.GetValueStr(dict.ConfigUpdateTime)
	return

}

// SetAuto
// @description: 设置自动更新策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 16:33
// @success:
func (ss *SysService) SetAuto(req *request.AutoUpdateConfig) (errCode int) {
	var cs ConfigService
	errCode = cs.SetValue(dict.ConfigAutoUpdate, req.AutoUpdate)
	if errCode != 0 {
		return
	}
	errCode = cs.SetValue(dict.ConfigUpdateRange, req.UpdateRange)
	if errCode != 0 {
		return
	}
	errCode = cs.SetValue(dict.ConfigUpdateTime, req.Time)
	go func() {
		AutoCron(req.UpdateRange, req.Time)
	}()
	return
}

// UpdateStateVersion
// @description: 是否升级完成，升级完成，更新当前版本
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/17 17:15
// @success:
func (ss *SysService) UpdateStateVersion() {
	var cs ConfigService
	go func() {
		//定时升级
		updateRange, _ := cs.GetValueStr(dict.ConfigUpdateRange)
		updateHour, _ := cs.GetValueStr(dict.ConfigUpdateTime)
		if updateHour != "" && updateRange != "" {
			AutoCron(updateRange, updateHour)
		}
	}()
	var updateFinished bool
	nowTime, allTimes := 0, 10
	for !updateFinished {
		if nowTime > allTimes {
			return
		}
		res, _ := assist.UpdateState()
		if res.State == assist.UpdateSuccess {
			//升级成功，更新当前版本

			latestVersion, errCode := cs.GetValueStr(dict.ConfigLatestVersion)
			if errCode != 0 {
				functions.AddErrLog(log.Fields{"msg": "获取最新版本错误", "errCode": errCode})
			}
			if latestVersion != "" {
				errCode = cs.SetValue(dict.ConfigVersion, latestVersion)
				if errCode != 0 {
					functions.AddErrLog(log.Fields{"msg": "更新当前版本错误", "errCode": errCode})
				}
			}

			updateFinished = true
		} else if res.State == assist.UpdateFailed {
			updateFinished = true
		} else {
			time.Sleep(time.Second * 2)
		}
		nowTime++
	}
	return
}
