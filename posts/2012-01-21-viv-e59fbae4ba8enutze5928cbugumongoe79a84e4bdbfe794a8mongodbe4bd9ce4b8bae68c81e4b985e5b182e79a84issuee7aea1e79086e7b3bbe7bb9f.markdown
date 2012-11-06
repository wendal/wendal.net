---
comments: true
date: 2012-01-21 15:44:09
layout: post
slug: viv-%e5%9f%ba%e4%ba%8enutz%e5%92%8cbugumongo%e7%9a%84%e4%bd%bf%e7%94%a8mongodb%e4%bd%9c%e4%b8%ba%e6%8c%81%e4%b9%85%e5%b1%82%e7%9a%84issue%e7%ae%a1%e7%90%86%e7%b3%bb%e7%bb%9f
title: Viv -- 基于Nutz和BuguMongo的,使用Mongodb作为持久层的Issue管理系统
permalink: '/360.html'
wordpress_id: 360
categories:
- Java
- mongodb
tags:
- bug
- git
- Java
- Javadoc
- js
- mongodb
- Nutz
---

项目地址: [https://github.com/wendal/viv](https://github.com/wendal/viv)

项目的核心思想,是基于Issue的Tag而非Issue的Status, 由zozoh提出来,我只是扩展并按照自己的项目设计了一套,跟github的issue系统类似.

终于赶在春节前,把主要功能点完成了(页面还没做,看看谁愿意帮忙弄一个,前后台通信用ajax/json)
用户登录,新增issue,上传附件,添加comments ... ...

**BuguMongo 项目主页** [http://code.google.com/p/bugumongo/](http://code.google.com/p/bugumongo/)
国产的小框架,自称"BuguMongo已在多个正式商业项目中使用，并取得了理想的效果。".我的使用感受是, 总体不错,但还很不成熟.
1. 代码中的瑕疵还是比较明显的,我个人比较在意的是出错时不打印堆栈,严重错误时也不抛出异常...
2. 文档还不够齐备,JavaDoc缺失严重

这代表着,作者自己用得非常爽,但社区的人会碰很多很多的钉子... 我对某些错误真是无语+无语
作为国内少有的mongodb框架,很希望它能持续发展下去.虽说Nutz的社区还不成熟,但发现BuguMongo的社区差不多等于0,基本上死寂 ... 跟当年nutz刚发布的时候差不多. 看来Nutz社区与BuguMongo的社区做一些交流还是很有必要的,呵呵


**回想自己参与开源事业这好几年,发现很多开源项目都需要突破一些坎:**

**1. 勇敢发布,并产生第一个使用者**
>> 很多人想,这不是很容易吗? 看上去是的,但只需要到googlecode/github逛逛,你就能发现,很多项目连一个版本都没发布过,就消亡了

**2. 收到第一个有用的Bug/Issue报告**
>> 有用户使用才能产生bug/issue报告,才能发现一些你没有考虑到的情况. 真正发起一个项目并得到第一个bug/issue报告,对很多开源项目来说,那是消亡之前都等不到的

**3. 开立社区(论坛,QQ群,等一切交流工具,并维持一定的人气)**
>> 维持一个论坛比一个QQ群累,但效果会比QQ群好,在我看来,论坛属于知识积累的一部分,持久化的, 而QQ是过程式,临时交流的结果.可惜nutz的论坛至今发展不起来,QQ群倒是非常热闹

**4. 找到第一个共同开发者**
>> 我已经深深感觉到找一个共同开发者对一个开源项目有多重要. 思想的碰撞,具体实现的差异,协助开发,等等,都非常有利于项目的延续.我自己也发起过,参与过不少的开源项目,很多很多,慢慢就变成个人项目,随之慢慢死去

各位,春节了哦,快回家吃饭吧@@@!!


