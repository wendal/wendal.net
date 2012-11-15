---
comments: true
date: 2011-10-01 08:37:57
layout: post
slug: '%e5%9c%a8linux%e4%b8%8b%e7%bc%96%e8%af%91qt-libs-%e8%83%bd%e4%b8%8d%e8%83%bd%e5%bf%ab%e7%82%b9%e5%91%a2'
title: 在Linux下编译Qt Libs -- 能不能快点呢?!
permalink: '/330.html'
wordpress_id: 330
categories:
- 工作
tags:
- Demo
- 下载
---

出于好玩,下载了Qt libs for linux, 发现是个源码包,好吧编译呗

敲入:

    ./configure -qt-sql-sqlite
    gmake
    gmake install
    
原本以为半个小时就万事,结果花了2个多小时!! 我吐血啊...

难道就不能快点吗? demo examples docs 能不能编译啊!!

好吧,谷哥告诉我,能这样:

    ./configure -qt-sql-sqlite -opensource -fast -no-qt3support -nomake demos -nomake docs -nomake examples -optimized-qmake -nomake tools
    gmake -j4
    gmake install
    
呵呵,半个小时搞定!!!
