---
title: 用NDK编译ssdb
date: '2014-05-13'
permalink: '/2014/05/13.html'
categories:
- 工作
tags:
- android
- ssdb
---

首先,需要准备一下环境
------------------------

ubuntu 14.04 x86 版 -- 当前最新啦,哈哈

android ndk r9b     -- 为啥用这个版本?因为我机器上有...

编译环境准备
----------------------

把ndk转为独立工具链, 这是今天获取的新技能,呵呵,以后不写Android.mk也能编译部分软件了.

```
/opt/android-ndk-r9b/build/tools/make-standalone-toolchain.sh --install-dir=/opt/ndk-armv7 --arch=armv7
```

下载源码,暂时从我fork出来的库里面取吧,不知道ideawu会不会合并这个修改, who knows ...

```
cd /opt/
wget https://github.com/wendal/ssdb/archive/android2.zip
unzip android2.zip
cd ssdb-android2
```

执行编译
-----------------------

需要指定3个变量 CXX CC TARGET_OS, 以使用对应的编译器及源码

CC 指定c源文件的编译器

CXX 指定cpp源文件的编译器

TARGET_OS 是ssdb的build.sh和leveldb构建脚本中,判断系统环境的属性,默认是通过uname -s获取,需要指定为OS_ANDROID_CROSSCOMPILE

```
CC=/opt/ndk-armv7/bin/arm-linux-androideabi-gcc CXX=/opt/ndk-armv7/bin/arm-linux-androideabi-g++ TARGET_OS=OS_ANDROID_CROSSCOMPILE make
```

与原版的差异
---------------------

1. 因为android下没有pthread_cancel, 所以, 修改了thread.h里面的WorkPool::stop方法,意味着线程池不可关闭.由于线程池仅在ssdb退出时关闭,所以没大影响
2. 没有使用jemalloc

性能测试
--------------------

测试环境, 某全志A10开发板, 单核1G内存, 8G nand

cpu参数

```
root@android:/ # cat /proc/cpuinfo
Processor       : ARMv7 Processor rev 2 (v7l)
BogoMIPS        : 238.54
Features        : swp half thumb fastmult vfp edsp neon vfpv3
CPU implementer : 0x41
CPU architecture: 7
CPU variant     : 0x3
CPU part        : 0xc08
CPU revision    : 2

Hardware        : sun4i
Revision        : 0000
Serial          : 07c1521b5654484880778253162367cb
root@android:/ #
```

内存状态

```
root@android:/ # free -m
             total         used         free       shared      buffers
Mem:           797          428          369            0            6
-/+ buffers:                421          376
Swap:            0            0            0
```

将ssdb-server和ssdb.conf传输到/dev/shm/ 内存目录下,开始执行

```
cd /dev/shm/
mkdir var
chmod 777 ssdb-server
./ssdb-server ssdb.conf
```

显示效果, 也许有人会问为啥数据文件夹也在/dev/shm下,那是因为, 放/sdcard跟/dev/shm测试结果没多少差别.

```
root@android:/dev/shm # ./ssdb-server ssdb.conf
ssdb 1.6.8.6
Copyright (c) 2012-2014 ideawu.com

```

跑ssdb-bench, 提醒一下,测试结果极其飘忽, qps从500到2500都出现过.

```
Copyright (c) 2013-2014 ideawu.com

========== set ==========
qps: 1196, time: 8.360 s
========== get ==========
qps: 1704, time: 5.866 s
========== del ==========
qps: 1268, time: 7.883 s
========== hset ==========
qps: 897, time: 11.141 s
========== hget ==========
qps: 772, time: 12.944 s
========== hdel ==========
qps: 635, time: 15.733 s
========== zset ==========
qps: 1251, time: 7.992 s
========== zget ==========
qps: 1584, time: 6.311 s
========== zdel ==========
qps: 1631, time: 6.129 s
========== qpush ==========
qps: 1263, time: 7.912 s
========== qpop ==========
qps: 1463, time: 6.831 s
```

用途
----------

Cubieboard可以跑ssdb了, 安卓机能跑ssdb了... 其他场景自行补充...