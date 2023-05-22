package main

import (
	"fmt"
	"github.com/simonmittag/punycoder"
	"os"
)

import (
	"flag"
)

func main() {
	flag.Usage = punycoderUsage
	modeAscii := true
	modeUnicode := flag.Bool("u", false, "convert to unicode")
	if *modeUnicode == true {
		modeAscii = false
	}
	modeVersion := flag.Bool("v", false, "print punycoder version")
	modeUsage := flag.Bool("h", false, "print usage")
	flag.Parse()

	if *modeVersion {
		printVersion()
	} else if *modeUsage {
		punycoderUsage()
	} else if *modeUnicode {
		host := parseHost(flag.Args())
		fmt.Println(punycoder.EncodeUnicode(host))
	} else if modeAscii {
		host := parseHost(flag.Args())
		fmt.Println(punycoder.EncodeAscii(host))
	} else {
		punycoderUsage()
	}
}

func punycoderUsage() {
	fmt.Fprintf(os.Stdout, "Usage: punycoder host | [-u] host | [-v] | [-h]\n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("punycoder[%s]\n", punycoder.Version)
}

func parseHost(args []string) string {
	var host = ""
	if len(args) >= 1 {
		host = args[0] //because of Ipv6
	} else {
		fmt.Scanf("%s", &host)
	}
	return host
}
