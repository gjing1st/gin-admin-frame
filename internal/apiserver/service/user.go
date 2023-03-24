// Path: internal/apiserver/service
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 15:54$

package service

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/dict"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/entity"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/model/request"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/mysql"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/functions"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils"
	"github.com/gjing1st/gin-admin-frame/internal/pkg/utils/gm"
	errcode2 "github.com/gjing1st/gin-admin-frame/pkg/errcode"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
}

var userMysql = mysql.UserMysql{}

// Create
// @description: 创建管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 15:58
// @success:
func (us *UserService) Create(req *request.UserCreate) (errCode int) {
	var cs ConfigService
	loginType, errCode1 := cs.GetValueStr(dict.ConfigLoginType)
	if errCode1 != 0 {
		return errCode1
	}
	if utils.Int(loginType) == dict.LoginTypeBackendUKey {

	}
	us.CreateUser(req)
	return
}

// CreateUser
// @description: 创建用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/17 14:26
// @success:
func (us *UserService) CreateUser(req *request.UserCreate) (errCode int) {
	user, errCode3 := userMysql.GetByName(req.Name)
	if errCode3 != 0 && errCode3 != errcode2.GafUserNotFoundErr {
		errCode = errCode3
		return
	}
	if user.ID != 0 {
		functions.AddErrLog(log.Fields{"msg": "创建管理员，该用户已存在", "userName": req.Name})
		errCode = errcode2.GafUserHasExist
		return
	}
	var data entity.User
	data.Name = req.Name
	data.NickName = req.Name
	data.UserSerialNum = req.Serial
	data.Password = gm.EncryptPasswd(data.Name, req.Pin)

	data.RoleId = req.RoleId
	_, errCode = userMysql.Create(nil, &data)
	return
}

// Login
// @description: 管理员登录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:05
// @success:
func (us *UserService) Login(req *request.UserLogin) (user *entity.User, errCode int) {
	user, errCode = userMysql.GetByName(req.Name)
	if errCode != 0 {
		return
	}
	ok := gm.CheckPasswd(req.Name, req.Password, user.Password)
	if !ok {
		errCode = errcode2.GafUserPasswdErr
		return
	}
	//删除之前的toekn，断点登录
	if user.Token != "" {
		gCache.Remove(user.Token)
	}
	var info entity.UserTokenInfo
	info.Id = user.ID
	info.Name = user.Name
	info.RoleId = user.RoleId

	token, errCode1 := TokenService{}.GenerateToken(&info)
	if errCode1 != 0 {
		errCode = errCode1
		return
	}
	user.Token = token
	//更新数据库token
	errCode = userMysql.UpdateToken(user.ID, user.Token)
	//u, _ := TokenService{}.GetInfo(user.Token)
	//fmt.Println("u===", u)
	return
}

// List
// @description: 用户列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:25
// @success:
func (us *UserService) List(req *request.UserList) (list interface{}, total int64, errCode int) {
	list, total, errCode = userMysql.List(req)
	return
}

// InfoByName
// @description: 操作员查询管理员列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 20:19
// @success:
func (us *UserService) InfoByName(name string) (list interface{}, total int64, errCode int) {
	user, errCode1 := userMysql.GetByName(name)
	if errCode1 != 0 {
		errCode = errCode1
		return
	}
	total = 1
	var list1 []*entity.User
	list = append(list1, user)
	return
}

// UKeyLogin
// @description: ueky登录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 15:40
// @success:
func (us *UserService) UKeyLogin(req *request.UKeyLogin) (user *entity.User, errCode int) {
	user, errCode = userMysql.GetByNameAndSerialNum(req.Name, req.Serial)
	if errCode != 0 || user.ID == 0 {
		return
	}
	//验证服务器ukey
	//CheckUserInfoIntegrity(user)
	//if !strings.Contains(user.LoginType, "3") {
	//	err = errors.New(global.ERR_NOT_EXIST)
	//	return
	//}

	//删除之前的toekn，断点登录
	if user.Token != "" {
		gCache.Remove(user.Token)
	}
	var info entity.UserTokenInfo
	info.Id = user.ID
	info.Name = user.Name
	info.RoleId = user.RoleId

	token, errCode2 := TokenService{}.GenerateToken(&info)
	if errCode2 != 0 {
		errCode = errCode2
		return
	}
	user.Token = token
	//更新数据库token
	errCode = userMysql.UpdateToken(user.ID, user.Token)
	return
}

// DeleteById
// @description: 删除指定id
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 11:26
// @success:
func (us *UserService) DeleteById(userid int) (errCode int) {
	errCode = userMysql.DeleteById(nil, userid)
	return
}

// DeleteUser
// @description: 删除管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/17 16:18
// @success:
func (us *UserService) DeleteUser(req *request.UserDelete) (errCode int) {
	var cs ConfigService
	loginType, errCode1 := cs.GetValueStr(dict.ConfigLoginType)
	if errCode1 != 0 {
		return errCode1
	}
	if utils.Int(loginType) == dict.LoginTypeBackendUKey {
		//后端key登录
	}
	errCode = userMysql.DeleteById(nil, req.ID)
	return
}
