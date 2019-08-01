package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var out io.Writer = os.Stdout

var (
	shatype = flag.String("type", "256", "hash type")
)

func main() {
	flag.Parse()

	if err := sha(os.Args); err != nil {
		log.Fatal(err)
	}
}

func sha(args []string) error {
	if len(args) < 2 {
		return errors.New("Invalid Argument")

	}
	switch *shatype {
	case "256":
		sb := sha256.Sum256([]byte(args[len(args)-1]))
		_, err := fmt.Fprintln(out, hex.EncodeToString(sb[:]))
		if err != nil {
			return err
		}
	case "512":
		sb := sha512.Sum512([]byte(args[len(args)-1]))
		_, err := fmt.Fprintln(out, hex.EncodeToString(sb[:]))
		if err != nil {
			return err
		}
	}

	return nil
}
