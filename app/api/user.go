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

// Logout godoc
// @Summary 登出
// @Description
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /Logout [post]
func (userApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		library.JsonExit(r, 500, err.Error(), nil)
	} else {
		r.Response.RedirectTo(service.Middleware.GetLoginUrl())
	}
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
func (userApi) Register(r *ghttp.Request) {
	var user *model.User
	if err := r.Parse(&user); err != nil {
		library.JsonExit(r, 1, err.Error(), nil)
	}
	if err := service.User.Register(r.Context(), user); err != nil {
		library.JsonExit(r, 500, err.Error(), nil)
	} else {
		library.JsonExit(r, 0, "注册成功", nil)
	}
}

// UploadAvatar godoc
// @Summary 注册
// @Description
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /Upload-avatar [post]
func (userApi) UploadAvatar(r *ghttp.Request) {
	sourcePath := "public/"
	filePath := "image/tmp/"
	files := r.GetUploadFiles("upload-file")
	names, err := files.Save(sourcePath + filePath)
	if err != nil {
		library.JsonExit(r, 1, err.Error())
	}
	library.JsonExit(r, 0, "", filePath+names[0])
}
