package main

import (
	"fmt"
	"github/simonmittag/punycoder"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		al := punycoder.EncodeAscii(os.Args[1])
		fmt.Printf(" %v\n", al)
	} else {
		fmt.Println("usage: punycoder a😀.com")
	}
}

