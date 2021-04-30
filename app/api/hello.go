package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct {
}

// Index godoc
// @Summary 登录成功后跳转主页页面
// @Description
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /login [post]
func (a *helloApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:   "欢迎你",
		MainTpl: "hello.html",
	})
}
