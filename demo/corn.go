package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main4() {
	// 创建一个cron实例
	c := cron.New()

	// 每整点30分钟执行一次
	c.AddFunc("30 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
	// 上午3-6点，下午8-11点的30分钟执行
	c.AddFunc("30 3-6,20-23 * * *", func() {
		fmt.Println(".. in the range 3-6am, 8-11pm")
	})
	// 东京时间4:30执行一次
	c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() {
		fmt.Println("Runs at 04:30 Tokyo time every day")
	})
	// 从现在开始每小时执行一次
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour, starting an hour from now")
	})
	// 从现在开始，每一个半小时执行一次
	c.AddFunc("@every 1h30m", func() {
		fmt.Println("Every hour thirty, starting an hour thirty from now")
	})

	// 启动cron
	c.Start()

	// 在cron运行过程中增加任务
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// 查看运行中的任务
	//inspect(c.Entries())

	// 停止cron的运行，优雅停止，所有正在运行中的任务不会停止。
	c.Stop()
}
