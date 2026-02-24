package ui

import (
	"fmt"
	"strings"

	"mdes-cli/internal/engine"
)

// RenderConfidenceBar creates a visual representation of the probability
func RenderConfidenceBar(probability float64) string {
	filled := int(probability * 20)
	bar := "["
	for i := 0; i < 20; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	return bar
}

// ExplainDiagnosis prints detailed information about the diagnosis
func ExplainDiagnosis(d engine.DiagnosisResult) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  DETAILED EXPLANATION")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Diagnosis: %s\n", d.Disease)
	fmt.Printf("Certainty Factor: %.2f\n", d.Confidence)
	fmt.Printf("Calculated Probability: %.2f%%\n", d.Probability*100)

	fmt.Println("\nSymptom Analysis:")
	for _, exp := range d.Explanation {
		if strings.HasPrefix(exp, "Present") {
			fmt.Printf("  ✓ %s\n", strings.TrimPrefix(exp, "Present: "))
		} else if strings.HasPrefix(exp, "Supporting") {
			fmt.Printf("  + %s (supporting)\n", strings.TrimPrefix(exp, "Supporting: "))
		}
	}
}
