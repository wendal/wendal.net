---
comments: true
date: 2010-10-12 22:16:53
layout: post
slug: '%e4%b8%banutlab%e6%b7%bb%e5%8a%a0%e4%b8%80%e4%b8%aa%e6%96%b0%e9%a1%b9%e7%9b%ae-nutz%e4%b8%8eopenid%e9%9b%86%e6%88%90'
title: 为NutLab添加一个新项目-Nutz与OpenID集成
wordpress_id: 94
categories:
- Java
tags:
- bug
- Demo
- io
- Nutz
- OpenID
- 配置
- 集成
---

忙乎了两天, 终于把项目正确运行起来了.

使用即将发布的Nutz 1.a.32 ,加 JOpenID 1.0.7

地址: [](http://code.google.com/p/nutzlab/source/browse/#svn/trunk/NutOpenID)

同时,发现JOpenID 1.0.7的一个Bug.
JOpenID 默认使用UTF-8来对参数进行getBytes. 当参数中还有非英文字符,且没有配置URIEncoding时,获取的byte[]是错误的.
具体代码:

    
    
    package org.expressme.openid;
    //......
    public class OpenIdManager {
    //......
    
    
        String getHmacSha1(String data, byte[] key) {
            SecretKeySpec signingKey = new SecretKeySpec(key, HMAC_SHA1_ALGORITHM);
            Mac mac = null;
            try {
                mac = Mac.getInstance(HMAC_SHA1_ALGORITHM);
                mac.init(signingKey);
            }
            catch(NoSuchAlgorithmException e) {
                throw new OpenIdException(e);
            }
            catch(InvalidKeyException e) {
                throw new OpenIdException(e);
            }
            try {
                byte[] rawHmac = mac.doFinal(data.getBytes("UTF-8")); //不一定,也许是ASCII
                return Base64.encodeBytes(rawHmac);
            }
            catch(IllegalStateException e) {
                throw new OpenIdException(e);
            }
            catch(UnsupportedEncodingException e) {
                throw new OpenIdException(e);
            }
        }
    
    //.....
    }
    


项目地址: [http://code.google.com/p/nutzlab/source/browse/#svn/trunk/NutOpenID](http://code.google.com/p/nutzlab/source/browse/#svn/trunk/NutOpenID)
