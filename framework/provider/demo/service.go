package demo

import (
	"fmt"
	"github.com/dll02/webgo/framework"
	"github.com/dll02/webgo/framework/contract"
)

// 具体的接口实例
type DemoService struct {
	// 实现接口
	contract.DemoService

	// 参数
	c framework.Container
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() contract.Foo {
	return contract.Foo{
		Name: "i am foo",
	}
}
