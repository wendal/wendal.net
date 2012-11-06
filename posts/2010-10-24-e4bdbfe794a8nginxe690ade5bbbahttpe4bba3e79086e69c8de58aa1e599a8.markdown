---
comments: true
date: 2010-10-24 20:42:46
layout: post
slug: '%e4%bd%bf%e7%94%a8nginx%e6%90%ad%e5%bb%bahttp%e4%bb%a3%e7%90%86%e6%9c%8d%e5%8a%a1%e5%99%a8'
title: 使用Nginx搭建Http代理服务器
wordpress_id: 108
categories:
- Nginx
- VPS/Linux
tags:
- io
- nginx
- 代理
- 配置
---

昨天, 折腾了一个下午,终于配好了.
配置如下:

    
    
        server {
            listen       8888;
                    client_body_timeout 60000;
                    client_max_body_size 1024m;
                    send_timeout       60000;
                    client_header_buffer_size 16k;
                    large_client_header_buffers 4 64k;
    
                    proxy_headers_hash_bucket_size 1024;
                    proxy_headers_hash_max_size 4096;
                    proxy_read_timeout 60000;
                    proxy_send_timeout 60000;
                    
            location / {
                resolver 8.8.8.8;
                proxy_pass http://$http_host$request_uri;
            }
        }
    


resolver 8.8.8.8; 代表使用Google DNS来解析域名
client_body_timeout , large_client_header_buffers 等设置,确保大的请求不会返回400错误.

但,这个代理服务器只支持Http请求, Https会报400错误.

