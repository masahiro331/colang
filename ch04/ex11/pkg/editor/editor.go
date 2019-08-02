package editor

import (
	"bufio"
	"fmt"
	"os"
)

func MessageInput(m string) (string, error) {
	fmt.Print(m)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	res := s.Text()
	return res, nil
}
