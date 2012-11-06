---
comments: true
date: 2011-01-06 10:17:32
layout: post
slug: alfresco-virtual-tomcat%e9%bb%98%e8%ae%a4%e4%b8%8d%e8%a7%a3%e6%9e%90el
title: Alfresco Virtual tomcat默认不解析EL
permalink: '/226.html'
wordpress_id: 226
categories:
- Alfresco
tags:
- alfresco
- bug
- el
- io
- js
- Tomcat
- XML
- 配置
---

我正在使用Alfresco 3.3 SP4 , 发现其Virtual tomcat竟然是构建在tomcat5.5上的.
在jsp页面中写el表达式,死活不出来.最后才发现,是tomcat5.5默认不解析EL !!

查原因, 在tomcat5.5.31的源码中:
1. 类名: org.apache.jasper.compiler.JspConfig
2. 代码:

    
    
    if (webApp == null
         || !"2.4".equals(webApp.findAttribute("version"))) {//直接字符串判断,晕!
         defaultIsELIgnored = "true";
         return;
    }
    //如果version不是2.4,那根本不走这里!
    TreeNode jspConfig = webApp.findChild("jsp-config");
    if (jspConfig == null) {
         return;
    }
    



解决方法: :
第一种,在JSP页面头部加入指令:

    
    
    < %@ page isELIgnored="false" %>
    


第二种,在web.xml中设置version为2.4 :
第三种,在web.xml中添加一个设置(这解决方法依赖于第二种!!故基本上无效!!!),配置el-ignored为false!

严重怀疑这个是bug !!
