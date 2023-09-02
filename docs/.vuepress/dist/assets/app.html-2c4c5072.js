import{_ as e,o as a,c as d,d as i}from"./app-2d690129.js";const n={},s=i(`<h1 id="运行" tabindex="-1"><a class="header-anchor" href="#运行" aria-hidden="true">#</a> 运行</h1><h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令" aria-hidden="true">#</a> 命令</h2><p>这里的运行是运行整个 app，这个 app 可以只包含后端，也可以只包含前端，但是后端也是隐藏在前端后面运行。具体可以参考 app/http/route.go</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>package http

import (
	&quot;github.com/gohade/hade/app/http/controller/demo&quot;
	&quot;github.com/gohade/hade/framework/gin&quot;
)

func Routes(r *gin.Engine) {
	r.Static(&quot;/dist/&quot;, &quot;./dist/&quot;)
	r.GET(&quot;/demo/demo&quot;, demo.Demo)
}

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>运行相关的命令为 app。</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>[~/Documents/workspace/hade_workspace/demo5]$ ./hade app
start app serve

Usage:
  hade app [flags]
  hade app [command]

Available Commands:
  restart     restart app server
  start       start app server
  state       get app pid
  stop        stop app server

Flags:
  -h, --help   help for app

Use &quot;hade app [command] --help&quot; for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="启动" tabindex="-1"><a class="header-anchor" href="#启动" aria-hidden="true">#</a> 启动</h2><p>可以使用 <code>./hade app start</code> 启动一个应用。</p><p>也可以使用 <code>./hade app start -d</code> 使用 deamon 模式启动一个应用。应用名称为 <code>hade app</code></p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>[~/Documents/workspace/hade_workspace/demo5]$ ./hade app start -d
app serve started
log file: /Users/Documents/workspace/hade_workspace/demo5/storage/log/app.log
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>app 应用的输出记录在 <code>/storage/log/app.log</code></p><p>进程 id 记录在 <code>/storage/pid/app.pid</code></p><h2 id="状态" tabindex="-1"><a class="header-anchor" href="#状态" aria-hidden="true">#</a> 状态</h2><p>当使用 deamon 模式启动的时候，需要查看当前应用是否有启动，如果启动了，进程号是多少，可以使用命令 <code>./hade app state</code></p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>[~/Documents/workspace/hade_workspace/demo5]$ ./hade app state
app server started, pid: 28170
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="重启" tabindex="-1"><a class="header-anchor" href="#重启" aria-hidden="true">#</a> 重启</h2><p>当使用 deamon 模式启动的时候，需要重启应用，可以使用命令 <code>./hade app restart</code></p><div class="custom-container tip"><p class="custom-container-title">TIP</p><p>如果程序还未启动，调用 restart 命令，效果和 start 命令一样，deamon 模式启动应用</p></div><h2 id="停止" tabindex="-1"><a class="header-anchor" href="#停止" aria-hidden="true">#</a> 停止</h2><p>当使用 deamon 模式启动的时候，需要关闭应用，可以使用命令 <code>./hade app stop</code></p>`,20),r=[s];function t(l,c){return a(),d("div",null,r)}const p=e(n,[["render",t],["__file","app.html.vue"]]);export{p as default};
