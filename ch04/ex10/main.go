package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masahiro331/colang/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	ltMonth := ""
	ltYear := ""
	later := ""
	for _, item := range result.Items {
		if t.AddDate(0, -1, 0).Before(item.CreatedAt) {
			ltMonth += fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		} else if t.AddDate(-1, 0, 0).Before(item.CreatedAt) {
			ltYear += fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		} else {
			later += fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%s\n", ltMonth)
	fmt.Println("-------- less than one Month ---------")
	fmt.Printf("%s\n", ltYear)
	fmt.Println("-------- less than one Year ----------")
	fmt.Printf("%s\n", later)
}
