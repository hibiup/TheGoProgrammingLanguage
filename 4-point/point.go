package point_sample


import (
	"fmt"
)

/*
    go 语言中的指针和 C 语言非常类似
*/
func main() {
	x := 1            // 定义一个 int 变量 x
	p := &x           // 指向 x 的指针。（和 C 语言类似）
	fmt.Println(*p)   // *p 得到 p 指针指向的值

	*p = 2            // 修改指针指向的内容为 2
	fmt.Println(x)    // 也就等同修改了 x
}