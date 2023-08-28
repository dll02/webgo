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