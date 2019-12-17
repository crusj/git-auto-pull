package service

import (
	"github.com/crusj/git-auto-pull/global"
	"github.com/crusj/logger"
	"github.com/valyala/fasthttp"
	"os"
	"os/exec"
)

var (
	gitPath = ""
)

func Pull(ctx *fasthttp.RequestCtx) {
	name := ctx.FormValue("project_name")
	if name == nil {
		logger.Channel("unknown").Info("空project_name")
	} else {
		logger.Info("项目名:", string(name))
		if path := global.Cfg.Section("project." + string(name)).Key("path").String(); path == "" {
			logger.Channel("unknown").Info("无效的project_name: %s", name)
		} else {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				logger.Channel("unknown").Info("项目: %s,路径: %s,不存在", name, path)
			} else {
				go func() {
					if gitPath == "" {
						gitPath = global.Cfg.Section("git").Key("path").String()
					}
					remoteRepo := global.Cfg.Section("project." + string(name)).Key("remoteRepo").String()
					if remoteRepo == "" {
						remoteRepo = "origin"
					}
					branch := global.Cfg.Section("project." + string(name)).Key("branch").String()
					if branch == "" {
						branch = "master"
					}
					cmd := exec.Command(gitPath+"/git", "pull", remoteRepo, branch)
					cmd.Dir = path
					out, err := cmd.Output()
					if err != nil {
						logger.Channel("warn").Warn("项目: %s,路径: %s,错误,%s，%s ", name, path, err, out)
					} else {
						logger.Channel("default").Info("项目: %s,路径: %s，pull success,%s", name, path, out)
					}
				}()
			}
		}
	}

}
