package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/net/ghttp"
)

var Login = loginApi{}

type loginApi struct {
}

// Do godoc
// @Summary 登录
// @Description 验证账号密码
// @Accept  json
// @Produce  json
// @Param passport formData string true "账号"
// @Param password formData string true "密码"
// @Success 200 object} library.JsonRes{Code:0}
// @Failure 400,404 {object} library.JsonRes{Code:1}
// @Failure 500 {object} library.JsonRes{Code:1}
// @Router /login/do [post]
func (loginApi) Do(r *ghttp.Request) {
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

// Index godoc
// @Summary 跳转登录页面
// @Description 验证账号密码
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /login [post]
func (loginApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "主页",
	})
}
