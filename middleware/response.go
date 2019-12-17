package middlware

import (
	"github.com/crusj/git-auto-pull/route"
	"github.com/valyala/fasthttp"
)

func Json(next route.HandleFunc) route.HandleFunc {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Content-Type", "Application/json; charset=utf-8")
		next(ctx)
	}
}
