package editor

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func openEditor(e string) (string, error) {
	if e == "" {
		e = "vim"
	}
	path := os.TempDir() + "/temp"
	tempF, err := os.Create(path)
	if err != nil {
		return "", err
	}
	tempF.Close()

	cmd := exec.Command(e, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	resByte, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(resByte), "\n"), nil
}

func MessageAndInput(message string) (string, error) {
	fmt.Print(message)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	res := s.Text()
	if res == "EDITOR" {
		r, err := openEditor(os.Getenv("EDITOR"))
		if err != nil {
			return "", err
		}
		res = string(r)
	}
	return res, nil
}
