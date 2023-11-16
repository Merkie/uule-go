package uule

import (
	"fmt"
	"testing"
	"time"
)

func TestLookup(t *testing.T) {
	startTime := time.Now()
	if CreateUULE("Houston", "Texas", "United States").LookupCanonicalName().CanonicalName != "Houston,Texas,United States" {
		t.Error("Houston, Texas, United States did not lookup correctly")
	}
	fmt.Printf("Lookup took %s\n", time.Since(startTime))
}
