---
title: 从源码编译Chrome(chromium)
date: '2015-02-04'
description: 在windows下编译Chrome
categories: 其他
permalink: '/2015/02/04.html'
tags: 
- chrome
- chromium

---

**Update: 20160724 最新版源码已经打包上传好,均为tar文件, 无历史记录和完整历史记录,两个版本. **

官网描述
------------------------------------

http://www.chromium.org/developers/how-tos/build-instructions-windows


为啥还要写这篇博客
------------------------------------

1. 太久没在这里写博客
2. Chrome编译的教程,网上太多太多了,但各种错误(也许对他们当时的版本来说并不是错误)
3. 好多博客没有把问题说清楚

源码打包下载(vbox的系统镜像文件)
-------------------------------

http://pan.baidu.com/s/1ntCHH1Z 密码：nbrm

如果失效请留言

第一步,修改系统语言
------------------------------------

切换系统语言为英文!!!!!!

官网原文: 

```
You must set your Windows system locale to English, or else you may get build errors about "The file contains a character that cannot be represented in the current code page."
```

不改?绝对的坑!! 报各种编码错误,最后我重新安装win7旗舰版!!

必须是x64系统!! 想想内存需求也应该明白!

```
You must have Windows 7 x64 or later. x86 OSs are unsupported.
```

再提醒一句, 安装所有重要的系统补丁, 用windows update服务安装!!! 里面包含IE11,必备. -- 这一步只是为了保险起见.

第二步,安装VS2013
-------------------------------------

官网的要求是VS2013, 不是2008,不是2010,不是2012, 当前最新的要求是2013!!

```
You must build with Visual Studio 2013 Update 4, no other versions are supported.
```

http://www.visualstudio.com/downloads/download-visual-studio-vs

网络安装或下载iso(6G左右) 均可, 只需要安装C++套装,其他一概取消.

提示: 最后一步安装update4补丁包的时候(也就是最后的阶段),会很慢很慢,很慢,不知道为啥,反正很久,等吧.

第三步,科学上网
-----------------------------------

往下的步骤都需要科学上网,稳定的科学上网,可靠的科学上网,别怪我没提醒你git clone是不支持断点续传的!!!

第四步,添加环境变量
----------------------------------

系统的环境变量加入 DEPOT_TOOLS_WIN_TOOLCHAIN 值为0


第五步,安装depot_tools
------------------------------------

下载页面: http://commondatastorage.googleapis.com/chrome-infra-docs/flat/depot_tools/docs/html/depot_tools_tutorial.html#_setting_up
下载地址: https://src.chromium.org/svn/trunk/tools/depot_tools.zip

解压到某个盘的根目录,别带中文,特殊字符等一切蛋疼的东西, 修改系统的环境变量, 把depot_tools的路径加入到PATH

启动cmd, 随便找个目录,执行

```
gclient
```

会自行下载python,git,svn等等依赖工具, 系统已经安装的python,git是不认的!!!


第六步,下载源码
------------------------------------

重新打开一个console(cmd或者ComEms均可)

在一个剩余空间60G以上的盘, 严重建议是SSD, 起码是SSD加速盘或混合硬盘

建一个文件夹,叫chrome_build, 或任何你喜欢的英文名,别中文啊啊啊啊,假设为 W:\chrome_build

```
W:
mkdir chrome_build
cd W:\chrome_build
```

经典做法, 直接fetch

```
fetch chromium #会很久很久
```

省流量的方法, 只下载最新的代码,没有历史记录

```
fetch --nohooks --no-history chromium
```

或者下载我的打包好的源码镜像文件

PS: 20151126,正在下载最新的,完成后打包上传. windows下的压缩包

第七步, 编译
--------------------------------------

编译可以说是最简单的一步

先生成各种文件(可以省略)
```
gclient runhooks
```

执行编译, out/Debug可以改成out/Release等等.
```
cd src
ninja -C out/Debug chrome
```

输出:

```
ninja: Entering directory `out/Debug'
[541/19418] RULE Assembling nacl_switch_unwind_win.asm to obj\native_client\sr...ice_runtime\arch\x86_64\service_runtime_x86_64.gen\nacl_switch_unwind_win.obj.
 Assembling: nacl_switch_unwind_win.asm
[19418/19418] STAMP obj\chrome\chrome.actions_rules_copies.stamp
```

19418个编译任务, 193xx的时候来开始链接,很慢, 我的笔记本电脑i7-3630, 8G内存, 32G SSD加速的普通机械硬盘, 编译了2小时.

编译完成后可以到 out/Debug目录下找到chrome.exe, 启动一下就是你编译的Chrome了.

具体怎么打包成安装文件,还没找到方法.


