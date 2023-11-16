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
