package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"

	"jsondemo"
	"os"
)

func main() {
	// ## read from get http
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	// if err != nil {
	// 	return
	// }

	// ## read from local file
	resp, err := ioutil.ReadFile("/Users/onebluesky882/local_files/Learning/golearn/golang_foundation/src/todo.html")
	if err != nil {
		return
	}

	todoDecoder := json.NewDecoder(bytes.NewReader(resp))
	todos := jsondemo.Todo{}
	todoDecoder.Decode(&todos)

	// indexTemplateBytes, err := ioutil.ReadFile("/Users/onebluesky882/local_files/Learning/golearn/golang_foundation/src/index.html")
	// if err != nil {
	// 	return
	// }
	indexTemplate := template.Must(template.New("index").Parse(string(resp)))
	indexTemplate.Execute(os.Stdout, todos)
}
