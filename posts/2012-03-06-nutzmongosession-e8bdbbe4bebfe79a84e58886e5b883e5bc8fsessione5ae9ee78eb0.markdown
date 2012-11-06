---
comments: true
date: 2012-03-06 15:20:40
layout: post
slug: 'nutz-mongo-sessoion'
title: NutzMongoSession -- 轻便的分布式Session实现
permalink: '/393.html'
wordpress_id: 393
categories:
- mongodb
tags:
- io
- js
- Nutz
- SSI
- XML
- 兼容
- 配置
---

作为Ngqa项目的其中一个功能, NutzMongoSession已经开发完成,并迁入zTask项目中,与NutzMongo的代码整合.

**1. 用法**
基本配置(web.xml),接管全部请求,并替换其获取会话的req.getSession方法:

    
    
        <filter>
            <filter-name>mongoSession</filter-name>
            <filter-class>org.nutz.mongo.session.MongoSessionFilter</filter-class>
        </filter>
        <filter-mapping>
            <filter-name>mongoSession</filter-name>
            <url-pattern>/*</url-pattern>
    		<dispatcher>REQUEST</dispatcher>
    		<dispatcher>FORWARD</dispatcher>
    		<dispatcher>INCLUDE</dispatcher>
        </filter-mapping>
    



在项目启动的代码中,加入下面的语句:

    
    
    new MongoSessionManager(dao).register(servletContext, null);
    //其中,dao是MongoDao的实例
    



**2. 与标准HttpSession的差异**
可存放的对象类型: 普通数据类型,MongoDao所管理的Pojo,一切可以顺利json序列化/反序列化的对象
由于是分布式的Session,必须考虑对象的状态托管问题,上代码:

    
    
    User user = new User();
    user.setName("wendal");
    session.setAttribute("me", user);
    user.setName("ABC"); //这个方法调用,改变了对象的属性,但没有再次调用session.setAttribute
    //....
    
    //另外一个方法里面
    String myName = ((User)session.getAttribute("me")).getName();
    //这里所得到的值,将是"wendal"而非"ABC"
    
    //------------------------------------------------------------------------------------
    
    //而,如果User是一个MongoDao所管理的bean的话,那么,它的
    session.setAttribute("me", user);
    user.setName("ABC"); //这个方法调用,改变了对象的属性,但没有再次调用session.setAttribute
    dao.save(user); 把user更新到mongo
    
    //另外一个方法里面
    String myName = ((User)session.getAttribute("me")).getName();
    //这里所得到的值,将是"ABC"
    
    



**3. 获取全部Session**
一直以来,我都以为无法通过标准的ServletAPI简单获取当前应用的全部Session,直至我完成了NutzMongoSession,才发现HttpSession接口有一个已经废弃的方法可以做到
好吧,我承认,我误导了很多很多童鞋....

    
    
    req.getSession().getSessionContext().getIds();
    //这样就能拿到全部Session的Id,然后通过Id,获取
    req.getSession().getSessionContext().getSession(id);
    //来获取具体的Session
    


为了与标准的HttpSession最大兼容,所以,NutzMongoSession也实现了这个方法.

**4. 不足之处**
清理过期Session的算法,还需要改进,因为在服务器端,循环执行new Date()还是有一定成本的

当前, NutMongoSession已经应用在Ngqa和zTask,效果还不错.
有了NutMongoSession, **应用做水平扩展**就会很方便了,呵呵
