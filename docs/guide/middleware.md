# 中间件

## 指南
webgo 的 HTTP 路由服务并没有自己开发，而是使用 gin。gin 生态已经有非常完善的(中间件体系)[https://github.com/gin-contrib]。

我们没有必要重新开发这些中间件。所以 webgo 框架提供了 middleware 命令来获取这些中间件。（不能直接使用go get 的方式来获取，因为 webgo 在 gin 基础上做了一些微调）

```
[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo middleware
webgo middleware

Usage:
  webgo middleware [flags]
  webgo middleware [command]

Available Commands:
  add         add middleware to app, https://github.com/gin-contrib/[middleware].git
  list        list all installed middleware
  remove      remove middleware from app

Flags:
  -h, --help   help for middleware

Use "webgo middleware [command] --help" for more information about a command.
```

## 安装

可以安装 https://github.com/gin-contrib/ 项目中的任何中间件，使用命令 `./webgo middleware add gzip`

命令会从 https://github.com/gin-contrib/gzip.git 项目中下载中间件，并且安装到 `app/http/middleware` 中。

## 查询

检查目前已经安装了哪些中间件，可以使用命令 `./webgo middleware list`

## 删除

删除某个中间件，可以使用命令 `./webgo middleware remove`