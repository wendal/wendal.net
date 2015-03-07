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

官网描述
------------------------------------

http://www.chromium.org/developers/how-tos/build-instructions-windows


为啥还要写这篇博客
------------------------------------

1. 太久没在这里写博客
2. Chrome编译的教程,网上太多太多了,但各种错误(也许对他们当时的版本来说并不是错误)
3. 好多博客没有把问题说清楚

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

省流量的方法

```
fetch chromium # 出现sync字样后, 终止(ctrl+c)

gclient sync --no-history --force # 需要比较长的时间了
# 如果断开了,重新执行gclient语句就可以了,不需要再执行fetch.
```

别以为sync只是下载源码, 最后还会下载nacl的sdk的!!

第6步补充说明, 关于下载Webkit
----------------------------------

正常下载的话, 下载webkit会很久很久,因为是整个历史都下载下来.

下面介绍的做法,需要修改DEPS文件, 比较折腾, 自行想象吧.


看到

```
[0:13:08]   src/third_party/WebKit
```

的时候,可以终止gclient

然后,执行:

```
# 用notepad++ 打开src/DEPS,找到webkit_revision的配置
cd src/third_party/
git clone --depth=10 https://chromium.googlesource.com/chromium/blink.git WebKit
#大概下载360mb

Cloning into 'WebKit'...
remote: Sending approximately 5.08 GiB ...
remote: Counting objects: 123641, done
remote: Finding sources: 100% (123641/123641)
remote: Total 123641 (delta 37772), reused 78823 (delta 37772)
Receiving objects: 100% (123641/123641), 360.61 MiB | 444.00 KiB/s, done.
Resolving deltas: 100% (37772/37772), done.
Checking connectivity... done.
Checking out files: 100% (144812/144812), done.

# 执行git rev-list找个可用的rev
git rev-list  HEAD
# 然后找出倒数第二个commit的sha1, 修改webkit_revision的值
# 回到根目录,重新开始gclient
gclient sync --no-history --force
```

继续长时间的等待, 真的很久很久, 洗洗睡觉吧.

PS: v8也很慢, 见仁见智吧.

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


