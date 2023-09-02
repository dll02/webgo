package test

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/provider/app"
	"github.com/dll02/webgo/framework/provider/env"
)

const (
	BasePath = "/Users/linglingdai/go/src/webgo"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewWebgoContainer()
	// 绑定App服务提供者
	container.Bind(&app.WebgoAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.WebgoTestingEnvProvider{})
	return container
}
