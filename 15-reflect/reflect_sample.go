package reflect_sample

import (
	"fmt"
	"reflect"
	"strings"
)

func change_value_by_reflect(x int) { // 假设 x = 2
	px := &x
	*px = 3
	fmt.Println(x) // x = 3

	/* 也可以用 reflect.ValueOf(&x).Elem() 来取得地址，对于有些不可取地址的变量也适用。 */
	d := reflect.ValueOf(&x).Elem()  // d refers to the variable x
	px = d.Addr().Interface().(*int) // 取得 d 后强制通过类型断言转换类型，等同于 px := &x
	*px = 4                          // x = 4
	fmt.Println(x)

	d.Set(reflect.ValueOf(5)) // 或者直接通过 set 来更新值，免去类型转换。
	fmt.Println(x)
}

func PrintMethods(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}
