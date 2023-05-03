package punycoder

import (
	"golang.org/x/net/idna"
)

func EncodeAscii(p string) string {
	al, _ := idna.ToASCII(p)
	return al
}