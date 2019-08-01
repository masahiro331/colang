package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		log.Fatal("Error Arguments")
	}

	if isAnagram(args[1], args[2]) {
		fmt.Println("Anagram!")
		return
	}
	fmt.Println("Not Anagram!")
	return
}

func isAnagram(str1, str2 string) bool {
	strArray1 := strings.Split(str1, "")
	strArray2 := strings.Split(str2, "")
	sort.Strings(strArray1)
	sort.Strings(strArray2)
	if strings.Join(strArray1, "") == strings.Join(strArray2, "") {
		return true
	}
	return false
}
