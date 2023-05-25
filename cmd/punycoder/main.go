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

// ParseFlags parses the command line args, allowing flags to be
// specified after positional args.
func parseFlags() error {
	return parseFlagSet(flag.CommandLine, os.Args[1:])
}

// ParseFlagSet works like flagset.Parse(), except positional arguments are not
// required to come after flag arguments.
func parseFlagSet(flagset *flag.FlagSet, args []string) error {
	var positionalArgs []string
	for {
		if err := flagset.Parse(args); err != nil {
			return err
		}
		// Consume all the flags that were parsed as flags.
		args = args[len(args)-flagset.NArg():]
		if len(args) == 0 {
			break
		}
		// There's at least one flag remaining and it must be a positional arg since
		// we consumed all args that were parsed as flags. Consume just the first
		// one, and retry parsing, since subsequent args may be flags.
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}
	// Parse just the positional args so that flagset.Args()/flagset.NArgs()
	// return the expected value.
	// Note: This should never return an error.
	return flagset.Parse(positionalArgs)
}
