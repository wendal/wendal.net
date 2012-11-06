---
comments: true
date: 2012-03-27 22:08:44
layout: post
slug: 'source-file-for-newLISP-ShellGame'
title: newLISP的ShellGame配套lsp源文件
permalink: '/406.html'
wordpress_id: 406
categories:
- lisp
tags:
- bug
- el
- git
- 视频
---

[newLISP](http://www.newlisp.org/)官方提供的[ShellGame](http://www.neglook.com/?Shell_Games)系列视频,并未提供lsp源文件,我看完了整个系列,顺便把源文件补了一部分,地址:
[https://github.com/wendal/learn-newLISP/tree/master/ShellGame](https://github.com/wendal/learn-newLISP/tree/master/ShellGame)

总共有23个视频,当前我完成了17个配套的lsp文件,剩下6个视频,没看懂,所以lsp文件还没写出来 -- 难道是比较高阶的?

附上第一个视频的lsp文件

    ;;视频原地址: http://www.neglook.com/
    ;;lsp文件由wendal创建 http://wendal.net
    
    ;; 用默认的context方法,创建数值生成器 -- 类似于自增
    
    (setq generator:acc 0)  ; 创建一个名叫generator的上下文,并添加一个符号acc
    
    (define (generator:generator) (inc generator:acc)) ;缺省的context方法(即与context同名的方法),其中的inc是方法,等同于++
    
    (generator)   ;;分号是单行注释的开始,而非语句的结束符
    
    (generator) 
    (generator)
    (generator)   ;;连续调用几次后,现在的值应该是acc应当等于4
    
    ;;费波那西数列（Fibonacci Sequence）
    
    (define (fibo:fibo) 
    	(if (not fibo:mem) 
    		(setq fibo:mem '(0 1))) 
    	(last (push (+ (fibo:mem -2) (fibo:mem -1)) fibo:mem -1)))
    
    (fibo)
    (fibo)
    (fibo)
    (fibo)
    (fibo)
    
    (println fibo:mem) ;;打印fibo上下文(context)中mem变量的值
    
    (exit) ;; 执行完毕,环境关闭
    
同时,我也发现一个bug,就是中文注释后的一行代码,会被无视...汗 ------- 已经提交bug report
