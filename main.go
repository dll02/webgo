package main

import (
	"time"
	"webgo/framework"
	"net/http"
)

const shortDuration = 1 * time.Millisecond

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
