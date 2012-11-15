---
comments: true
date: 2012-05-17 22:02:59
layout: post
slug: 'call-newlisp-in-luajit-by-ffi'
title: 在LuaJIT中通过FFI直接调用newlisp
permalink: '/424.html'
wordpress_id: 424
categories:
- lisp
tags:
- io
- lua
---

1. 首先,当然是编译newlisp,并拷贝到/usr/lib/libnewlisp.so
2. 编译luajit,启动之

上代码

    --载入ffi
    ffi = require("ffi")
    --载入newlisp
    newlisp = ffi.load("newlisp")
    --定义newlisp的公开API
    ffi.cdef[[
    char * newlispEvalStr(char * cmd);
    ]]
    
    --接下来,就是调用过程了
    newlisp_str = "()"
    tmp = ffi.new("char[2]") -- 因为newlispEvalStr的参数是char*,而newlisp_str是string,需要转一下
    ffi.copy(tmp, newlisp_str)
    
    --执行
    newlisp.newlispEvalStr(tmp)
    
    ---------------------------------------------
    --------------封装一下,做个库------------------
    ---------------------------------------------
    
    function newlisp(newlisp_str)
        local ffi = require("ffi")
        local newlisp = ffi.load("newlisp")
        ffi.cdef[[
           char * newlispEvalStr(char * cmd);
        ]]
        local tmp = ffi.new("char[" .. #newlisp_str .. "]")
        ffi.copy(tmp, newlisp_str)
        newlisp.newlispEvalStr(tmp)
    end
    
TODO: 改为[lua-newlisp](https://github.com/wendal/lua-newlisp)形式的调用
