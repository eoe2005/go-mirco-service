package server

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	Get("/a", func(u string) {
		fmt.Printf(" -- > %s \n", u)
	})
	Run()
}
