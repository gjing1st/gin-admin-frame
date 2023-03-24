// Path: internal/pkg/constant
// FileName: response_message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/28$ 18:21$

package global

const (
	OperateSuccess    = "操作成功"
	CreatedSuccess    = "创建成功"
	SaveSuccess       = "保存成功"
	UploadSuccess     = "添加附件成功"
	QuerySuccess      = "查询成功"
	DeleteSuccess     = "删除成功"
	LoginSuccess      = "登录成功"
	DeleteUserSuccess = "删除管理员成功"
	CreateKeySuccess  = "生成密钥成功"
	SetSuccess        = "设置成功"
)

const (
	OperateFailed    = "操作失败"
	CreatedFailed    = "创建失败"
	UploadFailed     = "上传文件失败，请重试"
	SaveFailed       = "保存失败，请重试"
	DeleteFailed     = "删除失败，请重试"
	QueryFailed      = "查询失败"
	LoginTypeErr     = "当前登录方式不允许"
	DataNotFound     = "数据不存在"
	LoginFail        = "登录失败，请检查用户名和密码是否输入正确"
	UserNotFound     = "用户不存在"
	UserHasAddFailed = "新增失败"
	UKeyHasAddAdmin  = "当前KEY已注册，新增失败"
	LoginTokenErr    = "登录超时，请重新登录"
	LoginTokenNull   = "请重新登录"
	Unauthorized     = "未认证"
	AuthForbidden    = "无权访问"
	InitDown         = "初始化已完成"
	DeleteUserFail   = "删除管理员失败，请重试"
	CreateKeyFail    = "生成密钥失败，请重试"
	SetFail          = "设置失败，请重试"
)
