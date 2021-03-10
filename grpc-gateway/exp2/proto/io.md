exp1

option (google.api.http) = {
    get:"/v1/{name=messages/*}"
};

message GetMessageRequest{
    string name=1; //name
}

// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456`  | `GetMessage(name: "messages/123456")`

exp2

option (google.api.http) = {
   get:"/v1/messages/{message_id}"
};

message GetMessageRequest{
  message SubMessage{
    string subfield =1;
  }
  string message_id =1;
  int64 revision =2;
  SubMessage sub=3;
}

// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456?revision=2&sub.subfield=foo` |
// `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield:
// "foo"))`


exp3

option (google.api.http) = {
  patch:"/v1/messages/{message_id}"
  body: "message"
};

message UpdateMessageRequest{
    message SubMessage{
        string subfield =1;
    }
    string message_id =1;
    int64 revision =2;
    SubMessage sub=3;
    Message message =4;
}

message Message{
    string text=1; //body response
}

// HTTP | gRPC
// -----|-----
// `PATCH /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id:
// "123456" message { text: "Hi!" })`


exp4

option (google.api.http) = {
    get: "/v1/messages/{message_id}"
    additional_bindings {
    get: "/v1/users/{user_id}/messages/{message_id}"
    }
};



// HTTP | gRPC
// -----|-----
// `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id:
// "123456")`