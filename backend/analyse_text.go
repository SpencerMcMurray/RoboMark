package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	var subscriptionKeyVar string = "TEXT_ANALYTICS_SUBSCRIPTION_KEY"
	if "" == os.Getenv(subscriptionKeyVar) {
		log.Fatal("Please set/export the environment variable " + subscriptionKeyVar + ".")
	}

	var subscriptionKey string = os.Getenv(subscriptionKeyVar)
	var endpointVar string = "TEXT_ANALYTICS_ENDPOINT"
	if "" == os.Getenv(endpointVar) {
		log.Fatal("Please set/export the environment variable " + endpointVar + ".")
	}
	var endpoint string = os.Getenv(endpointVar)

	const uriPath = "/text/analytics/v2.1/languages"

	var uri = endpoint + uriPath

	data := []map[string]string{
		{"id": "1", "text": "This is a document written in English."},
		{"id": "2", "text": "Este es un document escrito en Español."},
		{"id": "3", "text": "这是一个用中文写的文件"},
	}

	documents, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		return
	}

	r := strings.NewReader("{\"documents\": " + string(documents) + "}")

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var f interface{}
	json.Unmarshal(body, &f)

	jsonFormatted, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		fmt.Printf("Error producing JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonFormatted))
}
