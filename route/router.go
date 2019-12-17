package route

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/crusj/logger"
	"github.com/valyala/fasthttp"
)

//路由结构
type RS struct {
	Method     string
	Path       string
	HandleFunc HandleFunc
}

//处理器
type HandleFunc = func(ctx *fasthttp.RequestCtx)

//中间件
type MiddleWare func(next HandleFunc) HandleFunc

//路由函数
type routesFunc func() []*RS

type Routes struct {
	//中间件
	middles []MiddleWare
}

var (
	Router *fasthttprouter.Router = fasthttprouter.New()
	Route  *Routes                = &Routes{}
)

//中间件
func (r *Routes) Middleware(middles []MiddleWare) *Routes {
	for _, middleware := range middles {
		r.middles = append(r.middles, middleware)
	}
	return r
}

//分组
func (r *Routes) Group(routes routesFunc) {
	routeHandles := routes()
	middlewareCount := len(r.middles)
	for _, v := range routeHandles {
		t := v.HandleFunc
		//串联中间件
		if middlewareCount > 0 {
			for i := middlewareCount - 1; i >= 0; i-- {
				t = r.middles[i](t)
			}
		}
		switch v.Method {
		case "GET":
			logger.Trace("注册路由 GET %s\n", v.Path)
			Router.GET(v.Path, t)
		case "POST":
			Router.POST(v.Path, t)
		case "PUT":
			Router.PUT(v.Path, t)
		case "DELETE":
			Router.DELETE(v.Path, t)
		default:
			panic("")
		}
	}
}
