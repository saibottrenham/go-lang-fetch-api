package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// json payload example with struct below
// {
//   "to" : [
// 	{"email": "tobiasmahnert@gmail.com"},
// 	{"email": "tobiasmahnert@web.de"}
//   ],
// "from" : {
//   "email_address_id": 16426,
//   "name": "tobias"
// },
// "schedule": 1624166184,
// "subject": "test subject",
// "body": "test body"
// }
type payload_struct struct {
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	From struct {
		EmailAddressID int    `json:"email_address_id"`
		Name           string `json:"name"`
	} `json:"from"`
	Schedule int    `json:"schedule"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

func sender(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t payload_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	// change the name to uppercase
	t.From.Name = strings.ToUpper(t.From.Name)
	// encode struct to json
	json_payload, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	// send the new name to the receiver endpoint
	resp, err := http.Post("http://localhost:8082/receiver", "application/json", strings.NewReader(string(json_payload)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("receiver sent the following response ->", resp.Status, string(data))
}

func receiver(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	// check if payload matches the json payload example
	if err != nil {
		panic(err)
	}
	var t payload_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	// return response of t.From.Name which should be the uppercase name of the
	// oiginal payload
	rw.Write([]byte(t.From.Name))
}

func main() {
	http.HandleFunc("/sender", sender)
	http.HandleFunc("/receiver", receiver)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
