---
date: 2013-03-16
layout: post
title: Golang的坑之http读取大文件必须读完
permalink: '/2013/0316.html'
categories:
- go
tags:
- go
---

先上代码
--------

```
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://mirrors.ustc.edu.cn/opensuse/distribution/12.3/iso/openSUSE-12.3-GNOME-Live-i686.iso")
	if err != nil {
		panic(err)
	}
	fmt.Println("Resp code", resp.StatusCode)
	resp.Body.Close() // 注意,这里并不读取resp.Body, 而resp.Body有大概700mb未读取
}
```

你猜会怎样呢? 卡住了?! 

如果你的网速够快,你会发现, 几十秒之后, 程序自动退出了,但如果你很不幸是小水管,你会发现一直卡住...

原因是啥呢?
-----------

http包默认会重用连接,重用连接就需要先把前一个连接的数据读完

代码片段(net/http/transfer.go)
```
func (b *body) Close() error {
	if b.closed {
		return nil
	}
	defer func() {
		b.closed = true
	}()
	if b.hdr == nil && b.closing {
		return nil
	}
	
	if b.res != nil && b.res.requestBodyLimitHit {
		return nil
	}

	// 操,问题就在这了,读完整个body!!
	if _, err := io.Copy(ioutil.Discard, b); err != nil {
		return err
	}
	return nil
}
```

怎么解决呢?
-----------

按上面代码片段的逻辑, 需要提前返回nil,从而避免被读取

```
	// b.hdr 总是为nil,因为从不设置
	// 那b.closing什么时候为true呢?
	if b.hdr == nil && b.closing {
		return nil
	}
```

读源码可知, b.closing依赖于transferReader的Close值

而transferReader的Close值, 是根据shouldClose方法判断的

```
// 这里的header是resp的
func shouldClose(major, minor int, header Header) bool {
	if major < 1 {
		return true
	} else if major == 1 && minor == 0 {
		if !strings.Contains(strings.ToLower(header.Get("Connection")), "keep-alive") {
			return true
		}
		return false
	} else {
		// TODO: Should split on commas, toss surrounding white space,
		// and check each field.
		if strings.ToLower(header.Get("Connection")) == "close" {
			header.Del("Connection")
			return true
		}
	}
	return false
}
```

由于没法在这些代码之前修改resp的header,所以修改req的header,使服务器总是返回Connection: close

最终代码
--------

```
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://mirrors.ustc.edu.cn/opensuse/distribution/12.3/iso/openSUSE-12.3-GNOME-Live-i686.iso", nil)
	req.Header.Set("Connection", "close")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resp code", resp.StatusCode)
	resp.Body.Close()
}
```


一个月没写blog了, 心情欠佳+身体抱恙 ~_~ 哎,多事的3月
