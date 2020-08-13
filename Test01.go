package main

import (
	"fmt"

	"github.com/eoe2005/go-mirco-service/server"
)

func main() {

	server.Get("/a", func(u server.GData) {
		fmt.Printf(" -- > %s \n", u)
	})
	server.Any("/b", func(u server.GData) {
		fmt.Printf(" -- > %s \n", u)
	})
	server.Get("/c", func(u server.GData) {
		fmt.Printf(" -- > %s \n", u)
	})
	server.Run()
}
