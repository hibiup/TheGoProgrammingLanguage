package Expr

/**
一个计算表达式的例子：
一个表达式就是一个运算字面量，比如一个简单的表达式字面量: x+1。
表达式计算就是: 假如给定条件 x = 1.0，结果就得到 2.0。所以表达式计算由四个部分构成：
  表达式字面量(x+1)
  变量(x)
  变量值(1.0)
  运算符(+)

表达式字面量是一个字符串，变量条件用 Env 变量来传入
*/

import (
	"fmt"
	"math"
)

/**
下面定义基本数据类型。一个最基本的表达式是: 条件变量名(Var) = 变量值字面量(literal)，例如 x
*/
// 先定义一个存放变量的 map
type Env map[Var]float64

// 变量名, e.g., x.
type Var string

// 变量值字面量，浮点数, e.g., 3.141.
type literal float64

/**
定义表达式接口：
*/
// 一个表达式必须具有求值函数 Eval()。因为表达式会逐层分解，因此每一级求值都是递归的，直到最小表达式：变量和变量值字面量上。因此输入的是条件列表，最终返回结果。
type Expr interface {
	Eval(env Env) float64 // Eval 函数为env变量，返回对应的值
}

/**
定义对表达式的运算操作。运算操作具有层级关系，因此会实现递归，每一级递归的对象都是更小的表达式。
*/
// 单操作数表达式, e.g., -x.
type unary struct {
	op rune // 操作符 '+', '-'
	x  Expr // 作用的表达式
}

// 双操作数的表达式, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr // 作用的表达式
}

// 函数表达式: sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

/**
定义 Expr 接口函数，所有的东西，包括变量，变量值字面量，操作都是表达式的组成，所以都应该具有这个接口的 Eval 方法。
*/
// Var 根据条件列表返回变量的值
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// 根据变量值字面量得到实际值，因为直接根据字面量求值，所以参数可省
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// 一元运算的实现
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env) // 递归
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// 二元运算的实现
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env) // 递归
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// 函数表达式的实现
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
