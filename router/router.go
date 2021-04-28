package router

import (
	"gf-L/app/api"
	"gf-L/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/ping", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Hello)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Ctx)
		group.ALL("/", api.User)
	})
}
