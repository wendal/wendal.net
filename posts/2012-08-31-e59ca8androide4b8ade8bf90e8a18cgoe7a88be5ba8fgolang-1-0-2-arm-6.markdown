---
comments: true
date: 2012-08-31 22:14:37
layout: post
slug: 'run-golang-in-android'
title: 在Android中运行go程序(Golang 1.0.2, ARM 6)
permalink: '/454.html'
wordpress_id: 454
categories:
- go
tags:
- Android
- el
- golang
- 下载
---

原本以为很简单的,网上一堆什么5g啊5l啊, 下载go 1.0.2才发现, 我去,根本就没有5g和5l







难道是官方编译版本没带而已,我自己编译一个呗



    
    
    apt-get install gcc libc6-dev ercurial
    #yum install gcc libc6-devel mercurial
    
    #预先把变量设置好
    export GOROOT=$HOME/go
    export PATH=$PATH:$GOROOT/bin
    
    #获取go的源码
    cd $HOME
    hg clone -u release https://code.google.com/p/go
    cd go/src
    ./all.bash
    
    #这样就安装好适合当前系统的go,但还需要arm(即Android的低层环境)的版本
    CGO_ENABLED=0 GOARCH=arm GOOS=linux ./make.bash
    
    #验证一下,应该会显示有5g和5l
    go tool
    







接下来,就是写个hello world,然后编译






    
    
    CGO_ENABLED=0 GOARCH=arm go build hi.go
    ./hi
    #呵呵,自己试试吧
    





提醒一句, 虽然可以通过变通的方式用上cgo,但据说不推荐,所以暂时还是不要用了



参考文章: [https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/ESQ0_yxH130](https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/ESQ0_yxH130)
