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

## 5. Elasticsearch

### （1）新建data和config挂载目录

```bash
[root@go ~]# mkdir -p /data/elasticsearch/config
[root@go ~]# mkdir -p /data/elasticsearch/data

# 添加权限
[root@go ~]# chmod 777 -R /data/elasticsearch/
[root@go ~]# ll /data/elasticsearch/
总用量 0
drwxrwxrwx. 2 root root 31 10月 24 16:02 config
drwxrwxrwx. 2 root root  6 10月 24 15:58 data
```

### （2）写入配置到elasticsearch.yml

```bash
[root@go ~]# echo "http.host: 0.0.0.0" >> /data/elasticsearch/config/elasticsearch.yml
```

### （3）启动Elasticsearch

```bash
docker run --name shop_es --restart=always -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx128m" -v /data/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /data/elasticsearch/data:/usr/share/elasticsearch/data -v /data/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.17.6
```

> -e "discovery.type=single-node" 设置为单节点
>
> -e ES_JAVA_OPTS="-Xms64m -Xmx128m"  测试环境下，设置ES的初始内存和最大内存，否则导致过大启动不了ES

**测试启动成功**

```bash
[root@go ~]# curl localhost:9200
{
  "name" : "1a131380b19e",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "UKmfhXXjRAuEfgVSyXexdQ",
  "version" : {
    "number" : "7.17.6",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "f65e9d338dc1d07b642e14a27f338990148ee5b6",
    "build_date" : "2022-08-23T11:08:48.893373482Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

### （4）启动kibana

> 注意！kibana的版本号需要和es版本号一致

```bash
# docker run --name shop_kibana --restart=always -e ELASTICSEARCH_HOSTS=http:自己的IP地址:9200 -p 5601:5601 -d kibana:7.17.6
[root@go ~]# docker run --name shop_kibana --restart=always -e ELASTICSEARCH_HOSTS=http:192.168.139.130:9200 -p 5601:5601 -d kibana:7.17.6
```

**浏览器输入`ip:5601`进入web控制台**

选择==Dev Tools==

执行`GET /_cat/indices`查看所有index

### （5）集成ik中文分词器

下载对应版本ik

[https://github.com/medcl/elasticsearch-analysis-ik/releases/tag/v7.17.6](https://github.com/medcl/elasticsearch-analysis-ik/releases/tag/v7.17.6)

解压后重命名为ik

移动到挂载的`/data/elasticsearch/plugins`目录

## 6. rocketmq
使用install压缩包通过docker-compose安装
web控制台：[http://192.168.139.130:8080](http://192.168.139.130:8080)

## 7. jaeger
```shell
docker run -d --rm --name shop_jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
```
web控制台：[http://192.168.139.130:16686](http://192.168.139.130:16686)

## 8. kong
通过docker安装postgres
```shell
docker run -d --name kong-database --restart=always \
-p 5432:5432 \
-e "POSTGRES_USER=kong" \
-e "POSTGRES_DB=kong" \
-e "POSTGRES_PASSWORD=kong" postgres:12
```
初始化表
```shell
docker run --rm \
-e "KONG_DATABASE=postgres" \
-e "KONG_PG_HOST=192.168.139.130" \
-e "KONG_PG_PASSWORD=kong" \
-e "POSTGRES_USER=kong" \
-e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
kong kong migrations bootstrap
```
通过yum安装kong
```shell
curl -Lo kong-2.6.1.rpm $(rpm --eval "https://download.konghq.com/gateway-2.x-centos-7/Packages/k/kong-2.6.1.el7.amd64.rpm")
sudo yum install kong-2.6.1.rpm

systemctl stop firewalld.service
systemctl restart docker
# 启用防火请的话
firewall-cmd --zone=public --add-port=8001/tcp --permanent
firewall-cmd --zone=public --add-port=8000/tcp --permanent
firewall-cmd --reload

cp /etc/kong/kong.conf.default /etc/kong/kong.conf
vim /etc/kong/kong.conf
#修改如下内容
database = postgres
pg_host = 192.168.139.130
pg_port = 5432
pg_timeout = 5000

pg_user = kong
pg_password = kong
pg_database = kong

dns_resolver = 192.168.139.130:8600 #consul的dns
admin_listen = 0.0.0.0:8001 reuseport backlog=16384, 0.0.0.0:8444 http2 ssl reuseport backlog=16384
proxy_listen = 0.0.0.0:8000 reuseport backlog=16384, 0.0.0.0:8443 http2 ssl reuseport backlog=16384

###
kong migrations bootstrap up -c /etc/kong/kong.conf #初始化数据库
kong start -c /etc/kong/kong.conf #启动kong
```
通过docker安装konga
```shell
docker run -d --name shop_konga --restart=always -p 1337:1337 pantsel/konga
```
web控制台：[http://192.168.139.130:1337](http://192.168.139.130:1337)

用户名：fanb 密码：shop123456

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
## 4. order

### （1）order_srv.json

```json
{
    "name":"order_srv",
    "host":"192.168.1.105",
    "tags":[ "order", "srv" ],
    "mysql":{
        "host":"192.168.139.130",
        "port":3306,
        "db":"shop_order",
        "user":"root",
        "password":"shop123456"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    },
    "goods_srv":{
        "name": "goods_srv"
    },
    "inventory_srv":{
        "name":"inventory_srv"
    }
}
```

### （2）order_web.json
```json
{
    "name":"order_web",
    "host":"192.168.1.105",
    "tags":[ "order", "web" ],
    "port":8084,
    "goods_srv":{
        "name":"goods_srv"
    },
    "order_srv":{
        "name":"order_srv"
    },
    "inventory_srv":{
        "name":"inventory_srv"
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
    },
    "alipay":{
        "app_id":"",
        "private_key":"",
        "ali_public_key":"",
        "notify_url":"",
        "return_url":""
    }
}
```

## 5. userop

### （1）userop_srv.json

```json
{
    "name":"userop_srv",
    "host":"192.168.1.111",
    "tags":[ "userop", "srv" ],
    "mysql":{
        "host":"192.168.139.130",
        "port":3306,
        "db":"shop_userop",
        "user":"root",
        "password":"shop123456"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    }
}
```
### （2）userop_web.json
```json
{
    "name":"userop_web",
    "host":"192.168.1.111",
    "tags":[ "userop", "web" ],
    "port":8085,
    "goods_srv":{
        "name":"goods_srv"
    },
    "userop_srv":{
        "name":"userop_srv"
    },
    "jwt":{
        "key":"fanb"
    },
    "consul":{
        "host":"192.168.139.130",
        "port":8500
    }
}
```

# 三、细节
## 1. 使用consul服务发现时
```go
package initialize

import (
    _ "github.com/mbobakov/grpc-consul-resolver" // It's important
)
```