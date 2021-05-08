package service

import (
	"gf-L/app/model"
	"gf-L/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 中间件服务
var Middleware = middlewareStruct{
	loginUrl: "/login",
}

// 获取用户登录URL
func (s *middlewareStruct) GetLoginUrl() string {
	return s.loginUrl
}

type middlewareStruct struct {
	loginUrl string
}

// 初始化上下文 存在用户则放入上下文
func (middlewareStruct) Ctx(r *ghttp.Request) {
	session := r.Session
	ctx := &model.Context{
		Session: session,
		User:    nil,
		Data:    nil,
	}
	Context.Init(r, ctx)
	if session != nil {
		user := Session.GetUser(r.Context())
		if user != nil {
			var contextUser *model.ContextUser
			gconv.Struct(user, &contextUser)
			Context.SetUser(r.Context(), contextUser)
		}
	}

	r.Middleware.Next()
}

// 权限控制
func (m middlewareStruct) Auth(r *ghttp.Request) {
	user := Session.GetUser(r.Context())
	if user == nil {
		library.JsonRedirectExit(r, 1, "", "/login")
	}
	r.Middleware.Next()
}

func (m middlewareStruct) SessionManage(r *ghttp.Request) {
	g.Log().Println("sessionID:" + r.Session.Id() + "过来看看")
	r.Middleware.Next()
}
