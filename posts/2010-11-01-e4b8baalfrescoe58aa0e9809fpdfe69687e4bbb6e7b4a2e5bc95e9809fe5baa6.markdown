---
comments: true
date: 2010-11-01 17:14:41
layout: post
slug: '%e4%b8%baalfresco%e5%8a%a0%e9%80%9fpdf%e6%96%87%e4%bb%b6%e7%b4%a2%e5%bc%95%e9%80%9f%e5%ba%a6'
title: 为Alfresco加速PDF文件索引速度
permalink: '/114.html'
wordpress_id: 114
categories:
- Alfresco
- Java
tags:
- alfresco
- io
- Java
- PDF
- XML
- 索引
---

本文仅为 http://thinkalfresco.blogspot.com/2009/03/speeding-up-pdf-indexing-alfresco-hack.html 的 Alfresco 3.2版. 因为原文中的代码,已经无法在3.2版上使用.

新的代码:

    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE beans PUBLIC "-//SPRING//DTD BEAN//EN" "http://www.springframework.org/dtd/spring-beans.dtd">
    <beans>
            
      <bean id="transformer.PdfBox" class="java.lang.String"></bean>
            
      <bean id="transformer.complex.OpenOffice.PdfBox" class="java.lang.String"></bean>
    
      <bean id="transformer.PdfToTextTool" class="org.alfresco.repo.content.transform.RuntimeExecutableContentTransformerWorker">
                    <property name="mimetypeService">
             		<ref bean="mimetypeService"></ref>
          		</property>
                    <property name="transformCommand">
                            <bean name="transformer.pdftotext.Command" class="org.alfresco.util.exec.RuntimeExec">
                                    <property name="commandMap">
                                            <map>
                                                    <entry key=".*">
                                                            <value>/usr/bin/pdftotext -enc UTF-8 ${options} ${source} ${target}</value>
                                                    </entry>
                                            </map>
                                    </property>
                                    <property name="defaultProperties">
                                            <props><prop key="options"></prop></props>
                                    </property>
                            </bean>
                    </property>
                    
                    <property name="checkCommand">
                            <bean name="transformer.pdftotext.checkCommand" class="org.alfresco.util.exec.RuntimeExec">
                                    <property name="commandMap">
                                            <map>
                                                    <entry key=".*">
                                                            <value>chmod 777 /usr/bin/pdftotext</value>
                                                    </entry>
                                            </map>
                                    </property>
                                    <property name="defaultProperties">
                                            <props><prop key="options"></prop></props>
                                    </property>
                            </bean>
                    </property>
                    <property name="explicitTransformations">
             <list>
                <bean class="org.alfresco.repo.content.transform.ExplictTransformationDetails">
                   <property name="sourceMimetype">
                      <value>application/pdf</value>
                   </property>
                   <property name="targetMimetype">
                      <value>text/plain</value>
                   </property>
                </bean>
             </list>
          </property>
    </bean>
    
    <bean id="transformer.complex.OpenOffice.PdfToTextTool" parent="baseContentTransformer" class="org.alfresco.repo.content.transform.ComplexContentTransformer">
          <property name="transformers">
             <list>
                <ref bean="transformer.OpenOffice"></ref>
                <bean class="org.alfresco.repo.content.transform.ProxyContentTransformer">
                  	<property name="worker">
             		<ref bean="transformer.PdfToTextTool"></ref>
          		</property>
                </bean>
             </list>
          </property>
          <property name="intermediateMimetypes">
             <list>
                <value>application/pdf</value>
             </list>
          </property>
       </bean>
    </beans>
    
