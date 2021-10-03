package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}

	request(client, "GET", "http://127.0.0.1:8080/course/1", "")
	request(client, "POST", "http://127.0.0.1:8080/course", `{
		"name": "string",
		"expectedWorkload": 0
	}`)
	request(client, "PUT", "http://127.0.0.1:8080/course/1", `{
		"name": "string",
		"expectedWorkload": 0
	}`)
	request(client, "DELETE", "http://127.0.0.1:8080/course/1", "")
}

func request(client *http.Client, method string, url string, body string) {
	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = strings.NewReader(body)
	}

	fmt.Printf("Requesting [%s] %s\n", method, url)
	req, rErr := http.NewRequest(method, url, bodyReader)
	if rErr != nil {
		fmt.Println(rErr)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	fmt.Println()
}
