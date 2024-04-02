package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create channel for main routine to be able to communicate with its child routine(s)
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // create child routine each time this func is called
	}

	// for { // endless / infinite loop
	// 	// fmt.Println(<-c) // wait for a value to be sent into the channel, then log it out
	// 	go checkLink(<-c, c) // waiting for message then directly call the function and pass the URL back
	// }

	for l := range c { // wait channel to return value then assign it to variable l then starts the loop
		go func(link string) { // function literal
			time.Sleep(time.Second * 5) // sleep routine for every call
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down!")
		c <- link //send message to channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
