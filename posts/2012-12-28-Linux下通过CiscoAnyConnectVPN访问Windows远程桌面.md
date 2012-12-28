---
title: Linux下通过CiscoAnyConnectVPN访问Windows远程桌面
date: '2012-12-24'
description:
categories: [linux]
tag : [AnyConnect]
permalink: '/2012/1228.html'
---

Cisco的AnyConnect产品,好多大公司都在用,但貌似木有官方的Linux客户端(如果你有,请提供链接,无比感谢)

首先,需要一个开源的客户端openconnect
----------------------------------
	wget ftp://ftp.infradead.org/pub/openconnect/openconnect-4.07.tar.gz
	tar xvf openconnect-4.07.tar.gz
	cd openconnect-4.07
	./configure
	make
    
得到编译好的openconnect后, 连接服务器
--------------------------------------------------------------

	#root权限哦, 或者能添加tun的帐户也行
	./openconnect vpn.wendal.net
	#提示如下:
	Attempting to connect to 124.99.99.99:443
	SSL negotiation with vpn.wendal.net
	Connected to HTTPS on vpn.wendal.net
	GET https://vpn.wendal.net/
	Got HTTP response: HTTP/1.0 302 Object Moved
	SSL negotiation with vpn.wendal.net
	Connected to HTTPS on vpn.wendal.net
	GET https://vpn.wendal.net/+webvpn+/index.html
	Please enter your username and password.
	username: #输入帐户
	password: #输入密码

	#当然,你可以先指定user和password咯
	./openconnect -u wendal -p wendal vpn.wendal.net

登陆成功后, 查看本地地址
----------------------

	ifconfig tun0

接下来,就是远程桌面了
-------------------

	#安装rdesktop
	yum install -y rdesktop
	./rdesktop -z win.wendal.net
	#哈哈,你能看到界面了吗? 输入帐户密码就可以登陆了

	#-z是压缩参数
	#还可以指定用户名和密码实现自动登陆

看看成果
-------

<img src="{{urls.media}}/2012/12/vpn_remote_desktop.jpg">
	
