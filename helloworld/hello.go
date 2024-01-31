package main

import "fmt"

func main() {
	fmt.Println(hello("Arun"))
}

func hello(name string) string {
	return "Hello, " + name
}