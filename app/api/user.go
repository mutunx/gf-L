package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"time"
)

var User = userApi{}

type userApi struct {
}

// Register godoc
// @Summary 注册
// @Description
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /Register [post]
func (a userApi) Register(r *ghttp.Request) {
	// 重复登录校验
	token := r.Get("token")
	if b, err := gcache.Contains(token); err != nil {
		library.JsonExit(r, 1, err.Error(), nil)
	} else {
		if b {
			library.JsonExit(r, 1, gerror.New("重复提交").Error(), nil)
		} else {
			gcache.Set(token, true, 30*time.Minute)
		}
	}

	var user *model.User
	if err := r.Parse(&user); err != nil {
		library.JsonExit(r, 1, err.Error(), nil)
	}
	if id, err := service.User.Register(r.Context(), user); err != nil {
		library.JsonExit(r, 500, err.Error(), nil)
	} else {
		g.Log().Print("id:", id)
		if err = a.uploadAvatar(r, id); err != nil {
			library.JsonExit(r, 1, err.Error(), nil)
		} else {
			library.JsonExit(r, 0, "注册成功")
		}
	}

}

// 上传头像
func (userApi) uploadAvatar(r *ghttp.Request, id string) error {

	files := r.GetUploadFiles("upload-file")
	return service.Avatar.Upload(r.Context(), files, id)

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
