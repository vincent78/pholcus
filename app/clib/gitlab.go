package clib

import (
	// 基础包
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	//"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	// "github.com/henrylee2cn/pholcus/logs"           //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用

	// net包
	// "net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	//"strconv"
	// "strings"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	Gitlab.Register()
}

var Gitlab = &Spider{
	Name:        "公司内部Gitlab",
	Description: "公司内部Gitlab-游戏组相关的所有库信息",
	EnableCookie: true,
	RuleTree:&RuleTree{
		Root:  func(ctx *Context) {
			req := &request.Request{
				Url: "https://git.80021.xyz:39980/users/sign_in",
				Rule: "用户登陆",
				Method: "POST",
				PostData: "user[login]=spiderTest&user[password]=spider1234&user[remember_me]=1",
			}

			ctx.AddQueue(req)
		},
		Trunk: map[string]*Rule{
			"用户登陆": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					return nil
				},
				ParseFunc: func(ctx *Context) {
					ctx.FileOutput("gitlab.test.html")
				},
			},
		},
	},
}