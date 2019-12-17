# git-auto-pull
trigger git pull when push

## 配置
conf/app.ini
```ini
[listen]
#监听端口
http = 8080
[project]
#触发的项目
[project.circle]
#项目名称
name = circle
#项目路径
path = /tmp/testPull
```
## 执行
`nohup go run main/main.go&`
## 请求
`http -p Hh POST ip:8080/pull project_name==circle`