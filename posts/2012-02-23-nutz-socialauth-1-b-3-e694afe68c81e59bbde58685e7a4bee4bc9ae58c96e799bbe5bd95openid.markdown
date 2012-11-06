---
comments: true
date: 2012-02-23 10:08:55
layout: post
slug: nutz-socialauth-1-b-3-%e6%94%af%e6%8c%81%e5%9b%bd%e5%86%85%e7%a4%be%e4%bc%9a%e5%8c%96%e7%99%bb%e5%bd%95openid
title: Nutz-Socialauth 1.b.3 -- 支持国内社会化登录(OpenId)
wordpress_id: 376
categories:
- Java
tags:
- git
- Nutz
- OpenID
- 下载
- 提供商
---

作为[Ngqa](https://github.com/howe/ngqa)的子项目之一,开发了2周,终于能拿得出手了

这个项目算是[socialauth](http://code.google.com/p/socialauth/)的一个插件项目,加上本项目的代码,将支持以下的社会化登录(打星号的是本项目添加的登录方式,**率先支持github和BrowserID哦**):



	
  * google

	
  * yahoo

	
  * twitter

	
  * facebook

	
  * hotmail(msn)

	
  * ***QQ连接(qq)**

	
  * *新浪微博(sina)

	
  * *开心网(kaixin001)

	
  * ***github**

	
  * *豆瓣(douban)

	
  * *百度(baidu)

	
  * ***支付宝(alipay)**

	
  * *人人网(renren)

	
  * Foursquare(foursquare)

	
  * Yammer(yammer)

	
  * *腾讯微博(qqweibo)

	
  * *搜狐(sohu)

	
  * ***淘宝(taobao)**

	
  * LinkedIn(linkedin)

	
  * MySpace(myspace)

	
  * AOL(aol)

	
  * ***BrowserID**


有2个尚未验证的,由于提供商本身的原因,尚未解决
--> *网易(net163) -- 神经病机制,竟然要求自行返回原网站输入Code
--> *盛大(sdo) -- 成功过,然后又挂了,QQ群直接拒绝我的加入请求,无解中...

虽然部分提供商有官方/非官方的SDK,但质量非常参差不齐,而且很多都依赖一大堆额外的jar, 故本项目没有使用这些SDK.

大部分网站采用OAuth1/OAuth2登录,部分网站在请求其OpenAPI时需要额外提供签名参数.值得指出的是盛大连接,自行实现了一套不靠谱的规范,成功率低,行为可预测性低.

就登录而已,这次开发,也了解到国内各家提供商对社会化登录/OpenID的态度.
态度最好的,莫过于新浪,完整且可靠的文档,其次是QQ连接. 部分网站,如支付宝/盛大,提供是代码,压根没文档!!

**OAuth1与OAuth2的区别**:
OAuth1 是3次握手, 先检查网站的密钥的可靠性,然后转到用户登录界面,用户登录后再校验用户的返回
OAuth2 是2次握手, 网站引导用户到登录界面,用户登录后再校验用户的返回

国内有多家OAuth2提供商,部分提供商要求提供备案号,保密协议等等神奇的事,当然,在天朝,这都很正常.

**下载地址**: [https://github.com/howe/ngqa/downloads](https://github.com/howe/ngqa/downloads)

如果你有发现哪家网站提供了OAuth1/2登录的话,不妨提醒一下,很可能提供相应的实现哦
