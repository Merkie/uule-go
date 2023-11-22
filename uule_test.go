package uule

import (
	"testing"
)

func TestCreateUULEFromParams(t *testing.T) {
	uule, err := CreateUULE("Houston", "Texas", "United States")
	if err != nil {
		t.Error("Houston, Texas, United States did not encode correctly")
	}

	if uule != "w+CAIQICIbSG91c3RvbixUZXhhcyxVbml0ZWQgU3RhdGVz" {
		t.Error("Houston, Texas, United States did not encode correctly")
	}
}

func TestCreateUULEFromLocationString(t *testing.T) {
	uule, err := CreateUULEFromString("Houston,Texas,United States")
	if err != nil {
		t.Error("Houston,Texas,United States did not encode correctly")
	}

	if uule != "w+CAIQICIbSG91c3RvbixUZXhhcyxVbml0ZWQgU3RhdGVz" {
		t.Error("Houston,Texas,United States did not encode correctly")
	}
}

func TestDecodeUULEString(t *testing.T) {
	location, err := DecodeUULEString("w+CAIQICIbSG91c3RvbixUZXhhcyxVbml0ZWQgU3RhdGVz")
	if err != nil {
		t.Error("Houston,Texas,United States did not decode correctly")
	}

	if location != "Houston,Texas,United States" {
		t.Error("Houston,Texas,United States did not decode correctly")
	}
}
