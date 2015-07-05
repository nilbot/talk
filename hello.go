package main

import "fmt"

func main() {
	fmt.Println(show() + "world")
}

func show() string {
	return "Hello"
}
