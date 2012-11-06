---
comments: true
date: 2012-08-26 22:18:26
layout: post
slug: 'random-order-when-run-junit-testcase'
title: 换种顺序来执行TestCase(Junit适用)
permalink: '/453.html'
wordpress_id: 453
categories:
- Java
tags:
- Eclipse
- junit
- Nutz
---

Junit的TestCase,总是按固定的顺序执行的. 正如你在Eclipse中跑Run As Junit Test, 无论你跑多少次, TestCase的执行顺序都是一致的,可重复的. 这就导致一个问题, TestCase之间的独立性无法保证.

例如下面一个Test类中的2个TestCase:

    public class DaoTest {
    
        @Test
        public void test_count() {
            dao.insert(new User("root", "123456"));
            assertEquals(1, dao.count(User.class));
        }
    
        @Test
        public void test_insert() {
            dao.clear(User.class, null);
            dao.insert(new User("admin", "123456"));
            assertEquals(1, dao.count(User.class));
        }
    
    }
    
如果先执行test_count()然后执行test_insert(),两个TestCase都能通过.

但如果先执行test_insert(),然后执行test_count(),则test_count()会失败.

所以,有必要去打乱TestCase的默认执行顺序,以暴露出TestCase本身的问题. TestCase更可靠,才能让主代码更可靠.

我实现了一个简单的方式,使用的是Junit的公开API, 测试过4.3和4.8.2,均可使用:

            //得到所有带@Test的方法,这里用的是Nutz的资源扫描,反正你能得到全部Test类就行
            List<Class> list = Scans.me().scanPackage("org.nutz");
            List<request> reqs = new ArrayList<request>();
            Map<Request, Method> reqMap = new HashMap<Request, Method>();
            for (Class clazz : list) {
                Method[] methods = clazz.getMethods();
                for (Method method : methods) {
                    if (method.getAnnotation(Test.class) != null) {
                        //将单个TestCase(即一个Test Method),封装为Junit的Test Request
                        Request req = Request.method(clazz, method.getName());
                        reqs.add(req);
                        reqMap.put(req , method);//在最终打印测试结果时,方便查找具体出错的Method
                    }
                }
            }
    
            // 因为reqs 是一个List,我们可以按需调整TestCase的顺序
            // 正序 //nothing change.
            // 反序Collections.reverse(reqs)
            // 乱序Collections.shuffle(reqs)
    
            //把执行顺序保存下来,方便重现执行顺序
            try {
                FileWriter fw = new FileWriter("./test_order.txt");
                for (Request request : reqs) {
                    fw.write(reqMap.get(request).toString());
                    fw.write("\n");
                }
                fw.flush();
                fw.close();
            }
            catch (IOException e) {}
    
            //到这里, List已经按我们预期的方式排好,可以执行测试了
            final TestResult result = new TestResult();
            RunNotifier notifier = new RunNotifier();
            notifier.addListener(new RunListener() { //需要设置一个RunListener,以便收集测试结果
    
                public void testFailure(Failure failure) throws Exception {
                    result.addError(asTest(failure.getDescription()), failure.getException());
                }
                public void testFinished(Description description) throws Exception {
                    result.endTest(asTest(description));
                }
                public void testStarted(Description description) throws Exception {
                    result.startTest(asTest(description));
                }
                
                public junit.framework.Test asTest(Description description) {
                    return new junit.framework.Test() {
                        
                        public void run(TestResult result) {
                            throw Lang.noImplement();
                        }
                        
                        public int countTestCases() {
                            return 1;
                        }
                    };
                }
            });
            //来吧,执行之!!
            for (Request request : reqs) {
                request.getRunner().run(notifier);
            }
    
            //接下来,就是打印结果了.
            System.out.printf("Run %d , Fail %d , Error %d \n", result.runCount(), result.failureCount(), result.errorCount());
            
            if (result.failureCount() > 0) { //断言失败的TestCase
                Enumeration<testfailure> enu = result.failures();
                while (enu.hasMoreElements()) {
                    TestFailure testFailure = (TestFailure) enu.nextElement();
                    System.out.println("--Fail------------------------------------------------");
                    System.out.println(testFailure.trace());
                    testFailure.thrownException().printStackTrace(System.out);
                }
            }
            
            if (result.errorCount() > 0) { //抛异常的TestCase
                Enumeration<testfailure> enu = result.errors();
                while (enu.hasMoreElements()) {
                    TestFailure testFailure = (TestFailure) enu.nextElement();
                    System.out.println("--ERROR------------------------------------------------");
                    System.out.println(testFailure.trace());
                    testFailure.thrownException().printStackTrace(System.out);
                }
            }
    
来, 考验一下你的TestCase吧!! 让它在乱序中多次执行. Nutz按这种思路,已经爆出几个Bug(当然,我已经迅速fix了)

[https://github.com/nutzam/nutz/blob/master/test/org/nutz/AdvancedTestAll.java](https://github.com/nutzam/nutz/blob/master/test/org/nutz/AdvancedTestAll.java)
