---
comments: true
date: 2012-10-20 11:28:21
layout: post
slug: 'build-goexif-under-windows'
title: 在windows下编译goexif
wordpress_id: 462
categories:
- go
tags:
- git
- 下载
---

纯go版:



    
    
    go get github.com/rwcarlsen/goexif/exif
    







cgo版:



    
    
    # 1. 下载libexif源码,获取头文件
    # 2. 下载exif.dll, 记得下载cygwin下的版本, 直接google得到的版本不靠谱,版本太老
    go get github.com/gosexy/exif
    




[libexif-0.6.21](http://wendal.net/wp-content/uploads/2012/10/libexif-0.6.21.zip)
