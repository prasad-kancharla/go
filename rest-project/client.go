package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := &http.Client{}

	// resp, err := client.Get("https://www.example.com")
	// resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	resp, err := client.Get("https://swapi.dev/api/people/1")
	if err != nil {
		fmt.Println("Got error:", err)
		return
	}

	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Got error:", err)
		return
	}

	fmt.Println(string(resBytes))

	defer resp.Body.Close()
}
