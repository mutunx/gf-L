package service

import (
	"context"
	"gf-L/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文管理服务
var Context = contextStruct{}

type contextStruct struct {
}

// 初始化上下文 将请求中的上下文指针指向上下文对象,以便后期修改
func (c contextStruct) Init(r *ghttp.Request, ctx *model.Context) {
	r.SetCtxVar(model.ContextKey, ctx)
}

// 获取上下文对象 没有则返回nil
func (c contextStruct) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	// 类型转换
	if localContext, ok := value.(*model.Context); ok {
		return localContext
	}
	return nil
}

// 设置上下文的用户
func (c contextStruct) SetUser(ctx context.Context, user *model.ContextUser) {
	c.Get(ctx).User = user
}

// 设置上下文的数据
func (c contextStruct) SetData(ctx context.Context, data g.Map) {
	c.Get(ctx).Data = data
}
