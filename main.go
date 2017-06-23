package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/mgutz/ansi"
)

var client http.Client

var timeout string

func init() {
	flag.StringVar(&timeout, "timeout", "1s", "overall timeout for http requests")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [URL ...]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatal("Invalid timeout specified")
	}

	client = http.Client{
		Timeout: timeoutDuration,
	}

	results := make(chan string)
	var wg sync.WaitGroup

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			msg, err := check(u)
			if err != nil {
				results <- fmt.Sprintf("  %-30s : %s", u, ansi.Color(err.Error(), "red"))
			} else {
				results <- fmt.Sprintf("  %-30s : %s", u, ansi.Color(msg, "green"))
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

}

func check(url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return resp.Status, nil
}
