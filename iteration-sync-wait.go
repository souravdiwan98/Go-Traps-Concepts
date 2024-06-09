// What is the output of this program?
package main

import (
	"fmt"
	"sync"
)

func iterateSyncWait() {
	var wg sync.WaitGroup

	// Can't declare a "const" array, so this is a "var"
	var nephews = [...]string{"Huey", "Dewey", "Louie"}

	for _, nephew := range nephews {
		wg.Add(1)
		go greet(nephew, wg)
	}

	// Wait for all greetings to complete.
	wg.Wait()
}

func greet(name string, wg sync.WaitGroup) {
	fmt.Println("Hello", name)
	wg.Done()
}

// Note: sync.WaitGroup is a struct, so it should be passed by reference.
