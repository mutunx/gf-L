package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct {
}

// 登录 登录成功后调转之前地址(如果存在
func (userApi) DoLogin(r *ghttp.Request) {
	var loginReq *model.UserApiLoginReq
	if err := r.Parse(&loginReq); err != nil {
		library.JsonExit(r, 1, err.Error(), nil)
	}
	if err := service.User.Login(r.Context(), loginReq); err != nil {
		library.JsonExit(r, 1, err.Error(), nil)
	}
	referer := service.Session.GetLoginReferer(r.Context())
	if referer != "" {
		service.Session.RemoveLoginReferer(r.Context())
	}
	library.JsonRedirectExit(r, 0, "", referer, nil)

}

func (userApi) Login(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "主页",
	})
}

// 登出 跳回登录页面
func (userApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		library.JsonExit(r, 500, err.Error(), nil)
	} else {
		r.Response.RedirectTo(service.Middleware.GetLoginUrl())
	}
}
