---
title: ssh tunnel without shell
date: '2012-12-10'
description:
categories: [linux]
tags : [vps]
---

IT人士,必备翻墙梯
===============

在VPS创建无权限的用户
-------------------

	useradd -s /bin/false free2
    passwd free2 #创建密码

在本地访问之
-----------

	ssh -D 0.0.0.0:7070 -N -C free2@nutz.cn
    #输入密码,就可以了

前提?当然是你有自己的VPS了
=======================