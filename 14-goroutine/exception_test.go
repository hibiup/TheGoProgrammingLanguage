package goroutine_sample

import (
	"testing"
)

func TestException(t *testing.T) {
	Run(1)
	Run(5)
	Run(10)
	Run(20)
}
