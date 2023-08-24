package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webgo/framework"
	"webgo/framework/middleware"
)

const shortDuration = 1 * time.Millisecond

func main() {
	core := framework.NewCore()
	// 注册core 上的中间件
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
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
	// timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
