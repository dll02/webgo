package framework

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	// 注册
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	// 获取服务
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// WebgoContainer 是服务容器的具体实现
type WebgoContainer struct {
	Container
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// NewWebgoContainer 创建一个服务容器
func NewWebgoContainer() *WebgoContainer {
	return &WebgoContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (webgo *WebgoContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range webgo.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (webgo *WebgoContainer) Bind(provider ServiceProvider) error {
	webgo.lock.Lock()

	key := provider.Name()

	webgo.providers[key] = provider
	webgo.lock.Unlock()
	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(webgo); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(webgo)
		method := provider.Register(webgo)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		webgo.instances[key] = instance
	}
	return nil
}

func (webgo *WebgoContainer) IsBind(key string) bool {
	return webgo.findServiceProvider(key) != nil
}

func (webgo *WebgoContainer) findServiceProvider(key string) ServiceProvider {
	webgo.lock.RLock()
	defer webgo.lock.RUnlock()
	if sp, ok := webgo.providers[key]; ok {
		return sp
	}
	return nil
}

func (webgo *WebgoContainer) Make(key string) (interface{}, error) {
	return webgo.make(key, nil, false)
}

func (webgo *WebgoContainer) MustMake(key string) interface{} {
	serv, err := webgo.make(key, nil, false)
	if err != nil {
		panic("container not contain key " + key)
	}
	return serv
}

func (webgo *WebgoContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return webgo.make(key, params, true)
}

func (webgo *WebgoContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(webgo); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(webgo)
	}
	method := sp.Register(webgo)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// 真正的实例化一个服务
func (webgo *WebgoContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	webgo.lock.RLock()
	defer webgo.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := webgo.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return webgo.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := webgo.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := webgo.newInstance(sp, params)
	if err != nil {
		return nil, err
	}

	webgo.instances[key] = inst
	return inst, nil
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (webgo *WebgoContainer) NameList() []string {
	ret := []string{}
	for _, provider := range webgo.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}
