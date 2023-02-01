package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "Hippo"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Print(Hello("Hippo"))
}
