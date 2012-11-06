---
comments: true
date: 2012-03-26 14:41:59
layout: post
slug: 'difference-bettwen-while-1-and-while-True-in-python'
title: Python的while 1跟while True到底有什么区别?
permalink: '/405.html'
wordpress_id: 405
categories:
- Python
tags:
- io
- python
- 反编译
---

定义两个方法,分别使用while循环

    
    
    def w() :
      while 1 :
        pass
    
    def w2() :
      while True:
        pass
    



单从功能上说,两种无任何区别,那么,来看看字节码上的区别:

    
    
    import dis  #载入反编译模块,Python内置的
    
    dis.dis(w) #对应的是while 1,下面是输出
      2           0 SETUP_LOOP               3 (to 6)
    
      3     >>    3 JUMP_ABSOLUTE            3
            >>    6 LOAD_CONST               0 (None)
                  9 RETURN_VALUE
    
    dis.dis(w2) #对应的是while True,下面是输出
      2           0 SETUP_LOOP              10 (to 13)
            >>    3 LOAD_GLOBAL              0 (True)
                  6 POP_JUMP_IF_FALSE       12
    
      3           9 JUMP_ABSOLUTE            3
            >>   12 POP_BLOCK
            >>   13 LOAD_CONST               0 (None)
                 16 RETURN_VALUE
    



很明显, while 1的字节码只有while True的一半.
为什么呢? 因为Python2.x中True不是关键字,只是一个全局变量而已

更详细,更专业的分析,请看
[http://stackoverflow.com/questions/3815359/while-1-vs-for-whiletrue-why-is-there-a-difference](http://stackoverflow.com/questions/3815359/while-1-vs-for-whiletrue-why-is-there-a-difference)
