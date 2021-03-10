protoc --grpc-gateway_out=logtostderr=true:. hello_http/exp1.proto

protoc --swagger_out=logtostderr=true:. hello_http/exp1.proto

protoc -I . --openapiv2_out=logtostderr=true:. hello_http/exp1.proto


protoc -I . --go_out=. hello_http/exp1.proto 