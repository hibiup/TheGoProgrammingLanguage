package oop

import (
	"testing"
	"fmt"
)

func TestMapOOP(t *testing.T) {
	m := Values{"lang": {"en"}}     // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))      // "en"
	fmt.Println("<EMPTY>", m.Get("q"))         // ""
	fmt.Println(m.Get("item"))      // "1" (first value)
	fmt.Println(m["item"])          // "[1 2]" (direct map access)
	m = nil
	fmt.Println(m.Get("item"))      // ""
	m.Add("item", "3")
}