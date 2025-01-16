package main

import (
	"fmt"
	"net/http"
)

// Handlers(ResponseWriter, Request): Receive and writes to specified writer.
// By following this signature, the function implements the handler interface.
// The http.Request struct can be large; reference avoids overhead.
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n") // The HTTP connection output, which directs to browser, etc..
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// HandleFunc: Registers handler functions to the URL route.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	// ListenAndServe: Listens to specified port, with specified handler.
	// If handler is nil, the default ServeMux is used as handler.
	http.ListenAndServe(":8010", nil)
}

// http.Request
// .Method: HTTP method type (GET/POST)
// .URL: Request URL (sub divide into .Path, .Query())
// .Header: HTTP headers (Content-Type)
// .Body: Request body
// .Context(): Context for cancellation/deadlines/values
