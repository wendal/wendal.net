---
title: 在android上运行io.js
date: '2015-05-28'
description: 在android上运行io.js
permalink: '/2015/05/28.html'
categories:
-其他
tags:
-iojs
---

无意中看到io.js在4月份开始已经支持android编译,果断弄一个

环境
-------------------------------------

opensuse 11.4 x86

android 4.2.2 linux 3.2, 已root 

下载并安装NDK r10e
-------------------------------------

http://dl.google.com/android/ndk/android-ndk-r10e-linux-x86.bin

直接下载挺慢的,走百度云就很快

在opensuse切换到root,然后运行

```
cd /opt
./android-ndk-r10e-linux-x86.bin
```

即可自行解压出ndk

下载并编译io.js
------------------------------------

```
cd /opt
# 或者百度云下载也可以,更快
wget https://iojs.org/dist/v2.1.0/iojs-v2.1.0.tar.gz
tar xf iojs-v2.1.0.tar.gz
cd iojs-v2.1.0
./android-configure /opt/android-ndk-r10e
make -j8
```

无异常,无Error,顺利编译完成

传输到android上
-----------------------------------

```
adb remount
adb push out/Release/iojs /system/bin/iojs
adb shell chmod 777 /system/bin/iojs
```

尝试运行一下
------------------------------------

先弄个官网demo, 存为example.js

```
var http = require('http');

http.createServer(function (request, response) {
  response.writeHead(200, {'Content-Type': 'text/plain'});
  response.end('Hello World\n');
}).listen(8124);

console.log('Server running at http://127.0.0.1:8124/');
```

传到android上

```
adb shell mkdir /sdcard/iojs/
adb push example.js /sdcard/iojs/
adb shell /system/bin/iojs /sdcard/iojs/example.js
```

访问一下android设备上的服务,哈哈

```
curl -v http://192.168.72.107:8124/
```

输出如下内容,搞定

```
*   Trying 192.168.72.107...
* Connected to 192.168.72.107 (192.168.72.107) port 8124 (#0)
> GET / HTTP/1.1
> Host: 192.168.72.107:8124
> User-Agent: curl/7.42.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Thu, 28 May 2015 02:13:18 GMT
< Connection: keep-alive
< Transfer-Encoding: chunked
< 
Hello World
* Connection #0 to host 192.168.72.107 left intact
```