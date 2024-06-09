package main

import "fmt"

func securityFilterSolution() {
	greeting, err := filterGreet("BeagleBoy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Here is my greeting:", greeting)
}

func filterGreetSolution(name string) (greeting string, err error) {
	msg, err := greetFilterSolution(name)
	if name == "Huey" || name == "Dewey" || name == "Louie" {
		greeting = msg
	} else {
		err = fmt.Errorf("not a legit nephew : %q", name)
		greeting = fmt.Sprintf("No greeting because of error [%v]", err.Error())
	}
	return
}

func greetFilterSolution(name string) (greeting string, err error) {
	return fmt.Sprint("Hello ", name), nil
}
