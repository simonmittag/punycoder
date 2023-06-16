package punycoder

import (
	"golang.org/x/net/idna"
)

const Version string = "v0.2.2"

func EncodeAscii(p string) string {
	al, _ := idna.ToASCII(p)
	return al
}

func EncodeUnicode(p string) string {
	ul, _ := idna.ToUnicode(p)
	return ul
}
