package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	appHttp "github.com/dll02/goweb/app/http"
	"github.com/dll02/goweb/framework/gin"
	"github.com/dll02/goweb/framework/middleware"
	"github.com/dll02/goweb/framework/provider/demo"
	"github.com/dll02/goweb/framework/provider/app"
)

const shortDuration = 1 * time.Millisecond

func main() {

	core := gin.New()

		// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})
	core.Bind(&app.WebgoAppProvider{})

	// 配中间处理器
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	// 绑定路由处理器
	appHttp.Routes(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 启动服务的 Goroutine
	go func() {
		server.ListenAndServe()
	}()
	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
