package main

import (
	"webgo/framework"
)


func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 打印控制器名字
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
