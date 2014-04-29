---
title:  JDK8 Access to Parameter Names at Runtime
date: '2014-03-30'
permalink: '/2014/03/30.html'
categories:
- 其他
tags:
- java
- jdk8
---


JDK 8的新特性之一, 运行时获取方法参数的名称
-------------------------------------------

JDK8之前, Eclipse自带的ECJ编译器,同本地变量表,把方法参数的名字,放在最前面,使其编译出来的class的名字看推测.

而 JDK8把这种行为规范化(终于...) http://openjdk.java.net/jeps/118

参考文章 [Java 8 parameter name at runtime](http://www.java-allandsundry.com/2013/12/java-8-parameter-name-at-runtime.html)

演示代码
--------

```
package nutz_jdk8;

import java.lang.reflect.Constructor;
import java.lang.reflect.Parameter;

public class Bot {
    private final String name;
    private final String author;
    private final int rating;
    private final int score;

    public Bot(String name, String author, int rating, int score) {
        this.rating = rating; // 注意这里的顺序,并非按参数顺序逐一调用
        this.score = score;
        this.name = name;
        this.author = author;
    }

    public static void main(String[] args) throws NoSuchMethodException, SecurityException {
		Class<Bot> clazz = Bot.class;
		Constructor ctor = clazz.getConstructor(String.class, String.class, int.class, int.class);
		Parameter[] ctorParameters =ctor.getParameters();
		for (Parameter param: ctorParameters) {
		    System.out.println(param.isNamePresent() + ":" + param.getName());
		}
	}
}
```

输出的结果是:

```
true:name
true:author
true:rating
true:score
```

然后, 这个特性并未默认启用,javac需要额外的参数"-parameters"

```
javac -parameters nutz_jdk8\Bot.java
java -cp nutz_jdk8.Bot
```

如果是Eclipse 4.4 (正式支持JDK8的初始版本), 则需要手动在"Java Compiler"中启用之

<img src="{{urls.media}}/2014/03/30/jep118_eclipse.jpg" />

如果是maven,则需要这段

```
<plugin>
	<groupId>org.apache.maven.plugins</groupId>
	<artifactId>maven-compiler-plugin</artifactId>
	<version>3.1</version>
	<configuration>
		<source>1.8</source>
		<target>1.8</target>
		<compilerArgument>-parameters</compilerArgument>
	</configuration>
</plugin>
```

JDK8 , 开始学习啦
-----------------

开始学习新特性啦,哇哈哈