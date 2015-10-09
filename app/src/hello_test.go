package main

import "testing"
import "log"

func TestShowHello(t *testing.T) {
	testString := show()
	if testString != "Hello" {
		log.Fatalf("expected Hello, got %v", testString)
	}
}
