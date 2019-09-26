# Benchmark

## 测试环境
+ CPU: 1  Intel(R) Xeon(R) CPU E5-2676 v3 @ 2.40GHz
+ Memory: 1 G
+ OS: Linux version 4.14.138-114.102.amzn2.x86_64 , Red Hat 7.3.1-5
+ Go: 1.12.10

我们以`helloworld/gin/server`为服务端，`helloworld/gin/proto`为proto文件，对比相同环境下GRPC+Proto和HTTP+JSON两种交互方式的性能差距


## 测试过程

### 服务端与客户端部署在相同机器

#### Get请求
|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|GRPC|10000|0.143|
|2|GRPC|10000|0.138|
|3|GRPC|10000|0.141|
|4|GRPC|10000|0.145|
|5|GRPC|10000|0.144|
|平均|GRPC|10000|0.142|

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|HTTP|10000|0.189|
|2|HTTP|10000|0.277|
|3|HTTP|10000|0.167|
|4|HTTP|10000|0.333|
|5|HTTP|10000|0.252|
|平均|GRPC|10000|0.243|

背景：本地部署+Get请求




### 服务端与客户端部署在不同机器

