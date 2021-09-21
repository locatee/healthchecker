package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected URL as command-line argument")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Println(url)
	if resp, err := http.Get(url); err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
}
