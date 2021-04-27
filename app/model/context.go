package model

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文变量存储名 前后端共享
const ContextKey = "ContextKey"

type Context struct {
	Session *ghttp.Session //session管理对象
	User    *ContextUser   // 上下文用户信息
	Data    g.Map          // 数据信息
}

type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
	Avatar   string // 用户头像
	IsAdmin  bool   // 是否是管理员
}
