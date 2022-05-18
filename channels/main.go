package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link)
	}

	// this is equivalent to writing for {}. When iterating over a chan with `range`, we are just waiting for chan to return something and going to store that in l
	for l := range c {
		// check the link again after we get some message from it
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(c, link)
		}(l)
		// one important thing is that above, we pass in l into the function literal.
		// had we not done that (and just called checkLink(c, l)), then we would've run into issues
		// this is because we are using the same l that the main routine is updating. So when checkLink goes into its child goroutine and hits the blocking call,
		// the main routine continues the for loop and updates the l variable. When the first checkLink gets a response and prints out
		// "[link] is up", link is not the same one that we started out with - it was updated by the main routines for loop
	}
}

func checkLink(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Received error from: %v. This link might be down\n", link)
		c <- link
		return
	}
	fmt.Printf("%v is up\n", link)
	// we push the link back into the channel
	c <- link
}
