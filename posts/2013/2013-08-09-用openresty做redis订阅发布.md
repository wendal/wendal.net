---
date: 2013-08-09
layout: post
title: 用openresty做redis订阅发布
permalink: '/2013/08/09.html'
categories:
- nginx
tags:
- linux
- openresty
- redis
---

openresty依然很强大,先上服务器端的代码, 定义一个通用的订阅接口

```
    location =/subscribe {
        send_timeout 365h; #很久很久很久,呵呵
        content_by_lua "
        local redis = require \"resty.redis\"
        local red = redis:new()
        red:connect(\"127.0.0.1\", 6379)
        red:subscribe(ngx.var.arg_key) -- 严格来说,需要判断空值和url decode
        local res, err = red:read_reply()
        if res then 
            ngx.say(res) -- 其实是个数组.. 第3个值才是publish的值
        else
            ngx.status = 500
        end
        red:close()
        ";
    }
```

客户端? 简单啦,Nutz一行代码搞定

```
String str = Sender.create(url).setTimeout(60*60*1000).send().getContent();
```