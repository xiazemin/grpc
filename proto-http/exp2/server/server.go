package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	pb "google.golang.org/protobuf/proto"

	"proto-http/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myClient := proto.Client{}

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println(err)
		}
		if err := pb.Unmarshal(data, &myClient); err != nil {
			fmt.Println(err)
		}
		println(myClient.Id, ":", myClient.Name, ":", myClient.Email, ":", myClient.Country)

		for _, mail := range myClient.Inbox {
			fmt.Println(mail.RemoteEmail, ":", mail.Body)
		}
	})

	http.ListenAndServe(":3000", nil)
}
