---
comments: true
date: 2011-04-05 20:51:34
layout: post
slug: logger%e5%af%b9%e8%b1%a1%e6%98%af%e5%90%a6%e4%bc%9a%e9%87%8d%e5%a4%8d%e5%88%9b%e5%bb%ba%e5%91%a2-%e4%bb%a5log4j%e4%b8%ba%e4%be%8b
title: Logger对象是否会重复创建呢? (以Log4j为例)
wordpress_id: 263
categories:
- Java
tags:
- bug
- el
- io
- log4j
- logger
- Nutz
- SSI
- 下载
---

不久前有人问及 重复调用Logger.getLogger之类的语句,是否会拿到不同的Logger对象呢?
起初我并未在意,但现在觉得是应该探究一下.

故,先是设计这样的代码:

    
    
    import org.apache.log4j.Logger;
    public class LogTest {
    
    	public static void main(String[] args) {
    		Logger log = Logger.getLogger(LogTest.class);
    		Logger log2 = Logger.getLogger(LogTest.class);
    		System.out.println(log == log2);
    	}
    }
    


log4j.properties文件的内容:

    
    
    log4j.rootLogger=DEBUG,Console
    log4j.appender.Console=org.apache.log4j.ConsoleAppender
    log4j.appender.Console.layout=org.apache.log4j.PatternLayout
    log4j.appender.Console.layout.ConversionPattern=log4j: %d [%t] %-5p %c - %m%n
    
    


打印的结果是true,即同一个Log对象

恩,为什么呢?? 好,下载源码.

经过简单探寻,找到关键代码,位于org.apache.log4j.Hierarchy类:

    
    
      public
      Logger getLogger(String name, LoggerFactory factory) {
        //System.out.println("getInstance("+name+") called.");
        CategoryKey key = new CategoryKey(name);
        // Synchronize to prevent write conflicts. Read conflicts (in
        // getChainedLevel method) are possible only if variable
        // assignments are non-atomic.
        Logger logger;
    
        synchronized(ht) {
          Object o = ht.get(key);
          if(o == null) {
    	logger = factory.makeNewLoggerInstance(name);
    	logger.setHierarchy(this);
    	ht.put(key, logger);
    	updateParents(logger);
    	return logger;
          } else if(o instanceof Logger) {
    	return (Logger) o;
          } else if (o instanceof ProvisionNode) {
    	//System.out.println("("+name+") ht.get(this) returned ProvisionNode");
    	logger = factory.makeNewLoggerInstance(name);
    	logger.setHierarchy(this);
    	ht.put(key, logger);
    	updateChildren((ProvisionNode) o, logger);
    	updateParents(logger);
    	return logger;
          }
          else {
    	// It should be impossible to arrive here
    	return null;  // but let's keep the compiler happy.
          }
        }
      }
    


以上代码简单含义为,构建一个查询的key,然后获取特定的Logger对象,如果没有就创建它

问题又来啦,按照上述代码,获取Logger对象的过程,是一个同步过程,因为写着synchronized(ht) 呢!

另外说一句,按照当前的Nutz.Log实现,每次调用getLog都是返回新的Logger封装对象哦,不过这个对象只有一个变量,就是实际的logger,故,内存消耗依旧很小的.
