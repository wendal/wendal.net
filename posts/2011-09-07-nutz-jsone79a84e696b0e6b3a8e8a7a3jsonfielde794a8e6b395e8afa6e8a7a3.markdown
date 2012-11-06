---
comments: true
date: 2011-09-07 21:06:39
layout: post
slug: nutz-json%e7%9a%84%e6%96%b0%e6%b3%a8%e8%a7%a3jsonfield%e7%94%a8%e6%b3%95%e8%af%a6%e8%a7%a3
title: Nutz.Json的新注解@JsonField用法详解
wordpress_id: 319
categories:
- Java
tags:
- el
- js
- Nutz
---


长期以来Nutz的Json包,仅有一个注解@toJson,用于指定当前对象的通过什么方法进行Pojo-->String的转换

1.b.38版新增了@JsonField,在即将发布的1.b.40版将这个注解再次增强
正如其名,@JsonField是针对Json处理中字段级别的控制

首先是字段命名,看代码

    
    
    @JsonField("z-index")
    private String zIndex;
    


对应的Json将会是:

    
    
    {
        'z-index' : '10px'
    }
    


可以看到,无需再强求 json的key的名字与类属性名一致了

然后是"忽略",看代码

    
    
    @JsonField(ignore=true)
    private String password;
    


很清楚,就是不序列化这个字段

最后,是刚刚加入的by(生成器),这次用一个完整的例子:

    
    
    public class User {
    
        @JsonField(by="toString")
        private ObjectId objectId;
    
        @JsonField("username")
        private String name;
    
        @JsonField(by="net.wendal.helper.Md5#create")
        private String password;
    
        @JsonField(ignore=true)
        private String ip;
    }
    
    public final class Md5 {
    
        public static final String create(Object obj) {
            return XXXXX;
        }
    
    }
    


上面演示了by的两种写法,规则非常简单:
如果不包含#号,则代表调用自身的无参数方法,例子中的by="toString" 将调用objectId.toString()
如果包含#号,则代表调用指定类的指定静态方法,这个方法必须以Object作为参数,但可以返回任意类型的参数

@toJson加上@JsonField应该足够满足很大一部分的需要了

另外,这种by的语法,我正在积极考虑应用到Nutz.Dao的@Prev注解中,解决长期以来注解生成机制的缺失
