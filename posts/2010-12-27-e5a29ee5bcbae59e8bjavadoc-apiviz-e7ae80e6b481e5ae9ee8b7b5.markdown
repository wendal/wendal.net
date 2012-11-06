---
comments: true
date: 2010-12-27 13:40:55
layout: post
slug: '%e5%a2%9e%e5%bc%ba%e5%9e%8bjavadoc-apiviz-%e7%ae%80%e6%b4%81%e5%ae%9e%e8%b7%b5'
title: 增强型JavaDoc -- APIviz 简洁实践
wordpress_id: 222
categories:
- Java
- VPS/Linux
tags:
- Ant
- apiviz
- Classpath
- io
- Java
- Javadoc
- JBoss
- Nutz
- UML
- 下载
- 路径
---

一段时间之前,为Nutz的JavaDoc添加了这个插件,[APIviz](http://code.google.com/p/apiviz/),可以自动生成包依赖关系图,和简单的UML类图.
在国内还没见到公开的使用,还是介绍一下,[Nutz示例](http://build.sunfarms.net/nutz/lastest/api/)
第一步,安装[Graphviz](http://www.graphviz.org/),该软件可以运行在N多平台上,这里以Ubuntu为例

    
    
    apt-get install graphviz
    


第二步,当然就是下载APIviz了, 当前最新版为 1.3.1 GA
第三步, Ant调用:

    
    
    <javadoc doclet="org.jboss.apiviz.APIviz" encoding="UTF-8" sourcepath="src" docletpath="${basedir}/build/deps/apiviz-1.3.1.GA.jar" charset="utf-8" destdir="${javadoc-dir}" additionalparam="-author -version -sourceclasspath ${classes-dir-jdk6}" classpathref="nutz-classpath" docencoding="utf-8"></javadoc>
    


与普通的JavaDoc相比,添加了3个属性:

    
    
    #引用APIviz
    doclet="org.jboss.apiviz.APIviz" 
    #指向APIviz的jar包
    docletpath="${basedir}/build/deps/apiviz-1.3.1.GA.jar" 
    #声明类文件的路径,其余的-author -version是官网上建议添加的,非必需.
    #这里的${classes-dir-jdk6}就是Nutz编译好的class存放的地址
    additionalparam="-author -version -sourceclasspath ${classes-dir-jdk6}"
    


这样,就可以自动生成全部包依赖关系了,执行的时候,需要一点点耐心哦
顺带说一下,这APIviz应该是Jboss的产品
