package main

import (
	"testing"
	"fmt"
)

func TestHelloEmptyArg(t *testing.T) {
	input := ""
	result := hello(input)
	expected := hello("Dude")
	t.Logf("Input: '%v'\n", input)
	if result != expected {
		t.Errorf("hello(\"%v\") failed, expected '%v', got '%v'", input, expected, result)
	} else {
		t.Logf("hello(\"%v\") success, expected '%v', got '%v'", input, expected, result)
	}
}

func TestHelloValidArg(t *testing.T) {
	input := "Mike"
	result := hello(input)
	expected := fmt.Sprintf("Hello, %v!", input)
	t.Logf("Input: '%v'\n", input)
	if result != expected {
		t.Errorf("hello(\"%v\") failed, expected '%v', got '%v'", input, expected, result)
	} else {
		t.Logf("hello(\"%v\") success, expected '%v', got '%v'", input, expected, result)
	}
}

func TestHello(t *testing.T) {
	emptyResult := hello("")
	t.Logf("Input: '%v'\n", emptyResult)
	if emptyResult != "Hello, Dude!" {
		t.Errorf("hello(\"\") failed, expected '%v', got '%v'", "Hello, Mike!", emptyResult)
	}

	result := hello("Mike")
	t.Logf("Input: '%v'\n", result)
	if result != "Hello, Mike!" {
		t.Errorf("hello(\"Mike\") failed, expected '%v', got '%v'", "Hello, Mike!", result)
	}
}
