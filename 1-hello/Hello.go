package main

import (
	  "fmt"
		"os"
		"strings"
)

/**
  Main entrance
*/
func main() {
		say(strings.Join(os.Args[1:], " "))
}

func say(msg string) {
	  const greeting string = "Hello"
	  fmt.Println( greeting + ", " + msg )
}
