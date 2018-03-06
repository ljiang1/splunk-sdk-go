package main

import (
	"fmt"
	"net/http"
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
	//resp, err := http.Get("http://notexist:8000")
	resp, err := http.Get("http://splunkworld:8000")
	//resp, err := http.Get("http://google.com")

	fmt.Print(resp)
	fmt.Print(err)
	fmt.Printf("\n\n")

	if resp.Status != "200 OK" {
		fmt.Print(resp.Status)

		return false
	}

	return true
}
