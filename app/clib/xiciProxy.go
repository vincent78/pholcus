package clib

/**
西刺-免费代理
验证网址： https://staticscdn.oss-cn-shanghai.aliyuncs.com/ok.html

*/
import (
	// 基础包
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
)

func init() {
	XiciProxy.Register()
}

var XiciProxy = &Spider{
	Name:         "免费代理网站-西刺",
	Description:  "西刺免费代理网站",
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			req := &request.Request{
				Url:          "https://www.xicidaili.com/nn/",
				Rule:         "列表首页",
				Method:       "GET",
				EnableCookie: true,
			}
			ctx.AddQueue(req)
		},
		Trunk: map[string]*Rule{
			"列表首页": {
				ParseFunc: func(ctx *Context) {
					ctx.FileOutput("xici.nn.html")
				},
			},
		},
	},
}
