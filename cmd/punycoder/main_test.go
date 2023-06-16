package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"unicode"
)

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
