---
comments: true
date: 2012-04-30 16:10:55
layout: post
slug: 'using-ecj-in-ant-maven'
title: 在Ant/Maven中使用ecj编译器(Eclipse内置的Java编译器)
permalink: '/416.html'
wordpress_id: 416
categories:
- Java
tags:
- Ant
- bug
- Classpath
- Eclipse
- Java
- Maven
- plugin
- XML
- 下载
- 兼容
---

为什么要换ecj呢? JDK自带的java不够好吗? 是的, 尤其是debug信息. 那两种兼容吗? 完全兼容, ecj和javac一样是经过认证的哦, 事实上,如果你正在使用Eclipse,那么,你的java源码, 100%是ecj编译的呢(当然,是你自己写的那部分)

Ant换用ecj
1. 在build.xml中加入:

    
    
    <property name="build.compiler" value="org.eclipse.jdt.core.JDTCompilerAdapter"></property>
    


2. 下载独立的ecj.jar
[ECJ 3.7.2](http://mirrors.ustc.edu.cn/eclipse/eclipse/downloads/drops/R-3.7.2-201202080800/ecj-3.7.2.jar)
3. 将ecj-3.7.2.jar放入ant的lib文件夹中
4. 如果是eclipse中跑ant,那么,需要设置一下,     Run As -- Ant Build ... -- ClassPath ,加入ecj.jar

Maven换用ecj
1. 官网文档: [http://maven.apache.org/plugins/maven-compiler-plugin/non-javac-compilers.html](http://maven.apache.org/plugins/maven-compiler-plugin/non-javac-compilers.html)
2. 设置为plexus-compiler-eclipse即可

带来的好处: [](http://wendal.net/394.html)
事实证明, 只有ecj编译的class文件的debug信息会原样遵循方法参数的声明顺序, 悲催啊...
