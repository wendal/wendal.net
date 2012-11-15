---
comments: true
date: 2011-06-10 12:15:42
layout: post
slug: '%e5%9c%a8%e5%87%a0%e7%a7%8d%e7%bc%96%e7%a8%8b%e8%af%ad%e8%a8%80%e4%b8%ad%e6%8e%a2%e7%b4%a2javapythonluac'
title: 在几种编程语言中探索(Java/Python/Lua/C)
permalink: '/292.html'
wordpress_id: 292
categories:
- Java
- 未分类
tags:
- Android
- el
- Java
- js
- lua
- Nutz
- python
- 迁移
---

猛然发现自己已经快一个月没写blog,太懒惰了!!

5月,先是脚受伤,接着过敏... 

Java -- 为Nutz重写的Json解析器,最初的实现(zozoh)我看不懂,第二版(juqkai)改造得不够彻底,代码比较乱. 这次修改之后,解析将比之前严格,对于非法转义字符,不匹配的[}均直接抛出异常. 这将导致已经使用nutz的项目,其json文件可能会被报错. 然而,这并未100%的执行Json规范,主要是数字处理和map的key

Python -- 最近用得比较多,但至今未打动我. 很多情况下,Python依旧是Shell的替代品,呵呵. 不知道是不是习惯于一种思维: 编程语言=语法+API , 也许还没领会到Python的精髓与思维模式. 与Java相比, 多了些语法糖果,如in和缩进式的逻辑语句

Lua -- 呵呵,最近貌似很流行呢!! 最近玩的愤怒的小鸟,也是用Lua写的
成果: [RK29xx固件解包打包工具_v1.1测试版](http://www.teclast.com/bbs/forum.php?mod=viewthread&tid=157140&extra=) 当前唯一正确解包/打包RK2918方案固件的工具,完美解开台电T760的2.x固件,酷比1.x固件
Lua真的是一门神奇的语言, 很少的语法,极度精简的API!! 很值得学习的一门语言

C -- 处理字符串,真是痛苦啊...

刚刚把Nutz的代码迁移到Github了,准备写个简单教程,说说经验,呵呵
