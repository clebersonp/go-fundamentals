package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const SIGNATURE string = "signature"

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getGuestbook(path string) Guestbook {
	var signatures []string
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		log.Printf("There is not a file named %v\n", path)
		return Guestbook{}
	}
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		signatures = append(signatures, scanner.Text())
	}
	check(scanner.Err())
	return Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
}

func putSignature(path string, signature string) {
	if len(signature) > 0 {
		options := os.O_APPEND | os.O_CREATE | os.O_WRONLY
		file, err := os.OpenFile(path, options, os.FileMode(0644))
		check(err)
		_, err = fmt.Fprintln(file, signature)
		check(err)
		err = file.Close()
		check(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	guestbook := getGuestbook("./signatures.txt")
	log.Printf("%#v\n", guestbook)
	// text/template is less security than html/template
	// If there are any script in html and use text/template, it'll be injected to html page
	// Inserting script tags like this is just one of many ways unscrupulous users
	// can insert malicious code into your web pages. The html/template
	// package makes it easy to protect against this and many other attacks!
	html, err := template.ParseFiles("view.html")
	check(err)
	// A Template value’s Execute method takes a value that
	// satisfies the io.Writer interface, and a data value that can be
	// accessed within actions in the template
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		err := request.ParseForm()
		check(err)
		fmt.Printf("Post from website! request.PostForm = %v\n", request.PostForm)
		newSignature := request.FormValue(SIGNATURE)
		fmt.Printf("New Signature: %v\n", newSignature)
		putSignature("./signatures.txt", newSignature)
	}
	// redirect the browser to the main guestbook page
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}
