package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

/**
  接口方法
*/
type MyWriter interface {
	io.Writer
}

func TestByteCounterInterface(t *testing.T) {
	byteCounter := ByteCounter(0)
	var c MyWriter = &byteCounter              // 赋予 MyWriter 接口类型的指针变量。不能用非指针实例赋值，因为接口方法只实现在指针上。
	fmt.Fprintf(c, "Hello writer!")            // Fprintf() 函数接受的第一个参数是一个具有 io.Writer.Write() 接口的对象的指针，因为这个接口只实现在指针上。
	fmt.Println("byteCounter = ", byteCounter) // 13
}

func TestRecipterInterface(t *testing.T) {
	recipter := Recipter("Hello, ")
	var c MyWriter = &recipter
	fmt.Fprintf(c, "Stranger")
	fmt.Println("recipter = ", recipter) // Hello, Stranger
}

var w io.Writer // 一个接口的零值就是他的类型值(type=io.Writer)和动态值(data)为 nil。(P224 页 的描述可能有错)

func TestInterfaceCast(t *testing.T) {
	/*
	 这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式的使用 io.Writer(os.Stdout) 是等价的
	*/
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
}

func TestTypeAssert(t *testing.T) {
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	w = rw.(io.Writer)
	w = new(ByteCounter)
	if w_, err := w.(io.ReadWriter); err { // 如果存在第二个返回值 “err”，那么就不会发生 panic
		fmt.Println(w_)
	}
}

func TestNilTypeAssert(t *testing.T) {
	f := func(s interface{}) string {
		if s == nil {
			return "NULL"
		} else if _, ok := s.(int); ok {
			return fmt.Sprintf("%d", s)
		} else if b, ok := s.(bool); ok {
			if b {
				return "TRUE"
			}
			return "FALSE"
		} else {
			return fmt.Sprintf("%s", s)
		}
	}
	fmt.Println(f(nil)) // 任何东西都是 interface{}, 包括 nil
}

func TestSwitchType(t *testing.T) {
	// switch...(type) 是一个 type 关键词的固定用法，type 不能用在非 switch 装饰的环境中

	f1 := func(s interface{}) string {
		switch x := s.(type) {
		case nil: // 匹配的是 s.(type) 的值，不是 x
			return fmt.Sprintf("NULL")
		case int, uint:
			return fmt.Sprintf("%d", s)
		case bool:
			if x {
				return "TRUE"
			}
			return "FALSE"
		default:
			return fmt.Sprintf("%s", s)
		}
	}
	fmt.Println(f1(true))
}
