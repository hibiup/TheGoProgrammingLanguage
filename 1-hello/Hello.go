/*
  go doc can see here  
*/


// 每个源文件以包的声明语句 package 开始，说明该源文件是属于哪个包。包名可以自定义
package main

// 依赖包
import (
	  "fmt"
		"os"
		"strings"
)

/*
  然后是包一级的类型、变量、常量、函数的声明语句. 
  Go语言主要有四种类型的声明：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。
	包一级的各种类型的声明语句的顺序无关紧要
	
	零值初始化机制可以确保每个声明的变量总是有一个良好定义的值，因此在Go语言中不存在未初始化的变量。
*/

// 在包一级声明语句声明的名字（变量、常量、函数）可在整个包对应的每个源文件中访问，而不是仅仅在其声明语句所在的源文件中访问。
var greeting string = "Hello"

func main() {
		say(strings.Join(os.Args[1:], " "))
}

func say(msg string) {
		/* 局部声明的名字就只能在函数内部很小的范围被访问。在函数内部，有一种称为简短变量声明语句的形式可用于声明和初始化局部变量。
			 它以“名字:= 表达式”形式声明变量，变量的类型根据表达式来自动推导。

			 := 有两用。如果变量被声明过，那么 := 就是赋值，但是变量声明语句中必须至少要声明一个新的变量, 例如：
			   in, err := os.Open(infile)
				 out, err := os.Create(outfile)
			 但是下面不能通过：
			   f, err := os.Open(infile)
				 f, err := os.Create(outfile)   // false
		*/
		separater := ", "

	  fmt.Println( greeting + separater + msg )
}
