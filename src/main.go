package main

import (
	"encoding/json"
	"jsondemo"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}

	todoDecoder := json.NewDecoder(resp.Body)
	todos := jsondemo.Todo{}
	todoDecoder.Decode(&todos)

	todoEncoder := json.NewEncoder(os.Stdout)
	todoEncoder.Encode(todos)

}
