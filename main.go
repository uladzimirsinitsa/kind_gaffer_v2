package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"net/http"
)

type Stack []string

type Data struct {
	Code int8 `json:"code"`
	Url string `json:"url"`
	Status bool `json:"status"`
	Parsing_data string `json:"parsing_data"`
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

func create_stack() Stack {

}

