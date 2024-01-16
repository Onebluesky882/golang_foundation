# json http/get

url := "<https://jsonplaceholder.typicode.com/todos>"

resp, err := http.Get(url)
if err != nil {
return
}

body, err := io.ReadAll(resp.Body)
if err != nil {
return
}

dataStruct := []jsondemo.Todo{}
v := &dataStruct
json.Unmarshal([]byte(body), v)

// edit json
dataStruct[0].Title = "new title"

// dataStruct -> std.Output
result, err := json.MarshalIndent(dataStruct, "", " ")
if err != nil {
return
}

fmt.Println(string(result))
defer resp.Body.Close()
