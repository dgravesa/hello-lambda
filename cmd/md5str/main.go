package main

import (
	"fmt"
	"os"

	"github.com/dgravesa/hello-lambda/pkg/md5hex"
)

func main() {
	strings := os.Args[1:]

	if len(strings) == 0 {
		// use stdin
		hexSum, err := md5hex.ReadSum(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Println(hexSum)
	} else {
		// loop over arguments, printing sum of each
		for _, s := range strings {
			hexSum := md5hex.Sum(s)
			fmt.Println(hexSum)
		}
	}
}
