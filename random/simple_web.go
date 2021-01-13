package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)



func request() string {
	resp, err := http.Get("http://ip.jsontest.com/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	ip1 := ipResponse{}
	bodyErr := json.Unmarshal(body, &ip1)
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	return ip1.ip
}

func main() {
	fmt.Println(request())
}