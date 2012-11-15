---
comments: true
date: 2010-12-01 22:17:45
layout: post
slug: '%e5%bd%92%e8%bf%98%e6%8a%80%e6%9c%af%e5%80%ba%e5%8a%a1-alfresco%e5%8d%87%e7%ba%a7%e8%ae%b0'
title: 归还技术债务--Alfresco升级记
permalink: '/191.html'
wordpress_id: 191
categories:
- Alfresco
- Java
- 工作
tags:
- alfresco
- Dom
- Hibernate
- io
- Java
- Oracle
- 下载
- 升级
- 历史遗留
- 反编译
- 技术
- 技术债务
- 部署
- 配置
---

这周开始，按领导的意思，开始着力升级Alfresco到3.3.3 (发现已经发布3.3.3.7 改为升级最新3.x版本)

了解升级步骤，下载升级需要的war包。部署到测试环境，一切都似乎很顺利。
启动！！ 等待了3分钟，报错了！！ 大意是 之前的升级未完成！！

    13:12:29,186 ERROR [org.alfresco.repo.domain.schema.SchemaBootstrap] Schema auto-update failed
    org.alfresco.error.AlfrescoRuntimeException: 11010000 A previous schema upgrade failed or was not completed.  Revert to the original database before attempting the upgrade again.
    	at org.alfresco.repo.domain.schema.SchemaBootstrap.onBootstrap(SchemaBootstrap.java:1373)
    	at org.alfresco.util.AbstractLifecycleBean.onApplicationEvent(AbstractLifecycleBean.java:62)
    	at org.springframework.context.event.SimpleApplicationEventMulticaster$1.run(SimpleApplicationEventMulticaster.java:77)
    	at org.springframework.core.task.SyncTaskExecutor.execute(SyncTaskExecutor.java:49)
    
查阅 [Alfresco升级指南-Schema升级](http://wiki.alfresco.com/wiki/Schema_Upgrade_Scripts) 发现，多了一个表 alf_bootstrap_lock，删除之！
再次启动，Alfresco竟然开始建表了！！ 百思不得其解！！
没办法，出我的绝招，看源码/反编译！
研究了好久，尝试启动无数次后，发现，这Alfresco直接无视数据库中的表！！！ Why？？？！！！！
最后，发现Alfresco是通过Databasemeta来获取表信息的，其中的getTables方法，传入了catalog和schema的值。其中catalog我并未定义，为null，schema的值我设置为alfresco 。 难道。。。 难道。。。 大小写的问题！！！！ 猛然改为

    hibernate.default_schema=ALFRESCO
    
一直以来都无视这个，虽然[Alfresco数据库配置](http://wiki.alfresco.com/wiki/Database_Configuration)一直写的是大写，但我从未在意！
修改后，启动成功，自动开始打补丁！！哦也！！ 成功了！！

恩，为啥会多了一个alf_bootstrap_lock表呢？肯定是之前升级过，并且失败了。但我已经对比过表结构，并未异样。 但为啥一直以来都不报错呢？ 很久之前就写着

    db.schema.update=false
    
这个选项，屏蔽了这个错误！！ 印象中，很久很久以前，我因为看到数据库报错才加入这个选项，难道就是一个原因？？！！ 但是奇怪的是，如果没有添加

    hibernate.default_schema=alfresco
    
的话，一个Oracle实例，只安装一个Alfresco的话，也是不会出错的！ 难道是某个事情，在某个环境中使用了两个Alfresco，然后添加了这一语句，然后导致错误，进而添加禁止更新的选项？？

就这样，花了3天！！

**恩，这也许就是解释，这应该就是技术债务了！**
