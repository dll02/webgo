<template><div><hr>
<h2 id="lang-zh-cntitle-build-编译命令date-2023-09-02description-build-编译命令sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-build-编译命令date-2023-09-02description-build-编译命令sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: build 编译命令
date: 2023-09-02
description: build 编译命令
sidebar: auto</h2>
<h2 id="build-编译命令" tabindex="-1"><a class="header-anchor" href="#build-编译命令" aria-hidden="true">#</a> build 编译命令</h2>
<p>应用分为前端（frontend）和后端（backend），所以编译也分为三类</p>
<ul>
<li>编译前端</li>
<li>编译后端</li>
<li>自编译</li>
<li>同时编译</li>
</ul>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo build
build webgo command

Usage:
  webgo build [flags]
  webgo build [command]

Available Commands:
  all         build fronend and backend
  backend     build backend use go
  frontend    build frontend use npm
  self        build ./webgo command

Flags:
  -h, --help   help for build

Use "webgo build [command] --help" for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="编译前端" tabindex="-1"><a class="header-anchor" href="#编译前端" aria-hidden="true">#</a> 编译前端</h2>
<p>要求当前编译机器安装 npm 软件，并且当前项目已经运行了 npm install，安装完成前端依赖。</p>
<p>运行命令 <code v-pre>./webgo build frontend</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo build frontend

> webgo@0.1.0 build /Users/Documents/workspace/webgo_workspace/demo5
> vue-cli-service build


-  Building for production...
 DONE  Compiled successfully in 5012ms下午5:44:47

  File                                      Size             Gzipped

  dist/asset/js/chunk-vendors.222f9fef.j    117.96 KiB       42.76 KiB
  s
  dist/asset/js/index.0ed60f05.js           4.63 KiB         1.66 KiB
  dist/asset/css/index.fa7f2f34.css         0.33 KiB         0.23 KiB

  Images and other types of assets omitted.

 DONE  Build complete. The dist directory is ready to be deployed.
 INFO  Check out deployment instructions at https://cli.vuejs.org/guide/deployment.html

front end build success
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>编译后的前端文件在 dist 目录中</p>
<p>实际上 build 就是调用 <code v-pre>npm build</code> 来编译前端项目。</p>
<h2 id="编译后端" tabindex="-1"><a class="header-anchor" href="#编译后端" aria-hidden="true">#</a> 编译后端</h2>
<p>要求当前编译机器安装 go 软件，版本 &gt; 1.3。</p>
<p>运行命令： <code v-pre>./webgo build backend</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo build backend
build success please run ./webgo direct
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>在项目根目录下就看到生成的可执行文件 webgo。 后续可以通过 ./webgo 直接运行。</p>
<h2 id="自编译" tabindex="-1"><a class="header-anchor" href="#自编译" aria-hidden="true">#</a> 自编译</h2>
<p>在项目根目录下，webgo 可以通过 webgo 命令编译出 webgo 命令自己。通过命令 <code v-pre>webgo build self</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ webgo build self
build success please run ./webgo direct
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>在项目根目录下就看到生成的可执行文件 webgo。 后续可以通过 ./webgo 直接运行。</p>
<div class="custom-container tip"><p class="custom-container-title">TIP</p>
<p>其实自编译和后端编译是同样效果，但是为了命令语义化，增加了自编译的命令。</p>
</div>
<h2 id="同时编译" tabindex="-1"><a class="header-anchor" href="#同时编译" aria-hidden="true">#</a> 同时编译</h2>
<p>顾名思义，同时编译前端和后端，命令为 <code v-pre>./webgo build all</code></p>
</div></template>


