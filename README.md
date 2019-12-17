# git-auto-pull
trigger git pull when push

## 配置
conf/app.ini
```ini
[listen]
#http监听端口
http = 8080
[git]
# git所在目录
path = /usr/bin
[project]
#项目
[project.circle]
#项目名
name = circle
#项目路径
path = /tmp/testPull
# 远程仓库
remote_repo = origin
# 分支
branch = master
```
## 执行
`nohup go run main/main.go&`
## 请求
`http -p Hh POST ip:8080/pull project_name==circle`