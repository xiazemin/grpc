 
 https://github.com/leafduo/grpc-dev-proxy
 

 curl -X "POST" "http://localhost:2333/?user-id=1" \
     -H 'Target: localhost:50051' \
     -H 'Service: helloworld.Greeter' \
     -H 'Method: SayHello' \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "name": "world"
}'




