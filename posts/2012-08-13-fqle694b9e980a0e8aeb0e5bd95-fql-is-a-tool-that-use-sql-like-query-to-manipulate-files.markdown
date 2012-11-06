---
comments: true
date: 2012-08-13 16:52:41
layout: post
slug: 'fql-use-sql-like-query-to-manipulate-files'
title: fql改造记录 -- fql is a tool that use SQL like query to manipulate files
permalink: '/450.html'
wordpress_id: 450
categories:
- VPS/Linux
tags:
- bug
- el
- git
---

fql is a tool that use SQL like query to manipulate files

挺好玩的一个小工具, 用SQL语法来find文件,官网 <a href="https://github.com/dccmx/fql" title="fql is a tool that use SQL like query to manipulate files." target="_blank">https://github.com/dccmx/fql</a>

1. 添加readline支持, 原本是极其简单的fgets读取输入,那叫一个简陋啊


	//添加headers
	#if defined(HAVE_LIBREADLINE) && HAVE_LIBREADLINE==1
	# include <readline/readline.h>
	# include <readline/history.h>
	#endif

	//改造其获取输入的代码:
	#if defined(HAVE_LIBREADLINE) && HAVE_LIBREADLINE==1
      if (isatty(STDIN_FILENO)) { //如果是控制台输入,则输出提示符
          str = readline("> ");
          if( str && *str ) 
              add_history(str); //加入到readline历史记录
          else
              continue;
      } else {
          str = readline("");
          if (! str) break;
      }
	#else
      //老的,直接读取的方法, 不带历史记录,无法读取多行文本
      if (isatty(STDIN_FILENO)) printf("> ");
      char str[1024];
      fgets(str, 1024, stdin);
      if (feof(stdin)) break;
      if (!str || !strcmp("\n", str) || !strcmp("\r\n", str)) continue;
	#endif

2. 改为autoconf. 原项目是手写的Makefile,比较蛋疼(例如无法直接使用clang编译)

	autoscan
	mv configure.scan configure.in

	vim configure.in 
	#添加AM_INIT_AUTOMAKE,填上版本号,联系人等信息

	vim Makefile.am
	#写上fql_SOURCES bin_PROGRAMS等

	aclocal
	automake -a

	#搞定, 可以编译了
	./configure 
	make

3. 修正一个小bug--当文件夹或文件的uid或gid不合法时(指向一个不存在的用户),会发生段错误

    if (*ite == "uname") {
      struct passwd *pw = getpwuid(st.st_uid);
      if (pw)
          row.push_back(new String(pw->pw_name));
      else
          row.push_back(new String(""));
    } else if (*ite == "gname") {
      struct group *grp = getgrgid(st.st_gid);
      if (grp)
          row.push_back(new String(grp->gr_name));
      else
          row.push_back(new String(""));
    }


再fix掉clang编译时的一个小warning,搞定!

	CXX=clang CC=clang LDFLAGS="-Wall -lstdc++" ./configure
	make


最大的收获,就是知道原来平时命令行中,向上向下,查询历史命令,都是readline做的,一直以为是系统级的功能...