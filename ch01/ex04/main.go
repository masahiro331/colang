package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for _, file := range os.Args[1:] {
		result, err := dup(file)
		if err != nil {
			log.Fatal(err)
		}
		for line, n := range result {
			if n > 1 {
				_, err := fmt.Printf("%s\t%d\t%s\n", filepath.Base(file), n, line)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func dup(file string) (map[string]int, error) {
	counts := make(map[string]int)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		return nil, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			counts[line]++
		}
	}
	return counts, nil
}
