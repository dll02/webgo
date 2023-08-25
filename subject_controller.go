package main

import (
	"github.com/dll02/goweb/framework/gin"
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
	c.ISetOkStatus().IJson("ok, SubjectListController")
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
