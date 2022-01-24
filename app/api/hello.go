package api

import (
	"time"

	"github.com/fainc/gowork/library/jwt"
	"github.com/fainc/gowork/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	generate, err := jwt.Helper.Generate(jwt.GenerateParams{
		Uuid:     1,
		Scope:    "user",
		Duration: 1 * time.Hour,
		Secret:   "123",
	})
	if err != nil {
		response.Json.Success(r, err.Error())
	}
	response.Json.Success(r, generate)
}
