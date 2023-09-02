export const siteData = JSON.parse("{\"base\":\"/v1.0/\",\"lang\":\"zh-CN\",\"title\":\"webgo框架\",\"description\":\"一个支持前后端开发的基于协议的框架\",\"head\":[],\"locales\":{}}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  if (__VUE_HMR_RUNTIME__.updateSiteData) {
    __VUE_HMR_RUNTIME__.updateSiteData(siteData)
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(({ siteData }) => {
    __VUE_HMR_RUNTIME__.updateSiteData(siteData)
  })
}
