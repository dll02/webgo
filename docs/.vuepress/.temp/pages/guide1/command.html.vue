<template><div><hr>
<h2 id="lang-zh-cntitle-command-命令date-2023-09-02description-command-命令指南sidebar-auto" tabindex="-1"><a class="header-anchor" href="#lang-zh-cntitle-command-命令date-2023-09-02description-command-命令指南sidebar-auto" aria-hidden="true">#</a> lang: zh-CN
title: command 命令
date: 2023-09-02
description: command 命令指南
sidebar: auto</h2>
<h2 id="command-命令指南" tabindex="-1"><a class="header-anchor" href="#command-命令指南" aria-hidden="true">#</a> command 命令指南</h2>
<p>webgo 允许自定义命令，挂载到 webgo 上。并且提供了<code v-pre>./webgo command</code> 系列命令。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo command
all about commond

Usage:
  webgo command [flags]
  webgo command [command]

Available Commands:
  list        show all command list
  new         create a command

Flags:
  -h, --help   help for command

Use "webgo command [command] --help" for more information about a command.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul>
<li>list  // 列出当前所有已经挂载的命令列表</li>
<li>new   // 创建一个新的自定义命令</li>
</ul>
<h2 id="创建" tabindex="-1"><a class="header-anchor" href="#创建" aria-hidden="true">#</a> 创建</h2>
<p>创建一个新命令，可以使用 <code v-pre>./webgo command new</code></p>
<p>这是一个交互式的命令行工具。</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>[~/Documents/workspace/webgo_workspace/demo5]$ ./webgo command new
create a new command...
? please input command name: test
? please input file name(default: command name):
create new command success，file path: /Users/Documents/workspace/webgo_workspace/demo5/app/console/command/test.go
please remember add command to console/kernel.go
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>创建完成之后，会在应用的 app/console/command/ 目录下创建一个新的文件。</p>
<h2 id="自定义" tabindex="-1"><a class="header-anchor" href="#自定义" aria-hidden="true">#</a> 自定义</h2>
<p>webgo 中的命令使用的是 cobra 库。 https://github.com/spf13/cobra</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>package command

import (
        "fmt"

        "github.com/dll02/webgo/framework/cobra"
        "github.com/dll02/webgo/framework/command/util"
)

var TestCommand = &amp;cobra.Command{
        Use:   "test",
        Short: "test",
        RunE: func(c *cobra.Command, args []string) error {
                container := util.GetContainer(c.Root())
                fmt.Println(container)
                return nil
        },
}

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>基本上，我们要求实现</p>
<ul>
<li>Use // 命令行的关键字</li>
<li>Short // 命令行的简短说明</li>
<li>RunE // 命令行实际运行的程序</li>
</ul>
<p>更多配置和参数可以参考 <a href="https://github.com/spf13/cobra" target="_blank" rel="noopener noreferrer">cobra 的 github 页面<ExternalLinkIcon/></a></p>
<h2 id="挂载" tabindex="-1"><a class="header-anchor" href="#挂载" aria-hidden="true">#</a> 挂载</h2>
<p>编写完自定义命令后，请记得挂载到 console/kernel.go 中。</p>
<div class="language-golang line-numbers-mode" data-ext="golang"><pre v-pre class="language-golang"><code>func RunCommand(container framework.Container) error {
	var rootCmd = &amp;cobra.Command{
		Use:   &quot;webgo&quot;,
		Short: &quot;main&quot;,
		Long:  &quot;webgo commands&quot;,
	}

	ctx := commandUtil.RegiestContainer(container, rootCmd)

	webgoCommand.AddKernelCommands(rootCmd)

    rootCmd.AddCommand(command.DemoCommand)

	return rootCmd.ExecuteContext(ctx)
}

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></div></template>


