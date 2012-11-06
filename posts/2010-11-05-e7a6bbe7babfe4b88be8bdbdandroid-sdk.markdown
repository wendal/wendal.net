---
comments: true
date: 2010-11-05 23:04:22
layout: post
slug: '%e7%a6%bb%e7%ba%bf%e4%b8%8b%e8%bd%bdandroid-sdk'
title: 离线下载Android SDK
wordpress_id: 118
categories:
- Java
tags:
- Android
- io
- Java
- logger
- Sunfarm
- XML
- 下载
- 新版本
---

我实在搞不懂,为啥GFW要墙Android 的开发网站.


今天,把当前最新版本的Android SDK全部下载到 [http://build.sunfarms.net/android/repository/](http://build.sunfarms.net/android/repository/)
同时发现一个问题,即使把 http://build.sunfarms.net/android/repository/repository.xml 添加到SDK Manager ,与Google官网一样的内容,却不认平台下载.
仔细找了一下,原来写死在代码里面的 [点击查看](http://code.google.com/p/android-sdk-tool/source/browse/src/main/java/com/m11n/android/AndroidSdkTool.java?r=6426c47fe356e9d649fe612464563960a1ca7d74)
没办法. 只好自己编译一个了,哈哈.
Google源码截取:

    
    
            private String repositoryUrl = "http://dl-ssl.google.com/android/repository/";
            private String sdkUrl = "http://dl.google.com/android/";
            private String downloadDir = System.getProperty("java.io.tmpdir") + "/";
            private Boolean overwrite = true;
            private Boolean verbose = true;
            private DocumentBuilder builder;
    
            public AndroidSdkTool() throws ParserConfigurationException {
                    builder = DocumentBuilderFactory.newInstance().newDocumentBuilder();
            }
    
            public Repository downloadRepository() {
                    String file = "repository.xml";
                    try {
                            download(repositoryUrl + file, downloadDir + file, overwrite);
                            return parse(new FileInputStream(new File(downloadDir+file)));
                    }catch (Exception e) {
                            logger.error(e.getMessage(), e);
                    }
                    return null;
            }
    
