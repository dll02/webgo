package http

import (
	"github.com/dll02/webgo/app/http/middleware/auth"
	"github.com/dll02/webgo/app/http/middleware/cors"
	"github.com/dll02/webgo/app/http/module/qa"
	"github.com/dll02/webgo/app/http/module/user"

	"github.com/dll02/webgo/framework/contract"
	"github.com/dll02/webgo/framework/gin"
	ginSwagger "github.com/dll02/webgo/framework/middleware/gin-swagger"
	"github.com/dll02/webgo/framework/middleware/gin-swagger/swaggerFiles"
	"github.com/dll02/webgo/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	// 使用cors中间件
	r.Use(cors.Default())
	// 增加 autho 中间件
	r.Use(auth.AuthMiddleware())
	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 用户模块
	user.RegisterRoutes(r)
	// 问答模块
	qa.RegisterRoutes(r)
}
