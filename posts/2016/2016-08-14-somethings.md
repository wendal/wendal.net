---
title: 杂碎
date: '2016-08-14'
permalink: '/2016/08/14.html'
description: 关于nutz,maven,ffmpeg,人生的一些想法
categories:
- Linux
tags:
- nutz
- ffmpeg
- maven
---

## 关于Nutz的版本发布(群消息复制粘贴)

bug总会有的, 无论我多么希望发布一个完美无bug的版本.

每次发布时, 都尽可能的修正已知的bug,并添加新的testcase, 但总有一些未能覆盖的场景.

发布基本是这样: TestAll(顺序,反序,乱序)跑通,更新我能控制的各种服务,看看有无异常, 然后呼吁大家测试一下新版本(很少有反馈), 静候几天, 发布新版本. 

再然后, 大家开始更新版本, 出问题, 继续报issue -- 再fix,添加testcase, 让TestAll覆盖更多场景.

个人觉得nutz发布版本已经算严格的了(国内开源项目中对比)

## Maven的增量编译

nutzmore项目中的子项目众多,所以一直希望用mvn deploy直接发布.但总是报签名不正确.

后来深究了原因,是maven的增量编译(useIncrementalCompilation 参数)的缘故. 这个参数会导致gpg签名后,有一定概率重新打包jar

```xml

			<plugin>
				<artifactId>maven-compiler-plugin</artifactId>
				<version>3.3</version>
				<configuration>
					<source>1.8</source>
					<target>1.8</target>
					<compilerArgs>
						<arg>-parameters</arg>
					</compilerArgs>
					<useIncrementalCompilation>false</useIncrementalCompilation>
				</configuration>
			</plugin>
```

命令行下的配置

```cmd
mvn -Dmaven.compiler.useIncrementalCompilation=false clean package javadoc:jar source:jar gpg:sign deploy
```

官方说明: https://maven.apache.org/plugins/maven-compiler-plugin/compile-mojo.html

## ffmpeg

高版本的ffmpeg在输出h264时,如果源数据不是yuv420p且不指定pixfmt,会生成mp4文件在低版本的h264解码器上会失败

```cmd
ffmpeg -i xxx.mp4 -pix_fmt yuv420p out.mp4
```

## 人生的一些想法

已知的会觉得理所当然

200年前还不能记录并还原声音, 1000年前,能进行一百以内四则运算的人口比例不到1%

"娱乐至死",也许真的是统治的最高境界.

本届奥运会的娱乐化程度更高了,"今天头条/天天头条"推送的大多数是娱乐新闻/路边新闻.