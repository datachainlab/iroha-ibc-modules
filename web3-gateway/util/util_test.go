package util

import (
	"testing"
)

func TestAccountIDtoBurrowAddress(t *testing.T) {
	address, err := AccountIDToAddress("admin@test")
	if err != nil {
		t.Error(err)
	}

	expected := "f205c4a929072dd6e7fc081c2a78dbc79c76070b"
	if address != expected {
		t.Errorf("not equal %s:%s", address, expected)
	}
}
