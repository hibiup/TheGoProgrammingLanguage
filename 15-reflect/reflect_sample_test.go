package reflect_sample

import (
	"testing"
)

func TestChangeValueByReflect(t *testing.T) {
	change_value_by_reflect(2)
}

type Point struct {
	X, Y float64
}

func (p *Point) InnerScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleBy(factor float64) Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func TestPrintMethods(t *testing.T) {
	p := Point{1.0, 2.0}

	PrintMethods(p)
	PrintMethods(&p)
}
