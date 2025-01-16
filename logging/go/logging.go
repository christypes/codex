package main

import "log"

func main() {
	log.Println("standard logger")

	log.SetFlags((log.LstdFlags | log.Lmicroseconds))
	log.Println("with micro")
}
