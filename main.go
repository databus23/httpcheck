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
		fmt.Fprintf(os.Stderr, "Usage: httpcheck [URL ...]\n")
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
		// Don't follow redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	results := make(chan string)
	var wg sync.WaitGroup

	for _, url := range flag.Args() {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			start := time.Now()
			msg, err := check(u)
			elapsed := time.Since(start)
			if err != nil {
				results <- fmt.Sprintf("  %-30s : %s", u, ansi.Color(err.Error(), "red"))
			} else {
				results <- fmt.Sprintf("  %-30s : %s (%s)", u, ansi.Color(msg, "green"), elapsed)
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
