package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return
	}
}

func main() {
	Greet(os.Stdout, "Dog")
}
