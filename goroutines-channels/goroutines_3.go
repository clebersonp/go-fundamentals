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

type Page struct {
	URL  string
	Size int
}

func (p Page) String() string {
	return fmt.Sprintf("URL: %s, Size: %d", p.URL, p.Size)
}

func responseSizeWithChannel(url string, channel chan Page) {
	log.Printf("Request. Url: %s, current: %v\n", url, getGoRoutineIdWithChannel())
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
	// If you need a value from a goroutine, you’ll need to pass it a channel to send the value back on
	channel <- Page{
		URL:  url,
		Size: len(resBytes),
	} // sending the value to the channel will block until the main goroutine receive this value
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	start := time.Now().UnixMilli()
	log.Println("Starting...")

	// that is not possible return any value from function with key word go (goroutine)
	// because there’s no guarantee the return value will be ready before we attempt to use it
	// with channel is possible to communicate between goroutines and synchronize the execution order
	sizes := make(chan Page)
	urls := []string{
		"https://example.com/",
		"https://go.dev/talks/2012/splash.article#TOC_16",
		"https://go.dev/",
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://www.uol.com.br/",
		"https://www.terra.com.br/",
		"https://www.estadao.com.br/",
		"https://www.telerisco.com.br/",
	}
	for _, url := range urls {
		go responseSizeWithChannel(url, sizes)
	}

	log.Println("Main function:", getGoRoutineIdWithChannel())

	for i := 0; i < len(urls); i++ {
		log.Println(<-sizes) // receiving the channel value will release the sender goroutine
		//time.Sleep(2 * time.Second) // only for debug propose
	}
	log.Printf("Done with %v ms", time.Now().UnixMilli()-start)
}

func getGoRoutineIdWithChannel() string {
	stack := debug.Stack()
	fields := bytes.Fields(stack)
	return fmt.Sprintf("%v %v", string(fields[0]), string(fields[1]))
}
