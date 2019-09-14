package analysis

import (
	"fmt"
	"os/exec"
)

// AnalyseImage calls the Azure API to perform image analysis.
func AnalyseImage() {
	PyAnalyze := exec.Command("python3", "./server/analysis/AnalyzeImage.py")

	out, err := PyAnalyze.Output()
	if err != nil {
		panic(err)
	}

	// dataStr := strings.Replace(string(out), "  ", "", -1)

	fmt.Println(string(out))

}
