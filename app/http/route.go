package http

import (
	"github.com/dll02/webgo/app/http/module/demo"
	"github.com/dll02/webgo/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
