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


func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

type Stack []string

type JsonUrls struct {
	Urls []string `json:"urls"`
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

func (urls *Stack) Push(str string) {
	*urls = append(*urls, str)
}

func createTargetUrlsStack() Stack {
	URL_DB := os.Getenv("URL_DB")
	timeout := time.Duration(6 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(URL_DB)
	if err != nil {
        log.Println(err)
		}
	body, err := io.ReadAll(response.Body)
	if err != nil {
        log.Println(err)
		}
	var targetUrls JsonUrls
	err = json.Unmarshal(body, &targetUrls)
	if err != nil {
        log.Println(err)
		}
	defer response.Body.Close()
	urls := Stack(targetUrls.Urls)
	return urls
}

var urls = createTargetUrlsStack()

func thread() {
	for {
		url, _ := urls.Pop()
		if url == "" {
			break
		}
		log.Println(url)
	}
}

func main() {
	for i := 0; i < 36; i++	{
		go thread()
	}
	thread()
}