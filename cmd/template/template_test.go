package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
	t.Errorf("1 + 2 expected be 3, but %d got", 1)
}
