# 生成proto
pwd：grpc-example
<br>
protoc --proto_path=. --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./response/v1/response.proto 
<br>
protoc --proto_path=. --go_out=.  --go-grpc_out=. --grpc-gateway_out=. ./proto/v1/hello.proto 
<br>

# 运行
```
$go run server/server.go 
```

```
$go run client/client.go 
Name:"why"  Age:18
```
