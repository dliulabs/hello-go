package main

import (
	"fmt"
	"testing"
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
	if emptyResult != "Hello, Mike!" {
		t.Errorf("hello(\"\") failed, expected '%v', got '%v'", "Hello, Mike!", emptyResult)
	} else {
		t.Logf("hello(\"%v\") success, expected '%v', got '%v'", "", "Hello, Mike", emptyResult)
	}

	result := hello("Mike")
	expected := fmt.Sprintf("Hello, %v!", "Mike")
	if result != "Hello, Mike!" {
		t.Errorf("hello(\"Mike\") failed, expected '%v', got '%v'", "Hello, Mike!", result)
	} else {
		t.Logf("hello(\"%v\") success, expected '%v', got '%v'", "Mike", expected, result)
	}
}
