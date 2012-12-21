---
title: Android中以root权限调用本地程序
date: '2012-12-21'
description:
categories: [linux]
tag : [android]
permalink: '/2012/1221.html'
---

最近用到Android,需要执行一些本地程序,以前root机做过不少,但还没真正用Java代码调用过

今天,总算解决了
--------------

	public static boolean runAsRoot(String cmd) {
		try {
			Process p = Runtime.getRuntime().exec("su");
			OutputStream out = p.getOutputStream();
			out.write((cmd + "\n").getBytes());
			out.flush();
			out.close();
			if (p.waitFor() == 0) {
				return true;
			}
			return false;
		} catch (Exception e) {
			return false;
		}
	}

原理
---

1. 前提,当然是你的机器已经root
2. 所谓root过,就是能无限制地执行su
3. android上的su,就是改变当前进程的uid和gid,然后转为一个shell
4. 上述代码就是先执行su,然后将所需命令传入这个"shell"来执行