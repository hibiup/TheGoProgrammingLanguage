package oop

import (
	"testing"
	"fmt"
)

func TestStructOOP(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))    // "5", method call
}

func TestArrayOOP(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}

func TestStructPoint(t *testing.T) {
	r := &Point{1, 2}
	r.InnerScaleBy(2)
	fmt.Println(*r)  // 修改为 "{2, 4}"

	r1 := Point{1, 2}
	r2 := r1.ScaleBy(2)
	fmt.Println(r1)  // 不变 "{1, 2}"
	fmt.Println(r2)  // 返回值放映了修改结果 "{2, 4}"
}
