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
	Status             string              `json:"status"`
	RecognitionResults []RecognitionResult `json:"recognitionResults"`
}

// ImageAnalysisToInputDoc .
func ImageAnalysisToInputDoc(imageAnalysis ImageAnalysis) (textInput TextInput) {
	for i, recRes := range imageAnalysis.RecognitionResults {
		inputDoc := InputDocument{
			"en", string(i + 1), "",
		}
		for j, line := range recRes.Lines {
			if j > 0 {
				inputDoc.Text = inputDoc.Text + " "
			}
			inputDoc.Text = inputDoc.Text + line.Text
		}
		textInput.Documents = append(textInput.Documents, inputDoc)
	}
	return
}

// AnalyzeImage calls the Azure API to perform image analysis.
func AnalyzeImage(imageURL string) (analysis ImageAnalysis) {
	PyAnalyze := exec.Command("python3", "./server/analysis/AnalyzeImage.py", imageURL)

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
