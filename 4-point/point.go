package point_sample

/*
    go 语言中的指针和 C 语言非常类似
*/

func TestFunc1() *int {
	v := 1
	return &v			// 返回局部变量地址是安全的
}
