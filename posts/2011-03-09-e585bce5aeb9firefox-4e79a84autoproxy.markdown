---
comments: true
date: 2011-03-09 07:50:52
layout: post
slug: '%e5%85%bc%e5%ae%b9firefox-4%e7%9a%84autoproxy'
title: 兼容FireFox 4的AutoProxy
wordpress_id: 250
categories:
- 其他
tags:
- el
- Firefox，Autoproxy
- 下载
- 兼容
- 升级
- 配置
---

#########################################
##官方已经发布兼容FF4 版本!! [马上去下载安装吧!!](https://addons.mozilla.org/zh-CN/firefox/addon/autoproxy/versions/)
#########################################

前几天, FF4 RC1 放出, 作为FF粉丝怎么可以不升级?

安装后,启动,检查附件组件兼容性, 什么?!! Autoproxy不兼容?!! 这可不行,这是必备的!!!

到处查找更新版,官方发布的0.4b1也不支持FF4 RC1 !!!

好吧,根据其[issue 174](http://code.google.com/p/autoproxy/issues/detail?id=147),其实已经修改,只是还没有发布,只好直接跑到Autoproxy的官网下载源码进行编译(为此下载了Git For Windows,Perl,MS-DOS Zip)

经过简单试用,恩,一切正常!!

好东西不能独享,放出我自己的编译的版本给大家救急

如有更新,我也会及时跟进(前提是官方还未放出)
**下载:
---> autoproxy-0.4.0+.2011030823.xpi
[Google Docs 下载](https://docs.google.com/leaf?id=0B8hUXYDeoy_hY2JhMTZjNjUtODE2ZC00ODE0LThjMjYtZTZmNmNhMDNlOGJl&sort=name&layout=list&num=50) , [ 本地下载](http://build.sunfarms.net/download/autoproxy-0.4.0+.2011030823.xpi)**

**---> FF4 RC1 中文版(已经可以在官网公下载)
<del>[Google Docs 下载 ](https://docs.google.com/leaf?id=0B8hUXYDeoy_hMjhjNmQ3M2EtOGJkZi00MWNkLWE1M2YtZTVmMjU4ZjcxZTJh&hl=zh_CN)</del>**

**提醒: 安装后右下角无"福"字,需要进入附加组件管理器进行配置.
**

常出去走走,有益身心

附上访问Google Docs简单方法:
修改Host文件     C:\WINDOWS\system32\drivers\etc\hosts
添加:
209.85.225.101 docs.google.com
74.125.127.100 writely.google.com
72.14.203.100 spreadsheets.google.com
