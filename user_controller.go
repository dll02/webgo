package main

import (
	"fmt"
	"time"
	"webgo/framework"
)


func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	fmt.Println("start wait")
	// 等待10s才结束执行  
	time.Sleep(10 * time.Second)
	// 打印控制器名字
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
