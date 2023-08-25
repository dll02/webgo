package main

import (
	"github.com/dll02/goweb/framework/gin"
	"github.com/dll02/goweb/provider/demo"
)

func SubjectInfoNameController(c *gin.Context) {
	// 打印控制器名字
	c.ISetOkStatus().IJson("ok, SubjectInfoNameController")
}

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectGetController")
}

func SubjectListController(c *gin.Context) {
	// 获取 demo 服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)
	// 调用服务实例的方法
	foo := demoService.GetFoo()
	// 输出结果
	c.ISetOkStatus().IJson(foo)

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
