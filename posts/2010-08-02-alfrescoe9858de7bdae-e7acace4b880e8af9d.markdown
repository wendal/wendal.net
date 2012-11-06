---
comments: true
date: 2010-08-02 23:19:53
layout: post
slug: alfresco%e9%85%8d%e7%bd%ae-%e7%ac%ac%e4%b8%80%e8%af%9d
title: Alfresco配置 — 第一节
wordpress_id: 24
categories:
- Alfresco
- Java
tags:
- alfresco
- Hibernate
- io
- Java
- MySQL
- Tomcat
- Wiki
- 下载
- 路径
- 配置
---

**1. 有用的信息**

Alfresco下载地址 社区版 [_http://wiki.alfresco.com/wiki/Download_Community_Edition_](http://wiki.alfresco.com/wiki/Download_Community_Edition) 企业版可以试用30天,区别不大.

**2. 安装**

如果你下载的是Windows版的超大安装包,那么,基本上就是一路的next,中途填一下管理员密码就可以了.

启动之前,我建议你看看   安装文件夹/tomcat/shared/classes/alfresco-global.properties      ,你会看到不少有用的信息:


dir.root=./alf_data 非常非常核心的参数,务必使用绝对路径,能减少不必要的麻烦(例如移动文件夹后,启动报错)


db.开头的都是数据库配置, 其中db.url就是把部分参数合成jdbc url. 关于数据库,我的建议是使用数据源,Alfresco会默认查找jndi名为jdbc/dataSource的资源作为数据源.

如果你安装的时候使用默认密码admin,就会看到一个


alfresco_user_store.adminpassword=209c6174da490caeb422f3fa5a7ae634


提前说一个问题--忘记管理员密码咋办? 看看这篇文章

[http://wiki.alfresco.com/wiki/Security_and_Authentication#How_to_reset_the_admin_password](http://wiki.alfresco.com/wiki/Security_and_Authentication#How_to_reset_the_admin_password)

一般来说, 就是alf_node_properties表的第4或5行. 再提醒一句,只有使用alfresco/alfresco登录mysql才能看到alfresco数据库.

准备去启动Alfresco或者已经启动了? 不要急嘛, 如果你是*unix系统,请修改/etc/hosts, 查看你的主机名是否被解析为正确的ip,我可吃了不少苦头!!

**3. 故障排除**

启动失败? 不要惊慌,看看是不是以下错误:

PermGen XXXX ,哈哈,内存不够了? 修改alfresco.bat/alfresco.sh里面的JAVA_OPTS吧,调整-XX:MaxPermSize=160m为-XX:MaxPermSize=256m

Hibernate dialect Must set ,数据库连接出错啦!! 检查一下数据库是否已经启动,alfresco-global.properties里面填的数据库信息是否正确

切忌,不能安装到有空格的路径,最好连中文啥的都不要有.

据我的经验,启动alfresco需时60秒到90秒,第一次启动因为要初始化数据库,故需要的时间更长,有时候看上去停止了,只要cpu占用还不低,就没问题.

**4. 启动成功**

哦也,你终于启动成功了! (如果按默认安装的话,不能启动的概率=0)


访问 [http://localhost:8080/alfresco](http://localhost:8080/alfresco) 填入你的帐号密码(例如admin/admin) 就能看到强大的Alfresco了!
