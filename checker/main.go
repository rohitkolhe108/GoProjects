package main

//go routine (separate Thread), channel (communication between go routines), function literals (annonymous function)
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.golang.org",
		"http://www.facebook.com",
	}
	//channel uses to communicate between main and child routines
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
		//fmt.Println(<-c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// pause checker for 2 second, instead of choosing main or checkLink routine to postpone. We have created another routine with function literal (aka anonymous function)
		go func(link string) {
			time.Sleep(2 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "link might be down"
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	//c <- "link is up"
	c <- link
}
