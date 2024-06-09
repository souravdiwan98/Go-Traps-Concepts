// Inside the if else block, now i is not redeclared. It is just reassigned a new value.
// This is the correct way to reassign a new value to an existing variable.
// The original code redeclared i inside the if else block, which created a new variable with the same name.
package main

import "fmt"

func collatzConjectureSolution() {
	// If this loops forever, I might have broken the Collatz conjecture!!
	i := 1859
	fmt.Println(i)
	for i != 1 {
		if i%2 == 0 {
			i = i / 2
			fmt.Println("i/2\t=", i)
		} else {
			i = 3*i + 1
			fmt.Println("3i + 1\t=", i)
		}
	}
}
