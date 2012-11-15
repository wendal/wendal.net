---
comments: true
date: 2012-02-10 20:31:19
layout: post
slug: '%e7%ba%afjava%e5%a4%84%e7%90%86cmyk%e6%a0%bc%e5%bc%8f%e7%9a%84jpeg%e6%96%87%e4%bb%b6'
title: 纯Java处理CMYK格式(32位色深)的JPEG文件!!
permalink: '/369.html'
wordpress_id: 369
categories:
- Java
- 工作
tags:
- bug
- io
- Java
- Nutz
- 集成
---

无限折腾之后,终于找到一个能用的纯Java的解决方法:
1. 获取jpeg格式的ImageReader 
2. 通过ImageIO.createImageInputStream生成ImageInputStream
3. ImageReader和ImageInputStream协作,产生Raster
4. 使用createJPEG4通过色彩空间变化,生成BufferedImage
5. 重要: 把BufferedImage保存为临时的jpg文件,然后重新解析为BufferedImage

这样,最后得到的BufferedImage,将是一个普通的24位色深的RGB的jpg文件所对应的BufferedImage

注意, 第4步所产生的BufferedImage,进行某些操作时会报错,因为色彩空间是ColorSpace.CS_sRGB,所以,先保存到临时jpg文件,然后再生成标准的BufferedImage是非常重要的.

具体实现已经集成到[Nutz的Images类](https://github.com/nutzam/nutz/blob/master/src/org/nutz/img/Images.java)

感谢:
Sun java [http://bugs.sun.com/bugdatabase/view_bug.do?bug_id=4799903](http://bugs.sun.com/bugdatabase/view_bug.do?bug_id=4799903)
dsmart-30buy
 [http://dsmart-30buy.iteye.com/blog/1226969](http://dsmart-30buy.iteye.com/blog/1226969)

