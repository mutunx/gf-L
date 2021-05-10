package api

import (
	"fmt"
	"gf-L/app/model"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"time"
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
// @Router /login [post]
func (a loginApi) Do(r *ghttp.Request) {
	token := r.GetString("token")
	if err := a.validToken(token, "login"); err != nil {
		library.JsonExit(r, 1, err.Error())
	}
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
// @Router /login [get]
func (loginApi) Index(r *ghttp.Request) {
	// 生成唯一token 作为防止重复提交token 加入缓存
	crutime := fmt.Sprint(time.Now().Unix())
	token := gmd5.MustEncrypt(r.Session.Id() + crutime)
	// 由于主页包含注册和登录俩个功能,所以在一个token下缓存俩个变种用于验证一个页面下的登录和注册
	gcache.Set(gmd5.MustEncrypt("login"+token), true, 5*time.Minute)
	gcache.Set(gmd5.MustEncrypt("register"+token), true, 5*time.Minute)

	service.View.Render(r, model.View{
		Title: "主页",
		Token: token,
	})
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
func (loginApi) Logout(r *ghttp.Request) {
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
func (a loginApi) Register(r *ghttp.Request) {
	// 重复校验
	token := r.GetString("token")
	if err := a.validToken(token, "register"); err != nil {
		library.JsonExit(r, 1, err.Error())
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

func (a loginApi) validToken(tokens ...string) error {
	if len(tokens) == 0 {
		return gerror.New("没有参数")
	}
	token := tokens[0]
	var keyword string
	if len(tokens) > 1 {
		keyword = tokens[1]
	}
	realToken := gmd5.MustEncrypt(keyword + token)
	exist, err := gcache.Contains(realToken)
	if err != nil {
		return err
	}
	if exist {
		gcache.Remove(realToken)
	} else {
		err = gerror.New("无效token请刷新重试")
	}
	return err
}

// 上传头像
func (loginApi) uploadAvatar(r *ghttp.Request, id string) error {

	files := r.GetUploadFiles("upload-file")
	if len(files) == 0 {
		return nil
	}
	return service.Avatar.Upload(r.Context(), files, id)

}
