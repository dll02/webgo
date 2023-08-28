package trace

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

type WebgoTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *WebgoTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebgoTraceService
}

// Boot will called when the service instantiate
func (provider *WebgoTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *WebgoTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *WebgoTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *WebgoTraceProvider) Name() string {
	return contract.TraceKey
}
