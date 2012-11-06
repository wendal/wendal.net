---
comments: true
date: 2010-08-03 22:49:31
layout: post
slug: alfresco%e9%85%8d%e7%bd%ae-%e7%ac%ac%e4%ba%8c%e8%8a%82
title: Alfresco配置 — 第二节
wordpress_id: 29
categories:
- Alfresco
- Java
tags:
- alfresco
- Hibernate
- io
- Java
- SSI
- Tomcat
- Wiki
- XML
- 索引
- 连接池
- 部署
- 配置
---

如何关闭/调整部分功能(通过添加选项到alfresco-global.properties):


**1. 关闭OpenOffice连接 ooo.enabled=false**

Alfresco默认安装OpenOffice进行文件转换,不过,大部分时间是无需的,这部分功能会占用超过100M的内存,而且是JVM之外的内容

**2. 关闭CIFS和FTP cifs.enabled=false  ftp.enabled=false**

当你启动Alfresco后,你也许能通过 \\你的ip 访问到Alfresco的资源库,这对开发非常有用,但是一般情况下不太需要, ftp也是.

**3. 关闭用户空间配额限制 system.usages.enabled=false**

Alfresco允许你配置每个用户的空间占用,一般使用都是无需的,而且,要真正启用这个功能,你需要逐一配置每个用户的配额.

**4. 关闭自动创建用户空间 home.folder.creation.eager=false**

这个选项,是我在配置LDAP信息同步的时候遇到的,由于有好几千的用户信息同步到Alfresco,结果在User Space中对应地产生了好几千个子空间,虽然无害,但毕竟非常不雅!!

**5. 将索引恢复模式设置为自动 index.recovery.mode=AUTO**

其实这是默认值,但我仍然要单独提出来. 当你放了成千上万的文档时,你如果设置为FULL,启动Alfresco将非常漫长.如果你配置Alfresco集群,AUTO也绝对是最佳选项.除非你的索引已经被破坏,以致启动失败,那FULL才是你的选择.另外,我建议你每周做一次FULL,能提供索引的可靠性和减少体积.

**6. 调整连接池或者使用自定义的数据源**

先看看 [http://wiki.alfresco.com/wiki/Database_Configuration](http://wiki.alfresco.com/wiki/Database_Configuration)

当你使用WebLogic来部署Alfresco,那么请加上db.pool.statements.enable=false

我建议你使用自定义的数据源,替代Alfresco默认的DBCP,例如C3P0,Proxool,BoneCP

在Tomcat的server.xml添加一个全局的数据源,然后在content.xml引用它,并确保名字为jdbc/dataSource


某些故障排除




1. 曾经遇到一个情况,Alfresco启动时,读取完配置后就停住,cpu为0,假死, 后来发现是数据库服务器的内存耗尽,导致HibernateSessionFastory创建时一直等待.


2.**务必修改/etc/hosts文件,使其与当前ip匹配**

当服务器换ip后,没有改/etc/hosts文件里面的ip,导致启动是查找RMI端口时,长时间等待, 超时后报错停止.

如果使用默认的127.0.0.1, 当你把vtomcat放到其他机器上,启动连接到Alfresco时就会报127.0.0.2出错, windows下无此问题.

3. 无法添加新文件或文件夹

这个问题表现为页面上提示无法添加,后台提示无法创建XX资源. 原因是磁盘已满,用df看看磁盘占用情况,清理不需要的文件.


有个小小技巧


删除了重要文件,而且提交了修改,咋办??

不要惊慌,Alfresco并没有真正删除你的文件, 点击用户属性(上方第二个按钮),可以看到最下面有"已删除的资源" 点开,哈哈,看到你想找的文件没?
