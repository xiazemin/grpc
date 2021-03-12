https://github.com/nametake/protoc-gen-gohttp

go get -u github.com/nametake/protoc-gen-gohttp

cd exp1
protoc --go_out=plugins=grpc:. --gohttp_out=. greeter.proto


https://jacobmartins.com/2016/05/24/practical-golang-using-protobuffs/
cd exp2
protoc --go_out=. clientStructure.proto

 go mod init proto-http

 https://stackoverflow.com/questions/53407677/how-to-implement-googles-protocol-buffers-via-http

 https://www.cnblogs.com/baoshu/p/13574829.html