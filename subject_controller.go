package main

import (
	"webgo/framework"
)

func SubjectInfoNameController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, SubjectInfoNameController")
	return nil
}

func SubjectAddController(c *framework.Context) error {
	c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.Json(200, "ok, SubjectGetController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(200, "ok, SubjectListController")
	return nil
}
