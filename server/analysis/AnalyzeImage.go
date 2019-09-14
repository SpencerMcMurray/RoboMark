package analysis

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

// Word .
type Word struct {
	BoundingBox []int  `json:"boundingBox"`
	Text        string `json:"text"`
}

// Line .
type Line struct {
	BoundingBox []int  `json:"boundingBox"`
	Text        string `json:"text"`
	Words       []Word `json:"words"`
}

// RecognitionResult .
type RecognitionResult struct {
	ClockwiseOrientation float32 `json:"clockwiseOrientation"`
	Height               int     `json:"height"`
	Unit                 string  `json:"unit"`
	Width                int     `json:"width"`

	Lines []Line `json:"lines"`
	Page  int    `json:"page"`
}

// ImageAnalysis .
type ImageAnalysis struct {
	Status            string              `json:"status"`
	RecognitionResult []RecognitionResult `json:"recognitionResults"`
}

// AnalyzeImage calls the Azure API to perform image analysis.
func AnalyzeImage() (analysis ImageAnalysis) {
	PyAnalyze := exec.Command("python3", "./server/analysis/AnalyzeImage.py")

	out, err := PyAnalyze.Output()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	dataStr := strings.Replace(string(out), "'", "\"", -1)
	dataBytes := []byte(dataStr)

	json.Unmarshal(dataBytes, &analysis)
	return
}
