---
title: rm文件不再需要按y了,解脱啊
date: '2012-12-14'
description:
categories: [linux]
tags : [vps]
permalink: '/2012/1214.html'
---

终于解决了在某些linux下rm特定文件需要按y的问题
==============================================

原因就是在alias
--------------

	[root@MyVPS2923 ~]# alias
	alias cp='cp -i'
	alias l.='ls -d .* --color=tty'
	alias ll='ls -l --color=tty'
	alias ls='ls --color=tty'
	alias mv='mv -i'
	alias rm='rm -i'
	alias which='alias | /usr/bin/which --tty-only --read-alias --show-dot --show-tilde'

再找根源
-------

	[root@MyVPS2923 ~]# cat ~/.bashrc
	# .bashrc

	# User specific aliases and functions

	alias rm='rm -i'
	alias cp='cp -i'
	alias mv='mv -i'

	# Source global definitions
	if [ -f /etc/bashrc ]; then
        . /etc/bashrc
	fi

注释掉那3行alias,保存,重新登录, 哦也, 世界清静了!!
-----------------------------------------------
