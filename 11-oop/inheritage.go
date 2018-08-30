package oop

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct {
	Point // ColoredPoint 获得了 Point (定义在 oop.go 中) 的属性，但是这不是“继承”，它只是声明了一个匿名成员。
	Color color.RGBA
}

var (
	red  = color.RGBA{255, 0, 0, 255}
	blue = color.RGBA{0, 0, 255, 255}
)

func FillColoredPoint() {
	var cp ColoredPoint
	cp.X = 1                // 这是其实只是 go 的一个语法糖，这里没有继承，访问的其实是匿名成员的成员变量。
	fmt.Println(cp.Point.X) // 或显式地访问成员的属性
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{5, 4}, blue}
	ComputeDistance(p, q)
}

func ComputeDistance(p, q ColoredPoint) {
	// p 同样可以隐式访问成员的方法（定义在 oop.go 中）
	fmt.Println(p.Distance(q.Point)) // Distance() 接受 Point 类型，因此需要显示取得匿名成员。因为没有继承，因此不能强制转换。

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

func FunctionValue() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance                  // 绑定变量 p 并获得 Method value (此刻 p 的值会被复制并入栈，Method value 非常类似一个闭包！)
	fmt.Printf("%p\t%f\n", &p, distanceFromP(q)) // "5"

	p = Point{4, 6}                              // 修改 p 的值（相同地址）不会影响 Method value 因为 Method value 已经将原值入栈了
	fmt.Printf("%p\t%f\n", &p, distanceFromP(q)) // 依然返回 5，p 的改变不会改变 Method value 执行的结果

	// 相比以下两种方式可以在 method value 计算的时候再传入接收器，但是这样的话，就不能绑定接收器。
	p = Point{6, 6}
	prototype_distanceFromP := Point.Distance
	fmt.Printf("%p\t%f\n", &p, prototype_distanceFromP(p, q)) // 2.0
	p = Point{5, 6}
	fmt.Printf("%p\t%f\n", &p, prototype_distanceFromP(p, q)) // 1.0

	p = Point{1, 2}
	pointer_distanceFromP := (*Point).Distance
	fmt.Printf("%p\t%f\n", &p, pointer_distanceFromP(&p, q)) // 5.0
	p = Point{4, 6}
	fmt.Printf("%p\t%f\n", &p, pointer_distanceFromP(&p, q)) // 0.0
}
