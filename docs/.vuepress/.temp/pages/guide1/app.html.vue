<template><div><hr>
<h2 id="lang-zh-cntitle-app-运行命令date-2023-09-02description-app-命令指南sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-app-运行命令date-2023-09-02description-app-命令指南sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: app 运行命令
date: 2023-09-02
description: app 命令指南
sidebar: auto</h2>
<h1 id="app-运行命令" tabindex="-1"><a class="header-anchor" href="#app-运行命令" aria-hidden="true">#</a> app 运行命令</h1>
<p>这里的运行是运行整个 app，这个 app 可以只包含后端，也可以只包含前端，但是后端也是隐藏在前端后面运行。具体可以参考 app/http/route.go</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> http

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/dll02/webgo/app/http/controller/demo"</span>
	<span class="token string">"github.com/dll02/webgo/framework/gin"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">Routes</span><span class="token punctuation">(</span>r <span class="token operator">*</span>gin<span class="token punctuation">.</span>Engine<span class="token punctuation">)</span> <span class="token punctuation">{</span>
	r<span class="token punctuation">.</span><span class="token function">Static</span><span class="token punctuation">(</span><span class="token string">"/dist/"</span><span class="token punctuation">,</span> <span class="token string">"./dist/"</span><span class="token punctuation">)</span>
	r<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">"/demo/demo"</span><span class="token punctuation">,</span> demo<span class="token punctuation">.</span>Demo<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

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


