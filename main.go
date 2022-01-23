package main

import (
	_ "gf-start-up/boot"
	_ "gf-start-up/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
