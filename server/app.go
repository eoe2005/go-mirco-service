package main

import (
	"fmt"

	server "../server"
)

func main() {
	server.Get("/a", func(u string) {
		fmt.Printf(" -- > %s \n", u)
	})
	server.Run()
}
