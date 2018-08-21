package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

import "C"

// Go调用汇编和C: https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.1.html
func add(a, b uint64) uint64

// 多值返回: https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.2.html
func sub(x, y int)(z int) {
	z = x - y
	return
}

func main() {
	fmt.Printf("%T => %d", sub, sub(2, 1))
	fmt.Printf("%T => %d", add, add(2, 1))

	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
