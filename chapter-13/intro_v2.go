package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// In Go, concurrent tasks are called goroutines.
// Goroutines require less computer memory than threads, and less time to start up and stop.
// To start another goroutine, you use a 'go' (keyword) statement, which is just an ordinary function or method call
// with the go keyword in front of it.

type Page struct {
	URL  string
	Size int
}

func responseSizeV2(url string, channel chan Page) {
	start := time.Now().UnixMilli()
	fmt.Printf("Getting %v\n", url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("x-api-key", "WvDicDZfMk87ujnFixy5V21QpVe7b8854hol3ezO")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		fmt.Printf("Closing connection for %v\t", url)
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Time: %d ms\n", time.Now().UnixMilli()-start)
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Content: %v\n", string(body))
	channel <- Page{
		URL:  url,
		Size: len(body),
	} // instead of returning the page size, send it over the channel
}

func main() {
	start := time.Now().UnixMilli()
	endpoints := []string{
		"https://771rf71vn1.execute-api.us-east-1.amazonaws.com/mock/usuarios/57012525827/dispositivo/862536030196001/dados-cadastrais-temp",
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}

	// make a channel of int value to sync goroutines
	sizes := make(chan Page)

	for _, endpoint := range endpoints {
		// Goroutines allow for concurrency: pausing one task to work on others.
		// And in some situations they allow parallelism: working on multiple tasks simultaneously!
		// The main function of every Go program is started using a goroutine, so every Go program runs at least one goroutine
		// Go programs stop running as soon as the main goroutine (the goroutine that calls the main function) ends, even if other
		// goroutines are still running. Our main function completes before the code in the others responseSize functions has a chance to run
		// We need to keep the main goroutine running until the goroutines for the others responseSize functions can finish.
		// To do this properly, we’re going to need another feature of Go called channels
		// Go makes no guarantees about when it will
		// switch between goroutines, or for how long. This allows goroutines to run
		// in whatever way is most efficient. But if order your goroutines run in is
		// important to you, you’ll need to synchronize them using channels
		// we can’t use function return values in a go statement
		// Not only do channels allow you to send values from one goroutine to another, they
		//ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it
		// The only practical way to use a channel is to communicate from one goroutine to another goroutine
		// 1 - Create a channel.
		// 2 - Write a function that receives a channel as a parameter.
		// We’ll run this function in a separate goroutine, and use it to send values over the channel.
		// 3 - Receive sent values in our original goroutine
		// Each channel only carries values of a particular type, so you might have one
		// channel for int values, and another channel for values with a struct type.
		// To declare a variable that holds a channel, you use the chan keyword, followed by the type of values that channel will carry
		// Ex: var myChannel chan float64 => declare a variable to hold a channel.
		// Ex: myChannel = make(chan float64) => actually create the channel
		// To send a value on a channel, you use the <- operator (that’s a less-than
		// symbol followed by a dash). It looks like an arrow pointing from the value you’re sending to the channel you’re sending it on
		// Ex: myChannel <- 3.14
		// You also use the <- operator to receive values from a channel, but the positioning is different:
		// you place the arrow to the left of the channel you’re receiving from. (It kind of looks like you’re pulling a value out of the channel.)
		// Ex: <- myChannel
		// We mentioned that channels also ensure the sending goroutine has sent the
		// value before the receiving channel attempts to use it. Channels do this by
		// blocking—by pausing all further operations in the current goroutine. A
		// send operation blocks the sending goroutine until another goroutine
		// executes receive operation on the same channel. And vice versa: receive
		// operation blocks the receiving goroutine until another goroutine executes a
		// send operation on the same channel. This behavior allows goroutines to
		// synchronize their actions—that is, to coordinate their timing
		go responseSizeV2(endpoint, sizes)
	}

	for _, _ = range endpoints {
		// receives the values from channel
		page := <-sizes
		fmt.Printf("Receiving size '%d' value from %v\n", page.Size, page.URL)
	}
	fmt.Printf("Total time: %d ms\n", time.Now().UnixMilli()-start)
}
