package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router map[string]*Tree
	middlewares []ControllerHandler // 从core这边设置的中间件
}

// 初始化框架核心结构
func NewCore() *Core {
	// 定义二级map
	// getRouter := map[string]ControllerHandler{}
	// postRouter := map[string]ControllerHandler{}
	// putRouter := map[string]ControllerHandler{}
	// deleteRouter := map[string]ControllerHandler{}

	// 将二级map写入一级map
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 对应 method
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers :=append(c.middlewares,handlers...)

	if err :=c.router["GET"].AddRouter(url,allHandlers);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["POST"][upperUrl] = handler
	// 将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err :=c.router["POST"].AddRouter(url,allHandlers);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["PUT"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err :=c.router["PUT"].AddRouter(url,allHandlers);err!=nil {
		log.Fatal("add router error: ",err)
	}
}
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["DELETE"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err :=c.router["DELETE"].AddRouter(url,allHandlers);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

// ==== http method wrap end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}


// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		// 查找第二层map
		// if handler, ok := methodHandlers[upperUri]; ok {
		// 	return handler
		// }
		return methodHandlers.FindHandler(upperUri)
	}
	return nil
}

// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares =append(c.middlewares,middlewares...)
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)
	 
	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		ctx.Json(404,"not found")
		return
	}

	ctx.SetHandlers(handlers)

	if err:=ctx.Next();err!=nil{
		ctx.Json(500,"inner error")
		return
	}
}
