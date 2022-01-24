package boot

import (
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const VERSION = 100
const URL = "https://go.version.fainc.cn/check"
const APPID = "test"
const PASSWORD = "test"

func Update() error {
	//todo 集中update热更新控制
	passwordStr, _ := gmd5.Encrypt(PASSWORD)
	res := g.Client().ContentJson().PostContent(URL, g.Map{
		"VERSION":  VERSION,
		"APPID":    APPID,
		"PASSWORD": passwordStr,
	})
	if res != "" {
		return errors.New("请求更新服务器失败")
	}
	g.Dump(res)
	err := ghttp.RestartAllServer("./bin/main")
	if err != nil {
		return err
	}
	return nil
}
