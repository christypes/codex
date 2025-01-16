package main

import (
	"fmt"
	"net/http"
	"time"
)

// c.Context: Manages request-scoped values, cancellation signals, deadlines across API/goroutines.
// Often passed through functions to propogate cancellation/timeouts between functions.

// c.Context: HTTP specific context
// c.Background: Creates new context base
// c.WithCancel(c): Creates cancellable context
// c.WithTimeout(c, duration): Adds timeout to context
// c.WithValue(c, key, value): Attaches request-scoped data.

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context() // creates HTTP context.
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Blocks until
	select {
	case <-time.After(10 * time.Second): // 10 seconds passed
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): // or errored
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)  // on localhost"/hello"; run hello
	http.ListenAndServe(":8090", nil) //
}
