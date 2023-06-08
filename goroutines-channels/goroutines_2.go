package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func responseSize(url string) {
	fmt.Printf("Request. Url: %s, current: %v\n", url, getGoRoutineId())
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// execute this anonymous function as soon as possible before return responseSize function
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response. Url: %s, size: %d\n", url, len(resBytes))
}

func main() {
	// that is not possible return any value from function with key word go (goroutine)
	// because there’s no guarantee the return value will be ready before we attempt to use it
	go responseSize("https://example.com/")
	go responseSize("https://go.dev/talks/2012/splash.article#TOC_16")
	go responseSize("https://go.dev/")

	time.Sleep(1 * time.Second)
	fmt.Println(getGoRoutineId())
	fmt.Println("Done")
}

func getGoRoutineId() string {
	stack := debug.Stack()
	fields := bytes.Fields(stack)
	return fmt.Sprintf("%v %v", string(fields[0]), string(fields[1]))
}
