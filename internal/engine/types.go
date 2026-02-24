package engine

// Symptom represents a medical symptom
type Symptom struct {
	ID          string
	Name        string
	Description string
	Category    string
	Severity    string
}

// Disease represents a medical condition
type Disease struct {
	ID              string
	Name            string
	Description     string
	Category        string
	Conditions      []string
	Optional        []string
	CF              float64
	Exclusions      []string
	Severity        string
	Recommendations []string
}

// Rule represents an inference rule
type Rule struct {
	Disease    string
	Conditions []string
	Optional   []string
	Exclusions []string
	CF         float64
	Weight     map[string]float64
}

// DiagnosisResult holds inference output
type DiagnosisResult struct {
	Disease         string
	Probability     float64
	MatchedRequired int
	TotalRequired   int
	MatchedOptional int
	TotalOptional   int
	Confidence      float64
	Explanation     []string
	Severity        string
	Recommendations []string
}
