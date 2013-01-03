---
title: Ruhoh,Now!
date: '2012-11-06'
description:
categories: [linux]
tags : [blog]
---

正式转用Ruhoh静态博客引擎
-------------------------

* wordpress贴code实在太痛苦,还要排版
* 由于生成的是静态html,再也不必用神马php,世界清净了


博客内容完整迁移
----------------

* 很早就转用DISQUS,所以全部评论都完整保留
* 博客文章,转为markdown格式后,使用fix_wp_id_permalink.go和cleanup_blank_line.go修正为原链接及清理空行
* 本网站的源码存在github的[wendal.net库](http://github.com/wendal/wendal.net)
* 前端使用nginx 1.2.4, 配合git hook实现自动更新(待完成)


启用80vps的香港机房,弃用vpsee
-----------------------------

* 自从上传GFW发威, vpsee的机房就没快过
* 80vps的香港机房暂时看来还是很靠谱的,但峰值带宽只有1M
* 作为翻墙主要途径, vpsee的速度根本无法满足需求了


启用cdnzz,放弃cloudflare
------------------------

* 自从使用cloudflare,总用人投诉说博客无法访问,甚为不爽
* 当前使用cloudflare,也就是因为vpsee太慢
* cndzz收费,但1元/G,比较划算,按经验,每月流量也就2G