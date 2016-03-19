package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	var hashType = flag.String("sha", "sha256", "sha256|sha384|sha512")
	flag.Parse()

	var buf bytes.Buffer
	for idx, v := range os.Args[1:] {
		if idx != 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(v)
	}

	switch *hashType {
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(buf.Bytes()))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(buf.Bytes()))
	default:
		fmt.Printf("%x\n", sha256.Sum256(buf.Bytes()))
	}
}
