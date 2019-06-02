package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/masahiro331/colang/ch02/ex02/tempconv"
)

func main() {
	for _, arg := range inputParam() {
		f := tempconv.Fahrenheit(arg)
		c := tempconv.Celsius(arg)
		fmt.Printf("%s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

func inputParam() (result []float64) {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			param, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				log.Fatalf("Can not Atoi input parameter %v", err)
			}
			result = append(result, param)
		}
		return result
	} else {
		for _, arg := range os.Args[1:] {
			param, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Fatalf("Can not Atoi input parameter %v", err)
			}
			result = append(result, param)
		}
		return result
	}
	return result
}
