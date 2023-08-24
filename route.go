package main

import (
	"webgo/framework"
	"webgo/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("/foo", FooControllerHandler)
	// 需求1+2:HTTP方法+静态路由匹配
	// 在核心业务逻辑 UserLoginController 之外，封装一层 TimeoutHandler
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test2())
		subjectApi.Post("/add", SubjectAddController)
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test1(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
		// 多层嵌套通用前缀
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectInfoNameController)
		}
	}
}
