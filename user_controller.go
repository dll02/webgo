package main

import (
	"fmt"
	"time"

	"github.com/dll02/goweb/framework/gin"
)

func UserLoginController(c *gin.Context)  {
	foo, _ := c.DefaultQueryString("foo", "def")
	fmt.Println("start wait")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	// 打印控制器名字
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)
}
