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
	modeAscii := flag.Bool("a", false, "convert to ascii")
	modeUnicode := flag.Bool("u", false, "convert to unicode")
	modeVersion := flag.Bool("v", false, "print punycoder version")
	modeUsage := flag.Bool("h", false, "print usage")
	flag.Parse()

	host := parseArgs(flag.Args())

	if *modeVersion {
		printVersion()
	} else if *modeAscii {
		fmt.Println(punycoder.EncodeAscii(host))
	} else if *modeUnicode {
		fmt.Println(punycoder.EncodeUnicode(host))
	} else if *modeUsage {
		punycoderUsage()
	} else {
		punycoderUsage()
	}
}

func punycoderUsage() {
	fmt.Fprintf(os.Stdout, "Usage: punycoder [-a] host | [-u] host | [-v] | [-h]\n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("punycoder[%s]\n", punycoder.Version)
}

func parseArgs(args []string) string {
	if len(args) >= 1 {
		host := args[0] //because of Ipv6
		return host
	} else {
		return ""
	}
}
