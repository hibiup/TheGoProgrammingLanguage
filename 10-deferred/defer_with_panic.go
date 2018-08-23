package deferred

/************
 * panic defer recover 组合是 golang 中因对异常的机制，但是与其它语言的 try catch 不同，golang 开发的哲学并不希望程序继续处理 panic
 * 之后的代码，因此它即便被 recover ，程序也无法继续执行。
 * 
 * 和java动不动就抛出异常，然后捕获，然后转回正常处理不同，golang 不认为用异常来帮助程序的业务逻辑处理是正确的，
 * 相反它建议在程序中主动考虑错误处理，而不是借助异常机制来捕获，因此 panic被设计成不可恢复的，以强制用户只用它处理真正的异常。
 */
import (
	"fmt"
)

type bailout struct{
	Exception_code int
	Exception_message string
}

func defer_with_panic()(err error) {
	fmt.Println("start doing something!")
	defer func() {
		// 不能直接获得 deferred 函数的返回值，因此 deferred 函数一般以内嵌函数的方式定义，这样就可以共享堆栈(比如本例的 err 变量)。
		switch p := recover(); p.(type) {   // panic 返回类型为 interface{}，nil interface{} 不等于 null，因此 interface{}.(type) 是安全的
		case nil:
			fmt.Println("No panic!")
		case bailout:
    		// recover 唯一能和函数正常逻辑能产生关联的地方是它可以修改返回值，仅此而已。程序不会再回到中断之后的逻辑上去了。
			err = fmt.Errorf("internal error: %s (Code: %d)", p.(bailout).Exception_message, p.(bailout).Exception_code)
			fmt.Printf("defer func(): %s\n", err)
		default:
			fmt.Println("Unknown panic!")
		}
	}()

    panic(bailout{1,"Something happened"})      // func panic(interface{})
	// panic(nil)

	fmt.Println("defer_with_panic() ending!")   // panic 会终止程序的执行，即便被 recover也不能继续，因此本行不会被执行到。
	return
}
