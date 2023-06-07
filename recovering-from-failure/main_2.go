package main

import "fmt"

func tryRecovery() {
	// recover function is a builtin golang function
	// the recover function can return to the normal execution any panic
	// anything pass to panic function can be return in the recovery function
	// both recover and panic function return/passed any(interface{})
	// so, we can pass any type of data to panic e recovery that later with the recover function
	// and using type assertion we can get the value again
	// with this function executes without any error, recover will return nil instead
	p := recover()
	if p == nil {
		return
	}
	if err, ok := p.(error); ok {
		fmt.Println("Recovering: " + err.Error())
	} else {
		// if p it's not an error type, so call panic again with the original value
		fmt.Printf("Type of problem: %T\n", p)
		panic(p)
	}
}

func main() {
	// call later with delay
	// is like a try catch finally in java
	// but, but
	// It can even be said that using panic and recover is discouraged by the design of the language itself.
	// In a conference keynote in 2012, Rob Pike (one of the creators of Go) described panic and recover as
	// “intentionally clumsy.” That means that when designing Go, its creators did not try to make panic and recover
	// easy or pleasant to use, so that they’d be used less often
	// https://go.dev/talks/2012/splash.article#TOC_16
	defer tryRecovery()
	err := fmt.Errorf("there's an error")
	panic(err)
	panic("some error")
}
