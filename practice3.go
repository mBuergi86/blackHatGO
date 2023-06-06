package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (r *FooReader) Reader(p []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(p)
}

type FooWriter struct {
}

func (w *FooWriter) Writer(p []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader FooReader
		wrtier FooWriter
	)

	input := make([]byte, 4096)

	s, err := reader.Reader(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = wrtier.Writer(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
