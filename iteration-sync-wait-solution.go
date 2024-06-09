// The program will panic with the following error:
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc0000b2000)
//
//	/usr/local/go/src/runtime/sema.go:56 +0x42
//
// sync.(*WaitGroup).Wait(0xc0000b2000)
//
//	/usr/local/go/src/sync/waitgroup.go:130 +0x64
//
// main.main()
//
//	/tmp/sandbox282013366/main.go:17 +0x1b3
//
// exit status 2
// The issue is that the WaitGroup is being passed by value to the greet function.
// This means that each goroutine is working with a copy of the WaitGroup.
// When the greet function calls wg.Done(), it is decrementing the counter of the copy of the WaitGroup.
// The counter of the original WaitGroup is never decremented, so the main goroutine is waiting forever.
// To fix this issue, the WaitGroup should be passed by reference to the greet function.
// The greet function signature should be:
// func greet(name string, wg *sync.WaitGroup) {
// and the call to greet should be:
// go greet(nephew, &wg)
package main

import (
	"fmt"
	"sync"
)

// Can't declare a "const" array, so this is a "var"
var nephews = [...]string{"Huey", "Dewey", "Louie"}

func iterateSyncWaitSolution() {
	var wg sync.WaitGroup

	for _, nephew := range nephews {
		wg.Add(1)
		go greetSolution(nephew, &wg)
	}

	// Wait for all greetings to complete.
	wg.Wait()
}

func greetSolution(name string, wg *sync.WaitGroup) {
	fmt.Println("Hello", name)
	wg.Done()
}
