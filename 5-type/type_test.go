package tempconv

import (
	"testing"
	"fmt"
)

func TestTypeDefination(t *testing.T) {
	fmt.Printf("Brrrr! %v\n", AbsoluteZeroC) // "Brrrr! -273.15°C". "%v", AbsoluteZeroC 会调用 Celsius.String() 方法

	fmt.Printf("%g\n", BoilingC - FreezingC)        // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC))  // "180" °F

	// compile error: type mismatch !!
	//fmt.Printf("%g\n", boilingF - FreezingC)
}