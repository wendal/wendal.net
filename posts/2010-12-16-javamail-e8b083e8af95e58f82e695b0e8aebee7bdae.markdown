---
comments: true
date: 2010-12-16 09:17:24
layout: post
slug: javamail-%e8%b0%83%e8%af%95%e5%8f%82%e6%95%b0%e8%ae%be%e7%bd%ae
title: JavaMail 调试参数设置
permalink: '/214.html'
wordpress_id: 214
categories:
- Alfresco
- Java
tags:
- alfresco
- bug
- el
- io
- Java
- JavaMail
- SSI
- 升级
- 反编译
- 配置
---

这周折腾Alfresco升级,同时也发现Alfresco 3.3 SP4 的一个bug -- [mail.smtp.auth不起效](https://issues.alfresco.com/jira/browse/ALF-6186)

故顺便看看这个配置在代码中的位置,到处寻觅[JavaMail](http://www.oracle.com/technetwork/java/index-jsp-139225.html)的源码,终于在[kenai](http://kenai.com/projects/javamail/downloads)找到,不过其实我已经用[JD-GUI](http://java.decompiler.free.fr/?q=jdgui)反编译看了一下

找到以下代码:

    
    
    String str2 = this.jdField_session_of_type_JavaxMailSession.getProperty("mail." + this.name + ".auth");
    


另外找到一个比较详细的[JavaMail参数表](http://hi.baidu.com/jlhh/blog/item/823341434fdca71b9313c620.html),但是缺少了一个调试用的参数 -- mail.debug = true , 默认是false, 调试时加上,很多信息哦,O(∩_∩)O哈哈~

后续报道: [官方API](http://java.sun.com/products/javamail/javadocs/index.html)中已经有mail.debug , 看来有点多此一举了
