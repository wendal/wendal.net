---
date: 2013-06-25
layout: post
title: 在windows下编译OpenResty
permalink: '/2013/06/25.html'
categories:
- nginx
tags:
- nginx
- openresty
---

折腾了一天,终于解决了

首先,准备cygwin环境
===================

cygwin下载setup.exe,启动并开始安装,建议选163源或者日本的源,速度较快

需要的安装的包及其devel包: 
openssl zlib pcre

还有就是一些编译需要的工具:
gcc4 make perl lua  (不使用openresty内置的lua)

<img src="{{urls.media}}/2013/06/25/cygwin_install.jpg"></img>

<img src="{{urls.media}}/2013/06/25/cygwin_install2.jpg"></img>

下载openresty, [openresty官网](http://openresty.org/)

并解压到 C:\cygwin\tmp下

<img src="{{urls.media}}/2013/06/25/openresty_download.jpg"></img>

修正lua的C模块编译脚本
======================

共需要修正3个文件(其实就是3个模块),而且都是一样的修改. 版本号日新月异,自己搞定啦

打开 C:\cygwin\tmp\ngx_openresty-1.2.8.6\bundle\lua-cjson-1.0.3\Makefile, 加入 -llua5.1

<img src="{{urls.media}}/2013/06/25/cjson_fix.jpg"></img>

然后就是如法炮制,修正rds和redis处理模块

<img src="{{urls.media}}/2013/06/25/rds_fix.jpg"></img>

<img src="{{urls.media}}/2013/06/25/redis_fix.jpg"></img>



开始编译吧,童鞋们!
==================

启动cgywin

<img src="{{urls.media}}/2013/06/25/cygwin_startup.jpg"></img>

开始执行配置,注意,这里使用系统的lua,而非openresty内置的lua,原因就是cjson等模块会找不到内置的lua(配一下也可以,但麻烦)

```
cd /tmp/ngx_openresty-1.2.8.6

./configure --without-select_module --prefix=/opt/openresty --with-lua51=/usr
```

<img src="{{urls.media}}/2013/06/25/openresty_configure.jpg"></img>

开始编译(按你的实际情况设置并发数哦,不然很久很久的)

```
make -j8                  #8就是内核数,并行编译,按你的实际情况而定
```

<img src="{{urls.media}}/2013/06/25/openresty_make.jpg"></img>

编译完成

<img src="{{urls.media}}/2013/06/25/openresty_make_done.jpg"></img>

安装那点小事
============

如果你直接执行make install, 你会看到这些错误(也许?)

<img src="{{urls.media}}/2013/06/25/default_install_fail.jpg"></img>

这个我也纠结了一段时间,然后改成这样执行,注意是/opt2,而非原本的/opt

```
make install DSETDIR=/opt2
```

<img src="{{urls.media}}/2013/06/25/install_opt2.jpg"></img>

你以为完了?其实还没有,你需要把名字改回去

```
rm -fr /opt/openresty
mv /opt2/opt/openresty /opt/
```

先简单测试一下
=========

测试最基本的配置文件检查
--------------------------

```
/opt/openresty/nginx/sbin/nginx.exe -t
```

<img src="{{urls.media}}/2013/06/25/openresty_simple_test.jpg"></img>


然后就是测试核心的lua调用
-------------------------

打开nginx.conf文件,添加一个location

```
		location /lua/hi {
			content_by_lua 'ngx.say("LUA Here")' ;
		}
```

保存,启动nginx, 然后curl一下看看

```
/opt/openresty/nginx/sbin/nginx.exe
curl -v http://127.0.0.1/lua/hi
```

<img src="{{urls.media}}/2013/06/25/openresty_lua_hi.jpg"></img>

测试数据库连接
==============

数据库的resty.mysql库需要LuaBitOP库(汗,为啥openresty不包含?)

下载LuaBitOP库, [猛击下载地址](http://bitop.luajit.org/download.html),并解压到C:\cygwin\tmp下

惯例了,修正编译参数

<img src="{{urls.media}}/2013/06/25/bitop_fix.jpg"></img>

编译然后拷贝到openresty的lualib

```
cd /tmp/LuaBitOp-1.0.2
make
cp bit.so /opt/openresty/lualib
```

接下来,就是修改nginx.conf,加上官方的[测试例子](https://github.com/agentzh/lua-resty-mysql)了


```
		location /lua/mysql {
			content_by_lua '
			
            local mysql = require "resty.mysql"
            local db, err = mysql:new()
            if not db then
                ngx.say("failed to instantiate mysql: ", err)
                return
            end

            -- 省略1000字,自行到官网拷贝吧
			' ;
		}
```


让nginx重新加载配置,然后访问之

```
/opt/openresty/nginx/sbin/nginx.exe -s reload
curl -v http://127.0.0.1/lua/mysql
```

哦也,上截图

<img src="{{urls.media}}/2013/06/25/openresty_mysql_test.jpg"></img>

打完收工!! 但有啥不足呢?
========================

没有luajit,反正我没弄出来, configure阶段总是找不到库,不管了,windows就不那么追求性能了