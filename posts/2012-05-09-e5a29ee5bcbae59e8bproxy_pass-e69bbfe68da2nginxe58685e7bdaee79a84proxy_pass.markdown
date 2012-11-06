---
comments: true
date: 2012-05-09 10:33:34
layout: post
slug: 'enhand-proxypass-in-nginx'
title: 增强型Proxy_Pass - 替换nginx内置的proxy_pass
wordpress_id: 422
categories:
- Nginx
tags:
- el
- git
- io
- lua
- nginx
---

**项目地址: [https://github.com/wendal/lua-resty-http](https://github.com/wendal/lua-resty-http)**

这个项目是[https://github.com/liseen/lua-resty-http](https://github.com/liseen/lua-resty-http)的fork版本, 暂未被合并.

**nginx内置的proxy_pass有几个问题**:
1. 无法方便的调整后端host
2. 总是等待后端host把响应写完了,才开始向客户端写数据
3. proxy_next_upstream不灵活

**演示,实时proxy_pass,每读取1k就往浏览器写1k数据:**

    
    
    local url = 'http://'
    if ngx.var.http_host then
       url = url .. ngx.var.http_host 
    end
    url = url .. ngx.var.request_uri  -- 拼接完整的URL
    if ngx.var.args then
       url = url .. '?' .. ngx.var.args
    end
    local ok, code, headers, status, body  = hc:proxy_pass {
        url = url,
        fetch_size = 1024, -- 分段大小
        max_body_size = 100*1024*1024 ,  --响应体的最大大小.
        headers = ngx.req.get_headers(), -- 传递客户端的参数,可以根据需要进行修改哦.
        method = ngx.var.request_method, -- 真实还原客户端的请求方法,当然,你可以改!!
    }
    if not ok and not ngx.headers_sent then
        ngx.exit(502) -- 出错了哦? 这里只是简单遵循了nginx在后端报错时的响应,你完全可以实现自己的逻辑,进行错误处理
    else
        ngx.eof()
    end
    



核心扩展点,这是http.lua中的代码,我在这里附上中文注释:

    
    
    -- proxy_pass方法支持3种回调哦
    -- 提醒一句,回调里面,你可以调用任意ngx_lua的代码哦,就是说,你连ngx.exit(404)之类的中断请求的操作,也是完全可以的
    function proxy_pass(self, reqt)
        local nreqt = {}
        for i,v in pairs(reqt) do nreqt[i] = v end
    
        -- 响应回调,可以替代proxy_next_stream的功能哦,例如替换响应码,或者进行转向其他请求
        if not nreqt.code_callback then 
            nreqt.code_callback = function(code, ...)
                ngx.status = code
            end
        end
    
        -- header回调,可增减resp的header
        if not nreqt.header_callback then
            nreqt.header_callback = function (headers, ...)
                for i, v in pairs(headers) do
                    ngx.header[i] = v
                end
            end
        end
    
        -- body回调,注意chunked的情况哦
        if not nreqt.body_callback then
            nreqt.body_callback = function (data, chunked_header, ...)
                ngx.print(data)
                if chunked_header then
                    ngx.print('\r\n')
                end
            end
        end
        return request(self, nreqt)
    end
    



安装:
先编译[openresty](http://openresty.org/)
将lib/resty/url.lua和lib/resty/http.lua拷贝进openresty的lualib中

**注意: 回调的API尚未锁定, 将来可能根据需要添加更多参数**
