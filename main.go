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
var dataJSON schemeJSON

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

/*func thread() {
	for {
		url, _ := urls.Pop()
		if url == []byte("") {
			break
		}
		log.Println(url)
	}
}*/

func thread() {
	log.Println(dataJSON)
}

func main() {
	var envVariables map[string]string
	envVariables, _ = godotenv.Read()
	URL_DB := envVariables["URL_DB"]
	body, _ := makeRequest(URL_DB)
	dataJSON := serializeJSON(body)
	log.Println(dataJSON)
	//for i := 0; i < 2; i++	{
		//go thread()
	//}
	thread()
}