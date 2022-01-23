package api

import (
	"gf-start-up/middleware"
	"github.com/fainc/gowork/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	response.Json.Success(r, middleware.Jwt.GetUser(r))
}
