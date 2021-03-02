https://jergoo.gitbooks.io/go-grpc-practice-guide/content/chapter3/gateway.html

https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto
https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

go mod init grpc-gateway.http

cd proto 
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=github.com/jergoo/go-grpc-example/proto/google/api:. hello_http/*.proto


2021/03/01 10:50:50 WARNING: Deprecated use of 'go_package' option without a full import path in "hello_http/hello_http.proto", please specify:
        option go_package = "hello_http";
A future release of protoc-gen-go will require the import path be specified.
See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.


protoc --grpc-gateway_out=logtostderr=true:. hello_http/hello_http.proto


protoc --go_out=plugins=grpc:. hello_http/hello_http.proto


curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
{"message":"Hello gRPC-HTTP is working!."}


go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

 protoc -I/usr/local/include -I. -Igoogle/api --swagger_out=logtostderr=true:. hello_http/hello_http.proto

 https://segmentfault.com/a/1190000013513469

 https://github.com/envoyproxy/protoc-gen-validate