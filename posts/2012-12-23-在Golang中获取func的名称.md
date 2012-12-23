---
title: 在Golang中获取func的名称
date: '2012-12-23'
description:
categories: [go]
tag : [go, reflect]
permalink: '/2012/1223.html'
---

这个问题源之于群友SeanWu的一个提问

期望的效果
---------

	func ABC() {
	}

	func GetFuncName(fn func()) string {
		return //返回ABC
	}

stackoverflow上的方法
---------------------

[链接](http://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go)

	
	func GetFunctionName(i interface{}) string {
		return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	}

分析
----

我曾经尝试这种写法

	runtime.FuncForPC(reflect.ValueOf(i).Addr()).Name()

但func是不能执行Addr()的, 而

	funcPc,_,_ := runtime.Caller(0)
	runtime.FuncForPC(funcPc).Name()

只能在func被调用时才能获取到自身的名字