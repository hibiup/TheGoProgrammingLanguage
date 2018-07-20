package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main(){
	get(os.Args)
}

func get(args[] string) {
	start := time.Now()
	channel := make(chan string)

	for _, url := range args {
		go query(url, channel)
	}

	for range args {
		fmt.Println(<-channel)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func query(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}

	//b, err := ioutil.ReadAll(resp.Body)
	b, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, b, url)
}
