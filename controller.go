package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"webgo/framework"
)

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.Json(200, "ok, UserLoginController")
	return nil
}

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

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}

// 使用封装后的Context, 使用语义化高的接口函数
// 控制器
func Foo2(ctx *framework.Context) error {
	obj := map[string]interface{}{
		"data": nil,
	}
	// 从请求体中获取参数
	fooInt := ctx.FormInt("foo", 10)
	// 构建返回结构
	obj["data"] = fooInt
	// 输出返回结构
	return ctx.Json(http.StatusOK, obj)
}

// 未封装自定义 Context 的控制器   -> 精准控制 request & response
// 需要代码里自己设置 header 和 Body
func Foo1(request *http.Request, response http.ResponseWriter) {
	obj := map[string]interface{}{
		"data": nil,
	}
	// 设置控制器 response 的 header 部分
	response.Header().Set("Content-Type", "application/json")

	// 从请求体中获取参数
	foo := request.PostFormValue("foo")
	if foo == "" {
		foo = "10"
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回结构
	obj["data"] = fooInt
	byt, err := json.Marshal(obj)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回状态，输出返回结构
	response.WriteHeader(200)
	response.Write(byt)
	return
}
