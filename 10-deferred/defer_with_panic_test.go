package deferred

import (
	"testing"
	"fmt"
)

func TestDeferWithPanic(t *testing.T) {
	err := defer_with_panic()
	fmt.Printf("testcase: %s\n", err)
}
