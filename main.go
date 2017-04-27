package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/ulikunitz/xz/lzma"
)

var filename = flag.String("f", "", "file to extract")

func main() {
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	file.Seek(6, 0)
	r, err := lzma.NewReader(file)
	if err != nil {
		log.Fatalf("NewReader error %s", err)
	}
	if _, err = io.Copy(os.Stdout, r); err != nil {
		log.Fatalf("io.Copy error %s", err)
	}
}
