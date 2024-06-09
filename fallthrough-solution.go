// What is the output of this program? and Why?
package main

import "fmt"

func fallThroughSolution() {
	for _, name := range []string{
		"BeagleBoy",
		"Huey",
		"Dewey",
		"Louie",
		"Mickey",
	} {
		if checkSolution(name) {
			fmt.Println("Hello", name)
		} else {
			fmt.Println(name, ", you shall not pass!")
		}
	}
}

// Returns true if name is a legit nephew
func checkSolution(name string) bool {
	switch name {
	case "Huey":
		fallthrough
	case "Dewey":
		fallthrough
	case "Louie":
		return true
	}
	return false
}
