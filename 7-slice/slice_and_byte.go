package main

import (
	"fmt"
)

func main() {
	months := [...]string{
		1: "January",
		2: "Feb",   // 省略了 0 和 3，数组会预留两个空位
		4: "April",
		5: "May",
		6: "June",
	}
	fmt.Printf("Origin:\t\t%s\n", months)
	months[4] = "APRIL"
	fmt.Printf("Changed:\t%s\n", months)

	fab_to_april := months[2:5]   // 取 2 到 4 月, 3月是空的，5月不包含
	fmt.Printf("\nmonths:\t\t%s\nfab_to_april:\t%s\n", months, fab_to_april)
	// slice 的 capability 指当前数组的剩余空间。
	fmt.Printf("Len/Cap of months %d/%d\nLen/Cap of fab_to_april %d/%d\n", len(months),cap(months),len(fab_to_april),cap(fab_to_april))

	fab_to_april[1] = "MAR"                                                     // 向 slice 中添加 3 月
	fmt.Printf("\nmonths:\t\t%s\nfab_to_april:\t%s\n", months, fab_to_april)    // 会同时添加 months 数组，因为他们共用相同地址。

	fab_to_april = append(fab_to_april, "July")       // 向 slice 上 append，也会修改原 array，因此会发生覆盖
	fmt.Printf("\nmonths:\t\t%s\nfab_to_april:\t%s\n", months, fab_to_april)

	fab_to_april = append(fab_to_april, "Auguest")    // slice 到达当前数组空间的尽头, capability也就用完了
	fmt.Printf("Len/Cap of months %d/%d\nLen/Cap of fab_to_april %d/%d\n", len(months),cap(months),len(fab_to_april),cap(fab_to_april))

	fab_to_april = append(fab_to_april, "September")                           // 继续 append 超出原数组空间
	fmt.Printf("\nmonths:\t\t%s\nfab_to_april:\t%s\n", months, fab_to_april)   // 并不会改变原数组长度，但却修改了 slice 的长度。
	// (page:129) 因为 slice 被复制到了新的空间，并预留了新的 capability，和原数组分道扬镳
	fmt.Printf("Len/Cap of months %d/%d\nLen/Cap of fab_to_april %d/%d\n", len(months),cap(months),len(fab_to_april),cap(fab_to_april))

	// months = append(months,"SEP")  // append array 会失败，因为 array 长度不可变
	// may_to_july := months[5:7]     // 也不可得到超过长度的 slice。

	fab_to_april[0] = "FEB"                                                    // slice 超出后会复制到新的地址上，不再和 array 共享
	fmt.Printf("\nmonths:\t\t%s\nfab_to_april:\t%s\n", months, fab_to_april)   // 因此这次只有 slice 被更改，而 months 不变。
}
