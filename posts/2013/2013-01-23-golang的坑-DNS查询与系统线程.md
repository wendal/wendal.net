---
date: 2013-01-23
layout: post
title: Golang的坑之DNS查询与系统线程
permalink: '/2013/0123.html'
categories:
- go
tags:
- go
---

这个问题源之于我3个月前开发的一个缓存服务
-----------------------------------------

1. 代理Http请求
2. 缓存部分请求,减少外网访问

但是,在不稳定的网络环境下(例如3G网络),不时出现崩溃的情况
-------------------------------------------------------

纠结啊纠结啊...

1. 难道是因为我在32位系统下使用golang?
2. 难道是RP问题?

出错信息非常长(几百~上千个goroutine),就只贴头尾

```
runtime/cgo: pthread_create failed: Resource temporarily unavailable
SIGABRT: abort
PC=0xffffe424


goroutine 1 [chan receive]:
net.(*pollServer).WaitRead(0xd0170f0, 0xd02e850, 0xcfff5e0, 0xb)
	/opt/go/src/pkg/net/fd.go:268 +0x75
net.(*netFD).accept(0xd02e850, 0x80923d3, 0x0, 0xcfc0320, 0xcf7b178, ...)
	/opt/go/src/pkg/net/fd.go:622 +0x199
net.(*TCPListener).AcceptTCP(0xd030358, 0xcfc0d20, 0x0, 0x0)
	/opt/go/src/pkg/net/tcpsock_posix.go:320 +0x56
net.(*TCPListener).Accept(0xd030358, 0x0, 0x0, 0x0, 0x0, ...)
	/opt/go/src/pkg/net/tcpsock_posix.go:330 +0x39
net/http.(*Server).Serve(0xd0170c0, 0xcfff920, 0xd030358, 0x0, 0x0, ...)
	/opt/go/src/pkg/net/http/server.go:1029 +0x77
net/http.(*Server).ListenAndServe(0xd0170c0, 0xd0170c0, 0x40)
	/opt/go/src/pkg/net/http/server.go:1019 +0x9f
net/http.ListenAndServe(0x81f73c4, 0x5, 0xcfff840, 0xd0301f8, 0xd0301f8, ...)
	/opt/go/src/pkg/net/http/server.go:1091 +0x55
	
尾部
eax     0x0
ebx     0x231f
ecx     0x2357
edx     0x6
edi     0xb77ceff4
esi     0xb
ebp     0xafa112f8
esp     0xafa11050
eip     0xffffe424  //就这个,虽然不知道是什么,但很厉害的样子
eflags  0x202
cs      0x73
fs      0x0
gs      0x33
```

一直没找到原因

直至几天前,忽然灵机一动,难道是DNS的问题?
----------------------------------------

因为每次崩溃,总会带几个类似的goroutine

```
goroutine 1614 [syscall]:
net._C2func_getaddrinfo(0x87af730, 0x0)
	net/_obj/_cgo_defun.c:42 +0x32
net.cgoLookupIPCNAME(0xd20b36c, 0x1a, 0x0, 0x0, 0x0, ...)
	net/_obj/_cgo_gotypes.go:177 +0xe7
net.cgoLookupIP(0xd20b36c, 0x1a, 0x0, 0x0, 0x0, ...)
	net/_obj/_cgo_gotypes.go:223 +0x3d
net.cgoLookupHost(0xd20b36c, 0x1a, 0x0, 0x0, 0x0, ...)
	net/_obj/_cgo_gotypes.go:101 +0x43
net.lookupHost(0xd20b36c, 0x1a, 0x0, 0x0, 0x0, ...)
	/opt/go/src/pkg/net/lookup_unix.go:56 +0x3d
net.LookupHost(0xd20b36c, 0x1a, 0x0, 0x0, 0x0, ...)
	/opt/go/src/pkg/net/doc.go:10 +0x3d
net.hostPortToIP(0x81f6d58, 0x3, 0xd20b36c, 0x1f, 0x0, ...)
	/opt/go/src/pkg/net/ipsock.go:120 +0x183
net.ResolveTCPAddr(0x81f6d58, 0x3, 0xd20b36c, 0x1f, 0x0, ...)
	/opt/go/src/pkg/net/tcpsock.go:31 +0x37
net.resolveNetAddr(0x81f9c14, 0x4, 0x81f6d58, 0x3, 0xd20b36c, ...)
	/opt/go/src/pkg/net/dial.go:50 +0x35d
net._func_001(0xd4bf228, 0xd4bf230, 0xd4bf238, 0xd4bf240, 0x0, ...)
	/opt/go/src/pkg/net/dial.go:134 +0x44
created by net.DialTimeout
	/opt/go/src/pkg/net/dial.go:142 +0x13b
```

So,做了一个简单的DNS Cache,在执行net.Dail前,先自行解析域名

结果,在虚拟机上,模拟各种垃圾网络(断网,拔网线,拔路由器...),没有再出现崩溃

原因是什么呢?
---------------

看我10月份发送到讨论组的邮件 [runtime/cgo: pthread_create failed: Resource temporarily unavailable](https://groups.google.com/group/golang-china/browse_thread/thread/96e25b27abf9673b/98d271d98925fa98?lnk=gst&q=pthread_create#98d271d98925fa98)

待我解决问题后,马上有人解释原因了(汗... 为啥之前就没人回复呢? 问得太次?)

关键就是: *系统调用阻塞时大量生成内核级线程导致的*, 而cgo启用的情况下, 每一次DNS查询,都会起动一个系统线程!

我只能说,你妹啊!! 系统线程啊!! 难道就不能弄个线程池啊!! 再说,有pthread_cancel啊,为啥timeout不执行一下呢?!

写着写着, 忽然回想起之前遇到的一个情况,就是http.Get不返回(一直卡着,不往下执行),现在想起来,99%也是DNS的问题