---
comments: true
date: 2011-01-27 20:06:27
layout: post
slug: '%e7%b3%bb%e7%bb%9f%e5%8f%98%e9%87%8ffile-encoding%e5%af%b9java%e7%9a%84%e8%bf%90%e8%a1%8c%e5%bd%b1%e5%93%8d%e6%9c%89%e5%a4%9a%e5%a4%a7'
title: 系统变量file.encoding对Java的运行影响有多大?
permalink: '/232.html'
wordpress_id: 232
categories:
- Java
tags:
- el
- io
- Java
- Nutz
- 编码
---

这个话题来自: Nutz的[issue 361](http://code.google.com/p/nutz/issues/detail?id=361)

在考虑这个issue时, 我一直倾向于使用系统变量file.encoding来改变JVM的默认编码.

今天,我想到, 这个系统变量,对JVM的影响到底有多大呢?
我使用最简单的方法看看这个变量的影响--在JDK 1.6.0_20的src.zip文件中,查找包含file.encoding字眼的文件.
共找到4个, 分别是:
先上重头戏 java.nio.Charset类:

        public static Charset defaultCharset() {
            if (defaultCharset == null) {
    	    synchronized (Charset.class) {
    		java.security.PrivilegedAction pa =
    		    new GetPropertyAction("file.encoding");
    		String csn = (String)AccessController.doPrivileged(pa);
    		Charset cs = lookup(csn);
    		if (cs != null)
    		    defaultCharset = cs;
                    else 
    		    defaultCharset = forName("UTF-8");
                }
    	}
    	return defaultCharset;
        }
    
java.net.URLEncoder的静态构造方法,影响到的方法 java.net.URLEncoder.encode(String) 

        static {
    
    	dontNeedEncoding = new BitSet(256);
    	int i;
    	for (i = 'a'; i < = 'z'; i++) {
    	    dontNeedEncoding.set(i);
    	}
    	for (i = 'A'; i <= 'Z'; i++) {
    	    dontNeedEncoding.set(i);
    	}
    	for (i = '0'; i <= '9'; i++) {
    	    dontNeedEncoding.set(i);
    	}
    	dontNeedEncoding.set(' '); /* encoding a space to a + is done
    				    * in the encode() method */
    	dontNeedEncoding.set('-');
    	dontNeedEncoding.set('_');
    	dontNeedEncoding.set('.');
    	dontNeedEncoding.set('*');
    
        	dfltEncName = (String)AccessController.doPrivileged (
    	    new GetPropertyAction("file.encoding")
        	);
        }
    
com.sun.org.apache.xml.internal.serializer.Encoding的getMimeEncoding方法(209行起)

        static String getMimeEncoding(String encoding)
        {
    
            if (null == encoding)
            {
                try
                {
    
                    // Get the default system character encoding.  This may be
                    // incorrect if they passed in a writer, but right now there
                    // seems to be no way to get the encoding from a writer.
                    encoding = System.getProperty("file.encoding", "UTF8");
    
                    if (null != encoding)
                    {
    
                        /*
                        * See if the mime type is equal to UTF8.  If you don't
                        * do that, then  convertJava2MimeEncoding will convert
                        * 8859_1 to "ISO-8859-1", which is not what we want,
                        * I think, and I don't think I want to alter the tables
                        * to convert everything to UTF-8.
                        */
                        String jencoding =
                            (encoding.equalsIgnoreCase("Cp1252")
                                || encoding.equalsIgnoreCase("ISO8859_1")
                                || encoding.equalsIgnoreCase("8859_1")
                                || encoding.equalsIgnoreCase("UTF8"))
                                ? DEFAULT_MIME_ENCODING
                                : convertJava2MimeEncoding(encoding);
    
                        encoding =
                            (null != jencoding) ? jencoding : DEFAULT_MIME_ENCODING;
                    }
                    else
                    {
                        encoding = DEFAULT_MIME_ENCODING;
                    }
                }
                catch (SecurityException se)
                {
                    encoding = DEFAULT_MIME_ENCODING;
                }
            }
    
最后一个javax.print.DocFlavor类的静态构造方法:

        static {
    	hostEncoding = 
    	    (String)java.security.AccessController.doPrivileged(
                      new sun.security.action.GetPropertyAction("file.encoding"));
        }
    
可以看到,系统变量file.encoding影响到
1. Charset.defaultCharset() Java环境中最关键的编码设置
2. URLEncoder.encode(String) Web环境中最常遇到的编码使用
3. com.sun.org.apache.xml.internal.serializer.Encoding 影响对无编码设置的xml文件的读取
4. javax.print.DocFlavor 影响打印的编码

故,影响还是很大的哦, 可以说是Java中编码的一个关键钥匙!
