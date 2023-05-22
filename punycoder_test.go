package punycoder

import (
	"testing"
)

func TestConversion(t *testing.T) {
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
			name:   "ascii to chinese",
			in:     "xn--fiq8iy4u6s7b8bb.com",
			out:    "中国互联网.com",
			method: "u",
		},
		{
			name:   "emoji to ascii",
			in:     "1❤️.ws",
			out:    "xn--1-7iqv272g.ws",
			method: "a",
		},
		{
			name:   "ascii to emoji",
			in:     "xn--1-7iqv272g.ws",
			out:    "1❤️.ws",
			method: "u",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			if tt.method == "a" {
				got = EncodeAscii(tt.in)
			}
			if tt.method == "u" {
				got = EncodeUnicode(tt.in)
			}
			want := tt.out
			if got != want {
				t.Errorf("uh oh, punycoder want %v but got %v", want, got)
			} else {
				t.Logf("ok punycoder encode %v to %v method %v", tt.in, tt.out, tt.method)
			}
		})
	}
}
