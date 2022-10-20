# 一、shop项目所需的中间件

## 1. redis

```shell
docker run -d --name shop_redis --restart=always -p 6379:6379 redis
```

## 2. mysql

```shell
docker run -d --name shop_mysql --restart=always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=shop123456 mysql
```

账号root

密码shop123456

## 3. consul

服务注册与发现

```shell
docker run -d --name shop_consul --restart=always -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul agent -dev -client=0.0.0.0
```

web控制台 [http://192.168.139.130:8500](http://192.168.139.130:8500)

- 8300：集群内数据的读写和复制
- 8301：单个数据中心gossip协议通讯
- 8302：跨数据中心gossip协议通讯
- 8500：提供获取服务列表、注册服务、注销服务等HTTP接口；提供UI服务
- 8600：采用DNS协议提供服务发现功能

## 4. nacos

配置中心

```shell
docker run -d --name shop_nacos --restart=always -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 nacos/nacos-server
```

web控制台：[http://192.168.139.130:8848/nacos/index.html](http://192.168.139.130:8848/nacos/index.html)

账号nacos

密码nacos

# 二、nacos中配置文件

## 1. user

### （1）user_srv.json

```json
{
    "name":"user_srv",
    "host":"192.168.1.105",
    "tags":[ "user", "srv" ],
    "mysql":{
        "host":"192.168.139.130",
        "port":3306,
        "db":"shop_user",
        "user":"root",
        "password":"shop123456"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    }
}
```

### （2）user_web.json

```json
{
    "name":"user_web",
    "host":"192.168.1.105",
    "tags":[ "user", "web" ],
    "port":8081,
    "user_srv":{
        "name":"user_srv"
    },
    "jwt":{
        "key":"XXX"
    },
    "sms":{
        "key":"xxx",
        "secrect":"xxx",
        "template-code":"SMS_154950909",
        "sign-name":"阿里云短信测试",
        "region-id":"cn-zhangjiakou"
    },
    "redis":{
        "host":"192.168.139.130",
        "port":6379
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    }
}
```

## 2. goods

### （1）goods_srv.json

```json
{
    "name":"goods_srv",
    "host":"192.168.1.105",
    "tags":[ "goods", "srv" ],
    "mysql":{
        "host":"192.168.139.130",
        "port":3306,
        "db":"shop_goods",
        "user":"root",
        "password":"shop123456"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    },
    "es":{
        "host":"192.168.139.130",
        "port":9200
    }
}
```

### （2）goods_web.json

```json
{
    "name":"goods_web",
    "host":"192.168.1.105",
    "tags":[ "goods", "web" ],
    "port":8082,
    "goods_srv":{
        "name":"goods_srv"
    },
    "jwt":{
        "key":"fanb"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    },
    "jaeger":{
        "host":"192.168.139.130",
        "port":5775,
        "name":""
    }
}
```

### （3）oss_web.json

```json
{
    "name":"oss_web",
    "host":"192.168.1.105",
    "tags":["oss","web"],
    "port":8083,
    "jwt":{
        "key":"fanb"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    },
    "oss":{
        "key":"xxx",
        "secret":"xxx",
        "host":"http://xxx.aliyuncs.com",
        "callback_url":"http://jt2gxw.natappfree.cc/oss/v1/oss/callback",
        "upload_dir":"goods/"
    }
}
```

## 3. inventory

### （1）inventory_srv.json
```json
{
    "name":"inventory_srv",
    "host":"192.168.1.111",
    "tags":[ "inventory", "srv" ],
    "mysql":{
        "host":"192.168.139.130",
        "port":3306,
        "db":"shop_inventory",
        "user":"root",
        "password":"shop123456"
    },
    "redis":{
        "host":"192.168.139.130",
        "port":6379
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    }
}
```
