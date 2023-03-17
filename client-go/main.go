package main

import "testing"

func main() {
	clientSet()
	dynamicClient()
	discoveryCli()
}

// a function adds two numbers
func add(x, y int) int {
	return x + y
}

// Test add function
func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Errorf("add(1, 2) = %d; want 3", add(1, 2))
	}
}
