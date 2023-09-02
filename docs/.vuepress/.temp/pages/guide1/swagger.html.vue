<template><div><hr>
<h2 id="lang-zh-cntitle-swagger命令date-2023-09-02description-swagger命令指南sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-swagger命令date-2023-09-02description-swagger命令指南sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: swagger命令
date: 2023-09-02
description: swagger命令指南
sidebar: auto</h2>
<h2 id="swagger命令" tabindex="-1"><a class="header-anchor" href="#swagger命令" aria-hidden="true">#</a> swagger命令</h2>
<p>webgo 使用 <a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo<ExternalLinkIcon/></a> 集成了 swagger 生成和服务项目。并且封装了 <code v-pre>./webgo swagger</code> 命令。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo swagger
swagger operator

Usage:
  webgo swagger [flags]
  webgo swagger [command]

Available Commands:
  gen         generate swagger file, contain swagger.yaml, doc.go
  serve       use gen serve

Flags:
  -h, --help   help for swagger

Use "webgo swagger [command] --help" for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul>
<li>gen  // 生成swagger文件</li>
<li>serve // 提供swagger服务</li>
</ul>
<h2 id="注释" tabindex="-1"><a class="header-anchor" href="#注释" aria-hidden="true">#</a> 注释</h2>
<p>webgo 使用 <a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo<ExternalLinkIcon/></a> 来实现注释生成 swagger 功能。</p>
<p>全局注释在文件  <code v-pre>app/http/swagger.go</code> 中:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>// Package http API.
// @title webgo
// @version 1.0
// @description webgo测试
// @termsOfService https://github.com/swaggo/swag

// @contact.name yejianfeng
// @contact.email yejianfeng

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8989
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>接口注释请写在各自模块的 api.go 中</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>swagger 注释的格式和关键词可以参考：<a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo<ExternalLinkIcon/></a></p>
<h2 id="生成" tabindex="-1"><a class="header-anchor" href="#生成" aria-hidden="true">#</a> 生成</h2>
<p>使用命令 <code v-pre>./webgo swagger gen</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo swagger gen
2020/09/16 19:57:33 Generate swagger docs....
2020/09/16 19:57:33 Generate general API Info, search dir:./app/http/
2020/09/16 19:57:33 Generating demo.UserDTO
2020/09/16 19:57:33 create docs.go at /Users/Documents/workspace/webgo_workspace/demo5/app/http/swagger/docs.go
2020/09/16 19:57:33 create swagger.json at /Users/Documents/workspace/webgo_workspace/demo5/app/http/swagger/swagger.json
2020/09/16 19:57:33 create swagger.yaml at /Users/Documents/workspace/webgo_workspace/demo5/app/http/swagger/swagger.yaml
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>在目录 <code v-pre>app/http/swagger/</code> 下自动生成swagger相关文件。</p>
<h2 id="服务" tabindex="-1"><a class="header-anchor" href="#服务" aria-hidden="true">#</a> 服务</h2>
<p>可以使用命令 <code v-pre>./webgo swagger serve</code> 启动当前应用的 swagger ui 服务。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo swagger serve
swagger success: http://0.0.0.0:8069/swagger/index.html
if you want to replace, remember use command: swagger gen
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>如果你的 swagger 服务已经启动，更新 swagger 只需要重新运行 <code v-pre>./webgo swagger gen</code> 就能更新。因为 swagger 服务读取的是生成的 swagger.json 这个文件。</p>
</div>
<p>服务端口，我们也可以通过配置文件 <code v-pre>config/[env]/swagger.yaml</code> 中的配置</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>url: http://127.0.0.1:8069
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>来配置swagger serve 启动的服务。</p>
</div></template>


