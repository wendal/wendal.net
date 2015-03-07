---
title: Http协议及Json格式简介(给yeelink使用者的)
date: '2014-10-23'
permalink: '/2014/10/23.html'
categories:
- 其他
tags:
- arduino
---

高手,老鸟,请路过. 本说明忽略了Continue 100请求,代理,chunked等等高级话题.

Http协议基本:
----------------

分3部分, 请求行(request line), 头部键值对(header key-value), 请求体(body), 其中请求体是可选的, 尤其是GET/PUT请求

简单的GET请求

```
GET /v1.1/device/12825/sensor/20956/datapoints HTTP/1.1\r\n
Host: api.yeelink.net\r\n
Content-Length: 0\r\n
U-ApiKey: 121234132432143\r\n
\r\n
```

简单的POST请求

```
POST /v1.1/device/12825/sensor/20956/datapoints HTTP/1.1\r\n
Host: api.yeelink.net\r\n
Content-Length: 14\r\n
U-ApiKey: 121234132432143\r\n
\r\n
{"value":30.1}
```

请求行
------------------------

```
POST /v1.1/device/12825/sensor/20956/datapoints HTTP/1.1\r\n
```

格式为

```
$method $uri HTTP/1.1\r\n
```

其中:

*	$method是请求方法,可以是GET,POST,PUT,DELETE 等
*	$uri是请求的路径, 例如网址是 http://wendal.net/404.html, 那么$uri就是 /404.html
*	HTTP/1.1 是固定字符,为http协议版本,可以用HTTP/1.1或HTTP/1.0
*	\r\n 换行,标记请求行的结束

头部键值对(header key-value)
-----------------------------

```
Host: api.yeelink.net\r\n
Content-Length: 14\r\n
\r\n
```

注意, 这里总共3行,最后一个空行是headers的结束标志,必须有

header格式

```
$key: $value\r\n
```

其中:

*	$key 是"键", 例如Host代表主机名的键, Content-Length代表请求体的长度的键
*	$value 是"值", 例如主机名api.yeelink.net, 0等等
*	\r\n 换行,代表一个header的结束

header可以写很多很多行,但必须在所有header写完之后, 写入一个空行\r\n

有些header的值是严格限定的, 例如 Content-Length, 必须是请求体(body)的总字节数,不然服务器很有可能会拒绝.

请求体(body)
-----------------------

```
{"value":30.1}
```

约束:

*	请求体可以是一个字符串,一个数字,一个图片,一个压缩包... http协议本身并不限制body里面的格式
*	这部分对于GET/DELETE请求是不允许有的, 但对于POST/PUT,大部分情况下是必须的(不带body的POST请求在协议层面也是合法的)
*	请求体,在header的空行之后算起, 总长度需要填入heaader的Content-Length键值对.
*	在yeelink中, 非图片型传感器的上传数据, 是json字符串, 那么json字符串的字节长度,必须填入heaader的Content-Length键值对.
	而图片型传感器, 请求体是图片的二进制数据(别转成hex字符串了), 总字节数一样要heaader的Content-Length键值对.
*	再强调一次, Content-Length算的是字节长度, 是header空行之后的总字节数!!

json格式简介
--------------------------

json官网 http://json.org 里面有中文文档

json的基本格式是

```
{
 "value" : 31.0
}
```

上述数据以{开头, }结束, 代表一个键值对.

*	其中的键为 value, 必须用双引号包起来
*	这里演示的值是一个数值, 所以不需要双引号

值也是字符串的时候, 那么也需要字符串包起来:
```
{
 "value" : "I am ok"
}
```


代码怎么写?
---------------------------------

这里假设用了透传工具(wifi,gprs,网线,等等), 实际使用时,请删掉中文注释

一个POST的例子(arduino代码, 其他单片机就自行选用Serial.print等价的方法,例如printf)

```
// 首先,写入请求行
Serial.print("POST /v1.1/device/12825/sensor/20956/datapoints HTTP/1.1\r\n");
// 然后,写入headers
// 先写入域名
Serial.print("Host: api.yeelink.net\r\n");
// 再写入密钥
Serial.print("U-ApiKey: 121234132432143\r\n");
// 接着写入请求体的长度
Serial.print("Content-Length: ");
Serial.print("14");//这需要算好,算对哦, 下面写入的请求体是 {"value:"30.1}
Serial.print("\r\n");//别忘记换行了
// 必要的headers都写完了,其他都不写了, 作为headers结束,必须有个空行
Serial.print("\r\n");
// 接下来是请求体, 
Serial.print("{\"value\":"); // 这里的\"是转义,算一个字节.
Serial.print("30.1"); // 假设上传的是数值型数据, 值为30.1
Serial.print("}");// 注意看, 需要匹配哦, 键值对.
// 已经结束写入了,不需要再换行之类的操作,因为这是请求体, 只有服务器认识里面的内容,不是http协议约束的部分

delay(3000); //等待3秒, 网络畅顺的话几十毫秒就返回了

while (Serial.available()) {
	int c = Serial.read();
	// 千万别写 Serial.write(c)
	// 这时候我们需要把读出来的数据写入到另外一个串口(例如声明个软串口), 这样才不会透传模块冲突, 也方便调试.
	MySoftSerial.write(c);
}

delay(10000); // yeelink的上传间隔是10s哦.

```











