package init

import (
	"fmt"
	"github.com/crusj/git-auto-pull/global"
	. "github.com/crusj/git-auto-pull/global"
	middleware "github.com/crusj/git-auto-pull/middleware"
	. "github.com/crusj/git-auto-pull/route"
	. "github.com/crusj/git-auto-pull/service"
	"github.com/crusj/logger"
	"github.com/valyala/fasthttp"
	"gopkg.in/ini.v1"
)

func init() {
	initConf()
	initRoutesAndServe()
	initLog()
}
func initConf() {
	cfg, err := ini.Load("../conf/app.ini")
	if err != nil {
		logger.Channel("alert").Alert("加载配置文件失败：", err)
	} else {
		global.Cfg = cfg
	}
}

func initLog() {
	err := logger.SetLogger("../conf/log.json")
	if err != nil {
		logger.Channel("alert").Alert("日志配置文件加载失败,", err)
	} else {
		logger.Info("日志配置文件加载成功")
	}
}
func initRoutesAndServe() {
	Route.Middleware([]MiddleWare{middleware.Json}).Group(func() []*RS {
		//定义路由
		return []*RS{
			&RS{"POST", "/pull", Pull},
		}
	})
	logger.Info("starting http server")
	go func() {
		err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", Cfg.Section("listen").Key("http").String()), Router.Handler)
		if err != nil {
			logger.Channel("alert").Alert("HTTP服务器启动错误：", err)
		}
	}()
}
