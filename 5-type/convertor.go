package tempconv

import (
	"fmt"
)

/**
 给 Celsius 添加 String() 方法
**/
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }