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

func TestEncodeUULEString(t *testing.T) {
	uule, err := encodeUULEString("Houston,Texas,United States")
	if err != nil {
		t.Error("Houston,Texas,United States did not encode correctly")
	}

	if uule != "w+CAIQICIbSG91c3RvbixUZXhhcyxVbml0ZWQgU3RhdGVz" {
		t.Error("Houston,Texas,United States did not encode correctly")
	}
}
