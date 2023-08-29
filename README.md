# 官方网站

http://hade.funaio.cn/

# 框架特色：

## 基于协议

服务与服务间的协议是基于协议进行交互的。

## 前后端协同

前后端协同开发

## 命令行

有充分的命令行工具

## 集成定时服务

如果你需要启动定时服务，提供命令进行定时服务的启动

## 文档丰富

提供丰富的文档说明，提供丰富的文档说明

## 开发模式

在开发模式下进行前后端开发，极大提高了开发效率和开发体验

# 功能测试

## lesson-16: 配置服务
运行单元测试: 
* framework/provider/config/provider_test.go
* framework/provider/config/service_test.go
```shell
go build .
./webgo app start
curl http://localhost:8888/demo/demo
# 手动修改 database.mysql.password 的值观察是否热修改
./webgo  config get "database.mysql"
```

## lesson-17: 日志服务
运行单元测试:
* framework/provider/id/provier_test.go
```shell
go build .
./webgo app start
curl http://localhost:8888/demo/demo
# 检查日志文件
cat storage/log/coredemo.log.202*-*-*-*-*
```


## lesson-18: 集成前端 Vue
集成 vue 开发前端页面,实现前后端一体化开发功能,支持中小型业务需求,加快开发效率

```shell
cd webgo-vue
# 16及以上版本
node -v 
npm create vue@latest
npx: installed 1 in 1.204s

Vue.js - The Progressive JavaScript Framework
# 这里用了 ts 可以重新修改 vue 配置再覆盖前端部分 code
✔ Project name: … webgo-vue
✔ Add TypeScript? … No / Yes
✔ Add JSX Support? … No / Yes
✔ Add Vue Router for Single Page Application development? … No / Yes
✔ Add Pinia for state management? … No / Yes
✔ Add Vitest for Unit Testing? … No / Yes
✔ Add an End-to-End Testing Solution? › Nightwatch
✔ Add ESLint for code quality? … No / Yes
✔ Add Prettier for code formatting? … No / Yes

  cd webgo-vue

# copy 所有文件 除了 node_modules 到 webgo 目录下
cd webgo 
npm install
npm run format
npm run dev
go build .
# 同时编译前后端
./webgo build all
```

## lesson-19:20: 实时 dev调试模式 + 反向代理
在 devConfig中设置调试监听端口,注意问题
* 在webgo dev all 命令中启动前端时,要注意互相端口,host 之间能否访问成功
* 只配置了本地测试功能
```shell
go build .
# 如果代码里找不到 webgo 命令,要把 webgo 命令配进 path
export PATH="$PATH:$CODE_PATH/webgo"
./webgo dev all
# 改变 前后端code,检查是否实时热更新请求
curl  http://127.0.0.1:8070/
curl  http://127.0.0.1:8070/demo/demo
```

## lesson-21: 实现自动化命令
开发实现自动化命令,避免重复劳动
* 自动生成模版服务: provider
* 自动生成业务模版命令: command
* 自动生成或迁移gin的中间件: middleware
```shell
go build .
# 文件夹已生成
mkdir app/http/middleware
mkdir app/console/command
mkdir app/provider
# 执行命令测试相应的模版代码是否生成
./webgo middleware migrate
 # git@github.com:gin-contrib/cors.git
 > cors
./webgo command new  
./webgo provider new
```

## lesson-22: 实现自动生成项目脚手架

```shell
go build .
# 文件夹已生成
go build .
➜  webgo git:(main) ✗ ./webgo new                                                                                                                 
? 请输入目录名称： hade_web
? 请输入模块名称(go.mod中的module, 默认为文件夹名称)： github.com/hade_web
? 请输入版本名称(参考 https://github.com/gohade/hade/releases，默认为最新版本)： 
====================================================
开始进行创建应用操作
创建目录： /Users/***/go/src/webgo/hade_web
应用名称： github.com/hade_web
hade框架版本： v1.0.9
创建临时目录 /Users/***/go/src/webgo/template-hade-v1.0.9-1693286683
下载zip包到template.zip
解压zip包
删除临时文件夹 /Users/***/go/src/webgo/template-hade-v1.0.9-1693286683
删除.git目录
删除framework目录
更新文件:/Users/***/go/src/webgo/hade_web/app/grpc/kernel.go
更新文件:/Users/***/go/src/webgo/hade_web/app/grpc/service/helloworld/service.go
更新文件:/Users/***/go/src/webgo/hade_web/app/http/module/demo/api.go
更新文件:/Users/***/go/src/webgo/hade_web/app/http/module/demo/mapper.go
更新文件:/Users/***/go/src/webgo/hade_web/app/http/route.go
更新文件:/Users/***/go/src/webgo/hade_web/app/http/swagger.go
更新文件:/Users/***/go/src/webgo/hade_web/docs/guide/app.md
更新文件:/Users/***/go/src/webgo/hade_web/docs/guide/grpc.md
更新文件:/Users/***/go/src/webgo/hade_web/go.mod
更新文件:/Users/***/go/src/webgo/hade_web/main.go
创建应用结束
目录： /Users/***/go/src/webgo/hade_web

```

## lesson-23: 自动生成 swagger 文件和swagger-ui 服务
通过注释生成 swagger.json 文件
* 如果无法通过err := gen.New().Build(conf)生成对应的 swagger 文件,注意检查安装的包依赖
* 迁移gin-swagger时注意 gin 已经是本地修改过的 gin
> 需要设置模块swaggerFiles: 为本地的模块 github.com/dll02/webgo/framework/middleware/gin-swagger/swaggerFiles
```shell
# 生成 swagger
 go build .
./webgo swagger gen
# 运行 swagger
./webgo app start
curl http://localhost:8888/swagger/index.html
```
