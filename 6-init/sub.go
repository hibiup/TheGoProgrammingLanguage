package main

import (
    "fmt"
)

func init() {
	fmt.Printf("sub.init(): n=%d\n", n)
	n = 3
}

func getNumber() int {
    return n
}
