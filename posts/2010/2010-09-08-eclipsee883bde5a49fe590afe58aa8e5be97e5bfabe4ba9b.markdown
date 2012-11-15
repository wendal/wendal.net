---
comments: true
date: 2010-09-08 21:18:44
layout: post
slug: eclipse%e8%83%bd%e5%a4%9f%e5%90%af%e5%8a%a8%e5%be%97%e5%bf%ab%e4%ba%9b
title: Eclipse能够启动得快些
permalink: '/68.html'
wordpress_id: 68
categories:
- Java
tags:
- Eclipse
- Java
- 启动时间
---

昨天在Javaeye上看到一篇关于Eclipse调优的文件,真是当头一棒,之前咋就没想到呢?!!

自己调整了一下,得到以下参数:

    -Xms40m
    -Xmx256m
    -XX:MaxPermSize=128m
    -XX:ReservedCodeCacheSize=128m
    -Dfile.encoding=utf8
    -Xverify:none
    -XX:+DisableExplicitGC
    -XX:+UseParNewGC
    -Xnoclassgc
    -XX:+UseBiasedLocking
    -XX:+UseFastAccessorMethods
    
感觉上快了不少哦
