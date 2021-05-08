package service

import (
	"context"
	"gf-L/app/model"
	"github.com/gogf/gf/util/gconv"
)

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "sessionKeyUser"

	// Referer存储，当已存在该session时不会更新。
	// 用于用户未登录时引导用户登录，并在登录后跳转到登录前页面。
	sessionKeyLoginReferer = "sessionKeyLoginReferer"

	// 存放在Session中的提示信息，往往使用后则删除
	sessionKeyNotice = "sessionKeyNotice"
)

// session管理服务
var Session = sessionStruct{}

type sessionStruct struct {
}

// session用户设置
func (s sessionStruct) SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取用户
func (s sessionStruct) GetUser(ctx context.Context) *model.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		v := customCtx.Session.GetVar(sessionKeyUser)
		if !v.IsNil() {
			var user *model.User
			_ = gconv.Struct(v.Val(), &user)
			return user
		}
	}
	return nil
}

// 删除用户
func (sessionStruct) RemoveUser(ctx context.Context) error {
	return Context.Get(ctx).Session.Remove(sessionKeyUser)
}

// 设置LoginReferer.
func (s sessionStruct) SetLoginReferer(ctx context.Context, referer string) error {
	if s.GetLoginReferer(ctx) == "" {
		customCtx := Context.Get(ctx)
		if customCtx != nil {
			return customCtx.Session.Set(sessionKeyLoginReferer, referer)
		}
	}
	return nil
}

// 获取LoginReferer.
func (s sessionStruct) GetLoginReferer(ctx context.Context) string {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.GetString(sessionKeyLoginReferer)
	}
	return ""
}

// 删除LoginReferer.
func (s sessionStruct) RemoveLoginReferer(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyLoginReferer)
	}
	return nil
}

// 设置Notice
func (s sessionStruct) SetNotice(ctx context.Context, message *model.SessionNotice) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Set(sessionKeyNotice, message)
	}
	return nil
}

// 获取Notice
func (s sessionStruct) GetNotice(ctx context.Context) (*model.SessionNotice, error) {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		var message *model.SessionNotice
		err := customCtx.Session.GetVar(sessionKeyNotice).Struct(&message)
		return message, err
	}
	return nil, nil
}

// 删除Notice
func (s sessionStruct) RemoveNotice(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyNotice)
	}
	return nil
}
