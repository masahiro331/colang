package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		err := fetch(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch: %v\n", err)
	}
	io.Copy(writer, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("fetch: reading %s: %v\n", url, err)
	}
	return nil
}
