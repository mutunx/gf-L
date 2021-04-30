package main

import (
	_ "gf-L/boot"
	_ "gf-L/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func main() {
	server := g.Server()
	server.Plugin(&swagger.Swagger{})
	server.Run()
}
