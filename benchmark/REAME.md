# Benchmark

## 测试环境
+ CPU: 1  Intel(R) Xeon(R) CPU E5-2676 v3 @ 2.40GHz
+ Memory: 1 G
+ OS: Linux version 4.14.138-114.102.amzn2.x86_64 , Red Hat 7.3.1-5
+ Go: 1.12.10

我们以`helloworld/greeter/proto`作为简单结构体和`helloworld/gin/proto`作为复杂结构体作为参考，同时以`helloworld/gin/server`为服务端，分别在本地和远程启动服务对比相同环境下GRPC+Proto和HTTP+JSON两种交互方式的性能差距


## 测试过程

### 简单结构体

```proto
syntax = "proto3";

package greeter;

service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

#### 服务端与客户端部署在相同机器

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|GRPC|10000|0.146|
|2|GRPC|10000|0.141|
|3|GRPC|10000|0.149|
|4|GRPC|10000|0.142|
|5|GRPC|10000|0.142|
|平均|GRPC|10000|0.144|

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|HTTP|10000|0.223|
|2|HTTP|10000|0.183|
|3|HTTP|10000|0.182|
|4|HTTP|10000|0.177|
|5|HTTP|10000|0.177|
|平均|HTTP|10000|0.188|


#### 服务端与客户端部署在不同机器

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|GRPC|10000|0.437|
|2|GRPC|10000|0.444|
|3|GRPC|10000|0.449|
|4|GRPC|10000|0.464|
|5|GRPC|10000|0.423|
|平均|GRPC|10000|0.443|

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|HTTP|10000|0.471|
|2|HTTP|10000|0.461|
|3|HTTP|10000|0.461|
|4|HTTP|10000|0.439|
|5|HTTP|10000|0.498|
|平均|HTTP|10000|0.466|


### 复杂结构
```proto
syntax = "proto3";

package greeter;

// 定义Greeter服务
service Greeter {
  // 定义Greeter中SayHello方法
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// 定义请求结构
message HelloRequest {
  string name = 1;
}

// 定义响应接口
message HelloReply {
  string replyValue1 = 1;
  string replyValue2 = 2;
  string replyValue3 = 3;
  string replyValue4 = 4;
  string replyValue5 = 5;
  string replyValue6 = 6;
  string replyValue7 = 7;
  string replyValue8 = 8;
  string replyValue9 = 9;
  string replyValue10 = 10;
  string replyValue11 = 11;
  string replyValue12 = 12;
  string replyValue13 = 13;
  string replyValue14 = 14;
  string replyValue15 = 15;
  string replyValue16 = 16;
  string replyValue17 = 17;
  string replyValue18 = 18;
  string replyValue19 = 19;
  string replyValue20 = 20;
}
```
#### 服务端与客户端部署在相同机器

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


#### 服务端与客户端部署在不同机器

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|GRPC|10000|0.526|
|2|GRPC|10000|0.491|
|3|GRPC|10000|0.612|
|4|GRPC|10000|0.530|
|5|GRPC|10000|0.529|
|平均|GRPC|10000|0.538|

|次数|类型|次数|耗时ms|
|---|---|---|---|
|1|HTTP|10000|0.522|
|2|HTTP|10000|0.499|
|3|HTTP|10000|0.639|
|4|HTTP|10000|0.584|
|5|HTTP|10000|0.561|
|平均|HTTP|10000|0.561|