package env

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

type WebgoTestingEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *WebgoTestingEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebgoTestingEnv
}

// Boot will called when the service instantiate
func (provider *WebgoTestingEnvProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *WebgoTestingEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *WebgoTestingEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *WebgoTestingEnvProvider) Name() string {
	return contract.EnvKey
}
