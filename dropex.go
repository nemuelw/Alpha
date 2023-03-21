package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello friend")
}

func has_internet_access() bool {
	_, err := http.Get("https://www.google.com")
	return err == nil
}