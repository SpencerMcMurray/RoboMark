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

// AnalyseImage calls the Azure API to perform image analysis.
func AnalyseImage() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Add your Computer Vision subscription key and endpoint to your environment variables.
	subscriptionKey := os.Getenv("COMPUTER_VISION_SUBSCRIPTION_KEY")
	if subscriptionKey == "" {
		log.Fatal("\n\nSet the COMPUTER_VISION_SUBSCRIPTION_KEY environment variable.\n" +
			"**Restart your shell or IDE for changes to take effect.**\n")

	}

	endpoint := os.Getenv("COMPUTER_VISION_ENDPOINT")
	if "" == endpoint {
		log.Fatal("\n\nSet the COMPUTER_VISION_ENDPOINT environment variable.\n" +
			"**Restart your shell or IDE for changes to take effect.**")
	}

	const imageURL = "https://upload.wikimedia.org/wikipedia/commons/3/3c/Shaki_waterfall.jpg"
	const params = "?visualFeatures=Description&details=Landmarks&language=en"
	const imageURLEnc = "{\"url\":\"" + imageURL + "\"}"

	var uriBase = endpoint + "vision/v2.0/analyze"
	var uri = uriBase + params

	reader := strings.NewReader(imageURLEnc)

	// Create the HTTP client
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	// Create the POST request, passing the image URL in the request body
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		panic(err)
	}

	// Add request headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body
	// Note, data is a byte array
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON data from the byte array
	var f interface{}
	json.Unmarshal(data, &f)

	// Format and display the JSON result
	jsonFormatted, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(jsonFormatted))
}
