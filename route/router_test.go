package route

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	fmt.Println("进入测试....")

	router := &Routes{}
	router.Middleware([]MiddleWare{testMiddleWareOne, testMiddleWareTwo}).Group(func() []*RS {
		return []*RS{
			&RS{"GET", "/get", testHandleGet},
			&RS{"GET", "/", testHandleGet},
		}
	})
	go func() {
		fmt.Println("启动服务器....")
		err := fasthttp.ListenAndServe(":8080", Router.Handler)
		if err != nil {
			fmt.Println("启动服务器错误:", err)
		}

	}()
	writer, err := http.Get("http://127.0.0.1:8080/get")
	if err != nil {
		t.Errorf("request error: %v\n", err)
	}
	if writer.StatusCode != 200 {
		t.Errorf("Response code is %v\n", writer.StatusCode)
	}
	rsl, err := ioutil.ReadAll(writer.Body)
	if err != nil {
		t.Errorf("read body error: %v\n", err)
	}
	if string(rsl) != "onetwoHello world" {
		t.Errorf("return content error,return content is `%v`\n", string(rsl))
	}

}
func testHandleGet(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("Hello world"))
}
func testMiddleWareOne(next HandleFunc) HandleFunc {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Write([]byte("one"))
		next(ctx)
	}
}
func testMiddleWareTwo(next HandleFunc) HandleFunc {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Write([]byte("two"))
		next(ctx)
	}
}
