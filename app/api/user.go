package api

import (
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct {
}

func (a userApi) GetUser(r *ghttp.Request) {

	if data, err := service.User.GetUser(r.Context()); err != nil {
		library.JsonExit(r, 500, err.Error())
	} else {
		library.JsonExit(r, 0, "", data)
	}
}
func (a userApi) DelUser(r *ghttp.Request) {
	ids := r.GetArray("ids")
	if len(ids) == 0 {
		library.JsonExit(r, 0, "删除成功")
	}
	if err := service.User.DeleteUsers(r.Context(), ids); err != nil {
		library.JsonExit(r, 500, err.Error())
	} else {
		library.JsonExit(r, 0, "删除成功")
	}
}
