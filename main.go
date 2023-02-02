package main

import (
	"log"
	"net/http"
	"time"
	"io"
	"encoding/json"
	"github.com/joho/godotenv"
)


type Stack []string

type JsonUrls struct {
	Urls []string `json:"urls"`
}

var urls = createTargetUrlsStack()

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
	var EnvVariables map[string]string
	EnvVariables, _ = godotenv.Read()
	URL_DB := EnvVariables["URL_DB"]
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
	for i := 0; i < 2; i++	{
		go thread()
	}
	thread()
}