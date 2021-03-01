~/go/ghz/cmd/ghz/ghz --insecure \
--proto ./helloworld.proto \
--call helloworld.Greeter.SayHello \
-d '{"name":"Joe"}' \
-c 10 \
-n 20 \
localhost:50151

~/go/ghz/cmd/ghz/ghz --insecure --proto ./helloworld.proto --call helloworld.Greeter.SayHello -d '{"name":"Joe"}' localhost:50151  √

~/go/ghz/cmd/ghz/ghz --proto ./helloworld.proto --call helloworld.Greeter.SayHello -d '{"name":"Joe"}' localhost:50151