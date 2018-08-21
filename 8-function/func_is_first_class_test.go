package main

import (
	"testing"
	"fmt"
)

func TestFunctionIsFirstClass(t *testing.T) {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}