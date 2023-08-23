package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router map[string]*Tree
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
func (c *Core) Get(url string, handler ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["GET"][upperUrl] = handler
	if err :=c.router["GET"].AddRouter(url,handler);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["POST"][upperUrl] = handler
	if err :=c.router["POST"].AddRouter(url,handler);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["PUT"][upperUrl] = handler
	if err :=c.router["PUT"].AddRouter(url,handler);err!=nil {
		log.Fatal("add router error: ",err)
	}
}
func (c *Core) Delete(url string, handler ControllerHandler) {
	// upperUrl := strings.ToUpper(url)
	// c.router["DELETE"][upperUrl] = handler
	if err :=c.router["DELETE"].AddRouter(url,handler);err!=nil {
		log.Fatal("add router error: ",err)
	}
}

// ==== http method wrap end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}


// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
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

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)
	 
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404,"not found")
		return
	}

	if err:=router(ctx);err!=nil{
		ctx.Json(500,"inner error")
		return
	}
}
