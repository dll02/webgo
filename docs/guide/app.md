---
lang: zh-CN
title: app 运行命令
description: app 运行命令
---
# 运行

## 命令

这里的运行是运行整个 app，这个 app 可以只包含后端，也可以只包含前端，但是后端也是隐藏在前端后面运行。具体可以参考 app/http/route.go

```
package http

import (
	"github.com/dll02/webgo/app/http/controller/demo"
	"github.com/dll02/webgo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	r.GET("/demo/demo", demo.Demo)
}

```

运行相关的命令为 app。

```
[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app
start app serve

Usage:
  webgo app [flags]
  webgo app [command]

Available Commands:
  restart     restart app server
  start       start app server
  state       get app pid
  stop        stop app server

Flags:
  -h, --help   help for app

Use "webgo app [command] --help" for more information about a command.
```

## 启动

可以使用 `./webgo app start` 启动一个应用。

也可以使用 `./webgo app start -d` 使用 deamon 模式启动一个应用。应用名称为 `webgo app`

```
[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app start -d
app serve started
log file: /Users/Documents/workspace/webgo_workspace/demo5/storage/log/app.log
```

app 应用的输出记录在 `/storage/log/app.log`

进程 id 记录在 `/storage/pid/app.pid`

## 状态

当使用 deamon 模式启动的时候，需要查看当前应用是否有启动，如果启动了，进程号是多少，可以使用命令 `./webgo app state`

```
[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app state
app server started, pid: 28170
```

## 重启

当使用 deamon 模式启动的时候，需要重启应用，可以使用命令 `./webgo app restart`

::: tip
如果程序还未启动，调用 restart 命令，效果和 start 命令一样，deamon 模式启动应用
:::

## 停止

当使用 deamon 模式启动的时候，需要关闭应用，可以使用命令 `./webgo app stop`
