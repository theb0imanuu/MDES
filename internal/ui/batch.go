package ui

import (
	"fmt"
	"strings"

	"mdes-cli/internal/engine"
)

// RunBatchDiagnose allows testing with predefined symptom sets
func RunBatchDiagnose(es *engine.ExpertSystem, name string, symptoms []string) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("TEST CASE: %s\n", name)
	fmt.Printf("Symptoms: %v\n", symptoms)

	results := es.Diagnose(symptoms)

	if len(results) > 0 {
		fmt.Printf("Top Diagnosis: %s (%.1f%%)\n", results[0].Disease, results[0].Probability*100)
	} else {
		fmt.Println("No diagnosis possible")
	}
}
