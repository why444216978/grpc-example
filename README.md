# 生成proto
```
$pwd
$/home/users/weihaoyu/go/own/grpc-example
$protoc --proto_path=. --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./response/v1/response.proto 
$protoc --proto_path=. --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./proto/v1/hello.proto 
```

# 运行
```
$go run server/server.go 
```

```
$go run client/client.go 
Name:"why"  Age:18
```

<a href="https://success.blog.csdn.net/article/details/114959896">博文地址</a>
