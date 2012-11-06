---
comments: true
date: 2010-09-07 16:24:43
layout: post
slug: '%e9%81%87%e5%88%b0%e4%b8%80%e4%b8%aa%e5%85%b3%e4%ba%8eclasspath%e7%9a%84%e5%a5%87%e6%80%aa%e9%97%ae%e9%a2%98-classpath%e5%8f%aa%e8%ae%a4jar%e6%96%87%e4%bb%b6no'
title: 遇到一个关于ClassPath的奇怪问题 Classpath只认jar文件?No!
permalink: '/65.html'
wordpress_id: 65
categories:
- Java
tags:
- Classpath
- Java
- 路径
- 部署
- 配置
---

昨天在客户现场部署应用,解压后开始修改配置文件,以为一切顺利,结果发现程序根本就无视我的配置文件!!

启动代码是这样的:

	java -cp . -Djava.ext.dirs=. xxx.yyy.Main deploy.properties

在当前目录有 XXX.jar deploy.properties a.zip 还有就是一堆类文件在 org文件夹下, 整个文件夹的文件,就是a.zip的解压出来的.

deploy.properties就是我修改的配置文件,结果无论怎么改,程序都无视我的修改. 然后我怒了,把deploy.properties删除了,发现程序依旧运行!! 疯了,deploy.properties是启动该程序必须的!!怎么可能还能启动,这是使用spring properties holder 加载的, 写法是  classpath:deploy.properties

终于开始怀疑是否是zip压缩包的原因,删掉!! 结果,正确运行了!!

奇怪,为啥呢??!! 竟然zip文件都当成jar文件处理?? 找来一个有Main类的jar文件,并改名为XXXX.zip,执行:

	java -cp XXXX.zip xxx.yyy.Main     //结果正常启动了!!

再狠一点,改为后缀改为rar, 执行 java -cp XXXX.rar xxx.yyy.Main , 结果一样,照样运行!!

继续狠一下, 执行 

	java -Djava.ext.dirs=. xxx.yyy.Main        //没办法,照样运行!!!

再一次刷新我对Classpath的认识!!
