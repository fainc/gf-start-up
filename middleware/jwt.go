package middleware

import (
	"github.com/fainc/gowork/library/jwt"
	"github.com/fainc/gowork/library/response"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Jwt = jwtStruct{}

type jwtStruct struct{}

func (j *jwtStruct) JwtAuth(r *ghttp.Request) {
	scopes := g.SliceStr{
		"user",
	}
	whiteTables := g.SliceStr{
		"/hello/index",
	}
	Jwt.standardAuth(r, scopes, whiteTables, true)
	r.Middleware.Next()
}
func (*jwtStruct) standardAuth(r *ghttp.Request, scopes g.SliceStr, tables g.SliceStr, catch bool) {
	whiteTable := garray.NewStrArrayFrom(tables)
	uuid, scopeKey, err := jwt.Helper.Parse(jwt.ParseParams{
		Token:  r.GetHeader("Authorization"),
		Scopes: scopes,
		Secret: g.Cfg().GetString("jwt.secret"),
	})
	if err != nil {
		if !whiteTable.ContainsI(r.RequestURI) && catch {
			response.Json.Authorization(r, err.Error())
		} else {
			r.SetCtxVar("UUID", 0)
			r.SetCtxVar("SCOPE", "UNKNOWN")
		}
	} else {
		r.SetCtxVar("UUID", uuid)
		r.SetCtxVar("SCOPE", scopeKey)
	}
}

type user struct {
	UUID  int64
	SCOPE string
}

func (*jwtStruct) GetUser(r *ghttp.Request) *user {
	UUID := r.GetCtxVar("UUID", 0)
	SCOPE := r.GetCtxVar("SCOPE", "UNKNOWN")
	return &user{
		UUID:  gconv.Int64(UUID),
		SCOPE: gconv.String(SCOPE),
	}
}
