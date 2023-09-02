<template><div><h1 id="运行" tabindex="-1"><a class="header-anchor" href="#运行" aria-hidden="true">#</a> 运行</h1>
<h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令" aria-hidden="true">#</a> 命令</h2>
<p>这里的运行是运行整个 app，这个 app 可以只包含后端，也可以只包含前端，但是后端也是隐藏在前端后面运行。具体可以参考 app/http/route.go</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>package http

import (
	"github.com/dll02/webgo/app/http/controller/demo"
	"github.com/dll02/webgo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	r.GET("/demo/demo", demo.Demo)
}

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>运行相关的命令为 app。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app
start app serve

Usage:
  webgo app [flags]
  webgo app [command]

Available Commands:
  restart     restart app server
  start       start app server
  state       get app pid
  stop        stop app server

Flags:
  -h, --help   help for app

Use "webgo app [command] --help" for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="启动" tabindex="-1"><a class="header-anchor" href="#启动" aria-hidden="true">#</a> 启动</h2>
<p>可以使用 <code v-pre>./webgo app start</code> 启动一个应用。</p>
<p>也可以使用 <code v-pre>./webgo app start -d</code> 使用 deamon 模式启动一个应用。应用名称为 <code v-pre>webgo app</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app start -d
app serve started
log file: /Users/Documents/workspace/webgo_workspace/demo5/storage/log/app.log
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>app 应用的输出记录在 <code v-pre>/storage/log/app.log</code></p>
<p>进程 id 记录在 <code v-pre>/storage/pid/app.pid</code></p>
<h2 id="状态" tabindex="-1"><a class="header-anchor" href="#状态" aria-hidden="true">#</a> 状态</h2>
<p>当使用 deamon 模式启动的时候，需要查看当前应用是否有启动，如果启动了，进程号是多少，可以使用命令 <code v-pre>./webgo app state</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo app state
app server started, pid: 28170
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="重启" tabindex="-1"><a class="header-anchor" href="#重启" aria-hidden="true">#</a> 重启</h2>
<p>当使用 deamon 模式启动的时候，需要重启应用，可以使用命令 <code v-pre>./webgo app restart</code></p>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>如果程序还未启动，调用 restart 命令，效果和 start 命令一样，deamon 模式启动应用</p>
</div>
<h2 id="停止" tabindex="-1"><a class="header-anchor" href="#停止" aria-hidden="true">#</a> 停止</h2>
<p>当使用 deamon 模式启动的时候，需要关闭应用，可以使用命令 <code v-pre>./webgo app stop</code></p>
</div></template>


