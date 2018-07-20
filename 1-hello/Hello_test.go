package main

import "testing"

/**
  测试文件必须以 "_test.go" 结尾。"go build" 会忽略 _test.go 文件，但是 "go test" 指令会执行它们.append

  测试函数必须以 "Test" 开头加大写开头的后缀。
  
  "t" 参数是用于报告测试失败和附加的日志信息。
*/

var b string = "Test"

func TestHello(t *testing.T) {
    s := "Test"
    s = "Test again"
	  say(s)

    b = "Changed"
    say(b)
  }