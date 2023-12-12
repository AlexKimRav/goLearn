package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }

	checkWithSync(links)
}

func checkWithSync(links []string) {
	var wg sync.WaitGroup

	for _, url := range links {
		wg.Add(1)

		go func(link string) {
			checkLinkWithSync(link)
			wg.Done()
		}(url)
		wg.Wait()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}

func checkLinkWithSync(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	defer resp.Body.Close()
	fmt.Println(link, "is up")
}
