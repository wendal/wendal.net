---
title: arduino pro mini的电压问题
date: '2014-08-04'
permalink: '/2014/08/04.html'
categories:
- 其他
tags:
- arduino
---

arduino pro mini的MCU是 atmega328p
----------------------------------

http://www.atmel.com/devices/atmega328.aspx

电压从 1.8-5.5

工作频率最高20Mhz

电压与频率的关系
------------------------------------

3.3v 安全频率 约12.5Mhz
5v   安全频率 20Mhz

淘宝上绝大多数都是5V 16Mhz的, 但我手上的几块均能工作在3.3V

所以就很奇怪这5v和3.3v到底影响了什么.

由于3.3v下, 16Mhz高于安全频率, 所以mcu事实上工作在"超频"的状态
