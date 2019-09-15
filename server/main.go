package main

import (
	"fmt"
	"net/http"

	analysis "RoboMark/server/analysis"
	data "RoboMark/server/data"
)

const port = 8080
const root = "RoboMark Server Root"

// AnalyzeTextHandler .
func AnalyzeTextHandler(image analysis.ImageAnalysis) error {
	return nil
}

// AnalyzeImageHandler .
func AnalyzeImageHandler(writer http.ResponseWriter, request *http.Request) {
	result := analysis.AnalyzeImage("https://upload.wikimedia.org/wikipedia/commons/d/dd/Cursive_Writing_on_Notebook_paper.jpg")
	fmt.Println(result)
	err := AnalyzeTextHandler(result)
	if err != nil {
		panic(err)
	}
}

// Root .
func Root(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(root))
}

func main() {

	data.ViewTests(1)

	return
	fmt.Println("RoboMark server starting up.")
	http.HandleFunc("/", Root)
	http.HandleFunc("/image/", AnalyzeImageHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
