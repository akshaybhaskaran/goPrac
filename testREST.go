package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func main(){

	// Initial Get Request
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil{
		fmt.Printf("connection error")
	}else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// Serializing a JSON Object
	jsonData := map[string]string{"first":"Akshay", "last":"Bhaskaran"}
	jsonValue, _ := json.Marshal(jsonData)

	// Post request of Json Object
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil{
		fmt.Println("http request failed")
	} else{
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}