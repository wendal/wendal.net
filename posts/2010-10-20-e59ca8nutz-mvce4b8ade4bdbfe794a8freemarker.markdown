---
comments: true
date: 2010-10-20 12:27:06
layout: post
slug: '%e5%9c%a8nutz-mvc%e4%b8%ad%e4%bd%bf%e7%94%a8freemarker'
title: 在Nutz MVC中使用Freemarker
permalink: '/100.html'
wordpress_id: 100
categories:
- Java
tags:
- el
- Freemarker
- MVC
- Nutz
- XML
- 视图
- 配置
- 集成
---

大约一年前, axhack 发布了一篇文章 "<a href="http://axhack.javaeye.com/blog/542441">给 nutz 添加 freemarker 视图</a>" ,描述了如何集成nutz和freemarker.

前几天,我使用另外一种更简单的方法来实现(基于Nutz 1.a.33版新增的内部重定向视图),我使用的是Freemarker 2.3.16
首先, 在web.xml添加Freemarker官方文档描述的FreemarkerServlet, <a href="http://freemarker.sourceforge.net/docs/pgui_misc_servlet.html">查看原文描述</a>.

`
<servlet>
  <servlet-name>freemarker</servlet-name>
  <servlet-class>freemarker.ext.servlet.FreemarkerServlet</servlet-class>
  <init-param>
    <param-name>TemplatePath</param-name>
    <param-value>/</param-value>
  </init-param>
  <init-param>
    <param-name>NoCache</param-name>
    <param-value>true</param-value>
  </init-param>
  <init-param>
    <param-name>ContentType</param-name>
    <param-value>text/html; charset=UTF-8</param-value>
    <!-- 我觉得不需要了,如果是内部重定向的话, nutz已经设置了编码 -->
  </init>
  <init-param>
    <param-name>template_update_delay</param-name>
    <param-value>0</param-value><!-- 开发时才设置为0 -->
  </init-param>
  <init-param>
    <param-name>default_encoding</param-name>
    <param-value>UTF-8</param-value> <!-- 模板文件的编码 -->
  </init-param>
  <init-param>
    <param-name>number_format</param-name>
    <param-value>0.##########</param-value>
  </init-param>

  <load -on-startup>1</load>

<servlet-mapping>
  <servlet-name>freemarker</servlet-name>
  <url-pattern>*.ftl</url-pattern>
</servlet-mapping>
`

然后在需要Freemarker渲染的方法上,添加:


	@Ok("->:/forum/viewTip.ftl")

注意 1.a.33 才有内部重定向视图(->), 之前的版本,建议使用重定向视图(>>)代替
ftl后缀,就是web.xml配置的后缀.

这样,当方法正确返回时,就会使用Freemarker渲染, 返回值保存在 obj 变量中,你可以直接在Freemarker模板中调用.