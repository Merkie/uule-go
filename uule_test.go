package uule

import (
	"testing"
)

func TestEncodeUULE(t *testing.T) {
	uule := CreateUULE("Houston", "Texas", "United States")

	if uule.Region != "Texas" {
		t.Error("Region did not set correctly")
	}

	encodedUule, err := uule.Encode()
	if err != nil {
		t.Error("Houston, Texas, United States did not encode correctly")
	}

	if encodedUule != "w+CAIQICIbSG91c3RvbixUZXhhcyxVbml0ZWQgU3RhdGVz" {
		t.Error("Houston, Texas, United States did not encode correctly")
	}
}
