// Path: pkg/errcode
// FileName: user.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/3/24$ 13:15$

package errcode

// 用户相关的错误
// user
const (
	GafUserErr          = GafServer + GafUserCode + iota + 1 //用户错误
	GafUserLoginErr                                          //登录错误
	GafUserNotFoundErr                                       //用户不存在
	GafUserPasswdErr                                         //密码错误
	GafUserWithoutToken                                      //未携带token
	GafUserTokenExpired                                      //token过期
	GafUserRoleErr                                           // 用户的role_id错误
	GafUserForbiddenErr                                      // 用户没有权限
	GafUserHasExist                                          //用户已存在
)
