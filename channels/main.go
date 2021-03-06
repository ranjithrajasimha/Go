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
		"http://amazon.com",
		"http://golang.org",
		"http://youtube.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(lk string, c chan string) {
	_, er := http.Get(lk)

	if er != nil {
		fmt.Println(lk, "is down!")
		c <- lk
		return
	}

	c <- lk
	fmt.Println(lk, "is fine!")
}
