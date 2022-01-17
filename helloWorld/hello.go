package main

import "fmt"

func main() {
	fmt.Println(Hello("Paul"))
}

func Hello(name string) string {
	return "Hello, " + name
}
