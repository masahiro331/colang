package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		err := fetch(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fetch(url string) error {
	resp, err := http.Get(resolveUrl(url))
	if err != nil {
		return fmt.Errorf("fetch: %v\n", err)
	}
	fmt.Fprint(out, resp.Status+"\n\n")
	io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("fetch: reading %s: %v\n", url, err)
	}
	return nil
}

func resolveUrl(url string) string {
	if !strings.HasPrefix(url, "http://") {
		return "http://" + url
	}
	return url
}
