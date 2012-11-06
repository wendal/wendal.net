---
comments: true
date: 2011-04-01 21:04:15
layout: post
slug: '%e5%8f%91%e5%b8%83%e4%b8%80%e4%b8%aahttpsqs%e7%9a%84python%e5%ae%a2%e6%88%b7%e7%ab%af'
title: 发布一个Httpsqs的Python客户端
permalink: '/261.html'
wordpress_id: 261
categories:
- VPS/Linux
- 工作
tags:
- bug
- el
- httpsqs
- io
- Java
- js
- python
- 下载
---

Httpsqs是张宴的一款开源队列服务器,项目首页 [http://code.google.com/p/httpsqs/](http://code.google.com/p/httpsqs/)

这款软件有几种客户端, Java/Perl/C,却没有Python的客户端.

故,本人奉上一个实现,欢迎指正!!

直接去下载 [httpsqs-python-client-v1.zip](http://wendal-common.googlecode.com/files/httpsqs-python-client-v1.zip) 

代码:

    
    
    #Verion 1.0
    #Author wendal(wendal1985@gmail.com)
    #If you find a bug, pls mail me
    
    import sys,httplib
    
    ERROR = 'HTTPSQS_ERROR'
    
    GET_END = 'HTTPSQS_GET_END'
    
    PUT_OK = 'HTTPSQS_PUT_OK'
    PUT_ERROR = 'HTTPSQS_PUT_ERROR'
    PUT_END = 'HTTPSQS_PUT_END'
    
    RESET_OK = 'HTTPSQS_RESET_OK'
    RESET_ERROR = 'HTTPSQS_RESET_ERROR'
    
    MAXQUEUE_OK = 'HTTPSQS_MAXQUEUE_OK'
    MAXQUEUE_CANCEL = 'HTTPSQS_MAXQUEUE_CANCEL'
    
    SYNCTIME_OK = 'HTTPSQS_SYNCTIME_OK'
    SYNCTIME_CANCEL = 'HTTPSQS_SYNCTIME_CANCEL'
    
    class Httpsqs(object):
        def __init__(self,host,port=1218):
            self.host = host
            self.port = port
        
        def get(self,poolName):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=get&name;=" + poolName)
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                conn.close()
                return data
            return ''
    
        def put(self,poolName,data):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("POST", "/?opt=put&name;="+poolName,data)
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
    
        def status(self,poolName):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=status&name;="+poolName)
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
        
        def status_json(self,poolName):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=status_json&name;="+poolName)
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
    
        def reset(self,poolName):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=reset&name;="+poolName)
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
    
        def maxlen(self,poolName,num):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=maxqueue&name;="+poolName+"&num;="+str(num))
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
    
        def synctime(self,poolName,num):
            conn = httplib.HTTPConnection(self.host,self.port)
            conn.request("GET", "/?opt=synctime&name;="+poolName+"&num;="+str(num))
            r = conn.getresponse()
            if r.status == httplib.OK :
                data = r.read()
                return data
            return ''
    
    def isOK(data):
        if data is '' :
            return False
        if data is ERROR :
            return False
        if data is GET_END :
            return False
        if data is PUT_ERROR :
            return False
        if data is RESET_ERROR :
            return False
        if data is MAXQUEUE_CANCEL :
            return False
        if data is SYNCTIME_CANCEL :
            return False
        return True
    


测试代码就不贴,需要的话就下载zip包吧.

因为Httpsqs本身就是基于Http协议的,故各种客户端实现都只是封装一下,本python客户端也不例外.
