---
comments: true
date: 2012-01-26 22:54:17
layout: post
slug: mongodb%e5%ad%a6%e4%b9%a0%e5%b0%8f%e8%ae%b0
title: Mongodb学习小记
wordpress_id: 361
categories:
- Java
- mongodb
tags:
- io
- mongodb
- MVC
- Nutz
- OpenID
- SSI
- 下载
---

实现自增(跟oralce的序列是一个概念):

    
    
    public static Integer getAutoIncreaseID(String idName) {
    	BasicDBObject query = new BasicDBObject("name", idName);
    	BasicDBObject update = new BasicDBObject("$inc", new BasicDBObject("id", 1));
    	return (Integer) XX.getDB()
    			.getCollection("inc_ids")
    			.findAndModify(query, null, null, false, update, true, true)
    			.get("id");
    }
    




把元素添加到数组,仅当数组中没有这个值

    
    
    BasicDBObject query = new BasicDBObject("_id", new ObjectId("XXXXXXXXXXXXXXX"));
    BasicDBObject update = new BasicDBObject("$addToSet", new BasicDBObject("tags", tag));
    db.getCollection().update(query, update);
    



OpenId的登录信息,放入Mongodb中存放,使用JopenId

    
    
    package org.nutz.viv.module;
    
    @IocBean
    @InjectName
    @At("/user")
    public class UserModule {
    
        static final long _5min = 300000L;
        static final String ATTR_MAC = "openid_mac";
        static final String ATTR_ALIAS = "openid_alias";
    	
        private String enpoint = "Google";
    	
        private OpenIdManager manager = new OpenIdManager();
    	
        @At("/login")
        @Ok(">>:${obj}")
        public String login(HttpSession session) {
    	manager.setReturnTo(Mvcs.getReq().getRequestURL().toString() + "/callback");
    	manager.setRealm("http://"+Mvcs.getReq().getHeader("Host") + "/");
    	manager.setTimeOut(300 * 1000);
    	Endpoint endpoint = manager.lookupEndpoint(enpoint);
            Association association = manager.lookupAssociation(endpoint);
            session.setAttribute(ATTR_MAC, association.getRawMacKey());
            session.setAttribute(ATTR_ALIAS, endpoint.getAlias());
            return manager.getAuthenticationUrl(endpoint, association); //返回的是一个Google登录页面的地址
        }
    
        @At("/login/callback")
        public String returnPoint(HttpServletRequest request) {
    	//checkNonce(request.getParameter("openid.response_nonce"));
            // get authentication:
            byte[] mac_key = (byte[]) request.getSession().getAttribute(ATTR_MAC);
            String alias = (String) request.getSession().getAttribute(ATTR_ALIAS);
            Authentication authentication = manager.getAuthentication(request, mac_key, alias);
            authentication.getEmail();
            BasicDBObject query = new BasicDBObject();
            query.append("email", authentication.getEmail());
            query.append("openid", "Google");
            BasicDBObject update = new BasicDBObject();
            update.append("set", new BasicDBObject("lastLoginDate", new Date()));
            DBObject dbObject = userDao.getCollection().findAndModify(query, null, null, false, update, true, true); //最后一个参数,表示如果没有相应的记录,则插入一条新的记录
            UserBean user = (UserBean) MapperUtil.fromDBObject(UserBean.class, dbObject);
            request.getSession().setAttribute("me", user);
            return "Login success!";
        }
    }
    



findAndModify是个好东西,呵呵

最后,mark一下Mongodb手册的下载地址: [http://dl.mongodb.org/dl/docs/](http://dl.mongodb.org/dl/docs/)

