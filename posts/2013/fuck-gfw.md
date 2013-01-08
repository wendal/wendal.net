---
date: 2013-01-08
layout: post
title: Nutz福利之轻功
permalink: '/2013/0108.html'
categories:
- Nutz
tags:
- Nutz
---

作为挨踢人士,翻过伟大的强,总是那么频繁,Nutz新年福利之轻功,提供给大家一个免费的途径,方便大家访问讨论组,查找技术文章...
----------------------------------

首先,连接服务器(Linux下)
----------------------

	ssh -C -N -D 7070 nutz_xxxx@ci.wendal.net
	#输入密码即可, nutz_xxxx即为你的账号

首先,连接服务器(Windows下)
-------------------------

1. [下载putty](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html), 请使用官网地址,切勿使用所谓汉化版
2. 启动putty, 填入域名

	<img src="{{urls.media}}/2013/01/fuck_gfw_1.jpg"></img>

3. 设置不启动shell及启用压缩

	<img src="{{urls.media}}/2013/01/fuck_gfw_2.jpg"></img>
	
4. 添加tunnel

	<img src="{{urls.media}}/2013/01/fuck_gfw_3.jpg"></img>

5. 返回到session, 按"Save"保存设置,然后点击Open,启动连接,输入密码即可

然后就是浏览器设置了
------------------

1. Chrome用户,[安装ProxySwitcher](https://chrome.google.com/webstore/detail/proxy-switchy/caehdcpeofiiigpdhbabniblemipncjj), 使用127.0.0.1端口7070, sockt5协议
2. Firefox用户, 安装AutoProxy,选择ssh -D配置
3. IE,貌似IE对sockt5支持得不太好,不推荐

祝各位Nutzer轻功了得
==================