package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
	"unicode"
)

func TestMainFuncWithHelp(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append([]string{"-h"}, "-h")
	main()
}

func TestMainFuncWithVersion(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append([]string{"-v"}, "-v", "https://www.google.com")
	main()
}

func TestMainFuncToUnicode(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		out    string
		method string
	}{
		{
			name:   "chinese to ascii",
			in:     "中国互联网.com",
			out:    "xn--fiq8iy4u6s7b8bb.com",
			method: "a",
		},
		{
			name:   "emoji to ascii",
			in:     "🌐︎.com",
			out:    "xn--u86cs934b.com",
			method: "a",
		},
		{
			name:   "emoji 2 to ascii",
			in:     "🌐.com",
			out:    "xn--wg8h.com",
			method: "a",
		},
		{
			name:   "umlaut to ascii",
			in:     "ö.com",
			out:    "xn--nda.com",
			method: "a",
		},
		{
			name:   "Français to ascii",
			in:     "français.fr",
			out:    "xn--franais-xxa.fr",
			method: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runT(tt, t)
		})
	}
}

func runT(tt struct {
	name   string
	in     string
	out    string
	method string
}, t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	if tt.method == "u" {
		os.Args = append(os.Args, "-u")
	}
	os.Args = append(os.Args, tt.in)
	main()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	out = strings.TrimFunc(out, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})

	if tt.out != out {
		t.Errorf("wanted %v got %v", tt.out, out)
	}
}
