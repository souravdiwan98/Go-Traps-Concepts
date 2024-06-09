// What is the output of this program? and Why?
package main

import "fmt"

func fallThrough() {
	for _, name := range []string{
		"BeagleBoy",
		"Huey",
		"Dewey",
		"Louie",
		"Mickey",
	} {
		if check(name) {
			fmt.Println("Hello", name)
		} else {
			fmt.Println(name, ", you shall not pass!")
		}
	}
}

// Returns true if name is a legit nephew
func check(name string) bool {
	switch name {
	case "Huey":
	case "Dewey":
	case "Louie":
		return true
	}
	return false
}
