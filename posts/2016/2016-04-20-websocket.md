---
title: nutz websocket nginx maven 配合使用
date: '2016-04-20'
permalink: '/2016/04/20.html'
description: Nutz,websocket,nginx,maven配合使用
categories:
- 工作
tags:
- nginx
- nutz
- websocket
---

## 起因

很久以前,用jetty玩过websocket,为了nutz.cn的自动提醒功能,又实践了一次websocket.

websocket需要服务器端,nginx,页面端,同时配合,才能工作.

## 页面端

```js
        var ws;
		var WS_URL = window.location.host+"${base}/yvr/topic/socket";
		if (location.protocol == 'http:') { // 需要特别注意的地方,根据http/https协议选不同的websocket前缀
			ws = new WebSocket("ws://"+WS_URL); // 普通http用ws,全称WebSocket
		} else {
			ws = new WebSocket("wss://"+WS_URL); // https环境下需要wss协议,全称WebSocket Secure
		}
		
		ws.onmessage = function(event) { // onmessage 特指服务器端发送消息过来
		    var re = JSON.parse(event.data);
		    _replies_count = re.count;
		    var n = new Notification(re.data, re.options); // 这里使用了Chrome Notification API
			n.onclick = function() {
				location.reload();
			};
		};
		window.setInterval(function(){ // 定时查询
			ws.send(JSON.stringify({id:'${obj.topic.id}',replies:_replies_count}));
		}, 5000);
```

## 服务器端

```java
// 省略import

@ServerEndpoint("/yvr/topic/socket")
public class YvrTopicWebSocket {
    
    private static final Log log = Logs.get();
    
    @OnMessage
    public void onMessage(String message, Session session) {
        if (yvrService == null) // 从全局ioc容器取出需要的ioc bean
            yvrService = Mvcs.ctx().getDefaultIoc().get(YvrService.class);
        try {
            NutMap map = Json.fromJson(NutMap.class, message); // 看页面端,发送过来的是json字符串
            String topicId = map.getString("id");
            int replies = map.getInt("replies");
            Object re = yvrService.check(topicId, replies);
            if (re instanceof Map) {
			    // 消息反馈回去, 会调用js中的onmessage方法
                session.getBasicRemote().sendText(Json.toJson(re));
            }
        }
        catch (Throwable e) {
            log.debug("message="+message, e);
        }
    }
    
    YvrService yvrService; // 注意,并非注入
}
```

## Maven配置

编译依赖:

```xml
	<dependencies>
		<dependency>
			<groupId>javax.websocket</groupId>
			<artifactId>javax.websocket-api</artifactId>
			<version>1.1</version>
		</dependency>
	</dependencies>
```

jetty插件需要的依赖

必须加websocket-server,否则不扫描websocket相关的注解

```xml
			<plugin>
				<groupId>org.eclipse.jetty</groupId>
				<artifactId>jetty-maven-plugin</artifactId>
				<version>9.3.8.v20160314</version>
				<configuration>
					<jvmArgs>-Dfile.encoding=UTF-8</jvmArgs>
				</configuration>
				<dependencies>
					<dependency>
						<groupId>org.eclipse.jetty.websocket</groupId>
						<artifactId>websocket-server</artifactId>
						<version>9.3.8.v20160314</version>
					</dependency>
				</dependencies>
			</plugin>
```

## Nginx配置

这个我也是后来测试才发现,需要加点东西, websocket才能传到后端去

```txt
                location / {
                        proxy_http_version 1.1;  #版本也必须是http 1.1
                        client_max_body_size 10m;
                        proxy_pass http://nutz;
                        proxy_set_header Host $http_host;
                        proxy_set_header X-Forwarded-For $remote_addr;
                        proxy_redirect http:// https://;
                        #add_header Access-Control-Allow-Origin "*";
						# 下面两行是为了websocket特意加的.
                        proxy_set_header Upgrade $http_upgrade;
                        proxy_set_header Connection "upgrade";
                }

```

## 分析一下YvrTopicWebSocket可否注入

在其构造方法打印堆栈信息,显示如下, 而且是第一次访问时才打印(即第一次访问时才创建对象)

```java
java.lang.Throwable
	at net.wendal.nutzbook.module.yvr.YvrTopicWebSocket.<init>(YvrTopicWebSocket.java:21)
	at sun.reflect.NativeConstructorAccessorImpl.newInstance0(Native Method)
	at sun.reflect.NativeConstructorAccessorImpl.newInstance(NativeConstructorAccessorImpl.java:62)
	at sun.reflect.DelegatingConstructorAccessorImpl.newInstance(DelegatingConstructorAccessorImpl.java:45)
	at java.lang.reflect.Constructor.newInstance(Constructor.java:422)
	at java.lang.Class.newInstance(Class.java:442)
	at org.apache.tomcat.websocket.server.DefaultServerEndpointConfigurator.getEndpointInstance(DefaultServerEndpointConfigurator.java:36)
	at org.apache.tomcat.websocket.pojo.PojoEndpointServer.onOpen(PojoEndpointServer.java:50)
	at org.apache.tomcat.websocket.server.WsHttpUpgradeHandler.init(WsHttpUpgradeHandler.java:138)
	at org.apache.coyote.AbstractProtocol$AbstractConnectionHandler.process(AbstractProtocol.java:701)
	at org.apache.tomcat.util.net.NioEndpoint$SocketProcessor.doRun(NioEndpoint.java:1500)
	at org.apache.tomcat.util.net.NioEndpoint$SocketProcessor.run(NioEndpoint.java:1456)
	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
	at org.apache.tomcat.util.threads.TaskThread$WrappingRunnable.run(TaskThread.java:61)
	at java.lang.Thread.run(Thread.java:745)
```

可以看到是DefaultServerEndpointConfigurator这个类调用clazz.newInstance生成实例

该类的源码地址 http://grepcode.com/file/repo1.maven.org/maven2/org.apache.tomcat/tomcat-websocket/8.0.24/org/apache/tomcat/websocket/server/DefaultServerEndpointConfigurator.java/

依赖ServiceLoader加载的,但还没成功,明天再说了.

