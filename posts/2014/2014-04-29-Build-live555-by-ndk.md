---
title: 用NDK编译Live555
date: '2014-04-29'
permalink: '/2014/04/29.html'
categories:
- 工作
tags:
- android
---

针对 live555 2014.03.25和 live 2014.04.23 也就是当前最新咯.

准备工作
-------------

下载源码 http://www.live555.com/liveMedia/public/ 该地址经常被X,请问候非圆校长

解压到一个空文件夹, 并将目录名从live改成jni

建一个文件, 叫 Android.mk

文件内容在 https://gist.github.com/wendal/11399988

编译
---------------------------------------------

在jni目录下执行

```
ndk-build
```

如无意外,在 libs/armeabi/ 生成一个so文件.

