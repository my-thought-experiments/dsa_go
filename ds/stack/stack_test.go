package stack

import "testing"

func TestHello(t *testing.T) {
	result := hello()

	if result != "hello world" {
		t.Error("should have returned \"hello world\"")
	}
}
