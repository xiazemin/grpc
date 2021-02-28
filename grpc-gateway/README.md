 git submodule add https://github.com/googleapis/googleapis --depth=1

https://jergoo.gitbooks.io/go-grpc-practice-guide/content/chapter3/gateway.html

https://github.com/grpc-ecosystem/grpc-gateway

https://github.com/googleapis/googleapis

https://docs.bazel.build/versions/master/install-os-x.html


 go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc


protoc-gen-grpc-gateway
protoc-gen-openapiv2
protoc-gen-go
protoc-gen-go-grpc



$ git submodule add https://github.com/googleapis/googleapis

# 编译google.api
$ protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# 编译hello_http.proto
$ protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/jergoo/go-grpc-example/proto/google/api:. hello_http/*.proto

# 编译hello_http.proto gateway
$ protoc --grpc-gateway_out=logtostderr=true:. hello_http/hello_http.proto

注意这里需要编译google/api中的两个proto文件，同时在编译hello_http.proto时使用M参数指定引入包名，最后使用grpc-gateway编译生成hello_http_pb.gw.go文件，这个文件就是用来做协议转换的，查看文件可以看到里面生成的http handler，处理proto文件中定义的路由"example/echo"接收POST参数，调用HelloHTTP服务的客户端请求grpc服务并响应结果。

curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
{"message":"Hello gRPC-HTTP is working!."}
