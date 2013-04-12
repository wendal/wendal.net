---
date: 2013-04-12
layout: post
title: 在windows下编译nginx
permalink: '/2013/0412.html'
categories:
- nginx
tags:
- nginx
- 编译
---

又有人在windows下编译nginx -- 为什么那么多人喜欢自残呢?

官网教程
--------

[Building nginx on the Win32 platform with Visual C](http://nginx.org/en/docs/howto_build_on_win32.html)

本文基本上就是对着做,但需要对其进行微调 -- 不爽吗?咬我啊!!

准备工具
---------

系统: winxp sp3 32位, 例如你可以用个虚拟机什么的安装一个winxp

[MSYS-CN 2010-08-19 更新版本](https://msys-cn.googlecode.com/files/MSYS-Update.7z)

[zlib 1.2.7](http://zlib.net/zlib-1.2.7.tar.gz)

[pcre 8.32](http://nchc.dl.sourceforge.net/project/pcre/pcre/8.32/pcre-8.32.zip)

[openssl 1.0.1e](http://www.openssl.org/source/openssl-1.0.1e.tar.gz)

[VS2008 Express With SP1](http://download.microsoft.com/download/3/0/2/3025EAE6-2E15-4972-972A-F5B1ED248E85/VS2008ExpressWithSP1CHSX1504735.iso)

[Perl 5.16.3.1](http://strawberry-perl.googlecode.com/files/strawberry-perl-5.16.3.1-32bit.msi)

[Subversion](http://www.sliksvn.com/pub/Slik-Subversion-1.7.9-win32.msi) 或者你喜欢的svn客户端


安装必要的程序
---------------

VS2008,附带的sqlserver无需安装,也用不上

Perl,一直下一步即可

Slik-Subversion,一直就是下一步

将MSYS-CN解压到C盘

下载源码
--------

到C:\MSYS启动msys.bat,进入msys的bash

非常重要哦, 官网的tar包是不包含windows构建文件的!!

```
svn co svn://svn.nginx.org/nginx/tags/release-1.3.15
cd release-1.3.15
mkdir objs
mkdir objs/lib
```

然后, 把pcre/zlib/openssl的源码,均解压到C:\MSYS\home\UserName\release-1.3.15\objs\lib,
即上述语句所建立的文件夹,其中UserName是你的用户名.

生成构建脚本
-----------

依然在msys bash下.

```
cd release-1.3.15
auto/configure --with-cc=cl --builddir=objs --prefix= \
--conf-path=conf/nginx.conf --pid-path=logs/nginx.pid \
--http-log-path=logs/access.log --error-log-path=logs/error.log \
--sbin-path=nginx.exe --http-client-body-temp-path=temp/client_body_temp \
--http-proxy-temp-path=temp/proxy_temp \
--http-fastcgi-temp-path=temp/fastcgi_temp \
--with-cc-opt=-DFD_SETSIZE=1024 --with-pcre=objs/lib/pcre-8.32 \
--with-zlib=objs/lib/zlib-1.2.7 --with-openssl=objs/lib/openssl-1.0.1e \
--with-select_module --with-http_ssl_module --with-ipv6
```

语句比较长,可以写到build.bat中,然后执行 ./build.bat

编译
----

总有要编译啦,哇哈哈. 在开始菜单找VS2008的VS2008命令行,启动之

执行下面的语句

```
C:
cd \MSYS\home\UserName\release-1.3.15\
nmake -f objs/Makefile
```

你很快就会发现报错了,说找不到某某头文件. 

用你喜欢的编辑器打开 C:\MSYS\home\UserName\release-1.3.15\objs\lib\pcre-8.32\config.h

找到并注释掉(加//):

```
#ifndef HAVE_INTTYPES_H
#define HAVE_INTTYPES_H 1
#endif

#ifndef HAVE_STDINT_H
#define HAVE_STDINT_H 1
#endif
```

然后再执行就成功了:

```
nmake -f objs/Makefile
```

好吧,祝你好运!!
---------------

```
C:\MSYS\home\Administrator\release-1.3.15>nginx.exe -V
nginx version: nginx/1.3.15-http://wendal.net
TLS SNI support enabled
configure arguments: --with-cc=cl --builddir=objs --prefix= --conf-path=conf/ng
nx.conf --pid-path=logs/nginx.pid --http-log-path=logs/access.log --error-log-p
th=logs/error.log --sbin-path=nginx.exe --http-client-body-temp-path=temp/clien
_body_temp --http-proxy-temp-path=temp/proxy_temp --http-fastcgi-temp-path=temp
fastcgi_temp --with-cc-opt=-DFD_SETSIZE=1024 --with-pcre=objs/lib/pcre-8.32 --w
th-zlib=objs/lib/zlib-1.2.7 --with-openssl=objs/lib/openssl-1.0.1e --with-selec
_module --with-http_ssl_module
```