// Does the Collatz conjecture still hold for value 1859?
package main

import "fmt"

func collatzConjecture() {
	// If this loops forever, I might have broken the Collatz conjecture!!
	i := 1859
	fmt.Println(i)
	for i != 1 {
		if i%2 == 0 {
			i := i / 2
			fmt.Println("i/2\t=", i)
		} else {
			i := 3*i + 1
			fmt.Println("3i + 1\t=", i)
		}
	}
}
