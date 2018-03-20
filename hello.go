package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var i int = 0

	webrequest()
	fmt.Printf("hello world %d", i)
}

func avg(input1, input2 int) float32 {

	var x = input1 + input2
	return float32(x / 2)
}

func webrequest() bool {
	//workaround to wait splunk service to be ready. TODO: should in splunk service configuration
	time.Sleep(0 * time.Minute)
	//resp, err := http.Get("http://notexist:8000")
	//resp, err := http.Get("http://nginxworld")
	resp, err := http.Get("http://splunkworld:8000")
	//resp, err := http.Get("http://google.com")
	//time.Sleep(3 * time.Minute)

	fmt.Print(resp)
	fmt.Print(err)
	fmt.Printf("\n\n")

	if resp.Status != "200 OK" {
		fmt.Print(resp.Status)
		return false
	}

	return true
}
