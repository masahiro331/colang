package main

import (
	"log"
	"os"
	"strconv"

	"github.com/masahiro331/colang/ch04/ex11/pkg/github"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		log.Fatal("ArgumentError")
	}

	token := ""
	client := github.NewGithubClient(token)

	switch os.Args[1] {
	case "create":
		err = client.CreateIssue()
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		err = client.ListIssue()
		if err != nil {
			log.Fatal(err)
		}
	case "read":
		if len(os.Args) < 3 {
			log.Fatal("ArgumentError")
		}
		issueNumber := os.Args[2]
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(issueNumber)
		if err != nil {
			log.Fatal(err)
		}
		err = github.ReadIssue(number)
		if err != nil {
			log.Fatal(err)
		}
	case "close":
		if len(os.Args) < 3 {
			log.Fatal("ArgumentError")
		}
		issueNumber := os.Args[2]
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(issueNumber)
		if err != nil {
			log.Fatal(err)
		}
		err = client.CloseIssue(number)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("ArgumentError")

	}
}
