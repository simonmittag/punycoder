package punycoder

import (
	"golang.org/x/net/idna"
	"os"
)

func EncodeAscii(p string) string {
	al, _ := idna.ToASCII(os.Args[1])
	return al
}