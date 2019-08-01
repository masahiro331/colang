package main

import "fmt"

func main() {
	slice := []string{"apple", "beer", "beer", "apple"}
	fmt.Println(uniq(slice))
}

func uniq(slice []string) []string {
	prestr := ""
	for i, str := range slice {
		if i != 0 && prestr == str {
			slice = append(slice[:i], slice[i+1:]...)
		}
		prestr = str
	}
	return slice
}
