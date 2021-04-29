package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct {
}

// Index is a demonstration route handler for output "Hello World!".
func (a *helloApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:   "欢迎你",
		MainTpl: "hello.html",
	})
}
