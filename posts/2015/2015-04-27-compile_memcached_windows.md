---
title: 在windows上编译memcached v1.4.24 (用cygwin)
date: '2015-04-27'
description: 在windows上编译memcached v1.4.24 (用cygwin)
permalink: '/2015/04/27.html'
categories:
- 其他
tags:
- memcache
---

编译环境
===========================

cygwin x86 当前最新
win7 x64 sp1

成品的依赖关系(ldd输出)
==========================

```
$ ldd /usr/local/bin/memcached
        ntdll.dll => /cygdrive/c/Windows/SysWOW64/ntdll.dll (0x77780000)
        kernel32.dll => /cygdrive/c/Windows/syswow64/kernel32.dll (0x76fa0000)
        KERNELBASE.dll => /cygdrive/c/Windows/syswow64/KERNELBASE.dll (0x76b30000)
        ADVAPI32.DLL => /cygdrive/c/Windows/syswow64/ADVAPI32.DLL (0x75770000)
        msvcrt.dll => /cygdrive/c/Windows/syswow64/msvcrt.dll (0x75810000)
        sechost.dll => /cygdrive/c/Windows/SysWOW64/sechost.dll (0x769e0000)
        RPCRT4.dll => /cygdrive/c/Windows/syswow64/RPCRT4.dll (0x770d0000)
        SspiCli.dll => /cygdrive/c/Windows/syswow64/SspiCli.dll (0x75110000)
        CRYPTBASE.dll => /cygdrive/c/Windows/syswow64/CRYPTBASE.dll (0x75100000)
        cygwin1.dll => /usr/bin/cygwin1.dll (0x61000000)
        cyggcc_s-1.dll => /usr/bin/cyggcc_s-1.dll (0x6fdb0000)
        cygevent-2-0-5.dll => /usr/local/bin/cygevent-2-0-5.dll (0x63ec0000)
```

可以看到依赖了libevent

编译libevent
===========================

```
cd /tmp
wget https://sourceforge.net/projects/levent/files/libevent/libevent-2.0/libevent-2.0.22-stable.tar.gz
tar xf libevent-2.0.22-stable.tar.gz
cd libevent-2.0.22-stable
./configure --prefix=/usr/local
make all
make  install
```

全程无异常通过

编译memcached
==========================

```
cd /tmp
wget wget http://memcached.org/latest
tar xf latest
cd memcached-1.4.24/
chmod 777 configure
./configure
```

configure 执行完毕后,需要修改Makefile,不然编译会失败

大概是326行,删掉-Werror, 结果如下

```
CFLAGS = -g -O2 -pthread -pthread -Wall -pedantic -Wmissing-prototypes -Wmissing-declarations -Wredundant-decls
```

继续执行剩余的编译

```
make
make install
```

完成,启动一下
===========================

```
memcached -vv
```

输出一堆log,然后用telnet访问一下,正常,搞定.