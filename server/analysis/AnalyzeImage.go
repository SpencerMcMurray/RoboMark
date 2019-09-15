package analysis

import (
	"encoding/json"
	"os/exec"
	"strconv"
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
			"en", strconv.Itoa(i + 1), "",
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
func AnalyzeImage(file string) (ImageAnalysis, error) {
	PyAnalyze := exec.Command("python3", "./server/analysis/AnalyzeImage.py", file)

	out, err := PyAnalyze.Output()
	if err != nil {
		return ImageAnalysis{}, err
	}

	dataStr := strings.Replace(string(out), "'", "\"", -1)
	dataBytes := []byte(dataStr)

	var analysis ImageAnalysis
	json.Unmarshal(dataBytes, &analysis)

	return analysis, nil
}
