package main

import "fmt"

func main() {
	slice := []string{"apple", "beer", "beer", "apple"}
	fmt.Println(uniq(slice))
}

func uniq(slice []string) []string {
	i := 0
	for _, str := range slice {
		if i == 0 || str != slice[i-1] {
			slice[i] = str
			i++
		}
	}
	return slice[:i]
}
