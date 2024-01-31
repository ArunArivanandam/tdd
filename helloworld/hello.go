package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(hello("Arun"))
}

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
