package router

import (
	"gf-L/app/api"
	"gf-L/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/hello", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Ctx)
		group.POST("/login", api.Login.Do)
		group.GET("/login", api.Login.Index)
		group.GET("/logout", api.Login.Logout)
		group.ALL("/", api.User)
		group.ALL("/idCard", api.IdCard)
	})
}
