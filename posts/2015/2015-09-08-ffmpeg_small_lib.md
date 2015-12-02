---
title: ffmpeg编译一个仅带h264解码功能的库
date: '2015-09-08'
permalink: '/2015/09/08.html'
description: ffmpeg编译参数
categories:
- Linux
tags:
- ffmpeg
---

用的是ffmpeg当前最新的2.7.2

支持软解和vdpau硬解
支持解码文件和rstp

足够了, 静态链接之后,strip之后,目标程序小于4mb,压缩后不到2mb

一如既往上代码:

```
./configure --enable-nonfree --enable-vdpau --enable-gpl --enable-static \
	--disable-everything --enable-decoder=h264 --enable-decoder=aac  --enable-decoder=h264_vdpau \
	--prefix=/home/wendal/build --enable-parser=aac --enable-parser=h264 --enable-protocol=rstp \
	--enable-demuxer=h264 --enable-demuxer=aac --enable-vdpau --enable-protocol=file --enable-outdevs
```

因为不需要压缩h264, 所以无需x264

因为ffmpeg已经内置aac, 所以不需要额外添加
