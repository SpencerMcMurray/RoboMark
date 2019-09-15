package main

import (
	"RoboMark/server/data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	analysis "RoboMark/server/analysis"
)

const port = 8080
const root = "RoboMark Server Root"

const entities = "entities"
const keyPhrases = "keyPhrases"

const POST = "POST"

// AnalyzeTextHandler .
func AnalyzeTextHandler(image analysis.ImageAnalysis) error {
	log.Println("Analyzing text.")

	inputDoc := analysis.ImageAnalysisToInputDoc(image)
	result, err := analysis.AnalyzeText(keyPhrases, inputDoc)
	if err != nil {
		return err
	}

	log.Println("Finished analyzing text.")
	fmt.Println(result)

	return nil
}

// AnalyzeImageHandler .
func AnalyzeImageHandler() analysis.ImageAnalysis {
	log.Println("Analyzing image.")

	result, err := analysis.AnalyzeImage("https://upload.wikimedia.org/wikipedia/commons/d/dd/Cursive_Writing_on_Notebook_paper.jpg")
	if err != nil {
		panic(err)
	}

	log.Println("Finished analyzing image.")
	return result
}

// Analyze .
func Analyze(writer http.ResponseWriter, request *http.Request) {
	result := AnalyzeImageHandler()
	AnalyzeTextHandler(result)
}

// Root .
func Root(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(root))
}

// AddQuestion .
func AddQuestion(wrt http.ResponseWriter, req *http.Request) {
	if req.Method != POST {
		return
	}

	decoder := json.NewDecoder(req.Body)

	var question data.Question
	err := decoder.Decode(&question)
	if err != nil {
		panic(err)
	}

	log.Println("QUESTION", question)

	// INSERT QUESTION
}

func main() {

	fmt.Println("RoboMark server starting up.")
	http.HandleFunc("/", Root)
	http.HandleFunc("/analyze/", Analyze)
	http.HandleFunc("/question/", AddQuestion)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
