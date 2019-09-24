## helloword

### How To Start

#### Build Proto

```shell
protoc -I greeter/ greeter/proto/greeter.proto --go_out=plugins=grpc:greeter
```
After run this command you will see greeter.pb.go in greeter/proto

### Run Server
```$xslt
go run server/main.go
```

### Run Client
```$xslt
go run client/main.go
```


