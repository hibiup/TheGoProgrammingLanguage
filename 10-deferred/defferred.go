package deferred

import (
	"fmt"
)

func a_bunch_of_defers() int {
	fmt.Println("start doing something!")

	defer defer_1()
	fmt.Println("defer_1(): ended")

	defer defer_2()
	fmt.Println("defer_2(): ended")

	defer defer_3()
	fmt.Println("defer_3(): ended")

	defer defer_2()
	fmt.Println("defer_2(): ended")

	fmt.Println("a_bunch_of_defers() ending!")
	return 0
}

func defer_1() {
	fmt.Println("defer_1(): deferring")
	fmt.Println("defer_1() calls to defer_2(): ")
	defer defer_2()
	fmt.Println("defer_1() calls to defer_2() end")
}

func defer_2() {
	fmt.Println("defer_2(): deferring")
	//panic("defer_2() is panicing")
}

func defer_3() {
	fmt.Println("defer_3(): deferring")
	fmt.Println("defer_3() calls to defer_1(): ")
	defer defer_1()
	fmt.Println("defer_3() calls to defer_1() end")
}