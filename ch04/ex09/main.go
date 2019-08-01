package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		countWords(os.Args[i])
	}
}

func countWords(path string) {
	ws := make(map[string]int)
	f, _ := os.Open(path)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		ws[scanner.Text()]++
	}
	for w, n := range ws {
		fmt.Printf("%v: %d回\n", w, n)
	}
}
