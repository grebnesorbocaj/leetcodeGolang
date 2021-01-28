package main

import "testing"

func TestReverse(t *testing.T) {
	if reverse(-812) != -218 {
		t.Errorf("Expected response: -218, but got: %v", reverse(-812))
	}
	if reverse(0) != 0 {
		t.Errorf("Expected response: 0, but got: %v", reverse(0))
	}
	if reverse(1234) != 4321 {
		t.Errorf("Expected response: -218, but got: %v", reverse(1234))
	}
	if reverse(2147483647) != 0 {
		t.Errorf("Expected response: 0, but got: %v", reverse(2147483647))
	}
	if reverse(-2147483648) != 0 {
		t.Errorf("Expected response: 0, but got: %v", reverse(-2147483648))
	}
}
