package router

import (
	"fmt"

	"gf-start-up/app/api"
	"gf-start-up/boot"
	"gf-start-up/middleware"

	"github.com/fainc/gowork/library/cors"
	"github.com/fainc/gowork/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	err := boot.Update()
	fmt.Println(err.Error())
	s := g.Server()
	s.BindStatusHandler(404, response.Json.NotFound)
	s.BindStatusHandler(500, response.Json.ServerError)
	s.Use(cors.Standard)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Jwt.JwtAuth)
		group.ALL("/hello", api.Hello)
	})
	s.BindHandler("/update-hook", update) // hook被动触发更新
	s.BindHandler("/version", func(r *ghttp.Request) {
		response.Json.Success(r, boot.VERSION)
	})
}

func update(r *ghttp.Request) {
	err := boot.Update()
	if err != nil {
		response.Json.StandardError(r, err.Error())
	}
	response.Json.Success(r, "更新成功")
}
