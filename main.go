package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseHeaders struct {
	AcceptEncoding string // `json: Accept-Encoding`
	Host           string
	UserAgent      string // `json: User-Agent`
	XAmznTraceID   string // `json: X-Amzn-Trace-Id`
}

type resp struct {
	args map[string]string
	responseHeaders
	origin string
	url    string
}

// GetPage - return result from url
func GetPage(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func main() {

	urlPtr := flag.String("url", "https://httpbin.org/get", "url of content to fetch")
	flag.Parse()

	// fmt.Println("url: ", *urlPtr)

	page, err := GetPage(*urlPtr)
	if err != nil {
		fmt.Println("error retrieving url: ", err)
	}

	fmt.Println(page)
}

// sample response from test sent to http://httpbin.org
// {
// "args": {},
// "headers": {
//     "Accept-Encoding": "gzip",
//     "Host": "httpbin.org",
//     "User-Agent": "Go-http-client/2.0",
//     "X-Amzn-Trace-Id": "Root=1-5fb2dd0a-7db005a350e64f64552e1f09"
// },
// "origin": "38.75.198.14",
// "url": "https://httpbin.org/get"
// }
