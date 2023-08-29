package app

import (
	"path/filepath"

	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/util"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// WebgoApp 代表Webgo框架的App实现
type WebgoApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appId      string              // 表示当前这个app的唯一id, 可以用于分布式锁等
	configMap  map[string]string   // 配置加载
}

// Version 实现版本
func (h WebgoApp) Version() string {
	return "0.0.3"
}

// AppID 表示这个App的唯一ID
func (h WebgoApp) AppID() string {
	return h.appId
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h WebgoApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	h.baseFolder = util.GetExecDirectory()
	return h.baseFolder
}

// ConfigFolder  表示配置文件地址
func (app WebgoApp) ConfigFolder() string {
	if val, ok := app.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (app WebgoApp) LogFolder() string {
	if val, ok := app.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

func (app WebgoApp) HttpFolder() string {
	if val, ok := app.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(app.AppFolder(), "http")
}

func (app WebgoApp) ConsoleFolder() string {
	if val, ok := app.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(app.AppFolder(), "console")
}

func (app WebgoApp) StorageFolder() string {
	if val, ok := app.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

// Webgo 定义业务自己的服务提供者地址
func (app WebgoApp) Webgo() string {
	if val, ok := app.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(app.AppFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (app WebgoApp) MiddlewareFolder() string {
	if val, ok := app.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(app.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (app WebgoApp) CommandFolder() string {
	if val, ok := app.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app WebgoApp) RuntimeFolder() string {
	if val, ok := app.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (app WebgoApp) TestFolder() string {
	if val, ok := app.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

// NewWebgoApp 初始化WebgoApp
func NewWebgoApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	appId := uuid.New().String()
	configMap := map[string]string{}
	return &WebgoApp{baseFolder: baseFolder, container: container, appId: appId, configMap: configMap}, nil
}

// LoadAppConfig 加载配置map
func (app *WebgoApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		app.configMap[key] = val
	}
}

// AppFolder 代表app目录
func (app *WebgoApp) AppFolder() string {
	if val, ok := app.configMap["app_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "app")
}
