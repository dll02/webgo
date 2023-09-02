<template><div><hr>
<h2 id="lang-zh-cntitle-env-环境变量设置命令date-2023-09-02description-env-环境变量设置命令指南sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-env-环境变量设置命令date-2023-09-02description-env-环境变量设置命令指南sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: env 环境变量设置命令
date: 2023-09-02
description: env 环境变量设置命令指南
sidebar: auto</h2>
<h1 id="" tabindex="-1"><a class="header-anchor" href="#" aria-hidden="true">#</a> </h1>
<h2 id="env-环境变量设置命令" tabindex="-1"><a class="header-anchor" href="#env-环境变量设置命令" aria-hidden="true">#</a> env 环境变量设置命令</h2>
<p>webgo 支持使用应用默认下的隐藏文件 <code v-pre>.env</code> 来配置各个机器不同的环境变量。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>APP_ENV=development

DB_PASSWORD=mypassword
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>环境变量的设置可以在配置文件中通过 <code v-pre>env([环境变量])</code> 来获取到。</p>
<p>比如：</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>mysql:
    hostname: 127.0.0.1
    username: yejianfeng
    password:  env(DB_PASSWORD)
    timeout: 1
    readtime: 2.3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="应用环境" tabindex="-1"><a class="header-anchor" href="#应用环境" aria-hidden="true">#</a> 应用环境</h2>
<p>webgo 启动应用的默认应用环境为 development。</p>
<p>你可以通过设置 .env 文件中的 APP_ENV 设置应用环境。</p>
<p>应用环境建议选择：</p>
<ul>
<li>development // 开发使用</li>
<li>production // 线上使用</li>
<li>testing //测试环境</li>
</ul>
<p>应用环境对应配置的文件夹，配置服务会去对应应用环境的文件夹中寻找配置。</p>
<p>比如应用环境为 development，在代码中使用</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>configService := container.MustMake(contract.ConfigKey).(contract.Config)
url := configService.GetString("app.url")
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>查找文件为：<code v-pre>config/development/app.yaml</code></p>
<p>通过命令<code v-pre>./webgo env</code>也可以获取当前应用环境：</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo env
environment: development
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div></div></template>


