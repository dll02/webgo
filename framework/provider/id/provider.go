package id

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

type WebgoIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *WebgoIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebgoIDService
}

// Boot will called when the service instantiate
func (provider *WebgoIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *WebgoIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *WebgoIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *WebgoIDProvider) Name() string {
	return contract.IDKey
}
