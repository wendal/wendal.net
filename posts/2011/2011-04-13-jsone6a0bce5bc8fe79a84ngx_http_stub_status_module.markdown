---
comments: true
date: 2011-04-13 07:07:47
layout: post
slug: json%e6%a0%bc%e5%bc%8f%e7%9a%84ngx_http_stub_status_module
title: Json格式的ngx_http_stub_status_module
permalink: '/267.html'
wordpress_id: 267
categories:
- Nginx
tags:
- io
- js
- nginx
- 下载
---

原版的ngx_http_stub_status_module,查看到的状态信息是一个无格式文本,非常不好解析.

网上有另外一些实现,例如[nginx-json-status-module](https://github.com/drmingdrmer/nginx-json-status-module),但实现得极其垃圾,竟然直接调用malloc. 另外为了庆贺一下Nginx 1.0.0版本的发布,我做了一些修改,使其以Json格式返回.

上代码:

    /*
     * Copyright (C) Igor Sysoev
     * Modify by Wendal Chen
     */
    
    #include <ngx_config.h>
    #include <ngx_core.h>
    #include <ngx_http.h>
    
    static char *ngx_http_set_status(ngx_conf_t *cf, ngx_command_t *cmd,
                                     void *conf);
    
    static ngx_command_t  ngx_http_status_commands[] = {
    
        { ngx_string("stub_status"),
          NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_CONF_FLAG,
          ngx_http_set_status,
          0,
          0,
          NULL },
    
          ngx_null_command
    };
    
    static ngx_http_module_t  ngx_http_stub_status_module_ctx = {
        NULL,                                  /* preconfiguration */
        NULL,                                  /* postconfiguration */
    
        NULL,                                  /* create main configuration */
        NULL,                                  /* init main configuration */
    
        NULL,                                  /* create server configuration */
        NULL,                                  /* merge server configuration */
    
        NULL,                                  /* create location configuration */
        NULL                                   /* merge location configuration */
    };
    
    ngx_module_t  ngx_http_stub_status_module = {
        NGX_MODULE_V1,
        &ngx;_http_stub_status_module_ctx,      /* module context */
        ngx_http_status_commands,              /* module directives */
        NGX_HTTP_MODULE,                       /* module type */
        NULL,                                  /* init master */
        NULL,                                  /* init module */
        NULL,                                  /* init process */
        NULL,                                  /* init thread */
        NULL,                                  /* exit thread */
        NULL,                                  /* exit process */
        NULL,                                  /* exit master */
        NGX_MODULE_V1_PADDING
    };
    
    static ngx_int_t ngx_http_status_handler(ngx_http_request_t *r)
    {
        size_t             size;
        ngx_int_t          rc;
        ngx_buf_t         *b;
        ngx_chain_t        out;
        ngx_atomic_int_t   ap, hn, ac, rq, rd, wr;
    
        if (r->method != NGX_HTTP_GET && r->method != NGX_HTTP_HEAD) {
            return NGX_HTTP_NOT_ALLOWED;
        }
    
        rc = ngx_http_discard_request_body(r);
    
        if (rc != NGX_OK) {
            return rc;
        }
    
        ngx_str_set(&r-;>headers_out.content_type, "text/plain");
    
        if (r->method == NGX_HTTP_HEAD) {
            r->headers_out.status = NGX_HTTP_OK;
    
            rc = ngx_http_send_header(r);
    
            if (rc == NGX_ERROR || rc > NGX_OK || r->header_only) {
                return rc;
            }
        }
    
        size = sizeof("{accepted:,handled:,active:,requests:,reading:,writing:}") + 7 * NGX_ATOMIC_T_LEN;
    
        b = ngx_create_temp_buf(r->pool, size);
        if (b == NULL) {
            return NGX_HTTP_INTERNAL_SERVER_ERROR;
        }
    
        out.buf = b;
        out.next = NULL;
    
        ap = *ngx_stat_accepted;
        hn = *ngx_stat_handled;
        ac = *ngx_stat_active;
        rq = *ngx_stat_requests;
        rd = *ngx_stat_reading;
        wr = *ngx_stat_writing;
    
        b->last = ngx_sprintf(b->last, "{accepted:%uA,handled:%uA,active:%uA,requests:%uA,reading:%uA,writing:%uA}", 
                                                   ap,         hn,        ac,         rq,           rd,         wr);
    
        r->headers_out.status = NGX_HTTP_OK;
        r->headers_out.content_length_n = b->last - b->pos;
    
        b->last_buf = 1;
    
        rc = ngx_http_send_header(r);
    
        if (rc == NGX_ERROR || rc > NGX_OK || r->header_only) {
            return rc;
        }
    
        return ngx_http_output_filter(r, &out;);
    }
    
    static char *ngx_http_set_status(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
    {
        ngx_http_core_loc_conf_t  *clcf;
    
        clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
        clcf->handler = ngx_http_status_handler;
    
        return NGX_CONF_OK;
    }
    
使用方法
1. 直接替换原本的文件
2. 修改一下ngx_string("stub_status")为ngx_string("json_stub_status"),配合一下config文件一起作为额外模块进行编译:

    ngx_addon_name=ngx_http_json_stub_status_module
    HTTP_AUX_FILTER_MODULES="$HTTP_AUX_FILTER_MODULES ngx_http_json_stub_status_module"
    NGX_ADDON_SRCS="$NGX_ADDON_SRCS $ngx_addon_dir/ngx_http_json_stub_status_module.c"
    
[下载Zip包](https://docs.google.com/leaf?id=0B8hUXYDeoy_hMDJjYmUzNTktMTQyNi00MzdjLTk0YzYtMDcxYzY1NWU2MDA2)
