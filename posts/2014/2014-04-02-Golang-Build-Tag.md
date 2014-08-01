---
title: 为Golang程序打上编译标记
date: '2014-04-02'
permalink: '/2014/04/02.html'
categories:
- 其他
tags:
- golang
---

参考文章: http://stackoverflow.com/questions/11354518/golang-application-auto-build-versioning

昨天在查询怎么生成一个小体积的golang程序的时候,无意中发现这个文章.

对于固定的代码,及固定的golang版本,下面的命令总是得到一模一样的程序

```
go build
```

有时候需要为每个编译都打上标记,不然真的很乱啊

演示用的golang代码

```
package main

var _VERSION_ = "unknown"

func main() {
	print("http_su ver=" + _VERSION_ + "\n")
}
```

编译时,加入需要的版本号信息,而不是直接去改main.go的源码

```
export TAG=v1.b.50
go build -ldflags "-X main._VERSION_ '$TAG'"
```

运行结果:

```
> go build
> ./demo
http_su ver=unknown
> export TAG=v.1.b.50
> go build -ldflags "-X main._VERSION_ '$TAG'"
> ./demo
http_su ver=v.1.b.50
```

可以看到, 版本号根据编译参数的变化而变化了. 关键点是, 必须是 $package.$varName

本demo在linux/macos/windows/arm下测试通过.