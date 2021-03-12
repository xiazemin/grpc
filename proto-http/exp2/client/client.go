package main
import(
    "fmt"
    "net/http"
    "proto-http/proto"
    "bytes"
    pb "google.golang.org/protobuf/proto"
)
func main() {
	myClient := proto.Client{Id: 526, Name: "John Doe", Email: "johndoe@example.com", Country: "US"}
    clientInbox := make([]*proto.Client_Mail, 0, 20)
    clientInbox = append(clientInbox,
        &proto.Client_Mail{RemoteEmail: "jannetdoe@example.com", Body: "Hello. Greetings. Bye."},
        &proto.Client_Mail{RemoteEmail: "WilburDoe@example.com", Body: "Bye, Greetings, hello."})

    myClient.Inbox = clientInbox
	data, err := pb.Marshal(&myClient)
    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

    if err != nil {
        fmt.Println(err)
        return
    }
}
