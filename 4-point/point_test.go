package point_sample

import (
    "testing"
    "fmt"
)

func TestVariablePoint(t *testing.T) {
	x := 1				// 定义一个 int 变量 x
	p := &x				// 指向 x 的指针。（和 C 语言类似）
    fmt.Println(*p)		// *p 得到 p 指针指向的值
    fmt.Println(p)
    
    x = 2               // 变量是 mutable 的
    fmt.Println(*p)     // 证明该指针指向的值确实改变了。而不是像 python 那样是 immutable 的
    fmt.Println(&x)

	*p = 2				// 修改指针指向的内容为 2
	fmt.Println(x)		// 也就等同修改了 x
}

func TestFunctionPoint(t *testing.T) {
	p1 := TestFunc1()   // 函数可以返回指针
    fmt.Println(p1)
	fmt.Println(*p1)
    
	fmt.Println(p1 == TestFunc1())	// 但是每次可能会返回不同的值(指针)

	pf := TestFunc1 		// 获得函数指针
	fmt.Println(*pf())		// 执行函数指针，并访问返回的指针指向的值  "1"
}
