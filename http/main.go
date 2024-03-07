package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		fmt.Println(resp)
	}

	bs := make([]byte, 99999)  // create slice of empty byte to store the response from http request
	n, _ := resp.Body.Read(bs) // call read function from Reader interface
	fmt.Println(string(bs))
	fmt.Println(n)
}
