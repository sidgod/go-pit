package main

import (
	"time"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("Total time taken is %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	var fixedUrl = url
	if !strings.HasPrefix(url, "http://") {
		fixedUrl = "http://" + url
	}
	response, err := http.Get(fixedUrl)
	if err == nil {
		respDataLen, err := io.Copy(ioutil.Discard, response.Body)
		if err != nil {
			ch <- fmt.Sprintf("Failed to read response data for url %s", url)
		} else {
			fmt.Printf("Read content of size %d for url %s\n", respDataLen, url)
		}
		response.Body.Close()
	} else {
		ch <- fmt.Sprintf("Failed to get content for url %s", url)
	}
	ch <- fmt.Sprintf("Time taken is %.2fs", time.Since(start).Seconds())
	return
}
