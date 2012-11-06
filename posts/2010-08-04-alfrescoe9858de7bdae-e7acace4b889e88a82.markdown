---
comments: true
date: 2010-08-04 22:14:25
layout: post
slug: alfresco%e9%85%8d%e7%bd%ae-%e7%ac%ac%e4%b8%89%e8%8a%82
title: Alfresco配置 — 第三节
wordpress_id: 33
categories:
- Alfresco
- Java
tags:
- alfresco
- el
- io
- Java
- JBoss
- jbpm
- js
- logger
- Wiki
- XML
- 工作
- 工作流
- 配置
---

今天打算说说Alfresco里面的工作流,这个是我折腾了好几个星期的事.

Alfresco的工作流分为3部分, Define/Model/UI

其中Define就是普通的jBpm3.2工作流引擎,你可以使用Alfresco提供的AlfrescoScriptAction直接调用Alfresco的服务,最基本的就是logger

Model,其实不单单是工作流的配置,它使用你定义的命名空间,声明需要用户输入的属性,和需要显示的属性

UI, 属于web-client-config.xml的自定义版本 web-client-config-custom.xml,用于定义各Model中的type如何显示在页面上.

具体的工作流我就不打算详细说了,也许以后会贴出示例.

说说几个技巧:

1. 如果无法用AlfrescoScriptAction解决你遇到的问题,请毫不犹豫地使用自定义的ActionHandler,不过最好继承JBPMSpringActionHandler,以便获取Alfresco相应服务的bean, 而且,bean的name一般就是接口的首字母小写,例如节点服务 nodeService, 用户服务 personService,操作服务actionService(可以创建mail action).

2. 流程意外停止,无法继续正常流下去,咋办? 使用 admin/workflow-console.jsp , 然后使用相关的命令对付该流程.例如signal, delete, cancel

3. 开发工作流时,请先确保工作流本身是正确的,然后再调试Model/UI

哈哈,差点忘了, UI还需要properties< 文件,用于定义页面上显示的label

有用的链接:

[http://wiki.alfresco.com/wiki/WorkflowAdministration

[http://wiki.alfresco.com/wiki/Workflow_Console](http://wiki.alfresco.com/wiki/Workflow_Console)

[http://docs.jboss.com/jbpm/v3/userguide/](http://docs.jboss.com/jbpm/v3/userguide/)

[http://wiki.alfresco.com/wiki/Data_Dictionary_Guide](http://wiki.alfresco.com/wiki/Data_Dictionary_Guide)
