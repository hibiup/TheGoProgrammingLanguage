package main

import (
	"sort"
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {    // 不能使用 "func visitAll(items []string) {...}" 的形式, go 不支持有名函数嵌套(p185)。
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
