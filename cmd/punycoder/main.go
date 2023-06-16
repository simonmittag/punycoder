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
	flag.Usage = printUsage
	modeAscii := true
	modeUnicode := flag.Bool("u", false, "convert to unicode")
	if *modeUnicode == true {
		modeAscii = false
	}
	modeVersion := flag.Bool("v", false, "print punycoder version")
	modeUsage := flag.Bool("h", false, "print usage")
	err := parseFlags()

	if err != nil {
		printUsage()
	} else {
		if *modeVersion {
			printVersion()
		} else if *modeUsage {
			printUsage()
		} else if *modeUnicode {
			host := parseHost(flag.Args())
			fmt.Println(punycoder.EncodeUnicode(host))
		} else if modeAscii {
			host := parseHost(flag.Args())
			fmt.Println(punycoder.EncodeAscii(host))
		} else {
			printUsage()
		}
	}
}

func printUsage() {
	printVersion()
	fmt.Fprintf(os.Stdout, "Usage: punycoder host [-v] | [-h] | [-u] host \n")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("🌐 punycoder[%s]\n", punycoder.Version)
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

func parseFlags() error {
	return parseFlagSet(flag.CommandLine, os.Args[1:])
}

func parseFlagSet(flagset *flag.FlagSet, args []string) error {
	var positionalArgs []string
	for {
		if err := flagset.Parse(args); err != nil {
			return err
		}
		args = args[len(args)-flagset.NArg():]
		if len(args) == 0 {
			break
		}
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}
	return flagset.Parse(positionalArgs)
}
