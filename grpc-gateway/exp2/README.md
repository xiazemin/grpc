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


 https://cloud.tencent.com/developer/article/1627554
 https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto
 https://github.com/grpc-ecosystem/grpc-gateway

 https://github.com/protocolbuffers/protobuf-go


 go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc


google/api/annotations.proto
google/api/field_behaviour.proto
google/api/http.proto
google/api/httbody.proto


https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/

https://github.com/grpc-ecosystem/grpc-gateway/blob/master/protoc-gen-openapiv2/options/annotations.proto

https://github.com/grpc-ecosystem/grpc-gateway/blob/master/protoc-gen-openapiv2/options/openapiv2.proto