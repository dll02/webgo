package http

import (
	"github.com/dll02/webgo/app/http/module/demo"
	"github.com/dll02/webgo/framework/gin"
	"github.com/dll02/webgo/framework/middleware/static"
)

func Routes(r *gin.Engine) {

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	demo.Register(r)
}
