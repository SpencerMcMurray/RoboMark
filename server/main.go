package main

import (
	analysis "RoboMark/server/analysis"
)

func main() {
	// http.HandleFunc("/", hello)
	// http.ListenAndServe(":8080", nil)

	analysis.AnalyseImage()
	// analysis.AnalyseText()
}
