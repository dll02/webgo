import {defaultTheme} from "vuepress";

export default {
    lang: 'zh-CN',
    title: "webgo框架", // 设置网站标题
    description: "一个支持前后端开发的基于协议的框架", //描述
    dest: "./dist/", // 设置输出目录
    port: 2333, //端口
    base: "/v1.0/",
    theme: defaultTheme({
        navbar: [
            {
                text: 'Guide',
                link: '/guide/',
            },
            {
                text: 'Provider',
                link: '/provider/',
            },
        ],
        sidebar: {
            '/guide/': [
                {
                    text: '指南',
                    collapsable: false,
                    children: ["introduce",
                        "install",
                        "build",
                        "structure",
                        "app",
                        "env",
                        "dev",
                        "command",
                        "cron",
                        "middleware",
                        "swagger",
                        "provider",
                        "todo",],
                },
            ],
            '/provider/': [
                {
                    text: '服务提供者',
                    collapsable: false,
                    children: ["app",
                        "env",
                        "config",
                        "log",],
                },
            ],
        },
    }),
}