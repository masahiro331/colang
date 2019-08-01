package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/masahiro331/colang/ch04/ex11/pkg/editor"
)

const IssueURL = "https://api.github.com/repos/masahiro331/colang/issues/"

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	// Assignees []string `json:"assignees"`
	Labels []string `json:"labels"`
	State  string   `json:"state"`
	User   *User
}
type User struct {
	Login string
}

func CreateIssue(token string) error {
	issue, err := CreateIssueBody()
	if err != nil {
		return err
	}
	bodyJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", IssueURL, bytes.NewBuffer([]byte(bodyJSON)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	return nil
}

func CreateIssueBody() (Issue, error) {
	body := Issue{}

	title, err := editor.MessageAndInput("title: ")
	if err != nil {
		return Issue{}, err
	}
	body.Title = title

	body.Body, err = editor.MessageAndInput("body: ")
	if err != nil {
		return Issue{}, err
	}

	body.Title = title
	input, err := editor.MessageAndInput("asignees(, separated): ")
	if err != nil {
		return Issue{}, err
	}

	input, err = editor.MessageAndInput("labels(, separated): ")
	labels := strings.Split(input, ",")
	for _, label := range labels {
		if label != "" {
			body.Labels = append(body.Labels, label)
		}
	}
	return body, nil
}

func PatchIssue(token string, number int) error {
	url := IssueURL + strconv.Itoa(number)
	ReadIssue(token, number)
	issue, err := CreateIssueBody()
	if err != nil {
		return err
	}
	bodyJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(bodyJSON)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("patch issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}

func CloseIssue(token string, number int) error {
	url := IssueURL + strconv.Itoa(number)

	ReadIssue(token, number)
	issue := Issue{State: "closed"}
	bodyJSON, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(bodyJSON)))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("close issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return nil
}

func ListIssue(token string) error {
	req, err := http.NewRequest("GET", IssueURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("read issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	for _, issue := range result {
		fmt.Printf("#%04d, %50s, %s\n", issue.Number, issue.Title, issue.User.Login)
	}
	return nil
}

func ReadIssue(token string, number int) error {
	url := IssueURL + strconv.Itoa(number)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("read issue failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()
	fmt.Printf("#%05d: %s\n", result.Number, result.Title)
	fmt.Printf("%s\n", result.Body)
	return nil
}
