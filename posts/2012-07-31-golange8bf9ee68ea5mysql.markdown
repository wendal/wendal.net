---
comments: true
date: 2012-07-31 19:33:39
layout: post
slug: 'access-mysql-in-golang'
title: Golang连接Mysql
permalink: '/448.html'
wordpress_id: 448
categories:
- go
tags:
- el
- Java
- js
- MySQL
---

首先,安装golang-mysql库, 我这里选用是google上的[go-mysql-driver](http://code.google.com/p/go-mysql-driver/)

    
     
    go get code.google.com/p/go-mysql-driver/mysql
    #如果访问失败,请翻墙吧!! 需要mysql 4.1以上哦
    








然后,当然是我的最爱 -- 代码

    
    
    package main
    
    // 导入sql包, 跟java.sql类似的
    import "database/sql"
    import _ "code.google.com/p/go-mysql-driver/mysql"
    import "encoding/json"
    import "fmt"
    
    // 定义一个结构体, 需要大写开头哦, 字段名也需要大写开头哦, 否则json模块会识别不了
    // 结构体成员仅大写开头外界才能访问
    type User struct {
        User      string    `json:"user"`
        Password string `json:"password"`
        Host   string `json:"host"`
    }
    
    // 一如既往的main方法
    func main() {
        // 格式有点怪, @tcp 是指网络协议(难道支持udp?), 然后是域名和端口
        db, e := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mysql?charset=utf8")
        if e != nil { //如果连接出错,e将不是nil的
            print("ERROR?")
            return
        }
        // 提醒一句, 运行到这里, 并不代表数据库连接是完全OK的, 因为发送第一条SQL才会校验密码 汗~!
        _, e2 := db.Query("select 1")
        if e2 == nil {
            println("DB OK")
            rows, e := db.Query("select user,password,host from mysql.user")
            if e != nil {
                fmt.Print("query error!!%v\n", e)
                return
            }
            if rows == nil {
                print("Rows is nil")   
                return
            }
            for rows.Next() { //跟java的ResultSet一样,需要先next读取
                user := new(User)
                // rows貌似只支持Scan方法 继续汗~! 当然,可以通过GetColumns()来得到字段顺序
                row_err := rows.Scan(&user.User;,&user.Password;, &user.Host;)
                if row_err != nil {
                    print("Row error!!")
                    return
                }
                b, _ := json.Marshal(user)
                fmt.Println(string(b)) // 这里没有判断错误, 呵呵, 一般都不会有错吧
            }
            println("Done")
        }
    }
    









编译后, 体积高达2.5mb, 实在惊人. 运行速度也很不错, 0.012秒完成:

    
    
    linux-9rhn:/home/go # time ./test_mysql
    DB OK
    {"user":"root","password":"*6BB4837EB74329105EE4568DDA7DC67ED2CA2AD9","host":"localhost"}
    {"user":"root","password":"","host":"linux-9rhn"}
    {"user":"root","password":"","host":"127.0.0.1"}
    {"user":"","password":"","host":"localhost"}
    {"user":"","password":"","host":"linux-9rhn"}
    Done
    
    real 0m0.012s
    user 0m0.008s
    sys  0m0.000s
    









附上一句, 不用猜密码了,是123456, 写着呢
