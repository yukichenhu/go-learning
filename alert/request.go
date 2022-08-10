package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8081/test/a"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{})
	_ = json.Unmarshal(content, &result)
	fmt.Println(result["name"])
}
