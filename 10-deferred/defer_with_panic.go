package deferred

import (
	"fmt"
)

type bailout struct{
	Exception_code int
	Exception_message string
}

func defer_with_panic() {
	fmt.Println("start doing something!")
	defer defer_generator()()

    panic(bailout{1,"Something happened"})

	fmt.Println("a_bunch_of_defers() ending!")
}

func defer_generator() func() (err error) {
	fmt.Println("defer_generator()")

	return func() (err error){
		switch p := recover(); p.(type) {
		case nil:
			fmt.Println("No panic!")
		case bailout:
		  	err = fmt.Errorf("internal error: %s (Code: %d)", p.(bailout).Exception_message, p.(bailout).Exception_code)
			fmt.Printf("defer func(): %s\n", err)
		default:
			fmt.Println("Unknown panic!")
		}
		return
	}
}