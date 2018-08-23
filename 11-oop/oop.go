package oop

import (
	"math"
)

/** 例1) golang 允许给任意对象，包括 struct, list 等加上方法。
 * 左边括号的变量名叫“接收器”，可以随意定义
 */
type Point struct{ 
	X, Y float64
}

 func (orig Point)Distance(dest Point) float64 {
    return math.Hypot(orig.X-dest.X, orig.Y-dest.Y)
}

/** 例2) 通过 Point struct 得到一个 [] 数组 
 * 然后给数组也加上 Distance() 方法。
 */
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

/** 例3) 接收器如果是指针，效果就是直接修改了接收器的值
*/
func (p *Point)InnerScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 如果接收器不是指针，则应该返回一个拷贝。因为接收器本身不会被修改。
func (p Point)ScaleBy(factor float64)(Point) {
	p.X *= factor
	p.Y *= factor
	return p
}