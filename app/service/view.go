package service

import (
	"gf-L/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 视图管理
var View = viewService{}

type viewService struct {
}

// 将数据对象填写到模板,将模板作为response返回
func (viewService) RenderTpl(r *ghttp.Request, tpl string, data ...model.View) {
	// 生成数据对象
	var viewData = make(g.Map)
	var viewObj = model.View{}
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.MainTpl == "" {
		viewObj.MainTpl = "/login.html"
	}
	viewData = gconv.Map(viewObj)
	// 渲染模板
	r.Response.WriteTpl(tpl, viewData)
	r.Exit()
}

// 转发到默认登录模板
func (v viewService) Render(r *ghttp.Request, data ...model.View) {
	v.RenderTpl(r, g.Cfg().GetString("viewer.DefaultFile"), data...)
}

func (v viewService) Render404(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "资源不存在"
	}
	view.MainTpl = "/pages/404.html"
	v.Render(r, view)
}

func (v viewService) Render403(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = "/pages/403.html"
	v.Render(r, view)
}

func (v viewService) Render401(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "无访问权限"
	}
	view.MainTpl = "/pages/401.html"
	v.Render(r, view)
}

func (v viewService) Render500(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	if view.Title == "" {
		view.Title = "请求执行错误"
	}
	view.MainTpl = "/pages/500.html"
	v.Render(r, view)
}
