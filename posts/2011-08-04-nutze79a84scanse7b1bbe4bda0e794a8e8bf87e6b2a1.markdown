---
comments: true
date: 2011-08-04 08:40:25
layout: post
slug: nutz%e7%9a%84scans%e7%b1%bb%e4%bd%a0%e7%94%a8%e8%bf%87%e6%b2%a1
title: Nutz的Scans类,你用过没?
permalink: '/307.html'
wordpress_id: 307
categories:
- Java
tags:
- MVC
- Nutz
- 配置
---

你是否想得到某个package下的全部类呢?
你是否想得到某个文件夹下全部的配置文件?

你还在自己写这种的实现? 你还在拼命地问谷哥?

好吧,不吹水了,正式介绍Scans类:
org.nutz.resource.Scans

首先,必须先强调一点,在J2EE环境中,如果你没有使用NutzMVC的话,请在Filter/Servlet中加入这一句:

	Scans.me().init(config.getServletContext());

这个语句必须优先于其他任何Nutz相关语句

好,如何得到某个文件夹下全部的配置文件呢?即使打包成jar也无需改代码呢?

	//第一个参数是需要扫描的文件夹,第二个是文件名需要匹配正则表达式,可以为null
	List<nutresource> list = Scans.me().scan("config/","^.+\\.ini");

得到某个package下全部的类

	//第一个参数是需要扫描的文件夹,第二个是文件名需要匹配正则表达式,可以为null
	List<class <?>> list = Scans.me().scanPackage("net.wendal.mvc");


注意,你应该已经发现, 返回的是NutResource,而不是File之类的对象,why?
其实,这正是我当时做resource包的初衷 -- File对象无法表达jar里面的文件,而InputStream又不包含文件名之类的信息

NutResource类最重要的两个方法:

	public String getName();
	public InputStream getInputStream()
	
同时包含两中信息,满足大部分需求

Scans类的局限性:
1. 不允许进行根路径扫描 -- 就是说,你不能直接把文件放在classpath根下,必须放在文件夹中
2. 现在已经测试在Tomcat/Jetty/WebLogic下正常,在Maven环境下不正常(至今搞不懂Maven的ClassLoader在某些情形下竟然返回null)