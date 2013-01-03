---
title: Deploy Nutz as HttpAPI by Jetty 7
date: '2012-12-13'
description:
categories: [linux]
tags : [vps]
permalink: '/2012/1213.html'
---


将Nutz挂载到jetty上,作为HttpAPI
==============================

这里描述的,一个这样的web服务
---------------------------

1. 不需要jsp
2. 不需要静态资源,例如js/css
3. 仅挂载一个NutFilter,不需要其他jar

实现
----

	//新建一个Jetty Server,监听8080端口
	Server server = new Server(8080); 
	//创建一个Servlet容器,并映射在根路径
	ServletContextHandler ctx = new ServletContextHandler();
	ctx.setContextPath("/");
	
	//加入默认Servlet或者空Servlet类,否则Filter类无法访问NutFilter
    ctx.addServlet(DefaultServlet.class, "/*");
	//设置Session容器,否则Session不可以(Nutz会使用Session容器)
	ctx.setSessionHandler(new SessionHandler(new HashSessionManager()));
	
	//创建Filter持有者,也就是挂载NutFilter
	FilterHolder fh = new FilterHolder(NutFilter.class);
	//传入必需的参数modules,你还可以传入ignore之类的参数
	fh.setInitParameter("modules", "net.wendal.web.MainModule");
	ctx.addFilter(fh, "/*", null);
	
	server.setHandler(ctx);
	
	//启动服务
	server.start();

通过nutz-web来实现
------------------

	// 创建一个web.properties,填入
	app-root=.
	app-port=8080
	admin-port=8081
	#mainModuleClassName这个参数请查阅最新的nutz-web代码
	mainModuleClassName=net.wendal.web.MainModule
	
	//启动代码
	public static void main(String[] args) {
		org.nutz.web.WebLauncher.main(args);
	}

nutz-web项目简介
---------------

一个Jetty封装,外加几个NutMvc的View, (项目地址)[http://github.com/nutzam/nutz]

