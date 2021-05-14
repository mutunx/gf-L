package router

import (
	"fmt"
	"gf-L/app/api"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/hello", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.SessionManage)
		group.Middleware(service.Middleware.Ctx)
		group.POST("/login", api.Login.Do)
		group.GET("/login", api.Login.Index)
		group.GET("/logout", api.Login.Logout)
		group.POST("/register", api.Login.Register)
		group.ALL("/idCard", api.IdCard)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/", api.User)
		})
	})

	s.Group("/session", func(group *ghttp.RouterGroup) {
		group.ALL("/set", func(r *ghttp.Request) {
			r.Session.Set("hello", "you")
		})
		group.ALL("/get", func(r *ghttp.Request) {
			get := r.Session.Get("hello")
			library.JsonExit(r, 1, fmt.Sprint(get))
		})
	})
}
