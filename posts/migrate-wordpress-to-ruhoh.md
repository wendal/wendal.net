---
title: 从WordPress迁移到Ruhoh的经验
date: '2012-11-07'
description:
categories: 'linux'
tags: 迁移
---


备份旧博客
----------

* 备份mysql数据库中的wordpress数据
* 备份wordpress所在的目录

迁移评论
--------

* 由于ruhoh是静态博客引擎,只能使用社会化评论系统了
* wordpress的评论转到DISQUS最为方便, 到DISQUS注册好,wordpress安装插件,等待导入完成即可

导出并转换旧博客的文章
----------------------

* 由于ruhoh只认markdown格式,需要将进行转换
* 在wordpress控制界面导出wordpress.xml
* 安装ruby,以Fedora 14为例

	yum install libxml2-devel libtool gcc gcc-c++ make curl autoconf automake readline-devel
	curl -L get.rvm.io | bash -s stable

* 安装Jekyll,因为要用到其转换脚本

	gem install jekyll

* 下载[转换脚本](https://gist.github.com/1394128),并执行

	wget https://gist.github.com/raw/1394128/cc8a3113c76ab51ea262da517db533e43e7e8c5c/wordpressdotcom.rb
	ruby wordpressdotcom.rb wordpress.xml /tmp/output/
	#少量文章会转换失败,记下来,需要手工导入

安装ruhoh和git,并测试一下是否可用
---------------------------------

	yum install git
	gem install ruhoh
	ruhoh help

建新家
------

* 建文件夹,拷贝已转的文章

	mkdir -p /home/web/
	cd /home/web
	ruhoh new wendal.net
	cp /tmp/output/*.xml wendal.net/posts/
	cd wendal.net

* 编译一下,看看是否正常

	ruhoh compile

修正wordpress permalink, 因为我原本的permalink是 /400.html
----------------------------------------------------------------------------

* #含wordpress_id的文章,自动插入 permalink: '/450.html' 一类的设置

	wget https://github.com/wendal/wendal.net/raw/master/tools/fix_wp_id_permalink.go	
	go run fix_wp_id_permalink.go posts/     #这是一个golang小脚本

* 清理空行, 因为我发现自动转换后的文章带很多空行, 所以又写了一个脚本clean一下

	wget https://github.com/wendal/wendal.net/raw/master/tools/cleanup_blank_line.go
	go run cleanup_blank_line.go posts/

* 好了,清理完毕,再编译一次吧

	ruhoh compile
	find compiled/ #可以看到老文章全部都变回/450.html形式的文件名

设置DISQUS和google分析的账号,然后做些小配置
-------------------------------------------

* DISQUS的ID

	vim widgets/comments/config.yml #填入你老博客的ID
	#当文章的路径跟原博客中的路径相同,DISQUS就能无缝还原之前的屏幕

* google分析的账号

	vim widgets/analytics/config.yml #然后填入你自己的ID

* 关闭代码高亮的行号显示

	vim widgets/google_prettify/config.yml #设置为false

* 修改首页,里面有些ruhoh的信息,删掉前面那部分即可

	vim pages/index.html 

手工fix文章中的图片链接
-----------------------

* 将图片/附件,导入新博客

	cp -r wp所在目录/wp-content/uploads/* media/

* 在文章中查找 http://博客域名/wp-content/uploads,替换为

	\{\{ urls.media}} 并做适当的修正

* 修正文章中的错误排版 -- 纯体力了

* 把导入失败的文章,按照[markdown语法](http://wowubuntu.com/markdown/),手工转换为新格式

* 再编译一次吧

	ruhoh compile

安装并配置nginx
---------------

* 安装nginx,当然了,我建议自行编译

	yum install nginx

* 修改nginx的配置文件,在将location / {} 替换为

	# 这里是原本的feed地址, ruhoh下叫做rss.xml,需要映射一下
	location = /feed {
		root   /home/web/wendal.net/compiled;
		rewrite /feed /rss.xml;
	}

	#一起的分类目录,转到categories页面 -- 貌似不能直接跳到具体分类,原因不明
	location /category/ {
		rewrite /category/(.+)/ /categories/#$1-ref permanent;
	}

	#之前的标签页,转到tags页面
	location /tag/ {
		rewrite /tag/(.+)/ /tags#$1-ref permanent;
	}

	#直接指向compiled目录,并启用gzip,因为全是静态文件
	location / {
		root   /home/web/wendal.net/compiled;
		gzip             on;
		gzip_min_length  1024;
		index  index.html index.htm;
		add_header Cache-Control "max-age=3600, must-revalidate";
	}

启动nginx并测试之
-----------------

   	/usr/local/nginx/sbin/nginx -t
   	/usr/local/nginx/sbin/nginx

   	#访问一下
   	curl -I http://127.0.0.1/

呵呵,你已经搞定了,用浏览器访问一下你的新博客吧!
-----------------------------------------------


