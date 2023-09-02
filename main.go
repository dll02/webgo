package main

import (
	"github.com/dll02/webgo/app/console"
	"github.com/dll02/webgo/app/http"
	"github.com/dll02/webgo/app/provider/student"
	"github.com/dll02/webgo/app/provider/user"
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/config"
	"github.com/dll02/webgo/framework/provider/distributed"
	"github.com/dll02/webgo/framework/provider/env"
	"github.com/dll02/webgo/framework/provider/id"
	"github.com/dll02/webgo/framework/provider/kernel"
	"github.com/dll02/webgo/framework/provider/log"
	"github.com/dll02/webgo/framework/provider/orm"
	"github.com/dll02/webgo/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewWebgoContainer()
	// 绑定App服务提供者
	container.Bind(&app.WebgoAppProvider{})

	// 后续初始化需要绑定的服务提供者...
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&env.WebgoEnvProvider{})
	container.Bind(&config.WebgoConfigProvider{})
	container.Bind(&id.WebgoIDProvider{})
	container.Bind(&trace.WebgoTraceProvider{})
	container.Bind(&log.WebgoLogServiceProvider{})
	container.Bind(&orm.GormProvider{})

	// 绑定 app 自定义服务
	container.Bind(&student.SubjectProvider{})
	container.Bind(&user.UserProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.WebgoKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
