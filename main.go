package main

import (
	_ "gf-L/boot"
	_ "gf-L/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
