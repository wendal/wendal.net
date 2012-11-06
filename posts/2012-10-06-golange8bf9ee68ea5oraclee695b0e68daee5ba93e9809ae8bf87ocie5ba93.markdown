---
comments: true
date: 2012-10-06 11:17:54
layout: post
slug: "access-oracle-in-golang-by-oci"
title: Golang连接Oracle数据库(通过OCI库)
permalink: '/459.html'
wordpress_id: 459
categories:
- go
tags:
- Ant
- git
- golang
- Oracle
- 下载
- 视频
---

这是我对mattn/go-oci8的一个fork [https://github.com/wendal/go-oci8](https://github.com/wendal/go-oci8)

在Linux下的安装,应该是没啥难度的了,唯独蛋疼的Windows需要介绍一下:


	//假设的GOPATH指向C:\gohome
	0. 执行 go get github.com/wendal/go-oci8 ,然后肯定是报错了,没关系,代码会下载下来.
	1. 首先,你需要安装mingw到C:\mingw
	2. 然后,到Oracle官网,下载OCI及其SDK,解压到instantclient_11_2  -- 当前最新版
	3. 从我的go-oci8库的windows文件夹,拷贝pkg-config.exe到C:\mingw\bin\,拷贝oci8.pc到C:\mingw\lib\pkg-config\
	4. 设置环境变量 PATH           ,值为     原有PATH;C:\instantclient_11_2;C:\mingw\bin;
	5. 设置环境变量 PKG_CONFIG_PATH,值为     C:\mingw\lib\pkg-config
	6. 接下来,就最重要的,就是再执行一次,这次应该能成功的:  go get github.com/wendal/go-oci8
	7. 测试一下:
   		cd %GOPATH%/src/github.com/wendal/go-oci8/example
   		go run oracle.go
   		#提醒一句, oracle.go里面的写的密码是system/123456, 实例名XE  


[视频演示](http://www.tudou.com/programs/view/yet9OngrV_4/)

[下载视频及编译环境](https://github.com/wendal/go-oci8/downloads)
