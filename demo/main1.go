package main

import (
    "context"
    "fmt"
    "time"
)

func main1() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go handleRequest(ctx)

    // 模拟一些工作
    time.Sleep(5 * time.Second)
    fmt.Println("主函数完成")
}

func handleRequest(ctx context.Context) {
    dataCh := make(chan string)

    // 模拟数据库获取
    go fetchData(ctx, dataCh)

    select {
    case data := <-dataCh:
        fmt.Println("成功获取数据:", data)
    case <-ctx.Done():
        fmt.Println("在超时或取消前未能获取数据。原因:", ctx.Err())
    }
}

func fetchData(ctx context.Context, ch chan string) {
    // 模拟一些需要时间的数据库操作
    time.Sleep(3 * time.Second)

    // 在发送数据前检查context是否已完成
    if ctx.Err() == nil {
        ch <- "来自数据库的数据"
    }
}
