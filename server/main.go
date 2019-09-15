package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	analysis "RoboMark/server/analysis"
)

const port = 8080
const root = "RoboMark Server Root"

const entities = "entities"
const keyPhrases = "keyPhrases"

// POST protocol.
const POST = "POST"

// AnalyzeTextHandler .
func AnalyzeTextHandler(image analysis.ImageAnalysis) (analysis.TextOutput, error) {
	log.Println("Analyzing text.")

	inputDoc := analysis.ImageAnalysisToInputDoc(image)
	result, err := analysis.AnalyzeText(keyPhrases, inputDoc)
	if err != nil {
		return analysis.TextOutput{}, err
	}

	log.Println("Finished analyzing text.")
	return result, nil
}

// AnalyzeImageHandler .
func AnalyzeImageHandler(file string) analysis.ImageAnalysis {
	log.Println("Analyzing image.")

	result, err := analysis.AnalyzeImage(file)
	if err != nil {
		panic(err)
	}

	log.Println("Finished analyzing image.")
	return result
}

// Analyze .
func Analyze(writer http.ResponseWriter, request *http.Request) {

	out, err := os.Create("photo.jpg")
	defer out.Close()
	file, header, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	_, err = io.Copy(out, file)
	if err != nil {
		panic(err)
	}

	imageResult := AnalyzeImageHandler("photo.jpg")

	fmt.Println(imageResult)
	textResult, err := AnalyzeTextHandler(imageResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(textResult)

	// resp, err := http.Get("http://localhost:8081/api/questions?test=")

}

// Root .
func Root(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(root))
}

func main() {
	fmt.Println("RoboMark server starting up.")
	http.HandleFunc("/", Root)
	http.HandleFunc("/analyze/", Analyze)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
