package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	const (
		httpProtocol = "http://"
	)
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, httpProtocol) == false {
			url = httpProtocol + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// http.Get→resp.Bodyで簡単にクロール出来る
// やはりこのあたり強い
