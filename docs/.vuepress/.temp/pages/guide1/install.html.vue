<template><div><hr>
<h2 id="lang-zh-cntitle-install-安装可执行文件命令date-2023-09-02description-install-安装可执行文件命令指南sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-install-安装可执行文件命令date-2023-09-02description-install-安装可执行文件命令指南sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: install 安装可执行文件命令
date: 2023-09-02
description: install 安装可执行文件命令指南
sidebar: auto</h2>
<h1 id="安装" tabindex="-1"><a class="header-anchor" href="#安装" aria-hidden="true">#</a> 安装</h1>
<hr>
<h2 id="install-安装可执行文件" tabindex="-1"><a class="header-anchor" href="#install-安装可执行文件" aria-hidden="true">#</a> install 安装可执行文件</h2>
<p>我们有两种方式来获取可执行的webgo文件，第一种是直接下载对应操作系统的webgo文件，另外一种是下载源码自己编译</p>
<h3 id="直接下载" tabindex="-1"><a class="header-anchor" href="#直接下载" aria-hidden="true">#</a> 直接下载</h3>
<p>下载地址：
xxx</p>
<p>将生成的可执行文件 webgo 放到 $PATH 目录中：
<code v-pre>cp webgo /usr/local/bin/</code></p>
<h3 id="源码编译" tabindex="-1"><a class="header-anchor" href="#源码编译" aria-hidden="true">#</a> 源码编译</h3>
<p>下载 git 地址：<code v-pre>git@github.com/jianfengye/webgo:cloud/webgo.git</code> 到目录 webgo</p>
<p>在 webgo 目录中运行命令 <code v-pre>go run main.go build self</code></p>
<p>将生成的可执行文件 webgo 放到 $PATH 目录中：
<code v-pre>cp webgo /usr/local/bin/</code></p>
<h2 id="初始化项目" tabindex="-1"><a class="header-anchor" href="#初始化项目" aria-hidden="true">#</a> 初始化项目</h2>
<p>使用命令 <code v-pre>webgo new [app]</code> 在当前目录创建子项目</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace]$ webgo new --help
create a new app

Usage:
  webgo new [app] [flags]

Aliases:
  new, create, init

Flags:
  -f, --force        if app exist, overwrite app, default: false
  -h, --help         help for new
  -m, --mod string   go mod name, default: folder name
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>你可以通过 --mod 来定义项目名字的模块名字，默认项目名，目录名，模块名是相同的</p>
<p>接下来，可以通过命令 <code v-pre>go run main.go</code> 看到如下信息：</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ go run main.go
webgo commands

Usage:
  webgo [command]

Available Commands:
  app         start app serve
  build       build webgo command
  command     all about commond
  cron        about cron command
  deploy      deploy app by ssh
  dev         dev mode
  env         get current environment
  help        get help info
  middleware  webgo middleware
  new         create a new app
  provider    about webgo service provider
  swagger     swagger operator

Flags:
  -h, --help   help for webgo

Use "webgo [command] --help" for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>至此，项目安装成功。</p>
</div></template>


