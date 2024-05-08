package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#write-the-test-first-1
