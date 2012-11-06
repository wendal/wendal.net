---
comments: true
date: 2012-03-14 21:23:40
layout: post
slug: 'call-setAttr-at-end-of-login-action'
title: 登录入口的一个小小的细节,HttpSession.setAttribute的调用顺序
wordpress_id: 399
categories:
- Java
tags:
- io
- js
- MVC
- Nutz
- SSI
- 安全
---

今天与zozoh的交谈中,又把这个问题提了一下. 就是登录代码中,Session处于一个不可靠的状态.

这里用NutzMVC的代码来演示,但并不代表是nutz的问题,而是J2EE都面对的问题(PS: 如果你使用了SpringSecurity等安全框架,则可能它已经帮你出来了这个问题)


    
    
    //登录入口方法
    public View login(String userName, String passwd, HttpSession session) {
        session.setAttribute("me", user); //当代码运行到这里,会话中已经存在me这个attr
        session.setAttribute("rule", rule);
        return new JspView("/index.jsp");
    }
    
    //登录后才能访问的入口方法
    @Filters(@By(type = CheckSession.class, args = {"me", "/"})) //已经HttpSession中已经包含me这个attr,所以,这个过滤器会判定为已经登录
    public Object xxx(HttpSession session) {
        Rule myRule = session.get("rule"); //由于已经被@Filters标注,所以方法会认为会话是已经完整初始化的
        // .... ... .. 
    }
    



也就是说, HttpSession在登录入口方法完成前,处于一个不可靠的状态 -- 对于检测是否登录的代码来说,这个会话是已经登录的,但事实上这个会话并未完成逻辑上的初始化.
当恶意进行类似的访问(多线程),那么其他入口方法要么报NPE,要么执行不可预测的逻辑.如果入口方法需要设置多个attr,那么HttpSession将处于多个不同的半初始化状态.

要避免这个问题,是将登录的标记,放在入口方法的最后.

    
    
    public View login(String userName, String passwd, HttpSession session) {
        session.setAttribute("rule", rule);
        session.setAttribute("me", user); //由于jsp中不包含逻辑,所以在这里设置登录的识别信息,完成会话初始化.
        return new JspView("/index.jsp");
    }
    



小小细节,轻者NPE,重者,谁知道呢,呵呵
