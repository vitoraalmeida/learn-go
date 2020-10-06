package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fr *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fw *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)
	/*
		input := make([]byte, 4096)

		s, err := reader.Read(input)
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		fmt.Printf("Read %d bytes from stdin\n", s)

		s, err = writer.Write(input)
		if err != nil {
			log.Fatalln("Unable to write data")
		}
		fmt.Printf("Wrote %d bytes to stdout\n", s)
	*/
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}
