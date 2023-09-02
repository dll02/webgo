import{_ as e,o as n,c as a,d}from"./app-2d690129.js";const i={},s=d(`<h1 id="安装" tabindex="-1"><a class="header-anchor" href="#安装" aria-hidden="true">#</a> 安装</h1><hr><h2 id="可执行文件" tabindex="-1"><a class="header-anchor" href="#可执行文件" aria-hidden="true">#</a> 可执行文件</h2><p>我们有两种方式来获取可执行的hade文件，第一种是直接下载对应操作系统的hade文件，另外一种是下载源码自己编译</p><h3 id="直接下载" tabindex="-1"><a class="header-anchor" href="#直接下载" aria-hidden="true">#</a> 直接下载</h3><p>下载地址： xxx</p><p>将生成的可执行文件 hade 放到 $PATH 目录中： <code>cp hade /usr/local/bin/</code></p><h3 id="源码编译" tabindex="-1"><a class="header-anchor" href="#源码编译" aria-hidden="true">#</a> 源码编译</h3><p>下载 git 地址：<code>git@github.com/jianfengye/hade:cloud/hade.git</code> 到目录 hade</p><p>在 hade 目录中运行命令 <code>go run main.go build self</code></p><p>将生成的可执行文件 hade 放到 $PATH 目录中： <code>cp hade /usr/local/bin/</code></p><h2 id="初始化项目" tabindex="-1"><a class="header-anchor" href="#初始化项目" aria-hidden="true">#</a> 初始化项目</h2><p>使用命令 <code>hade new [app]</code> 在当前目录创建子项目</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>[~/Documents/workspace/hade_workspace]$ hade new --help
create a new app

Usage:
  hade new [app] [flags]

Aliases:
  new, create, init

Flags:
  -f, --force        if app exist, overwrite app, default: false
  -h, --help         help for new
  -m, --mod string   go mod name, default: folder name
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>你可以通过 --mod 来定义项目名字的模块名字，默认项目名，目录名，模块名是相同的</p><p>接下来，可以通过命令 <code>go run main.go</code> 看到如下信息：</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>[~/Documents/workspace/hade_workspace/demo5]$ go run main.go
hade commands

Usage:
  hade [command]

Available Commands:
  app         start app serve
  build       build hade command
  command     all about commond
  cron        about cron command
  deploy      deploy app by ssh
  dev         dev mode
  env         get current environment
  help        get help info
  middleware  hade middleware
  new         create a new app
  provider    about hade service provider
  swagger     swagger operator

Flags:
  -h, --help   help for hade

Use &quot;hade [command] --help&quot; for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>至此，项目安装成功。</p>`,18),l=[s];function r(c,o){return n(),a("div",null,l)}const m=e(i,[["render",r],["__file","install.html.vue"]]);export{m as default};
