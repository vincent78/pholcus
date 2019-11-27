package clib

import (
	"bytes"
	// 基础包
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	"net/url"

	"github.com/henrylee2cn/pholcus/common/goquery" //DOM解析
	// "github.com/henrylee2cn/pholcus/logs"           //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
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
	Name:         "git.80021.xyz:39980",
	Description:  "公司内部Gitlab-游戏组相关的所有库信息",
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			req := &request.Request{
				Url:          "http://git.80021.xyz:39980",
				Rule:         "用户登陆页",
				Method:       "GET",
				EnableCookie: true,
			}
			ctx.AddQueue(req)
		},
		Trunk: map[string]*Rule{
			"用户登陆页": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					return nil
				},
				ParseFunc: func(ctx *Context) {
					println("ddd %v", ctx.Response.Header.Get("Set-Cookie"))
					var csrfParam = ""
					var csrfToken = ""
					ctx.GetDom().Find("head").Find("meta").Each(func(i int, s *goquery.Selection) {
						if s.AttrOr("name", "") == "csrf-param" {
							csrfParam = s.AttrOr("content", "")
						} else if s.AttrOr("name", "") == "csrf-token" {
							csrfToken = s.AttrOr("content", "")
						}
					})
					var tmp = bytes.Buffer{}
					tmp.WriteString(csrfParam)
					tmp.WriteString("=")
					tmp.WriteString(csrfToken)
					tmp.WriteString("&")
					tmp.WriteString("user[login]=spiderTest&user[password]=spider1234&user[remember_me]=1")
					postData := url.QueryEscape(tmp.String())
					postData = "utf8=%E2%9C%93&" + postData
					req := &request.Request{
						Url:          "http://git.80021.xyz:39980/users/sign_in",
						Rule:         "用户登陆",
						Method:       "POST",
						EnableCookie: true,
						PostData:     postData,
					}
					ctx.AddQueue(req)
				},
			},
			"用户登陆": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					return nil
				},
				ParseFunc: func(ctx *Context) {

					ctx.FileOutput("gitlab.login.result.html")
				},
			},
		},
	},
}
