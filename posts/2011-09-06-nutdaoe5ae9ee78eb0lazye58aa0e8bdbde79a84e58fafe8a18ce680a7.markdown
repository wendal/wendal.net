---
comments: true
date: 2011-09-06 23:33:16
layout: post
slug: nutdao%e5%ae%9e%e7%8e%b0lazy%e5%8a%a0%e8%bd%bd%e7%9a%84%e5%8f%af%e8%a1%8c%e6%80%a7
title: NutDao实现Lazy加载的可行性
permalink: '/317.html'
wordpress_id: 317
categories:
- Java
tags:
- el
- Hibernate
- io
---

不知道多少人问过@Many @One等注释能否实现懒加载或者自动加载

从代码上说,NutDao实现的是不加载,你需要的时候自行调用 -- 其实也就是有点不方便,外加效率较低

我这里讨论的是NutDao做到真正懒加载的可能性与实现方式
焦点集中在NutDao.query方法,因为大部分DB-->Pojo都是走这个方法的,fetch方法其实就是封装了一下query方法

由于实现的是懒加载,不可避免的使用Aop,以拦截@Many/@One的getter
使用Aop会导致一个问题,就是query方法返回的对象的类改变了,变成用户请求的类的一个定制化的子类的对象,这本身就导致一些问题:
    1. 使用query出来的对象进行update/delete时,由于其并非用户期望的类(实际上是其子类),导致getEntity需要改造,因为@Table等注解是不被继承的,@Column也是,这样,写在@Many/@One字段上的信息会被隐藏掉
    2. 用户代码可能会需要修改,当然,都是那些不好的写法,例如Class对象进行equal比较

Aop实现的关键点:
    1. 获取Dao对象,并存为其一个属性的值
    2. 记录当前getter是否已经获取过,或者setter已经被执行过
    3. 跳过Hibernate式的事务问题,平等对待getter调用时的事务问题

可能的改造方法:
    1. Aop整个NutDao,对于query以外的方法,进行对象解包/类解包,以移除query方法带出Aoped类的影响
    2. 改造/继承AnnotationEntityMaker,让其生成对象时,使用Aop改造过的类

恩,应该OK了
