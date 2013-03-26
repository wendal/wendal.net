---
title: Reflect Of Golang
date: '2012-11-30'
description: 玩玩Golang的反射
categories: [go]
tags : [go,反射]
permalink: '/2012/1130.html'
---

正在重新实现mustache for golang
-----------------------------

* mustache作为模板引擎,必然需要用到反射啦
* 官方的实现,就一个go源文件,几百行,蛋碎
* 官方实现只返回一个string类型,出错就返回空字符串!!
* 原本打算在上面改的,结果还是算了,重新实现一个更好,现在已经完成60%
* [mustache from wendal](https://github.com/wendal/mustache)

访问Map
------

	// 这里的参数和返回值都用了reflect.Value,是因为这是最下层的实现
	// 在此之上,我们可以封装为 Get(_map interface{}, key string)之类的形式
	func GetMapValue(value reflect.Value, key reflect.Value) (rs reflect.Value) {
		//进行任何反射操作之前,判断其可用性很重要
		if !value.IsVaild() {
			return
		}
		//判断其真实类型
		//注意,这里的真实,是指最终的类型,例如
		// type AAA map[string]string, 那么仍会得到map,而非AAA
		// 如果想得到AAA,那么应该使用 value.Type().Name()或者全路径value.Type().String()
		if value.Type().Kind() != reflect.Map {
			return 
		}
		
		//reflect包很多方法都是针对具体类型的,不合乎就panic
		//例如MapIndex,如果value不是map,就直接panic了
		rs = value.MapIndex(key)
		return
	}

访问数组/切片
-----------

	// 在取值方法,数组和切片的规则是一样的,提供索引值即可
	func GetArrayValue(value reflect.Value, index int) (rs reflect.Value) {
		if !value.IsVaild() {
			return
		}
		if value.Type().Kind() != reflect.Array ||
			 value.Type().Kind() != reflect.Slice {
			return 
		}
		
		// value.Len()仅限于array和slice,map,string哦
		if 0 <= index && index < value.Len() {
			rs = value.Index(index)
		}
		return
	}
	
访问结构体及其指针
--------------

	// 这算是最复杂的了吧
	// 这里演示一下把T和*T当成Map用,呵呵
	// 也就是mustache模板引擎中Section节点
	func GetStructValue(value reflect.Value, key string) (rs reflect.Value) {
		if !value.IsVaild() {
			return
		}
		if value.Type().Kind() == reflect.Ptr {
			//value.Elem()可以得到指针所指向的对象
			if value.Elem().Kind() != reflect.Struct {
				return
			}
		} else if value.Type().Kind() == reflect.Struct {
			return
		}
		
		//好了,来取Struct的Field吧!
		
		//首先,我们把*T还原为T
		//如果本来就是Struct,那么只是简单返回而已
		//指针类型是不能获取Field的
		v := reflect.Indirect(ctx.value)
		field := v.FieldByName(key)
		if field.IsValid() { //字段存在时返回true
			rs = field
			return
		}
		
		//接下来,看看有米有对应的Method
		//注意,如果是*T,那么全部方法都能拿到
		//如果是T,那么只能获取那些非指针的方法哦
		//我也很纠结这个,尝试突破但没有成功
		t := value.Type()
		method, ok := t.MethodByName(key)
		if !ok { //没找到
			return
		}
		
		//输入的参数必须为1,也就是当前value,当然,如果你知道其他参数,也可以是传参的,也就一个数组嘛
		//输出的参数不为0就好了,我们只需要取第一个
		if method.Func.Type().NumIn() != 1 || method.Func.Type().NumOut() == 0 {
			return	
		}
		//调用之
		rs = method.Func.Call([]reflect.Value{value})[0] //最后的[0]就是取第一个返回值
		return
	}
	
总结一下
-------

* 在任何reflect的func调用前,判断IsVaild
* 判断具体类型,然后再调用相应的反射方法,不然分分钟会panic
* 如果传入的是T,那么是无法访问指针类的方法的