package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// 1. Issues Get request
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 2. Prints response status based on returned response
	fmt.Println("Response status:", resp.Status)

	// 3. Scans the body 5 lines.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
