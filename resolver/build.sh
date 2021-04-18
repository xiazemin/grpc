go mod init resolver
go: creating new go.mod: module resolver

mkdir helloworld

%protoc -I . --go_out=plugins=grpc:./helloworld helloworld.proto

%protoc -I . --go_out=plugins=grpc:. helloworld.proto 

%go run main.go