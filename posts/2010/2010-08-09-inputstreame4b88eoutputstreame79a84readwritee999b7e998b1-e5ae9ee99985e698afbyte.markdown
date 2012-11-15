---
comments: true
date: 2010-08-09 22:05:45
layout: post
slug: inputstream%e4%b8%8eoutputstream%e7%9a%84readwrite%e9%99%b7%e9%98%b1-%e5%ae%9e%e9%99%85%e6%98%afbyte
title: InputStream与OutputStream的read/write陷阱–实际是byte
permalink: '/41.html'
wordpress_id: 41
categories:
- Java
tags:
- io
- Java
- Javadoc
- 陷阱
---

先不要说这是标题党,我觉得这是很多人都已经在使用的误区,即误以为InputStream.read()返回的值真的是int,而OutputStream.write()接受的参数的确为int. 事实上,它们返回或接受的参数是 byte, 即一个字节, 务必仔细读清楚其JavaDoc, 明确说明是读出一个字节,而非int

我觉得这是Java核心API其中一个极度容易被人误用的地方,哈哈
    
    int read();
    void write(int data);
    
