---
date: 2013-02-05
layout: post
title: 获取goroutine的id
permalink: '/2013/0205.html'
categories:
- go
tags:
- go
---

获取goroutine的id? 官方不支持的!!
---------------------------------

人家官方说了:

	"This, among other reasons, to prevent programmers 
	for simulating thread local storage using the goroutine id as a key. "
	
就为了避免咱们当成ThreadLocal的key!! 这是为了神马?为神马?!!

方法还是有的嘛,改动一下源码
---------------------------

神马?!改源码这么大件事?! 对的,但只是添加,不修改不覆盖,不影响其他功能

文件一, $GOPATH/src/pkg/runtime/runtime.c, 在最后面添加一个方法

```
void
runtime·GetGoId(int32 ret)
{
        ret = g->goid;
        USED(&ret);
}
```

文件二, $GOPATH/src/pkg/runtime/extern.go 在最后面导出这个方法

```
func GetGoId() int
```

然后,就是重新编译golang了

```
cd $GOROOT/src
./make.bash
```

好了,测试一下吧
---------------

写一个main.go

```
package main

import "fmt"
import "runtime"

func main() {
        fmt.Println("Id =", runtime.GetGoId())
}
```

编译并运行之

```
go run main.go
```

结果是

```
Id = 1
```

恩,暴力到此结束哦,保重啦各位...
==============================