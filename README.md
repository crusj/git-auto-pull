# git-auto-pull
trigger git pull when push

## 安装

`git clone github/crusj/git-auto-pull`

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
```
cd main/
nohup go run main/main.go&
```
## 请求
`http -p Hh POST ip:8080/pull project_name==circle`

## 日志

 在`conf/log.json`中可配置filename,日志所在目录必须存在否则无法创建日志文件
 
 日志文件有三中
 * app.log 记录正确pull的日志对应的项目名和路径
 * error.log记录pull失败的日志
 * unknown.log 项目名或项目路径不存在



