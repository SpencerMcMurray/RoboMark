package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// OutputDocument .
type OutputDocument struct {
	ID         string   `json:"id"`
	KeyPhrases []string `json:"keyPhrases"`
}

// TextOutput .
type TextOutput struct {
	Documents []OutputDocument `json:"documents"`
	Errors    []interface{}    `json:"errors"`
}

// InputDocument .
type InputDocument struct {
	Language string `json:"language"`
	ID       string `json:"id"`
	Text     string `json:"text"`
}

// TextInput .
type TextInput struct {
	Documents []InputDocument `json:"documents"`
}

// AnalyzeText .
func AnalyzeText(extension string, input TextInput) (TextOutput, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return TextOutput{}, err
	}

	var subscriptionKeyVar string = "TEXT_ANALYTICS_SUBSCRIPTION_KEY"
	if "" == os.Getenv(subscriptionKeyVar) {
		log.Fatal("Please set/export the environment variable " + subscriptionKeyVar + ".")
		return TextOutput{}, nil
	}

	var subscriptionKey string = os.Getenv(subscriptionKeyVar)
	var endpointVar string = "TEXT_ANALYTICS_ENDPOINT"
	if "" == os.Getenv(endpointVar) {
		log.Fatal("Please set/export the environment variable " + endpointVar + ".")
	}
	var endpoint string = os.Getenv(endpointVar)

	uriPath := "text/analytics/v2.1/" + extension
	var uri = endpoint + uriPath

	data := input

	documents, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		return TextOutput{}, err
	}

	r := strings.NewReader(string(documents))
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return TextOutput{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on request: %v\n", err)
		return TextOutput{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return TextOutput{}, err
	}

	var analysis TextOutput
	json.Unmarshal(body, &analysis)

	return analysis, nil
}
