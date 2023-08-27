package env

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

type WebgoEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *WebgoEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebgoEnv
}

// Boot will called when the service instantiate
func (provider *WebgoEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *WebgoEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *WebgoEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *WebgoEnvProvider) Name() string {
	return contract.EnvKey
}
