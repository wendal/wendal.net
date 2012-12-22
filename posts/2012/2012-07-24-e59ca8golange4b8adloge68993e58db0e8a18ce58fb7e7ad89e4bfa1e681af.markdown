---
comments: true
date: 2012-07-24 19:11:08
layout: post
slug: 'show-linenumber-in-golang-log'
title: 在Golang中,Log打印行号等信息
permalink: '/446.html'
wordpress_id: 446
categories:
- go
tags:
- Dom
- golang
- 路径
---

做个小笔记, 默认情况下,log模块的只打印日期和时间, 没具体行号,比较不爽,嘿嘿

    package main
    /*
    #include <stdlib.h>
    */
    import "C"
    import "log"
    
    func main() {
        log.SetFlags(log.Lshortfile | log.LstdFlags)
        log.Println( C.random())
    }
    
打印结果:

    2012/07/24 19:27:55 X.cgo1.go:14: 1804289383
    
其中, log.Lshortfile 还可以设置为log.Llongfile 即完整文件路径

获取当前行数,文件名,函数名(方法名):

    package main
    
    import (
           "runtime"
           "fmt"
    )
    
    func main() {
            funcName, file, line, ok := runtime.Caller(0)
            if ok {
				fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
                fmt.Printf("file: %s    line=%d\n", file, line)
            }
    }
    
