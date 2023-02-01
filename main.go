package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"net/http"
	"time"
	"io"
	"encoding/json"
)


type TargetUrlsStack []string

type Data struct {
	Code int8 `json:"code"`
	Url string `json:"url"`
	Status bool `json:"status"`
	Parsing_data string `json:"parsing_data"`
}


type UrlsForTargetUrlsStack struct {
	Urls []string `json:"urls"`
}

func init() {
	godotenv.Load()
}

func (urls *TargetUrlsStack) IsEmpty() bool {
	return len(*urls) == 0
}

func (urls *TargetUrlsStack) Pop() (string, bool) {
	if urls.IsEmpty() {
		return "", false
	} else {
		index := len(*urls) - 1
		url := (*urls)[index]
		*urls = (*urls)[:index]
		return url, true
	}
}

func createTargetUrlsStack() UrlsForTargetUrlsStack {
	URL_DB := os.Getenv("URL_DB")
	timeout := time.Duration(6 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(URL_DB)
	if err != nil {
        log.Println(err)
		}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
        log.Println(err)
		}
	var targetUrls UrlsForTargetUrlsStack
	err = json.Unmarshal(body, &targetUrls)
	if err != nil {
        log.Println(err)
		}
	log.Print(targetUrls)
	return targetUrls
}

func main() {
	createTargetUrlsStack()
}