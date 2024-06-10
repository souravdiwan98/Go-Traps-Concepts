package main

import "fmt"

var a = 234
var b = 123
var c = 012

func addition() {
	fmt.Println(a + b + c)
}

// C is prefixed with 0, it is considered as octal number.
// 012 = 0 * 8^2 + 1 * 8^1 + 2 * 8^0 = 10
