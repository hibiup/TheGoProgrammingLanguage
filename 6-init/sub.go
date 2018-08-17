package main

import (
    "fmt"
)

func init() {
	fmt.Printf("sub.init(): n=%d\n", n)
	n = n+1
}

func getNumber() int {
    return n
}
