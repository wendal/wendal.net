---
comments: true
date: 2011-01-28 14:28:12
layout: post
slug: hibernate%e7%9a%84connectionprovider%e6%98%af%e6%80%8e%e4%b9%88%e4%b8%80%e5%9b%9e%e4%ba%8b
title: Hibernate的ConnectionProvider是怎么一回事?
wordpress_id: 233
categories:
- Java
tags:
- Ant
- bug
- el
- Hibernate
- io
- 连接池
- 配置
---

这次讨论一下Hibernate的ConnectionProvider接口, 因为我看到某些Hibernate项目是这样配置的:

    
    
    
    <property name="c3p0.min_size">5</property>
    <property name="c3p0.max_size">30</property>
    <property name="c3p0.time_out">1800</property>
    <property name="c3p0.max_statement">50</property>
    


问题是, 这样就配置好了吗? hibernate.connection.provider_class到底需不需要呢?
来看看Hibernate 3.3.2 GA的源码
ConnectionProviderFactory类的newConnectionProvider方法.

    
    
    public static ConnectionProvider newConnectionProvider(Properties properties, Map connectionProviderInjectionData) throws HibernateException {
    	ConnectionProvider connections;
    	String providerClass = properties.getProperty(Environment.CONNECTION_PROVIDER);
    	if ( providerClass!=null ) {
    		try {
    			log.info("Initializing connection provider: " + providerClass);
    			connections = (ConnectionProvider) ReflectHelper.classForName(providerClass).newInstance();
    		}
    		catch ( Exception e ) {
    			log.error( "Could not instantiate connection provider", e );
    			throw new HibernateException("Could not instantiate connection provider: " + providerClass);
    		}
    	}
    	else if ( properties.getProperty(Environment.DATASOURCE)!=null ) {
    		connections = new DatasourceConnectionProvider();
    	}
    	else if ( properties.getProperty(Environment.URL)!=null ) {
    		connections = new DriverManagerConnectionProvider();
    	}
    	else {
    		connections = new UserSuppliedConnectionProvider();
    	}
    
    	if ( connectionProviderInjectionData != null && connectionProviderInjectionData.size() != 0 ) {
    		//inject the data
    		try {
    			BeanInfo info = Introspector.getBeanInfo( connections.getClass() );
    			PropertyDescriptor[] descritors = info.getPropertyDescriptors();
    			int size = descritors.length;
    			for (int index = 0 ; index < size ; index++) {
    				String propertyName = descritors[index].getName();
    				if ( connectionProviderInjectionData.containsKey( propertyName ) ) {
    					Method method = descritors[index].getWriteMethod();
    					method.invoke( connections, new Object[] { connectionProviderInjectionData.get( propertyName ) } );
    				}
    			}
    		}
    		catch (IntrospectionException e) {
    			throw new HibernateException("Unable to inject objects into the conenction provider", e);
    		}
    		catch (IllegalAccessException e) {
    			throw new HibernateException("Unable to inject objects into the conenction provider", e);
    		}
    		catch (InvocationTargetException e) {
    			throw new HibernateException("Unable to inject objects into the conenction provider", e);
    		}
    	}
    	connections.configure(properties);
    	return connections;
    	}
    


上面的代码有点长, 精简出核心部分:

    
    
    ConnectionProvider connections;
    String providerClass = properties.getProperty(Environment.CONNECTION_PROVIDER);
    if ( providerClass!=null ) {
    	connections = (ConnectionProvider) ReflectHelper.classForName(providerClass).newInstance();
    }else if ( properties.getProperty(Environment.DATASOURCE)!=null ) {
    	connections = new DatasourceConnectionProvider();
    }else if ( properties.getProperty(Environment.URL)!=null ) {
    	connections = new DriverManagerConnectionProvider();
    }else {
    	connections = new UserSuppliedConnectionProvider();
    }
    /**
    Environment.CONNECTION_PROVIDER 的定义:
    public static final String CONNECTION_PROVIDER ="hibernate.connection.provider_class";
    Environment.DATASOURCE 的定义:
    public static final String DATASOURCE ="hibernate.connection.datasource";
    Environment.URL 的定义:
    public static final String URL ="hibernate.connection.url";
    */
    



可以看到, 如果hibernate.connection.provider_class和hibernate.connection.datasource都没有定义,就会使用内置的连接池,OK,那继续看默认的连接池DriverManagerConnectionProvider,只贴精华部分:

    
    
    /*连接池就是一个ArrayList !!*/
    private final ArrayList pool = new ArrayList();
    /*获取连接*/
    public Connection getConnection() throws SQLException {
    	synchronized (pool) {
    		if ( !pool.isEmpty() ) {
    			int last = pool.size() - 1;
    			Connection pooled = (Connection) pool.remove(last);
    			if (isolation!=null) pooled.setTransactionIsolation( isolation.intValue() );
    			if ( pooled.getAutoCommit()!=autocommit ) pooled.setAutoCommit(autocommit);
    			return pooled;
    		}
    	}
    
    	log.debug("opening new JDBC connection");
    	Connection conn = DriverManager.getConnection(url, connectionProps);
    	return conn;
    }
    /*释放连接*/
    public void closeConnection(Connection conn) throws SQLException {
    	synchronized (pool) {
    		int currentSize = pool.size();
    		if ( currentSize < poolSize ) {
    			pool.add(conn);
    			return;
    		}
    	}
    	conn.close();
    }
    


用一个简单ArrayList做出来的默认连接池,就是这样简单!!! sorry,是简陋!!! 无比简陋!! 性能能有多好?!
你的Hibernate还在用默认连接池? 你还没有配**hibernate.connection.provider_class**属性? 快去看看吧!!


