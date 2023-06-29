package main

import (
	"testing"
)

func TestHelloService_Hello(t *testing.T) {
	helloService := &HelloService{}

	var reply string

	// Test case 1: request is empty
	err := helloService.Hello("", &reply)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
	if reply != "Hello: " {
		t.Errorf("Expected 'Hello: ', but got '%s'", reply)
	}

	// Test case 2: request is a whitespace string
	err = helloService.Hello("   ", &reply)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
	if reply != "Hello:    " {
		t.Errorf("Expected 'Hello:    ', but got '%s'", reply)
	}

	// Test case 3: request contains special characters
	err = helloService.Hello("!@#^&*()", &reply)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
	if reply != "Hello: !@#^&*()" {
		t.Errorf("Expected 'Hello: !@#^&*()', but got '%s'", reply)
	}

	// Test case 4: request contains non-ASCII characters
	err = helloService.Hello("你好，世界！", &reply)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
	if reply != "Hello: 你好，世界！" {
		t.Errorf("Expected 'Hello: 你好，世界！', but got '%s'", reply)
	}

	// Test case 5: request is very long
	longRequest := "a"
	for i := 0; i < 100000; i++ {
		longRequest += "a"
	}
	err = helloService.Hello(longRequest, &reply)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}

	expectedReply := "Hello: " + longRequest
	if reply != expectedReply {
		t.Errorf("Expected '%s', but got '%s'", expectedReply, reply)
	}
}
