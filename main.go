package main

import (
	"time"
	"webgo/framework"
	"webgo/framework/middleware"
	"net/http"
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
	server.ListenAndServe()
}
