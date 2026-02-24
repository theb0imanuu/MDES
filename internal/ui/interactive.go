package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"mdes-cli/internal/engine"
)

type InteractiveUI struct {
	es     *engine.ExpertSystem
	reader *bufio.Reader
}

func NewInteractiveUI(es *engine.ExpertSystem) *InteractiveUI {
	return &InteractiveUI{
		es:     es,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ui *InteractiveUI) Run() {
	patientSymptoms := []string{}

	ui.printHeader()
	ui.printInstructions()

	categories := ui.getSymptomsByCategory()
	priorityOrder := []string{"General", "Pain", "Respiratory", "Cardiovascular", "Neurological", "Gastrointestinal", "Dermatological", "Infectious", "Viral", "Endocrine", "Autoimmune", "Oncological", "Psychiatric", "Psychological"}

	for _, category := range priorityOrder {
		symptoms, exists := categories[category]
		if !exists || len(symptoms) == 0 {
			continue
		}

		fmt.Printf("\n【%s SYMPTOMS】\n", strings.ToUpper(category))
		fmt.Println(strings.Repeat("-", 40))

		for _, sym := range symptoms {
			if ui.askSymptom(sym) {
				patientSymptoms = append(patientSymptoms, sym.ID)
				fmt.Printf("  ✓ Added: %s\n", sym.Name)
			}
		}
	}

	ui.performDiagnosis(patientSymptoms)
	ui.waitForExit()
}

func (ui *InteractiveUI) waitForExit() {
	fmt.Println("\nPress Enter to exit...")
	ui.reader.ReadString('\n')
}

func (ui *InteractiveUI) askSymptom(sym engine.Symptom) bool {
	fmt.Printf("Do you have %s? (%s) [y/n/s]: ", sym.Name, sym.Description)
	input, _ := ui.reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" {
		return true
	} else if input == "s" {
		fmt.Println("  Skipping remaining in category...")
		return false // This needs proper handling to skip category
	}
	return false
}

func (ui *InteractiveUI) performDiagnosis(symptoms []string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  ANALYZING SYMPTOMS...")
	fmt.Println(strings.Repeat("=", 60))

	diagnoses := ui.es.Diagnose(symptoms)

	if len(diagnoses) == 0 {
		ui.printNoDiagnosis()
		return
	}

	ui.printResults(diagnoses)
}

func (ui *InteractiveUI) printHeader() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  MEDICAL DIAGNOSTIC EXPERT SYSTEM (MDES) v2.0")
	fmt.Println("  Large Knowledge Base Implementation")
	fmt.Println(strings.Repeat("=", 60))
}

func (ui *InteractiveUI) printInstructions() {
	fmt.Println("\nINSTRUCTIONS:")
	fmt.Println("• Answer 'y' for yes, 'n' for no")
	fmt.Println("• Be accurate for reliable diagnosis")
	fmt.Println("• Preliminary assessment only - consult a doctor")
	fmt.Println(strings.Repeat("-", 60))
}

func (ui *InteractiveUI) printNoDiagnosis() {
	fmt.Println("\n⚠️  No matching conditions found.")
	fmt.Println("\nPossible reasons:")
	fmt.Println("• Symptoms don't match known patterns")
	fmt.Println("• Rare condition not in knowledge base")
	fmt.Println("• Insufficient symptom information")
	fmt.Println("\nRecommendation: Consult healthcare provider.")
}

func (ui *InteractiveUI) printResults(diagnoses []engine.DiagnosisResult) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  DIAGNOSTIC RESULTS")
	fmt.Println(strings.Repeat("=", 60))

	limit := 5
	if len(diagnoses) < limit {
		limit = len(diagnoses)
	}

	for i := 0; i < limit; i++ {
		d := diagnoses[i]
		ui.printDiagnosis(i+1, d)
	}

	if len(diagnoses) > 1 && diagnoses[0].Probability-diagnoses[1].Probability < 0.2 {
		fmt.Println("\n⚠️  DIFFERENTIAL DIAGNOSIS ADVISED")
	}

	ui.printDisclaimer()
}

func (ui *InteractiveUI) printDiagnosis(rank int, d engine.DiagnosisResult) {
	bar := renderConfidenceBar(d.Probability)

	fmt.Printf("\n%d. %s\n", rank, strings.ToUpper(d.Disease))
	fmt.Printf("   Probability: %.1f%% %s\n", d.Probability*100, bar)
	fmt.Printf("   Severity: %s\n", d.Severity)
	fmt.Printf("   Match: %d/%d required, %d/%d optional\n",
		d.MatchedRequired, d.TotalRequired, d.MatchedOptional, d.TotalOptional)

	if rank == 1 {
		fmt.Println("\n   🏥 RECOMMENDATIONS:")
		for _, rec := range d.Recommendations {
			fmt.Printf("   • %s\n", rec)
		}
	}
}

func (ui *InteractiveUI) printDisclaimer() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("DISCLAIMER: Decision support tool only.")
	fmt.Println("Final diagnosis requires clinical examination.")
	fmt.Println(strings.Repeat("=", 60))
}

func (ui *InteractiveUI) getSymptomsByCategory() map[string][]engine.Symptom {
	categories := make(map[string][]engine.Symptom)
	for _, sym := range ui.es.Symptoms {
		categories[sym.Category] = append(categories[sym.Category], sym)
	}
	return categories
}

func renderConfidenceBar(probability float64) string {
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
