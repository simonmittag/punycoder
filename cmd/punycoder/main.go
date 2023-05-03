package main

import (
	"fmt"
	"github.com/simonmittag/punycoder"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		al := punycoder.EncodeAscii(os.Args[1])
		fmt.Printf(" %v\n", al)
	} else {
		var pl string
		fmt.Scanf("%s", &pl)
		al := punycoder.EncodeAscii(pl)
		fmt.Printf(" %v\n", al)
	}
}

