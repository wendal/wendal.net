---
comments: true
date: 2010-08-05 22:57:29
layout: post
slug: alfresco%e9%85%8d%e7%bd%ae-%e7%ac%ac%e5%9b%9b%e8%8a%82
title: Alfresco配置 — 第四节
permalink: '/35.html'
wordpress_id: 35
categories:
- Alfresco
- Java
tags:
- alfresco
- Ant
- io
- Java
- XML
- 下载
- 部署
- 配置
---

# 讲讲如何在WebLogic上部署Alfresco 3.2, 以EAR方式

## 创建一个文件夹, 名为 WL_Alfresco

## 将alfresco.war解压到WL_Alfresco/alfresco.war

## 创建WL_Alfresco/META-INF, 放入两个文件: application.xml 和 weblogic-application.xml, 内容分别是:

		<application>
			<display -name>Alfresco</display>
			<description>Alfresco</description>
			<module>
				<web></web>
				<web-uri>alfresco.war</web-uri>
				<context-root>alfresco</context-root>
			</module>
		</application>

		<?xml version="1.0" encoding="UTF-8"?>
		<weblogic -application xmlns="http://www.bea.com/ns/weblogic/90">
			<prefer-application-packages>
				<package-name>org.mozilla.*</package-name>
				<package-name>antlr.*</package-name>
			</prefer-application-packages>
		</weblogic>
		
## 将以下jar放到JAVA_HOME/jre/lib/endorsed 文件夹内: serializer.jar xalan.jar , 这两个jar可以到apache上下载.

## 然后按标准的方法添加到WebLogic的部署中去即可.
