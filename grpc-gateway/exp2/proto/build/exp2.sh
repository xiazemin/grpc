protoc --grpc-gateway_out=logtostderr=true:. hello_http/exp2.proto
 protoc -I . --openapiv2_out=logtostderr=true:. hello_http/exp2.proto 