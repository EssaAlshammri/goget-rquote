package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// calling the api witha get method
	data := map[string]interface{}{}
	resp, err := http.Get("http://api.forismatic.com/api/1.0/?method=getQuote&key=457653&format=json&lang=en")
	defer resp.Body.Close()
	// here it handle if the Get return an error
	if err != nil {
		os.Exit(0)
	} else {
		// here we should read the response and print it out
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			os.Exit(0)
		}
		// bounding data to the variable called data ^^
		json.Unmarshal(body, &data)
		if len(data) < 1 {
			os.Exit(0)
		}
		fmt.Print(data["quoteText"])
		if data["quoteAuthor"] != "" {
			fmt.Println(" ----", data["quoteAuthor"])
		}
	}
}
