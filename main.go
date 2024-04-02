package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Id        int    `json:"id"`
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Title     string `json:"title"`
	ESRB      string `json:"esrb"`
	Release   int    `json:"release"`
}

func main() {
	var quote Quote
	response, err := http.Get("https://ultima.rest/api/random")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &quote)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", quote.Character)
	fmt.Println(quote.Quote)
	fmt.Printf("%s, %d\n", quote.Title, quote.Release)
}
