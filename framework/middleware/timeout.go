package middleware
 
import (
	"context"
	"fmt"
	"log"
	"time"
  "webgo/framework"
)


func TimeoutHandler(  d time.Duration) framework.ControllerHandler {
  // 使用函数回调
  return func(c *framework.Context) error {

    finish := make(chan struct{}, 1)
    panicChan := make(chan interface{}, 1)

    // 执行业务逻辑前预操作：初始化超时 context
    durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
    defer cancel()

    

    go func() {
      defer func() {
        if p := recover(); p != nil {
          panicChan <- p
        }
      }()
			// 使用next执行具体的业务逻辑
			c.Next()

      finish <- struct{}{}
    }()
    // 执行业务逻辑后操作
    select {
    case p := <-panicChan:
      log.Println(p)
      c.SetStatus(500).Json("inner error")
    case <-finish:
      fmt.Println("finish")
    case <-durationCtx.Done():
      c.SetHasTimeout()
      c.SetStatus(500).Json("time out")
    }
    return nil
  }
}
