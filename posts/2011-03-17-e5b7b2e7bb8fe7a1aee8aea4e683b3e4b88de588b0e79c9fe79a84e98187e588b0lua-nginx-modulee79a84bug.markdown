---
comments: true
date: 2011-03-17 15:32:15
layout: post
slug: '%e5%b7%b2%e7%bb%8f%e7%a1%ae%e8%ae%a4%e6%83%b3%e4%b8%8d%e5%88%b0%e7%9c%9f%e7%9a%84%e9%81%87%e5%88%b0lua-nginx-module%e7%9a%84bug'
title: '[已经确认]想不到真的遇到Lua-nginx-module的bug'
permalink: '/255.html'
wordpress_id: 255
categories:
- Nginx
tags:
- bug
- el
- io
- lua
- nginx
- 配置
---

难以置信啊, agentzh大牛已经确认是个bug,对应的版本:
Lua-nginx-module 0.16rc2 + Nginx 0.8.54

期待修复,哈哈

原本一直以为是自己代码的原因,因为实在很诡异,附上我的原文:

    
    
    我得到的最小集合就是:
    1. nginx.conf
    location /test {
        root html;
        content_by_lua_file "conf/test.lua";
    }
    server/http等配置按默认的, event module用的是 epoll
    
    2. test.lua文件仅2行:
    ngx.location.capture('/1.html')
    
    ngx.exec("/1.html")
    
    我尝试过,无论这两句话是否请求同一个文件,结果都一样.
    
    3. 1.html文件里面仅有几个字母,我已经试过不同的文件大小,结果一样
    
    
    我遇到的情况是这样的:
    1. 通过wget/curl/Firefox来访问 localhost/test 都能正常显示1.html中的内容
    2. 使用ab访问 localhost/1.html是正常的,能够pass
    3. 使用ab进行测试,总是timeout , 我使用的语句是 ab -v 5 localhost/test 
    Benchmarking localhost (be patient)...INFO: POST header ==
    ---
    GET /down2 HTTP/1.0
    Host: localhost
    User-Agent: ApacheBench/2.3
    Accept: */*
    
    
    ---
    LOG: header received:
    HTTP/1.1 200 OK
    Server: nginx/0.8.54
    Date: Tue, 15 Mar 2011 07:22:30 GMT
    Content-Type: text/plain
    Content-Length: 4
    
    Last-Modified: Fri, 11 Mar 2011 09:32:28 GMT
    Connection: close
    Accept-Ranges: bytes
    
    ABC
    
    
    LOG: Response code = 200
    apr_poll: The timeout specified has expired (70007)
    
    其中的ABC就是1.html的内容, 非常抱歉我之前写错了.
    
    对于0A0D的描述,仅仅是我的猜测,请无视之.
    
    单独写 ngx.exec("/1.html") 也是能够通过ab测试的.
    
    环境:
    Ubuntu 10.10
    Luajit-5.1-dev
    Nginx 0.8.54
    Lua-nginx-module 0.16rc2
    
    Thanks,
    Wendal Chen
    


agentzh的回应是:

    
    
    I've reproduced it on my side. This is indeed a bug. When ngx.exec()
    is used after ngx.location.capture() or ngx.location.capture_multi(),
    nginx 0.8.11+ will not close the client connection due to leaked
    request reference counter (r->main->count). A hacky work-around is to
    disable nginx http keepalive and rely on the browser (and other http
    clients) to actively close the connection. And that's why wget, curl,
    firefox, and other well-written http clients worked for you.
    
    Nginx 0.7.68 (and older) is confirmed to work in this context just
    because older nginx does not use reference counting.
    
    I'll attempt fix in the next few days. Thank you for reporting this
    and sorry about this issue :)
    
    Cheers,
    -agentzh
    


大概的意思是: 
0.8.11+才会有这个问题,0.7.68以下的版本,因为没有使用相关特性而没有问题.
解决方法: 关闭nginx的keeplive 且 客户端主动关闭连接.

我得出几个结论:
1. taobao不是用这种版本组合,或者不用这种写法
**2. ab不会主动释放连接, 与wget/curl的行为不一致!! 这个比较问题大**
3. 淘宝里面如果没有这种写法,那我要检讨一下是否应该这样写

恩,继续努力,奋斗在 nginx+lua+C指针 的苦海中
