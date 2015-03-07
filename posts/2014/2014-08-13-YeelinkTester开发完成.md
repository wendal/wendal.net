---
title: YeelinkTester 总算弄好了
date: '2014-08-13'
permalink: '/2014/08/13.html'
categories:
- 其他
tags:
- arduino
---

直接上程序截图
---------------

<img src="{{urls.media}}/2014/08/13/ABC.png"></img>

源码:
-------

https://github.com/wendal/yeelink_tester

用到的东西
-----------

* python 2.7.8
* pyqt 4.11.1
* mqtt
* pyserial

功能
--------------

* 从串口读取数据,上传到yeelink
* 从yeelink接收mqtt通知,写到串口
* 从串口读取指令,然后会写数据(当前实现了查询最新的值)
* 一个API测试对话框, 支持文件上传
* 一个Yeelink Http API请求分析器(API虚拟服务器),用于排查api调用的错误
