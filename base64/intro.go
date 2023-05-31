package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
)

func encodingBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func decodingBase64(textBase64 string) string {
	decoded, err := base64.StdEncoding.DecodeString(textBase64)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}
	return string(decoded)
}

func decodeBase64(str string) string {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}
	dst = dst[:n]
	return string(dst)
}

func newEncoderPrintln(str string) {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	if _, err := encoder.Write([]byte(str)); err != nil {
		fmt.Println("encode error:", err)
	}
	_ = encoder.Close()
}

func encodingWithLenAnPrint(str string) string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	base64.StdEncoding.Encode(dst, []byte(str))
	return string(dst)
}

func urlDecoder(urlEncoded string) string {
	urlDecoded, err := url.QueryUnescape(urlEncoded)
	if err != nil {
		fmt.Println("URL decode error:", err)
		return ""
	}
	return urlDecoded
}

func urlEncoder(urlDecoded string) string {
	return url.QueryEscape(urlDecoded)
}

func main() {
	msg := "37485296354:999990"
	fmt.Println("Before:", msg)
	encoded := encodingBase64(msg)
	fmt.Println("Encoded:", encoded)
	decoded := decodingBase64(encoded)
	fmt.Println("After decoded  :", decoded)

	fmt.Println()
	fmt.Printf("%q\n", decodeBase64("SGVsbG8sIHdvcmxkIQ=="))
	fmt.Printf("%q\n", decodeBase64("c29tZSBkYXRhIHdpdGggACBhbmQg77u/"))
	fmt.Printf("%q\n", decodeBase64(encodingBase64("foo\x00bar")))

	newEncoderPrintln("foo\x00bar")
	fmt.Println()
	fmt.Printf("Hello, world! %v\n", encodingWithLenAnPrint("Hello, world!"))

	urlEncoded := "https%3A%2F%2Fwww.google.com%2Fsearch%3Fq%3Dtradutor%26oq%3Dtradutor%26sourceid%3Dchrome%26ie%3DUTF-8"
	urlDecoded := urlDecoder(urlEncoded)
	fmt.Println(urlDecoded)
	fmt.Println(urlEncoder(urlDecoded))

}
