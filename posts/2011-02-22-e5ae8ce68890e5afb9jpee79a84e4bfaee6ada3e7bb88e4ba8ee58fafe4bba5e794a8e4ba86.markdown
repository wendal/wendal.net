---
comments: true
date: 2011-02-22 16:08:20
layout: post
slug: '%e5%ae%8c%e6%88%90%e5%af%b9jpe%e7%9a%84%e4%bf%ae%e6%ad%a3%e7%bb%88%e4%ba%8e%e5%8f%af%e4%bb%a5%e7%94%a8%e4%ba%86'
title: 完成对JBE的修正,终于可以用了.
wordpress_id: 244
categories:
- Java
tags:
- Ant
- el
- io
- Java
- SSI
- 下载
- 反编译
---

先放出下载地址:
[jbe-0.1b-Fixed-By-WendalChen.zip](https://docs.google.com/leaf?id=0B8hUXYDeoy_hNDczMTQ5OWUtYjJlNC00NTE2LWI1MWItODk4NGIwNDUyZjMz&hl=zh_CN)

原版下载地址: [JBE - Java Bytecode Editor](http://www.cs.ioc.ee/~ando/jbe/)

原版存在的问题:
对方法进行修改后,点击Save Method, 如果存在 invokeinterface 的行, 将无法保存!! 后台报错.

以下是我在压缩包中写的描述:
This JBE had been modify by Wendal Chen , base on JBE 0.1b

Fix:
when save method whit "invokeinterface" , app fail.

File change:
//---------------------------------------------------------------
JAsmParser.java line 169~170:
int arg1 = getMethodConstRef(instrElems, cpg, labels);
int arg2 = Utility.methodSignatureArgumentTypes(getDescrFromFullMethod(instrElems[1])).length;

//---------------------------------------------------------------
ConstantInterfaceMethodrefInfo.java Override getVerbose() :

public String getVerbose() throws InvalidByteCodeException {
    ConstantNameAndTypeInfo nameAndType = getNameAndTypeInfo();

    return classFile.getConstantPoolEntryName(classIndex) + "/" +
               classFile.getConstantPoolEntryName(nameAndType.getNameIndex())
               +classFile.getConstantPoolEntryName(nameAndType.getDescriptorIndex());
}

I check my class file, it work. I hope it work for you too.

通过修改其中两个类来修复这个问题.
