<div align="center">
<h1>Doge-CDN-Refresh</h1>

[![Auth](https://img.shields.io/badge/Auther--eryajf-ff69b4.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAACXBIWXMAAAsTAAALEwEAmpwYAAADZElEQVR4nO2ZX2iPURjHP/7/aZN/E21DaZvtwoVYyQUuGXLB/LtkLRcUhSJMSVwg3KCUJPJvLmRZtMQFLvwZhUJk/saGLWaYV6eet06n9/3tfd+9531/sW89td9z3vOc8z3nPOc8zzPoQQ+yAnlAA/AbcCzKR2ClTSL7LRNwNPkB5Noi0pggEQeYbovIh4SJLLBBojfwK2EiVTaIjEyYhANsskGkNAUi+2wQmZECkRM2iCxMgUi9DSKrxPhS7GOZjHXPhvGtYlwdMduYKWO9tmH8oBgvxj6KZawOoFfcxk+L8RzsI0fzkyFhO28EvqXg0FHlK1DtReRzFkzOCSmfvIhczIKJOSHlmheRgcAK4E4WTNDpQp4Ba4P4zxRgO3ArgSQqqDQDx4AKCWBDYwSwRDP4JWECG4CpQB9iwCTNsBsJl8srvxk4CtQBN4Enkre0GKQ7RaekCXgAXAcuAHuB1bLaZcAgrd+2uK78ai0PsfLSBshG70u9IDJmyUq2AYvlbD5N4Dg9l/EnArdFV9sdIlfEiLrNzFAlqPwE2kP2OaKNlyfH8Q9QFJVIixgepukqIuTdRRI3Be1TaczjnOgXRSXyTgwUarrBIVb4htbvcMA+nR7+UCdtc6MSOS8G9hj6+gjlnHEBd+WuMVaZ9FMyJiqRydrgB4B80VcFmFBtxALfFm3nlwPv48rfK7WoWDkc8jgpYq0eE/kOnAGG+4RBh3z6qTGOyzfIJeG2nQL6EQMmaEaTgtNdv+jKcFJwbI33XxEpAM76+IArrXIbFqVFpE0MuzeXF4nmEI9ecwZbhVo6GzsaxLhnniw74YQUVdDwwjppv2yzaPbIJ7HJdJz8RIX5JgYAL6V9vg0ifYEXMsAaj3Z3cpfkmPmhQAs5vHxgh+gbbdS0XMyTQdrkbdHhTiwTCdMHTCLl8hCqeGsalnFSO2JDNb3XxOZINqhkttFmfq8W4I1PbGcFuVr21qCloe7EarRvmzT9K01fYxAZJVmgA1yNKxwh4NF4q0Wr+UZs5JIxndskoVLoEskI1e/HRu6TCEpklR1ZebPcqk/YT9cu1UL190NgNCmhQGpfTjelLo2dMNEf2BkylXVF7eL6qAU3WyiVvCHIv7A7pBY2nizGWGC3FOr0XeoQZ96VIcbqAf8K/gLNGaTJ3vwbFgAAAABJRU5ErkJggg==)](https://github.com/eryajf)
[![Eryajf HitCount](https://views.whatilearened.today/views/github/eryajf/eryajf.svg)](https://github.com/eryajf)
[![Eryajf Blog](https://img.shields.io/badge/%E5%8D%9A%E5%AE%A2-%E4%BA%8C%E4%B8%AB%E8%AE%B2%E6%A2%B5-d7b1bf?logo=Blogger)](https://wiki.eryajf.net)
[![Eryajf WeChat](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E8%BF%90%E7%BB%B4%E8%89%BA%E6%9C%AF-71f9fe?logo=WeChat)](https://y.gtimg.cn/music/photo_new/T053M000003iCCnF30PTi3.jpg)
[![Eryajf Awesome Stars](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🆕 多吉云CDN缓存刷新插件 </p>

<img src="https://cnb.cool/66666/resource/-/git/raw/main/img/hengtiao.gif" width="100%" height="3">
</div><br>

## ℹ️ 项目简介

本项目为 [多吉云](https://www.dogecloud.com/) CDN 缓存刷新插件，支持刷新目录和 URL 缓存。访问：[插件市场](https://docs.cnb.cool/zh/plugins.html)

## 🗣️ 用法介绍

刷新CDN 目录缓存：

```yaml
main:
  push:
    - imports:
        - https://cnb.cool/eryajf/build-env/-/blob/main/env.yaml
      stages:
        - name: test dcr path
          image: docker.cnb.cool/znb/doge-cdn-refresh/dcr
          settings:
            ak: "${DOGE_AK}"
            sk: "${DOGE_SK}"
            rtype: "path"
            urls:
              - "https://jenkinsguide.opsre.top/"
```

刷新CDN URL缓存：

```yaml
main:
  push:
    - imports:
        - https://cnb.cool/eryajf/build-env/-/blob/main/env.yaml
      stages:
        - name: test dcr url
          image: docker.cnb.cool/znb/doge-cdn-refresh/dcr
          settings:
            ak: "${DOGE_AK"
            sk: "${DOGE_SK}"
            rtype: "url"
            urls:
              - "https://wiki.eryajf.net/about/"
              - "https://wiki.eryajf.net/pages/b2f34c/"
```

## 📑 参数说明

| 参数 | **必须**/**可选** | 类型 |          说明          |
| :--: | :---------------: | :--------------------: | :--------------------- |
|  ak  |     **必须**      |     string     | 访问多吉云的Access Key |
|  sk  |     **必须**      |    string | 访问多吉云的Secret Key |
| rtype |     **必须**      |    string |        刷新类型，接受 `path` 或 `url`  |
| urls  |     **必须**      |     array |        刷新URL，一个或多个          |


## 📇 项目地址

可选择你熟悉的平台浏览源码：

|   服务商   |                   地址                   |
| :------: | :------------------------------------------: |
|  `CNB`  | <https://cnb.cool/znb/doge-cdn-refresh>  |
| `GitHub` | <https://github.com/eryajf/doge-cdn-refresh> |