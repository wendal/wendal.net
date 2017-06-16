---
title: Webapp执行reload后内存泄漏之SSLSocketFactory
date: '2017-06-16'
permalink: '/2017/06/18.html'
description: 分析webapp重载后内存不释放的原因之一SSLSocketFactory
categories:
- Java
tags:
- nutz
- tomcat
- jetty
---

## tomcat/jetty是怎么做reload的呢?

首先, webapp的WEB-INF/lib目录并不在jvm的classpath内, javaee容器(tomcat/jetty/jboss)是通过自定义的ClassLoader来加载它们的.

而这个ClassLoader,通常的名字就叫做WebappClassLoader, 容器会为每个webapp的每次启动,都创建一个新的ClassLoader.

"每个webapp",保证了不同webapp之间的类隔离, 例如有A/B两个webapp,都使用了DEF类的XXX静态属性,那么在JVM里面就有2份DEF类,两份XXX静态属性.

"每次启动", 是因为容器会先执行一次unload,再执行load,相当于一个新的webapp加载进来.

## 为什么reload有泄漏?

首先,什么是泄漏? 就是你创建了某些对象/数据, 期望它会被GC, 但事实上没有.

那为啥不被GC呢? 那肯定是被引用了.

```
虚无的ROOT --> 根ClassLoader --> 一些类的静态属性 --> 对象(webapp里面创建的对象) --> 对象的类 --> WebappClassLoader --> 类 --> 静态属性
虚无的ROOT --> 线程组 --> 线程 --> 对象(webapp里面创建的对象) --> 对象的类 --> WebappClassLoader --> 其他类 --> 静态属性
```

两条路径:

* webapp内创建的对象,赋值到了根ClassLoader加载的一个类的静态属性上.
* webapp内创建的线程,reload之后也没有stop

一时半刻想不出其他路径了 -_-

## 静态属性的实例

一个非常非常经典的写法, 忽略Https的无效证书(通常是自签名证书)

```java
            SSLContext sc = SSLContext.getInstance("SSL");
            TrustManager[] tmArr = {new X509TrustManager() {
                // 全是空实现,不写出来了.
            }};
            sc.init(null, tmArr, new SecureRandom());
            HttpsURLConnection.setDefaultSSLSocketFactory(sc.getSocketFactory());
```

HttpsURLConnection并非WebappClassLoader加载,而是由根ClassLoader加载.

然后, new X509TrustManager(){}所创建的匿名内部类对象的类,是由WebappClassLoader加载的.

所以呢, 上述代码就把一个 WebappClassLoader所加载的类的实例,赋值给根ClassLoader加载的类的一个静态属性.

最后, 当Webapp被reload时, 老的WebappClassLoader不会被GC, 直至上述代码再被执行,属性值被覆盖.

在一年前, nutz的Http也是这个写法, 后来改成HttpsURLConnection的实例方法setSSLSocketFactory

类似的代码存在于很多需要访问第三方网站的java库里面, 例如jpush, socialauth

## 线程创建导致的泄漏

这种就只能靠自律了... 例如dubbo里面就建立一堆线程池,然而没有提供销毁的方法...