package kernel

import (
	"net/http"

	"github.com/dll02/goweb/framework/gin"
)

// 引擎服务
type WebgoKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewWebgoKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &WebgoKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *WebgoKernelService) HttpEngine() http.Handler {
	return s.engine
}
