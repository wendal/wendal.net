---
comments: true
date: 2011-09-26 23:03:32
layout: post
slug: '%e5%88%9d%e8%af%95jetty%e4%bd%bf%e7%94%a8mongodb%e4%bd%9c%e4%b8%basession%e7%ae%a1%e7%90%86%e5%99%a8'
title: 初试Jetty使用Mongodb作为Session管理器
permalink: '/326.html'
wordpress_id: 326
categories:
- Java
- 工作
tags:
- Eclipse
- el
- io
- Java
- mongodb
- Nutz
- SSI
---

闲话少说,直接上代码:

	package net.wendal.jetty.mongodb;

	import org.eclipse.jetty.nosql.mongodb.MongoSessionIdManager;
	import org.eclipse.jetty.nosql.mongodb.MongoSessionManager;
	import org.eclipse.jetty.server.Server;
	import org.eclipse.jetty.server.SessionManager;
	import org.eclipse.jetty.server.session.SessionHandler;
	import org.eclipse.jetty.webapp.WebAppContext;

	public abstract class TestMongodb4Jetty {

		public static void main(String[] args) throws Throwable {
			Server server = new Server(9090);
			WebAppContext webAppContext = new WebAppContext();
			webAppContext.setWar("E:\\NutzQuickStart.war"); //经典的nutz入门例子
			MongoSessionManager msm = new MongoSessionManager();
			SessionHandler sessionHandler = new SessionHandler();
			sessionHandler.setSessionManager(msm);
			webAppContext.setSessionHandler(sessionHandler);
			MongoSessionIdManager idMgr = new MongoSessionIdManager(server);
			idMgr.setWorkerName("wendal-mongodb-worker");
			idMgr.setScavengeDelay(60);
			msm.setSessionIdManager(idMgr);
			server.setHandler(webAppContext);
			server.start();
		}
	}

启动mongod,启动jetty,访问登录页面,登录,看到后台的log

	2011-09-26 23:12:37,470 MongoSessionManager:save:org.eclipse.jetty.nosql.NoSqlSession:wendal-mongodb-worker426chatn2460pobw6a4a6m14@1179468258
	2011-09-26 23:12:37,473 MongoSessionManager:save:db.sessions.update({ "id" : "wendal-mongodb-worker426chatn2460pobw6a4a6m14" , "valid" : true},{ "$inc" : { "context.::*.__metadata__.version" : 1} , "$set" : { "accessed" : 1317049954912 , "context.::*.org%2Enutz%2Equickstart%2Eauth%2Ebean%2EUser" : <Binary Data>}},true)


可以看到,默认情况下,使用的是标准的Java序列化,效率当然是不太高的了,呵呵

打开mongodb控制台,敲入下面的命令,也能看到登录后的session信息

	use HttpSessions
	db.sessions.find()

So esay!! 不过,也用掉了我2个小时,呵呵!!