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


type Stack []string

type Data struct {
	Code int8 `json:"code"`
	Url string `json:"url"`
	Status bool `json:"status"`
	Parsing_data string `json:"parsing_data"`
}


type UrlsForStack struct {
	Urls []string `json:"urls"`
}

func init() {
	godotenv.Load()
}

func (urls *Stack) IsEmpty() bool {
	return len(*urls) == 0
}

func (urls *Stack) Pop() (string, bool) {
	if urls.IsEmpty() {
		return "", false
	} else {
		index := len(*urls) - 1
		url := (*urls)[index]
		*urls = (*urls)[:index]
		return url, true
	}
}

func createStack() {
	URL_DB := os.Getenv("URL_DB")
	timeout := time.Duration(6 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(URL_DB)
	if err != nil {
        log.Println(err)
		}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var data UrlsForStack
	err = json.Unmarshal(body, &data)
	log.Print(data)

}

func main() {
	create_stack()
}