---
date: 2013-04-06
layout: post
title: Golang下通过syscall调用win32的api
permalink: '/2013/0406.html'
categories:
- go
tags:
- 低层调用
---

源于golang群中再次提到windows下获取磁盘空间的方法

由于golang的api并非完全跨平台, golang本身并没有直接提供windows下的方式

syscall.Syscall系列方法
-----------------------

当前共5个方法

```
syscall.Syscall
syscall.Syscall6
syscall.Syscall9
syscall.Syscall12
syscall.Syscall15
```

分别对应 3个/6个/9个/12个/15个参数或以下的调用

参数都形如

```
syscall.Syscall(trap, nargs, a1, a2, a3)
```

第二个参数, nargs 即参数的个数,一旦传错, 轻则调用失败,重者直接APPCARSH

多余的参数, 用0代替

调用示例
------------

获取磁盘空间

```
//首先,准备输入参数, GetDiskFreeSpaceEx需要4个参数, 可查MSDN
dir := "C:"
lpFreeBytesAvailable := int64(0) //注意类型需要跟API的类型相符
lpTotalNumberOfBytes := int64(0)
lpTotalNumberOfFreeBytes := int64(0)

//获取方法的引用
kernel32, err := syscall.LoadLibrary("Kernel32.dll") // 严格来说需要加上 defer syscall.FreeLibrary(kernel32)
GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

//执行之. 因为有4个参数,故取Syscall6才能放得下. 最后2个参数,自然就是0了
r, _, errno := syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)
			
// 注意, errno并非error接口的, 不可能是nil
// 而且,根据MSDN的说明,返回值为0就fail, 不为0就是成功
if r != 0 {
	log.Printf("Free %dmb", lpTotalNumberOfFreeBytes/1024/1024)
}
```

简单点的方式? 用syscall.Call
----------------------------

跟Syscall系列一样, Call方法最多15个参数. 这里用来Must开头的方法, 如不存在,会panic.

```
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("F:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	if r2 != 0 {
		log.Println(r2, err, lpFreeBytesAvailable/1024/1024)
	}
```

小提示
------

传struct不是个好想法, 不同语言之间的差异不好磨合
