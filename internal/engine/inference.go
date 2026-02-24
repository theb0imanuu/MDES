package engine

import (
	"fmt"
	"sort"
)

func (es *ExpertSystem) Diagnose(patientSymptoms []string) []DiagnosisResult {
	symptomSet := make(map[string]bool)
	for _, s := range patientSymptoms {
		symptomSet[s] = true
	}

	var results []DiagnosisResult

	for _, rule := range es.Rules {
		// Check exclusions
		excluded := false
		for _, excl := range rule.Exclusions {
			if symptomSet[excl] {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}

		requiredMatched := 0
		optionalMatched := 0
		explanation := []string{}
		totalWeight := 0.0
		matchedWeight := 0.0

		// Required conditions
		for _, cond := range rule.Conditions {
			totalWeight += rule.Weight[cond]
			if symptomSet[cond] {
				requiredMatched++
				matchedWeight += rule.Weight[cond]
				explanation = append(explanation, fmt.Sprintf("Present: %s", es.Symptoms[cond].Name))
			} else {
				explanation = append(explanation, fmt.Sprintf("Absent: %s", es.Symptoms[cond].Name))
			}
		}

		// Optional conditions
		for _, opt := range rule.Optional {
			totalWeight += rule.Weight[opt] * 0.5
			if symptomSet[opt] {
				optionalMatched++
				matchedWeight += rule.Weight[opt] * 0.5
				explanation = append(explanation, fmt.Sprintf("Supporting: %s", es.Symptoms[opt].Name))
			}
		}

		if requiredMatched > 0 {
			coverage := float64(requiredMatched) / float64(len(rule.Conditions))
			confidence := matchedWeight / totalWeight
			probability := coverage * confidence * rule.CF

			if coverage >= 0.75 {
				probability *= 1.1
			}
			if probability > 0.99 {
				probability = 0.99
			}

			disease := es.Diseases[rule.Disease]
			results = append(results, DiagnosisResult{
				Disease:         disease.Name,
				Probability:     probability,
				MatchedRequired: requiredMatched,
				TotalRequired:   len(rule.Conditions),
				MatchedOptional: optionalMatched,
				TotalOptional:   len(rule.Optional),
				Confidence:      confidence,
				Explanation:     explanation,
				Severity:        disease.Severity,
				Recommendations: disease.Recommendations,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Probability > results[j].Probability
	})

	return results
}
