package main

import (
	"github.com/dll02/webgo/app/console"
	"github.com/dll02/webgo/app/http"
	student "github.com/dll02/webgo/app/provider/student"
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/distributed"
	"github.com/dll02/webgo/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewWebgoContainer()
	// 绑定App服务提供者
	container.Bind(&app.WebgoAppProvider{})

	// 后续初始化需要绑定的服务提供者...
	container.Bind(&student.SubjectProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.WebgoKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
