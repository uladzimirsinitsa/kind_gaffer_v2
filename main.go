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


type schemeJSON struct {
	Urls []string `json:"urls"`
	Record string `json:"record"`
}


//var urls = createTargetUrlsStack()

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

/*
func createTargetUrlsStack() Stack {
	var envVariables map[string]string
	envVariables, _ = godotenv.Read()
	URL_DB := envVariables["URL_DB"]
	timeout := time.Duration(6 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(URL_DB)
	if err != nil {
        log.Println(err)
		}
	body, err := io.ReadAll(response.Body)
	// FIX io.ReadAll(response.Body)
	if err != nil {
        log.Println(err)
		}
	var dataJSON schemeJSON
	err = json.Unmarshal(body, &dataJSON)
	if err != nil {
        log.Println(err)
		}
	defer response.Body.Close()
	urls := Stack(dataJSON.Urls)
	return urls
}
*/

func makeRequest(url string) ([]byte, bool) {
	timeout := time.Duration(6 * time.Second)
	client := http.Client{Timeout: timeout}
	response, err := client.Get(url)
	if err != nil {
		log.Print("GET error:", err)
		return []byte(""), false
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return body, true
}


func serializeJSON(body []byte) schemeJSON {
	var dataJSON schemeJSON
	err := json.Unmarshal(body, &dataJSON)
	if err != nil {
        log.Println(err)
		}
	return dataJSON
}


/* TODO
func createWorkersStack()
*/

/* TODO
func updateDB()
*/

/* TODO
func reportError()
*/

func thread() {
	for {
		url, _ := urls.Pop()
		if url == []byte("") {
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