package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

const (
	MatchTitle      = 10
	MatchAlt        = 5
	MatchTranscript = 1
)

type XKCD struct {
	URL        string
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Title      string `json:"title"`
}

type XKCDSlice struct {
	Entries []*XKCD `json: entries`
}

func main() {
	if !isExists("index.json") {
		makeIndex()
	}
	file, err := os.Open("index.json")
	if err != nil {
		log.Fatal(err)
	}

	var xkcds XKCDSlice
	json.NewDecoder(file).Decode(&xkcds)
	input, err := MessageAndInput("> ")
	if err != nil {
		log.Fatal(err)
	}
	urlScore := searchTerms(strings.Fields(input), xkcds.Entries)

	type kv struct {
		Key   string
		Value int
	}
	urlScoreList := make([]kv, 0)
	for k, v := range urlScore {
		urlScoreList = append(urlScoreList, kv{k, v})
	}
	sort.Slice(urlScoreList, func(i, j int) bool {
		return urlScoreList[i].Value > urlScoreList[j].Value
	})

	urlContent := makeURLContentMap(xkcds.Entries)

	for i, kv := range urlScoreList {
		url := kv.Key
		score := kv.Value
		xkcd := urlContent[url]
		fmt.Println("#######################################")
		fmt.Printf("> url: %s (score %d)\n", url, score)
		fmt.Printf("> title: %s\n", xkcd.Title)
		fmt.Println("#######################################")
		fmt.Printf("%s\n", xkcd.Transcript)
		if i >= 5 {
			break
		}
	}
}

func searchTerms(terms []string, xkcds []*XKCD) map[string]int {
	score := make(map[string]int)
	for _, term := range terms {
		for _, xkcd := range xkcds {
			if strings.Index(strings.ToLower(xkcd.Title), strings.ToLower(term)) != -1 {
				score[xkcd.URL] += MatchTitle
			}
			if strings.Index(strings.ToLower(xkcd.Alt), strings.ToLower(term)) != -1 {
				score[xkcd.URL] += MatchAlt
			}
			if strings.Index(strings.ToLower(xkcd.Transcript), strings.ToLower(term)) != -1 {
				score[xkcd.URL] += MatchTranscript
			}
		}
	}
	return score
}

func makeURLContentMap(xkcds []*XKCD) map[string]*XKCD {
	urlContent := make(map[string]*XKCD)
	for _, xkcd := range xkcds {
		urlContent[xkcd.URL] = xkcd
	}
	return urlContent
}

func makeIndex() {
	var xkcdSlice XKCDSlice
	for i := 100; i < 500; i++ {
		xkcd, err := fetch(i)
		if err == nil {
			xkcdSlice.Entries = append(xkcdSlice.Entries, &xkcd)
		}
	}
	data, err := json.Marshal(xkcdSlice)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("index.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(data)
}

func fetch(num int) (XKCD, error) {
	fmt.Printf("curl %d\n", num)
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		return XKCD{}, err
	}

	var res XKCD
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return XKCD{}, err
	}
	resp.Body.Close()
	res.URL = url

	return res, nil
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func MessageAndInput(message string) (string, error) {
	fmt.Print(message)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	res := s.Text()
	return res, nil
}
