---
comments: true
date: 2011-10-16 22:22:36
layout: post
slug: nut-test-%e6%96%b9%e4%be%bf%e6%b5%8b%e8%af%95%e4%bd%bf%e7%94%a8nutz%e7%9a%84%e4%bb%a3%e7%a0%81%e8%bf%9b%e8%a1%8c%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95
title: Nut.Test -- 使用Nutz一样能方便地进行单元测试
permalink: '/332.html'
wordpress_id: 332
categories:
- Java
tags:
- io
- js
- MVC
- Nutz
---

首先,感谢Jay提出一个需求.

这个功能很有可能出现在1.b.40中

当前Nut.Test代码,主要解决以下几个问题:
1. 测试方法使用NutzDao操作数据库,希望在测试方法执行完成后进行回滚
2. Ioc支持,方便获取与实际生产环境类似的Ioc注入功能,方便进行功能测试

这个测试Runner, 很好地平衡了侵入性与灵活性, 让用户以很低的成本与耦合度测试Nutz代码

所以,我新建org.nutz.test包, 核心类NutzJUnit4ClassRunner,这个类本身非常简单

    public class NutzJUnit4ClassRunner extends TestClassRunner {
        public NutzJUnit4ClassRunner(final Class klass) throws InitializationError {
            super(klass, new NutTestClassMethodsRunner(klass));
        }
    }
    
核心代码位于NutTestClassMethodsRunner

    public class NutTestClassMethodsRunner extends TestClassMethodsRunner {
    	
        //其他辅助方法,属性
    
        protected void invokeTestMethod(final Method method, final RunNotifier notifier) {
            //处理事务回滚问题,判断当前方法是否需要自动回滚
            ... ...
            //检查Ioc支持,主要是判断当前类及父类是否标注了@IocBy
    	... ...
    		
    	//具体执行
            //如果不需要自动回滚,那么,直接调用父类的方法,按原生步骤执行
    
            //事务自动回滚的实现,就具体的执行过程,以事务模板包裹,并确保抛出异常
            try {
                Trans.exec(new Atom(){
                    public void run() {
                        NutTestClassMethodsRunner.super.invokeTestMethod(method, notifier);
                        throw JustRollback.me();//这样,无论原方法是否跑异常,事务模板都能收到异常,并回滚
                    }
                });
            } catch (JustRollback e) {}
        }
    
        //如果包含Ioc支持,并且当前类是一个IocBean的话,就可以从Ioc中获取对象
        protected Object createTest() throws Exception {
            if (NutTestContext.me().ioc != null && klass.getAnnotation(IocBean.class) != null)
                return NutTestContext.me().ioc.get(klass);
            return super.createTest();
        }
    }
    
整个实现,共用到4个注解 @NutTest @IocBy @IocBean @Inject
这里的@IocBy的具体行为,与Mvc中的@IocBy有轻微不同,因为没有web上下文
@Aop,声明式Aop均可生效,效果与MVC中的效果一致

使用示例

    @RunWith(value=NutzJUnit4ClassRunner.class)
    @IocBy(type=ComboIocProvider.class,args={"*org.nutz.ioc.loader.json.JsonLoader","ioc/",
    	  "*org.nutz.ioc.loader.annotation.AnnotationIocLoader","net.wendal"})
    @IocBean
    public class AuthServiceTest {
    
        @Inject
        private AuthDao authDao;
    
        public void setAuthDao(AuthDao authDao) {this.authDao = authDao;}
    
        @Test
        public void test_login() {
            User user = authDao.fetch("admin","wendal.net")
            assertNotNull(user);
       }
    }
    
局限性:
1. 不应该用于测试Action层
2. 自动回滚,必须注意隐式事务提交和多线程事务问题

Enjoy it!!
