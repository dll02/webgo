package config

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
	"path/filepath"
)

type WebgoConfigProvider struct {
	c      framework.Container
	folder string
	env    string

	envMaps map[string]string
}

// Register registe a new function for make a service instance
func (provider *WebgoConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebgoConfig
}

// Boot will called when the service instantiate
func (provider *WebgoConfigProvider) Boot(c framework.Container) error {
	provider.folder = c.MustMake(contract.AppKey).(contract.App).ConfigFolder()
	provider.envMaps = c.MustMake(contract.EnvKey).(contract.Env).All()
	provider.env = c.MustMake(contract.EnvKey).(contract.Env).AppEnv()
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *WebgoConfigProvider) IsDefer() bool {
	return true
}

// Params define the necessary params for NewInstance
func (provider *WebgoConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

/// Name define the name for this service
func (provider *WebgoConfigProvider) Name() string {
	return contract.ConfigKey
}
