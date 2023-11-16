package uule

import (
	"testing"
)

func TestGetPrefix(t *testing.T) {
	prefix, exists := getLengthPrefix(15)
	if !exists {
		t.Errorf("Expected prefix to exist")
	}

	if prefix != "P" {
		t.Errorf("Expected prefix to be P, got %s", prefix)
	}
}
