package app

import (
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

// WebgoAppProvider 提供App的具体实现方法
type WebgoAppProvider struct {
	BaseFolder string
}

// Register 注册HadeApp方法
func (h *WebgoAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewWebgoApp
}

// Boot 启动调用
func (h *WebgoAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *WebgoAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *WebgoAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *WebgoAppProvider) Name() string {
	return contract.AppKey
}
