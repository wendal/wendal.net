---
date: 2013-08-28
layout: post
title: 让golang的get命令兼容gitlab
permalink: '/2013/08/28.html'
categories:
- golang
tags:
- linux
- git
---

我们有些什么呢?
================

Gitlib已经发布了6.0,号称是企业级的版本了,作为私有git库的首选,自然越来越多人用

假设 有这样一个golang 的库,URL是 http://git.wendal.net/wendal/gofly

如果尝试执行下面的语句去获取这个库的话

```
go get git.wendal.net/wendal/gofly

#会输出
package git.wendal.net/wendal/gofly: unrecognized import path "git.wendal.net/wendal/gofly"
```

然而,如果我们输入

```
go get git.wendal.net/wendal/gofly.git

#会输出
fatal: repository 'git.wendal.net/wendal/gofly' does not exist
package git.wendal.net/wendal/gofly.git: exit status 128
```

为什么呢?貌似go get不支持自定义的库地址啊(git的)
===============================================

且看 $GOROOT/src/cmd/go/vcs.go里面的一段代码

```
	// General syntax for any server.
	{
		re:   `^(?P<root>(?P<repo>([a-z0-9.\-]+\.)+[a-z0-9.\-]+(:[0-9]+)?/[A-Za-z0-9_.\-/]*?)\.(?P<vcs>bzr|git|hg|svn))(/[A-Za-z0-9_.\-]+)*$`,
		ping: true,
	},
```

可以看到, 对于未知的库地址(非github/Google Code/Bitbucket/Launchpad),都按这里的配置进行设置

按上述的正则表达式,输入git.wendal.net/wendal/gofly可以得到

```
root = git.wendal.net/wendal/gofly
repo = git.wendal.net/wendal/gofly
vcs  = git
```

What? 当使用git进行clone的时候,其实就执行了

```
git clone $repo $GOPATH/src/$root
#展开之后(假设GOPATH=/opt/gopath)
git clone git.wendal.net/wendal/gofly /opt/gopath/src/git.wendal.net/wendal/gofly
```

git库的地址当成本地路径了,不出错才怪呢

怎么解决呢? 添加自定义的库地址,跟github类似
===========================================

首先,拷贝一份github的配置

```
	// Github
	{
		prefix: "github.com/",
		re:     `^(?P<root>github\.com/[A-Za-z0-9_.\-]+/[A-Za-z0-9_.\-]+)(/[A-Za-z0-9_.\-]+)*$`,
		vcs:    "git",
		repo:   "https://{root}",
		check:  noVCSSuffix,
	},
```

改成

```
	// git.xwoods.org
	{
		prefix: "git.wendal.net/",
		re:   `^(?P<root>git\.wendal\.net/(?P<p>.[A-Za-z0-9_.\-]+/[A-Za-z0-9_.\-]+))(/[A-Za-z0-9_.\-]+)*))$`,
		ping:   false,
		repo:   "git@git.wendal.net:{p}.git",
		vcs :   "git",
	},
```

注意 re和repo,做了特别处理哦,多一个p变量, 这样repo就把凑成标准的ssh式git地址,自动使用密钥(哈哈,私有库嘛)

最后,还需要把golang编译

```
# linux/mac 下
$GOROOT/src/make.bash

#windows下
cd %GOROOT%\src\cmd\go
go install
```