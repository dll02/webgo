package http

import (
	"github.com/dll02/goweb/app/http/module/demo"
	"github.com/dll02/goweb/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
