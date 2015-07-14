package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Challenge struct {
	Follow string
}

func fixURL(url string) string {
	return strings.Replace(url, "challenge", "challenge.json", -1)
}

func followURL(url string) Challenge {
	destURL := fixURL(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", destURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	v := Challenge{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatalln(err)
	}

	return v
}
func main() {
	url := "http://letsrevolutionizetesting.com/challenge"

	for {
		archiveURL := url
		resp := followURL(url)
		fmt.Println(resp.Follow)
		if resp.Follow != archiveURL {
			followURL(resp.Follow)
			url = resp.Follow
		}
	}
}
