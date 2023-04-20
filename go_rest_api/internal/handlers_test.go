package server

import "testing"

func TestAlwaysReturnsSuccess(t *testing.T) {
	a := 10
	b := 20

	if a+b < 0 {
		t.Error("a + b cannot be a negative value")
	}
}
