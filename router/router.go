package router

import (
	"gf-start-up/app/api"
	"gf-start-up/middleware"
	"github.com/fainc/gowork/library/cors"
	"github.com/fainc/gowork/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindStatusHandler(404, response.Json.NotFound)
	s.BindStatusHandler(500, response.Json.ServerError)
	s.Use(cors.Standard)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Jwt.JwtAuth)
		group.ALL("/hello", api.Hello)
	})
}
