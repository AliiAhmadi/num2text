/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AliiAhmadi/num2text"
)

var (
	INP string
)

func flags() {
	flag.StringVar(&INP, "-num", "0", "number you want to convert")
	flag.Parse()
}

func main() {
	flags()

	r, err := num2text.Convert(INP)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v", r)
}
